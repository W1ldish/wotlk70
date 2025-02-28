package warlock

import (
	"strconv"
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
)

func (warlock *Warlock) registerImmolateSpell() {
	actionID := core.ActionID{SpellID: 47811}

	warlock.Immolate = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolFire,
		ProcMask:    core.ProcMaskSpellDamage,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.17,
			Multiplier: 1 - []float64{0, .04, .07, .10}[warlock.Talents.Cataclysm],
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * (2000 - 100*time.Duration(warlock.Talents.Bane)),
			},
		},

		BonusCritRating: 0 +
			warlock.masterDemonologistFireCrit +
			core.TernaryFloat64(warlock.Talents.Devastation, 5*core.CritRatingPerCritChance, 0),
		DamageMultiplierAdditive: 1 +
			warlock.GrandFirestoneBonus() +
			0.03*float64(warlock.Talents.Emberstorm) +
			0.1*float64(warlock.Talents.ImprovedImmolate) +
			core.TernaryFloat64(warlock.HasSetBonus(ItemSetDeathbringerGarb, 2), 0.1, 0) +
			core.TernaryFloat64(warlock.HasSetBonus(ItemSetGuldansRegalia, 4), 0.1, 0),
		CritMultiplier:   warlock.SpellCritMultiplier(1, float64(warlock.Talents.Ruin)/5),
		ThreatMultiplier: 1 - 0.1*float64(warlock.Talents.DestructiveReach),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := 327 + 0.2*spell.SpellPower()
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			if result.Landed() {
				warlock.ImmolateDot.Apply(sim)
			}
			spell.DealDamage(sim, result)
		},
	})

	fireAndBrimstoneBonus := 0.02 * float64(warlock.Talents.FireAndBrimstone)

	warlock.ImmolateDot = core.NewDot(core.Dot{
		Spell: warlock.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolFire,
			ProcMask:    core.ProcMaskSpellDamage,

			BonusCritRating: warlock.Immolate.BonusCritRating,
			DamageMultiplierAdditive: 1 +
				warlock.GrandSpellstoneBonus() +
				0.03*float64(warlock.Talents.Emberstorm) +
				0.03*float64(warlock.Talents.Aftermath) +
				0.1*float64(warlock.Talents.ImprovedImmolate) +
				core.TernaryFloat64(warlock.HasMajorGlyph(proto.WarlockMajorGlyph_GlyphOfImmolate), 0.1, 0) +
				core.TernaryFloat64(warlock.HasSetBonus(ItemSetDeathbringerGarb, 2), 0.1, 0) +
				core.TernaryFloat64(warlock.HasSetBonus(ItemSetGuldansRegalia, 4), 0.1, 0),
			CritMultiplier:   warlock.SpellCritMultiplier(1, float64(warlock.Talents.Ruin)/5),
			ThreatMultiplier: warlock.Immolate.ThreatMultiplier,
		}),
		Aura: warlock.CurrentTarget.RegisterAura(core.Aura{
			Label:    "Immolate-" + strconv.Itoa(int(warlock.Index)),
			ActionID: actionID,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				warlock.ChaosBolt.DamageMultiplierAdditive += fireAndBrimstoneBonus
				warlock.Incinerate.DamageMultiplierAdditive += fireAndBrimstoneBonus
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				warlock.ChaosBolt.DamageMultiplierAdditive -= fireAndBrimstoneBonus
				warlock.Incinerate.DamageMultiplierAdditive -= fireAndBrimstoneBonus
			},
		}),
		NumberOfTicks: 5 + warlock.Talents.MoltenCore,
		TickLength:    time.Second * 3,

		OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
			dot.SnapshotBaseDamage = 615/5 + 0.2*dot.Spell.SpellPower()
			attackTable := dot.Spell.Unit.AttackTables[target.UnitIndex]
			dot.SnapshotCritChance = dot.Spell.SpellCritChance(target)
			dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(attackTable)
		},
		OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
			dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeSnapshotCrit)
		},
	})
}
