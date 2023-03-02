package hunter

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/common"
	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

var TalentTreeSizes = [3]int{26, 27, 28}

const ThoridalTheStarsFuryItemID = 34334

func RegisterHunter() {
	core.RegisterAgentFactory(
		proto.Player_Hunter{},
		proto.Spec_SpecHunter,
		func(character core.Character, options *proto.Player) core.Agent {
			return NewHunter(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Hunter)
			if !ok {
				panic("Invalid spec value for Hunter!")
			}
			player.Spec = playerSpec
		},
	)
}

type Hunter struct {
	core.Character

	Talents  *proto.HunterTalents
	Options  *proto.Hunter_Options
	Rotation *proto.Hunter_Rotation

	pet *HunterPet

	AmmoDPS                   float64
	AmmoDamageBonus           float64
	NormalizedAmmoDamageBonus float64

	currentAspect *core.Aura

	// Used for deciding when we can use hawk for the rest of the fight.
	manaSpentPerSecondAtFirstAspectSwap float64
	permaHawk                           bool

	// The most recent time at which moving could have started, for trap weaving.
	mayMoveAt time.Duration

	AspectOfTheHawk  *core.Spell
	AspectOfTheViper *core.Spell

	AimedShot       *core.Spell
	ArcaneShot      *core.Spell
	BlackArrow      *core.Spell
	ChimeraShot     *core.Spell
	ExplosiveShotR4 *core.Spell
	ExplosiveShotR3 *core.Spell
	ExplosiveTrap   *core.Spell
	KillCommand     *core.Spell
	//KillShot        *core.Spell
	MultiShot     *core.Spell
	RapidFire     *core.Spell
	RaptorStrike  *core.Spell
	ScorpidSting  *core.Spell
	SerpentSting  *core.Spell
	SilencingShot *core.Spell
	SteadyShot    *core.Spell
	Volley        *core.Spell

	// Fake spells to encapsulate weaving logic.
	TrapWeaveSpell *core.Spell

	AspectOfTheHawkAura    *core.Aura
	AspectOfTheViperAura   *core.Aura
	ImprovedSteadyShotAura *core.Aura
	LockAndLoadAura        *core.Aura
	RapidFireAura          *core.Aura
	ScorpidStingAuras      core.AuraArray
	TalonOfAlarAura        *core.Aura

	CustomRotation *common.CustomRotation
}

func (hunter *Hunter) GetCharacter() *core.Character {
	return &hunter.Character
}

func (hunter *Hunter) HasMajorGlyph(glyph proto.HunterMajorGlyph) bool {
	return hunter.HasGlyph(int32(glyph))
}
func (hunter *Hunter) HasMinorGlyph(glyph proto.HunterMinorGlyph) bool {
	return hunter.HasGlyph(int32(glyph))
}

func (hunter *Hunter) GetHunter() *Hunter {
	return hunter
}

func (hunter *Hunter) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	if hunter.Talents.TrueshotAura {
		raidBuffs.TrueshotAura = true
	}
	if hunter.Talents.FerociousInspiration == 3 && hunter.pet != nil {
		raidBuffs.FerociousInspiration = true
	}
}
func (hunter *Hunter) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (hunter *Hunter) Initialize() {
	// Update auto crit multipliers now that we have the targets.
	hunter.AutoAttacks.MHConfig.CritMultiplier = hunter.critMultiplier(false, false)
	hunter.AutoAttacks.OHConfig.CritMultiplier = hunter.critMultiplier(false, false)
	hunter.AutoAttacks.RangedConfig.CritMultiplier = hunter.critMultiplier(false, false)

	hunter.registerAspectOfTheHawkSpell()
	hunter.registerAspectOfTheViperSpell()

	multiShotTimer := hunter.NewTimer()
	arcaneShotTimer := hunter.NewTimer()
	fireTrapTimer := hunter.NewTimer()

	hunter.registerAimedShotSpell(multiShotTimer)
	hunter.registerArcaneShotSpell(arcaneShotTimer)
	hunter.registerBlackArrowSpell(fireTrapTimer)
	hunter.registerChimeraShotSpell()
	hunter.registerExplosiveShotSpell(arcaneShotTimer)
	hunter.registerExplosiveTrapSpell(fireTrapTimer)
	//hunter.registerKillShotSpell()
	hunter.registerMultiShotSpell(multiShotTimer)
	hunter.registerRaptorStrikeSpell()
	hunter.registerScorpidStingSpell()
	hunter.registerSerpentStingSpell()
	hunter.registerSilencingShotSpell()
	hunter.registerSteadyShotSpell()
	hunter.registerVolleySpell()

	hunter.registerKillCommandCD()
	hunter.registerRapidFireCD()

	hunter.DelayDPSCooldownsForArmorDebuffs(time.Second * 10)

	hunter.CustomRotation = hunter.makeCustomRotation()
	if hunter.CustomRotation == nil {
		hunter.Rotation.Type = proto.Hunter_Rotation_SingleTarget
	}

	if hunter.Options.UseHuntersMark {
		hunter.RegisterPrepullAction(0, func(sim *core.Simulation) {
			huntersMarkAura := core.HuntersMarkAura(hunter.CurrentTarget, hunter.Talents.ImprovedHuntersMark, hunter.HasMajorGlyph(proto.HunterMajorGlyph_GlyphOfHuntersMark))
			huntersMarkAura.Activate(sim)
		})
	}
}

func (hunter *Hunter) Reset(sim *core.Simulation) {
	hunter.mayMoveAt = 0
	hunter.manaSpentPerSecondAtFirstAspectSwap = 0
	hunter.permaHawk = false
}

