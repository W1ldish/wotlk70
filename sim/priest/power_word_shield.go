package priest

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
)

func (priest *Priest) registerPowerWordShieldSpell() {
	actionID := core.ActionID{SpellID: 48066}
	coeff := 0.8057 + 0.08*float64(priest.Talents.BorrowedTime)

	wsDuration := time.Second*15 -
		core.TernaryDuration(priest.HasSetBonus(ItemSetGladiatorsInvestiture, 4), time.Second*2, 0) -
		core.TernaryDuration(priest.HasSetBonus(ItemSetGladiatorsRaiment, 4), time.Second*2, 0)

	cd := core.Cooldown{}
	if !priest.Talents.SoulWarding {
		cd = core.Cooldown{
			Timer:    priest.NewTimer(),
			Duration: time.Second * 4,
		}
	}

	var glyphHeal *core.Spell

	priest.PowerWordShield = priest.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolHoly,
		ProcMask:    core.ProcMaskSpellHealing,
		Flags:       core.SpellFlagHelpful,

		ManaCost: core.ManaCostOptions{
			BaseCost: 0.23,
			Multiplier: 1 -
				[]float64{0, .04, .07, .10}[priest.Talents.MentalAgility] -
				core.TernaryFloat64(priest.Talents.SoulWarding, .15, 0),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: cd,
		},

		DamageMultiplier: 1 *
			(1 + .05*float64(priest.Talents.ImprovedPowerWordShield)) *
			(1 +
				.01*float64(priest.Talents.TwinDisciplines) +
				.02*float64(priest.Talents.FocusedPower) +
				.02*float64(priest.Talents.SpiritualHealing)) *
			core.TernaryFloat64(priest.HasSetBonus(ItemSetCrimsonAcolytesRaiment, 4), 1.05, 1),
		ThreatMultiplier: 1 - []float64{0, .07, .14, .20}[priest.Talents.SilentResolve],

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			weakenedSoul := priest.WeakenedSouls[target.UnitIndex]
			if weakenedSoul.IsActive() {
				panic("Cannot cast PWS on target with Weakened Soul!")
			}

			shieldAmount := 2230.0 + coeff*spell.HealingPower(target)
			shield := priest.PWSShields[target.UnitIndex]
			shield.Apply(sim, shieldAmount)

			weakenedSoul.Duration = wsDuration
			weakenedSoul.Activate(sim)

			if glyphHeal != nil {
				glyphHeal.Cast(sim, target)
			}
		},
	})

	priest.PWSShields = core.NewAllyShieldArray(
		&priest.Unit,
		core.Shield{
			Spell: priest.PowerWordShield,
		},
		core.Aura{
			Label:    "Power Word Shield",
			ActionID: priest.PowerWordShield.ActionID,
			Duration: time.Second * 30,
		})

	priest.WeakenedSouls = make([]*core.Aura, len(priest.Env.AllUnits))
	for _, unit := range priest.Env.AllUnits {
		if !priest.IsOpponent(unit) {
			priest.WeakenedSouls[unit.UnitIndex] = priest.makeWeakenedSoul(unit)
		}
	}

	if priest.HasMajorGlyph(proto.PriestMajorGlyph_GlyphOfPowerWordShield) {
		glyphHeal = priest.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{ItemID: 42408},
			SpellSchool: core.SpellSchoolHoly,
			ProcMask:    core.ProcMaskSpellHealing,
			Flags:       core.SpellFlagHelpful,

			// Talent effects are combined differently in this spell compared to PWS, for some reason.
			DamageMultiplier: 0.2 *
				(1 + .01*float64(priest.Talents.BlessedResilience)) *
				(1 + .02*float64(priest.Talents.FocusedPower)) *
				(1 +
					.05*float64(priest.Talents.ImprovedPowerWordShield) +
					.01*float64(priest.Talents.TwinDisciplines)) *
				core.TernaryFloat64(priest.HasSetBonus(ItemSetCrimsonAcolytesRaiment, 4), 1.05, 1),
			ThreatMultiplier: 1 - []float64{0, .07, .14, .20}[priest.Talents.SilentResolve],

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				baseHealing := 2230 + coeff*spell.HealingPower(target)
				spell.CalcAndDealHealing(sim, target, baseHealing, spell.OutcomeAlwaysHit)
			},
		})
	}
}

func (priest *Priest) makePWSShield(target *core.Unit) *core.Shield {
	return core.NewShield(core.Shield{
		Spell: priest.PowerWordShield,
		Aura: target.GetOrRegisterAura(core.Aura{
			Label:    "Power Word Shield",
			ActionID: priest.PowerWordShield.ActionID,
			Duration: time.Second * 30,
		}),
	})
}

func (priest *Priest) makeWeakenedSoul(target *core.Unit) *core.Aura {
	return target.GetOrRegisterAura(core.Aura{
		Label:    "Weakened Soul",
		ActionID: core.ActionID{SpellID: 6788},
		Duration: time.Second * 15,
	})
}

func (priest *Priest) CanCastPWS(sim *core.Simulation, target *core.Unit) bool {
	return priest.PowerWordShield.IsReady(sim) &&
		priest.WeakenedSouls[target.UnitIndex] != nil &&
		!priest.WeakenedSouls[target.UnitIndex].IsActive()
}
