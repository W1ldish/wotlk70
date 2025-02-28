package shaman

import (
	"math"
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
)

// Variables that control the Fire Elemental.
const (
	// 7.5 CPM
	maxFireBlastCasts = 15
	maxFireNovaCasts  = 15
)

type FireElemental struct {
	core.Pet

	FireBlast *core.Spell
	FireNova  *core.Spell

	FireShieldAura *core.Aura

	shamanOwner *Shaman
}

func (shaman *Shaman) NewFireElemental() *FireElemental {
	fireElemental := &FireElemental{
		Pet:         core.NewPet("Greater Fire Elemental", &shaman.Character, fireElementalPetBaseStats, shaman.fireElementalStatInheritance(), nil, false, true),
		shamanOwner: shaman,
	}
	fireElemental.EnableManaBar()
	fireElemental.EnableAutoAttacks(fireElemental, core.AutoAttackOptions{
		MainHand: core.Weapon{
			BaseDamageMin:  1,  // Estimated from base AP
			BaseDamageMax:  24, // Estimated from base AP
			SwingSpeed:     2,
			SwingDuration:  time.Second * 2,
			CritMultiplier: 2, // Pretty sure this is right.
			SpellSchool:    core.SpellSchoolFire,
		},
		AutoSwingMelee: true,
	})
	fireElemental.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritRatingPerCritChance/212)
	fireElemental.OnPetEnable = fireElemental.enable
	fireElemental.OnPetDisable = fireElemental.disable

	shaman.AddPet(fireElemental)

	return fireElemental
}

func (fireElemental *FireElemental) enable(sim *core.Simulation) {
	fireElemental.FireShieldAura.Activate(sim)
}

func (fireElemental *FireElemental) disable(sim *core.Simulation) {
	fireElemental.FireShieldAura.Deactivate(sim)
}

func (fireElemental *FireElemental) GetPet() *core.Pet {
	return &fireElemental.Pet
}

func (fireElemental *FireElemental) Initialize() {
	fireElemental.registerFireBlast()
	fireElemental.registerFireNova()
	fireElemental.registerFireShieldAura()
}

func (fireElemental *FireElemental) Reset(sim *core.Simulation) {
}

func (fireElemental *FireElemental) OnGCDReady(sim *core.Simulation) {
	/*
		TODO this is a little dirty, can probably clean this up, the rotation might go through some more overhauls,
		the random AI is hard to emulate.
	*/
	target := fireElemental.CurrentTarget
	fireBlastCasts := fireElemental.FireBlast.SpellMetrics[0].Casts
	fireNovaCasts := fireElemental.FireNova.SpellMetrics[0].Casts

	if fireBlastCasts == maxFireBlastCasts && fireNovaCasts == maxFireNovaCasts {
		fireElemental.DoNothing()
		return
	}

	if fireElemental.FireNova.BaseCost > fireElemental.CurrentMana() {
		fireElemental.WaitForMana(sim, fireElemental.FireNova.BaseCost)
		return
	}

	random := sim.RandomFloat("Fire Elemental Pet Spell")

	//Melee the other 30%
	if random >= .65 {
		if !fireElemental.TryCast(sim, target, fireElemental.FireNova, maxFireNovaCasts) {
			fireElemental.TryCast(sim, target, fireElemental.FireBlast, maxFireBlastCasts)
		}
	} else if random >= .35 {
		if !fireElemental.TryCast(sim, target, fireElemental.FireBlast, maxFireBlastCasts) {
			fireElemental.TryCast(sim, target, fireElemental.FireNova, maxFireNovaCasts)
		}
	}

	if !fireElemental.GCD.IsReady(sim) {
		return
	}

	fireElemental.WaitUntil(sim, sim.CurrentTime+time.Second)
}

func (fireElemental *FireElemental) TryCast(sim *core.Simulation, target *core.Unit, spell *core.Spell, maxCastCount int32) bool {
	if maxCastCount == spell.SpellMetrics[0].Casts {
		return false
	}

	if !spell.IsReady(sim) {
		return false
	}

	if !spell.Cast(sim, target) {
		return false
	}
	// all spell casts reset the elemental's swing timer
	fireElemental.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime+spell.CurCast.CastTime, false)
	return true
}

var fireElementalPetBaseStats = stats.Stats{
	stats.Mana:        1789,
	stats.Health:      994,
	stats.Intellect:   147,
	stats.Stamina:     327,
	stats.SpellPower:  995,  //Estimated
	stats.AttackPower: 1369, //Estimated

	// TODO : Log digging and my own samples this seems to be around the 5% mark.
	stats.MeleeCrit: (5 + 1.8) * core.CritRatingPerCritChance,
	stats.SpellCrit: 2.61 * core.CritRatingPerCritChance,
}

func (shaman *Shaman) fireElementalStatInheritance() core.PetStatInheritance {
	return func(ownerStats stats.Stats) stats.Stats {
		ownerSpellHitChance := math.Floor(ownerStats[stats.SpellHit] / core.SpellHitRatingPerHitChance)
		spellHitRatingFromOwner := ownerSpellHitChance * core.SpellHitRatingPerHitChance

		ownerHitChance := ownerStats[stats.MeleeHit] / core.MeleeHitRatingPerHitChance
		hitRatingFromOwner := math.Floor(ownerHitChance) * core.MeleeHitRatingPerHitChance
		return stats.Stats{
			stats.Stamina:     ownerStats[stats.Stamina] * 0.75,
			stats.Intellect:   ownerStats[stats.Intellect] * 0.30,
			stats.SpellPower:  ownerStats[stats.SpellPower] * 0.5218,
			stats.AttackPower: ownerStats[stats.SpellPower] * 4.45,

			// TODO tested useing pre-patch lvl 70 stats need to confirm in WOTLK at 80.
			stats.MeleeHit: hitRatingFromOwner,
			stats.SpellHit: spellHitRatingFromOwner,

			/*
				TODO working on figuring this out, getting close need more trials. will need to remove specific buffs,
				ie does not gain the benefit from draenei buff.
			*/
			stats.Expertise: math.Floor(spellHitRatingFromOwner * 0.79),
		}
	}
}