func NewHunter(character core.Character, options *proto.Player) *Hunter {
	hunterOptions := options.GetHunter()

	hunter := &Hunter{
		Character: character,
		Talents:   &proto.HunterTalents{},
		Options:   hunterOptions.Options,
		Rotation:  hunterOptions.Rotation,
	}
	core.FillTalentsProto(hunter.Talents.ProtoReflect(), options.TalentsString, TalentTreeSizes)
	hunter.EnableManaBar()

	hunter.PseudoStats.CanParry = true

	rangedWeapon := hunter.WeaponFromRanged(0)

	// Passive bonus (used to be from quiver).
	hunter.PseudoStats.RangedSpeedMultiplier *= 1.15

	if hunter.HasRangedWeapon() && hunter.GetRangedWeapon().ID != ThoridalTheStarsFuryItemID {
		switch hunter.Options.Ammo {
		case proto.Hunter_Options_TimelessArrow:
			hunter.AmmoDPS = 53
		case proto.Hunter_Options_MysteriousArrow:
			hunter.AmmoDPS = 46.5
		case proto.Hunter_Options_AdamantiteStinger:
			hunter.AmmoDPS = 43
		case proto.Hunter_Options_BlackflightArrow:
			hunter.AmmoDPS = 32
		}
		hunter.AmmoDamageBonus = hunter.AmmoDPS * rangedWeapon.SwingSpeed
		hunter.NormalizedAmmoDamageBonus = hunter.AmmoDPS * 2.8
	}

	hunter.EnableAutoAttacks(hunter, core.AutoAttackOptions{
		// We don't know crit multiplier until later when we see the target so just
		// use 0 for now.
		MainHand: hunter.WeaponFromMainHand(0),
		OffHand:  hunter.WeaponFromOffHand(0),
		Ranged:   rangedWeapon,
		ReplaceMHSwing: func(sim *core.Simulation, _ *core.Spell) *core.Spell {
			return hunter.TryRaptorStrike(sim)
		},
		AutoSwingRanged: true,
	})
	hunter.AutoAttacks.RangedConfig.ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
		baseDamage := hunter.RangedWeaponDamage(sim, spell.RangedAttackPower(target)) +
			hunter.AmmoDamageBonus +
			spell.BonusWeaponDamage()
		spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)
	}

	hunter.pet = hunter.NewHunterPet()

	hunter.AddStatDependency(stats.Strength, stats.AttackPower, 1)
	hunter.AddStatDependency(stats.Agility, stats.AttackPower, 1)
	hunter.AddStatDependency(stats.Agility, stats.RangedAttackPower, 1)
	hunter.AddStat(stats.AttackPower, -20)
	hunter.AddStat(stats.RangedAttackPower, -10)
	hunter.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritRatingPerCritChance/40)

	return hunter
}

func init() {
	const basecrit = -1.53 * core.CritRatingPerCritChance
	//const basespellcrit =
	const basehealth = 3568
	const basemana = 3383
	const baseap = core.CharacterLevel * 2

	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Health:    basehealth,
		stats.Strength:  61,
		stats.Agility:   153,
		stats.Stamina:   108,
		stats.Intellect: 80,
		stats.Spirit:    81,
		stats.Mana:      basemana,

		stats.AttackPower:       baseap,
		stats.RangedAttackPower: baseap,
		stats.MeleeCrit:         basecrit,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Health:    basehealth,
		stats.Strength:  65,
		stats.Agility:   148,
		stats.Stamina:   108,
		stats.Intellect: 77,
		stats.Spirit:    85,
		stats.Mana:      basemana,

		stats.AttackPower:       baseap,
		stats.RangedAttackPower: baseap,
		stats.MeleeCrit:         basecrit,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Health:    basehealth,
		stats.Strength:  69,
		stats.Agility:   147,
		stats.Stamina:   109,
		stats.Intellect: 76,
		stats.Spirit:    82,
		stats.Mana:      basemana,

		stats.AttackPower:       baseap,
		stats.RangedAttackPower: baseap,
		stats.MeleeCrit:         basecrit,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Health:    basehealth,
		stats.Strength:  60,
		stats.Agility:   155,
		stats.Stamina:   108,
		stats.Intellect: 77,
		stats.Spirit:    83,
		stats.Mana:      basemana,

		stats.AttackPower:       baseap,
		stats.RangedAttackPower: baseap,
		stats.MeleeCrit:         basecrit,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Health:    basehealth,
		stats.Strength:  67,
		stats.Agility:   148,
		stats.Stamina:   109,
		stats.Intellect: 74,
		stats.Spirit:    85,
		stats.Mana:      basemana,

		stats.AttackPower:       baseap,
		stats.RangedAttackPower: baseap,
		stats.MeleeCrit:         basecrit,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Health:    basehealth,
		stats.Strength:  69,
		stats.Agility:   147,
		stats.Stamina:   109,
		stats.Intellect: 73,
		stats.Spirit:    85,
		stats.Mana:      basemana,

		stats.AttackPower:       baseap,
		stats.RangedAttackPower: baseap,
		stats.MeleeCrit:         basecrit,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll, Class: proto.Class_ClassHunter}] = stats.Stats{
		stats.Health:    basehealth,
		stats.Strength:  65,
		stats.Agility:   153,
		stats.Stamina:   108,
		stats.Intellect: 73,
		stats.Spirit:    84,
		stats.Mana:      basemana,

		stats.AttackPower:       baseap,
		stats.RangedAttackPower: baseap,
		stats.MeleeCrit:         basecrit,
	}
}

// Agent is a generic way to access underlying hunter on any of the agents.
type HunterAgent interface {
	GetHunter() *Hunter
}
