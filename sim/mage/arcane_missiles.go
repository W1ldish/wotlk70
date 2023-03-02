package mage

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
)

func (mage *Mage) registerArcaneMissilesSpell() {
	spellCoeff := 1/3.5 + 0.03*float64(mage.Talents.ArcaneEmpowerment)

	mage.ArcaneMissiles = mage.RegisterSpell(core.SpellConfig{
		ActionID:     core.ActionID{SpellID: 38704},
		SpellSchool:  core.SpellSchoolArcane,
		ProcMask:     core.ProcMaskSpellDamage,
		Flags:        SpellFlagMage | core.SpellFlagChanneled,
		MissileSpeed: 20,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.31,
			Multiplier: 1 - .01*float64(mage.Talents.ArcaneFocus),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:         core.GCDDefault,
				ChannelTime: time.Second * 5,
			},
		},

		BonusHitRating:   float64(mage.Talents.ArcaneFocus) * core.SpellHitRatingPerHitChance,
		DamageMultiplier: 1 + .04*float64(mage.Talents.TormentTheWeak),
		DamageMultiplierAdditive: 1 +
			core.TernaryFloat64(mage.HasSetBonus(ItemSetTempestRegalia, 4), .05, 0),
		CritMultiplier:   mage.SpellCritMultiplier(1, mage.bonusCritDamage+core.TernaryFloat64(mage.HasMajorGlyph(proto.MageMajorGlyph_GlyphOfArcaneMissiles), .25, 0)),
		ThreatMultiplier: 1 - 0.2*float64(mage.Talents.ArcaneSubtlety),

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "ArcaneMissiles",
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					// TODO: This check is necessary to ensure the final tick occurs before
					// Arcane Blast stacks are dropped. To fix this, ticks need to reliably
					// occur before aura expirations.
					dot := mage.ArcaneMissiles.Dot(aura.Unit)
					if dot.TickCount < dot.NumberOfTicks {
						dot.TickOnce(sim)
					}
					mage.ArcaneBlastAura.Deactivate(sim)
				},
			},

			NumberOfTicks:       5,
			TickLength:          time.Second,
			AffectedByCastSpeed: true,

			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				baseDamage := 280 + spellCoeff*dot.Spell.SpellPower()
				result := dot.Spell.CalcDamage(sim, target, baseDamage, dot.Spell.OutcomeMagicHitAndCrit)
				dot.Spell.WaitTravelTime(sim, func(sim *core.Simulation) {
					dot.Spell.DealDamage(sim, result)
				})
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHit)
			if result.Landed() {
				spell.SpellMetrics[target.UnitIndex].Hits--
				spell.Dot(target).Apply(sim)
			}
			spell.DealOutcome(sim, result)
		},
	})
}
