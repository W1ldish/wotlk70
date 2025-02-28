package hunter

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
)

func (hunter *Hunter) registerArcaneShotSpell(timer *core.Timer) {
	hasGlyph := hunter.HasMajorGlyph(proto.HunterMajorGlyph_GlyphOfArcaneShot)
	var manaMetrics *core.ResourceMetrics
	if hasGlyph {
		manaMetrics = hunter.NewManaMetrics(core.ActionID{ItemID: 42898})
	}

	hunter.ArcaneShot = hunter.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 49045},
		SpellSchool: core.SpellSchoolArcane,
		ProcMask:    core.ProcMaskRangedSpecial,
		Flags:       core.SpellFlagMeleeMetrics,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.05,
			Multiplier: 1 - 0.03*float64(hunter.Talents.Efficiency),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    timer,
				Duration: time.Second*6 - time.Millisecond*200*time.Duration(hunter.Talents.ImprovedArcaneShot),
			},
		},

		BonusCritRating: 0 +
			2*core.CritRatingPerCritChance*float64(hunter.Talents.SurvivalInstincts),
		DamageMultiplierAdditive: 1 +
			.03*float64(hunter.Talents.FerociousInspiration) +
			.05*float64(hunter.Talents.ImprovedArcaneShot),
		DamageMultiplier: 1 *
			hunter.markedForDeathMultiplier(),
		CritMultiplier:   hunter.critMultiplier(true, true),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := 492 + 0.15*spell.RangedAttackPower(target)
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)
			if hasGlyph && result.Landed() && (hunter.SerpentSting.Dot(target).IsActive() || hunter.ScorpidStingAura.IsActive()) {
				hunter.AddMana(sim, 0.2*hunter.ArcaneShot.DefaultCast.Cost, manaMetrics, false)
			}
			spell.DealDamage(sim, result)
		},
	})
}
