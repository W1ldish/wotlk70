package mage

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
)

func (mage *Mage) registerScorchSpell() {
	hasImpScorch := mage.Talents.ImprovedScorch > 0
	procChance := float64(mage.Talents.ImprovedScorch) / 3.0

	if hasImpScorch {
		mage.ScorchAura = core.ImprovedScorchAura(mage.CurrentTarget)
		mage.CritDebuffCategory = mage.ScorchAura.ExclusiveEffects[0].Category
	}

	mage.Scorch = mage.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 42859},
		SpellSchool: core.SpellSchoolFire,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagMage | HotStreakSpells,

		ManaCost: core.ManaCostOptions{
			BaseCost: 0.08,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
		},

		BonusCritRating: 0 +
			float64(mage.Talents.Incineration+mage.Talents.CriticalMass)*2*core.CritRatingPerCritChance +
			float64(mage.Talents.ImprovedScorch)*1*core.CritRatingPerCritChance,
		DamageMultiplierAdditive: 1 +
			.02*float64(mage.Talents.SpellImpact) +
			.02*float64(mage.Talents.FirePower) +
			core.TernaryFloat64(mage.HasMajorGlyph(proto.MageMajorGlyph_GlyphOfScorch), 0.2, 0),
		CritMultiplier:   mage.SpellCritMultiplier(1, mage.bonusCritDamage),
		ThreatMultiplier: 1 - 0.1*float64(mage.Talents.BurningSoul),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(382, 451) + (1.5/3.5)*spell.SpellPower()
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			if hasImpScorch && result.Landed() && sim.Proc(procChance, "Improved Scorch") {
				mage.ScorchAura.Activate(sim)
			}
			spell.DealDamage(sim, result)
		},
	})
}
