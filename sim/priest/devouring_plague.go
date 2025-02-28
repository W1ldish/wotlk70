package priest

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
)

func (priest *Priest) registerDevouringPlagueSpell() {
	actionID := core.ActionID{SpellID: 48300}
	priest.DpInitMultiplier = 8 * 0.1 * float64(priest.Talents.ImprovedDevouringPlague)

	priest.DevouringPlague = priest.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagDisease,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.25,
			Multiplier: 1 - []float64{0, .04, .07, .10}[priest.Talents.MentalAgility],
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		BonusHitRating: float64(priest.Talents.ShadowFocus) * 1 * core.SpellHitRatingPerHitChance,
		BonusCritRating: 0 +
			3*float64(priest.Talents.MindMelt)*core.CritRatingPerCritChance +
			core.TernaryFloat64(priest.HasSetBonus(ItemSetCrimsonAcolyte, 2), 5, 0)*core.CritRatingPerCritChance,
		DamageMultiplier: 1 +
			0.02*float64(priest.Talents.Darkness) +
			0.01*float64(priest.Talents.TwinDisciplines) +
			0.05*float64(priest.Talents.ImprovedDevouringPlague) +
			core.TernaryFloat64(priest.HasSetBonus(ItemSetConquerorSanct, 2), 0.15, 0),
		CritMultiplier:   priest.DefaultSpellCritMultiplier(),
		ThreatMultiplier: 1 - 0.05*float64(priest.Talents.ShadowAffinity),

		Dot: core.DotConfig{
			Spell: priest.RegisterSpell(core.SpellConfig{
				ActionID:    actionID,
				SpellSchool: core.SpellSchoolShadow,
				ProcMask:    core.ProcMaskSpellDamage,
				Flags:       core.SpellFlagDisease,

				BonusHitRating: float64(priest.Talents.ShadowFocus) * 1 * core.SpellHitRatingPerHitChance,
				BonusCritRating: 0 +
					3*float64(priest.Talents.MindMelt)*core.CritRatingPerCritChance +
					core.TernaryFloat64(priest.HasSetBonus(ItemSetCrimsonAcolyte, 2), 5, 0)*core.CritRatingPerCritChance,
				DamageMultiplier: 1 +
					float64(priest.Talents.Darkness)*0.02 +
					float64(priest.Talents.TwinDisciplines)*0.01 +
					float64(priest.Talents.ImprovedDevouringPlague)*0.05 +
					core.TernaryFloat64(priest.HasSetBonus(ItemSetConquerorSanct, 2), 0.15, 0),
				CritMultiplier:   priest.SpellCritMultiplier(1, 1),
				ThreatMultiplier: 1 - 0.05*float64(priest.Talents.ShadowAffinity),
			}),
			Aura: core.Aura{
				Label: "DevouringPlague",
			},

			NumberOfTicks:       8,
			TickLength:          time.Second * 3,
			AffectedByCastSpeed: priest.Talents.Shadowform,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, _ bool) {
				dot.SnapshotBaseDamage = 1376/8 + 0.1849*dot.Spell.SpellPower()
				dot.SnapshotCritChance = dot.Spell.SpellCritChance(target)
				dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(dot.Spell.Unit.AttackTables[target.UnitIndex])
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				if priest.Talents.Shadowform {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeSnapshotCrit)
				} else {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			var result *core.SpellResult
			if priest.DpInitMultiplier == 0 {
				result = spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
			} else {
				baseDamage := (1376/8 + 0.1849*spell.SpellPower()) * priest.DpInitMultiplier
				result = spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			}

			if result.Landed() {
				priest.AddShadowWeavingStack(sim)
				spell.Dot(target).Apply(sim)
			}
		},
		ExpectedDamage: func(sim *core.Simulation, target *core.Unit, spell *core.Spell, useSnapshot bool) *core.SpellResult {
			if useSnapshot {
				dot := spell.Dot(target)
				if priest.Talents.Shadowform {
					return dot.CalcSnapshotDamage(sim, target, dot.OutcomeExpectedMagicSnapshotCrit)
				} else {
					return dot.CalcSnapshotDamage(sim, target, spell.OutcomeExpectedMagicAlwaysHit)
				}
			} else {
				baseDamage := 1376/8 + 0.1849*spell.SpellPower()
				if priest.Talents.Shadowform {
					return spell.CalcPeriodicDamage(sim, target, baseDamage, spell.OutcomeExpectedMagicCrit)
				} else {
					return spell.CalcPeriodicDamage(sim, target, baseDamage, spell.OutcomeExpectedMagicAlwaysHit)
				}
			}
		},
	})
}
