package rogue

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
)

func RegisterRogue() {
	core.RegisterAgentFactory(
		proto.Player_Rogue{},
		proto.Spec_SpecRogue,
		func(character core.Character, options *proto.Player) core.Agent {
			return NewRogue(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Rogue)
			if !ok {
				panic("Invalid spec value for Rogue!")
			}
			player.Spec = playerSpec
		},
	)
}

const (
	SpellFlagBuilder  = core.SpellFlagAgentReserved2
	SpellFlagFinisher = core.SpellFlagAgentReserved3
	AssassinTree      = 0
	CombatTree        = 1
	SubtletyTree      = 2
)

var TalentTreeSizes = [3]int{27, 28, 28}

const RogueBleedTag = "RogueBleed"

type Rogue struct {
	core.Character

	Talents  *proto.RogueTalents
	Options  *proto.Rogue_Options
	Rotation *proto.Rogue_Rotation

	priorityItems      []roguePriorityItem
	rotationItems      []rogueRotationItem
	assassinationPrios []assassinationPrio
	subtletyPrios      []subtletyPrio
	bleedCategory      *core.ExclusiveCategory

	sliceAndDiceDurations [6]time.Duration
	exposeArmorDurations  [6]time.Duration

	allMCDsDisabled bool

	maxEnergy float64

	BuilderPoints    int32
	Builder          *core.Spell
	Backstab         *core.Spell
	BladeFlurry      *core.Spell
	DeadlyPoison     *core.Spell
	FanOfKnives      *core.Spell
	Feint            *core.Spell
	Garrote          *core.Spell
	Ambush           *core.Spell
	Hemorrhage       *core.Spell
	GhostlyStrike    *core.Spell
	HungerForBlood   *core.Spell
	InstantPoison    [3]*core.Spell
	WoundPoison      [3]*core.Spell
	Mutilate         *core.Spell
	Shiv             *core.Spell
	SinisterStrike   *core.Spell
	TricksOfTheTrade *core.Spell
	Shadowstep       *core.Spell
	Preparation      *core.Spell
	Premeditation    *core.Spell
	ShadowDance      *core.Spell
	ColdBlood        *core.Spell
	MasterOfSubtlety *core.Spell
	Overkill         *core.Spell

	Envenom      *core.Spell
	Eviscerate   *core.Spell
	ExposeArmor  *core.Spell
	Rupture      *core.Spell
	SliceAndDice *core.Spell

	lastDeadlyPoisonProcMask    core.ProcMask
	deadlyPoisonProcChanceBonus float64
	instantPoisonPPMM           core.PPMManager
	woundPoisonPPMM             core.PPMManager

	AdrenalineRushAura   *core.Aura
	BladeFlurryAura      *core.Aura
	EnvenomAura          *core.Aura
	ExposeArmorAuras     core.AuraArray
	HungerForBloodAura   *core.Aura
	KillingSpreeAura     *core.Aura
	OverkillAura         *core.Aura
	SliceAndDiceAura     *core.Aura
	TricksOfTheTradeAura *core.Aura
	MasterOfSubtletyAura *core.Aura
	ShadowstepAura       *core.Aura
	ShadowDanceAura      *core.Aura
	DirtyDeedsAura       *core.Aura
	HonorAmongThieves    *core.Aura

	masterPoisonerDebuffAuras []*core.Aura
	savageCombatDebuffAuras   []*core.Aura
	woundPoisonDebuffAuras    []*core.Aura

	QuickRecoveryMetrics *core.ResourceMetrics

	costModifier               func(float64) float64
	finishingMoveEffectApplier func(sim *core.Simulation, numPoints int32)
}

func (rogue *Rogue) GetCharacter() *core.Character {
	return &rogue.Character
}

func (rogue *Rogue) GetRogue() *Rogue {
	return rogue
}

func (rogue *Rogue) AddRaidBuffs(raidBuffs *proto.RaidBuffs)    {}
func (rogue *Rogue) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {}

func (rogue *Rogue) finisherFlags() core.SpellFlag {
	flags := SpellFlagFinisher
	if rogue.Talents.SurpriseAttacks {
		flags |= core.SpellFlagCannotBeDodged
	}
	return flags
}

func (rogue *Rogue) ApplyFinisher(sim *core.Simulation, spell *core.Spell) {
	numPoints := rogue.ComboPoints()
	rogue.SpendComboPoints(sim, spell.ComboPointMetrics())
	rogue.finishingMoveEffectApplier(sim, numPoints)
}

func (rogue *Rogue) HasMajorGlyph(glyph proto.RogueMajorGlyph) bool {
	return rogue.HasGlyph(int32(glyph))
}

func (rogue *Rogue) HasMinorGlyph(glyph proto.RogueMinorGlyph) bool {
	return rogue.HasGlyph(int32(glyph))
}

func (rogue *Rogue) Initialize() {
	// Update auto crit multipliers now that we have the targets.
	rogue.AutoAttacks.MHConfig.CritMultiplier = rogue.MeleeCritMultiplier(false)
	rogue.AutoAttacks.OHConfig.CritMultiplier = rogue.MeleeCritMultiplier(false)

	if rogue.Talents.QuickRecovery > 0 {
		rogue.QuickRecoveryMetrics = rogue.NewEnergyMetrics(core.ActionID{SpellID: 31245})
	}

	rogue.costModifier = rogue.makeCostModifier()

	rogue.registerBackstabSpell()
	rogue.registerDeadlyPoisonSpell()
	rogue.registerPoisonAuras()
	rogue.registerEviscerate()
	rogue.registerExposeArmorSpell()
	rogue.registerFanOfKnives()
	rogue.registerFeintSpell()
	rogue.registerGarrote()
	rogue.registerHemorrhageSpell()
	rogue.registerInstantPoisonSpell()
	rogue.registerWoundPoisonSpell()
	rogue.registerMutilateSpell()
	rogue.registerRupture()
	rogue.registerShivSpell()
	rogue.registerSinisterStrikeSpell()
	rogue.registerSliceAndDice()
	rogue.registerThistleTeaCD()
	rogue.registerTricksOfTheTradeSpell()
	rogue.registerAmbushSpell()
	rogue.registerEnvenom()

	rogue.finishingMoveEffectApplier = rogue.makeFinishingMoveEffectApplier()
	rogue.DelayDPSCooldownsForArmorDebuffs(time.Second * 14)
}

func (rogue *Rogue) getExpectedEnergyPerSecond() float64 {
	const finishersPerSecond = 1.0 / 6
	const averageComboPointsSpendOnFinisher = 4.0
	bonusEnergyPerSecond := float64(rogue.Talents.CombatPotency) * 3 * 0.2 * 1.0 / (rogue.AutoAttacks.OH.SwingSpeed / 1.4)
	bonusEnergyPerSecond += float64(rogue.Talents.FocusedAttacks)
	bonusEnergyPerSecond += float64(rogue.Talents.RelentlessStrikes) * 0.04 * 25 * finishersPerSecond * averageComboPointsSpendOnFinisher
	return (core.EnergyPerTick*rogue.EnergyTickMultiplier)/core.EnergyTickDuration.Seconds() + bonusEnergyPerSecond
}

func (rogue *Rogue) ApplyEnergyTickMultiplier(multiplier float64) {
	rogue.EnergyTickMultiplier += multiplier
}

func (rogue *Rogue) getExpectedComboPointPerSecond() float64 {
	const criticalPerSecond = 1
	honorAmongThievesChance := []float64{0, 0.33, 0.66, 1.0}[rogue.Talents.HonorAmongThieves]
	return criticalPerSecond * honorAmongThievesChance
}

func (rogue *Rogue) Reset(sim *core.Simulation) {
	for _, mcd := range rogue.GetMajorCooldowns() {
		mcd.Disable()
	}
	rogue.allMCDsDisabled = true
	rogue.lastDeadlyPoisonProcMask = core.ProcMaskEmpty
	// Vanish triggered effects (Overkill and Master of Subtlety) prepull activation
	if rogue.Rotation.OpenWithGarrote || rogue.Options.StartingOverkillDuration > 0 {
		length := rogue.Options.StartingOverkillDuration
		if rogue.OverkillAura != nil {
			if rogue.Rotation.OpenWithGarrote {
				length = 20
			}
			rogue.OverkillAura.Activate(sim)
			rogue.OverkillAura.UpdateExpires(sim.CurrentTime + time.Second*time.Duration(length))
		}
		if rogue.MasterOfSubtletyAura != nil {
			if rogue.Rotation.OpenWithGarrote {
				length = 6
			}
			rogue.MasterOfSubtletyAura.Activate(sim)
			rogue.MasterOfSubtletyAura.UpdateExpires(sim.CurrentTime + time.Second*time.Duration(length))
		}
	}
	rogue.setPriorityItems(sim)
}

func (rogue *Rogue) MeleeCritMultiplier(applyLethality bool) float64 {
	primaryModifier := rogue.preyOnTheWeakMultiplier(rogue.CurrentTarget)
	var secondaryModifier float64
	if applyLethality {
		secondaryModifier += 0.06 * float64(rogue.Talents.Lethality)
	}
	return rogue.Character.MeleeCritMultiplier(primaryModifier, secondaryModifier)
}
func (rogue *Rogue) SpellCritMultiplier() float64 {
	primaryModifier := rogue.preyOnTheWeakMultiplier(rogue.CurrentTarget)
	return rogue.Character.SpellCritMultiplier(primaryModifier, 0)
}

func NewRogue(character core.Character, options *proto.Player) *Rogue {
	rogueOptions := options.GetRogue()

	rogue := &Rogue{
		Character: character,
		Talents:   &proto.RogueTalents{},
		Options:   rogueOptions.Options,
		Rotation:  rogueOptions.Rotation,
	}
	core.FillTalentsProto(rogue.Talents.ProtoReflect(), options.TalentsString, TalentTreeSizes)

	// Passive rogue threat reduction: https://wotlk.wowhead.com/spell=21184/rogue-passive-dnd
	rogue.PseudoStats.ThreatMultiplier *= 0.71
	rogue.PseudoStats.CanParry = true
	maxEnergy := 100.0
	if rogue.Talents.Vigor {
		maxEnergy += 10
	}
	if rogue.HasMajorGlyph(proto.RogueMajorGlyph_GlyphOfVigor) {
		maxEnergy += 10
	}
	if rogue.HasSetBonus(ItemSetGladiatorsVestments, 4) {
		maxEnergy += 10
	}
	rogue.maxEnergy = maxEnergy
	rogue.EnableEnergyBar(maxEnergy, rogue.OnEnergyGain)
	rogue.ApplyEnergyTickMultiplier([]float64{0, 0.08, 0.16, 0.25}[rogue.Talents.Vitality])

	rogue.EnableAutoAttacks(rogue, core.AutoAttackOptions{
		MainHand:       rogue.WeaponFromMainHand(0), // Set crit multiplier later when we have targets.
		OffHand:        rogue.WeaponFromOffHand(0),  // Set crit multiplier later when we have targets.
		AutoSwingMelee: true,
	})
	rogue.applyPoisons()

	rogue.AddStatDependency(stats.Strength, stats.AttackPower, 1)
	rogue.AddStatDependency(stats.Agility, stats.AttackPower, 1)
	rogue.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritRatingPerCritChance/40)

	return rogue
}

func (rogue *Rogue) ApplyCutToTheChase(sim *core.Simulation) {
	if rogue.Talents.CutToTheChase > 0 && rogue.SliceAndDiceAura.IsActive() {
		procChance := float64(rogue.Talents.CutToTheChase) * 0.2
		if sim.Proc(procChance, "Cut to the Chase") {
			rogue.SliceAndDiceAura.Duration = rogue.sliceAndDiceDurations[5]
			rogue.SliceAndDiceAura.Activate(sim)
		}
	}
}

func (rogue *Rogue) CanMutilate() bool {
	return rogue.Talents.Mutilate &&
		rogue.HasMHWeapon() && rogue.HasOHWeapon() &&
		rogue.GetMHWeapon().WeaponType == proto.WeaponType_WeaponTypeDagger &&
		rogue.GetOHWeapon().WeaponType == proto.WeaponType_WeaponTypeDagger
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  92,
		stats.Agility:   160,
		stats.Stamina:   89,
		stats.Intellect: 42,
		stats.Spirit:    56,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  100,
		stats.Agility:   154,
		stats.Stamina:   90,
		stats.Intellect: 38,
		stats.Spirit:    57,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceGnome, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  90,
		stats.Agility:   160,
		stats.Stamina:   89,
		stats.Intellect: 42,
		stats.Spirit:    58,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  95,
		stats.Agility:   158,
		stats.Stamina:   89,
		stats.Intellect: 39,
		stats.Spirit:    58,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  91,
		stats.Agility:   162,
		stats.Stamina:   89,
		stats.Intellect: 39,
		stats.Spirit:    58,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  98,
		stats.Agility:   155,
		stats.Stamina:   90,
		stats.Intellect: 36,
		stats.Spirit:    60,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  96,
		stats.Agility:   160,
		stats.Stamina:   89,
		stats.Intellect: 35,
		stats.Spirit:    59,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Health:    3704,
		stats.Strength:  94,
		stats.Agility:   156,
		stats.Stamina:   89,
		stats.Intellect: 37,
		stats.Spirit:    63,

		stats.AttackPower: 140,
		stats.MeleeCrit:   0 * core.CritRatingPerCritChance,
		stats.SpellCrit:   0 * core.CritRatingPerCritChance,
	}
}

// Agent is a generic way to access underlying rogue on any of the agents.
type RogueAgent interface {
	GetRogue() *Rogue
}
