package warrior

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
)

func (warrior *Warrior) registerMortalStrikeSpell(cdTimer *core.Timer) {
	warrior.MortalStrike = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 47486},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagIncludeTargetBonusDamage,

		RageCost: core.RageCostOptions{
			Cost:   30,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: time.Second*6 - time.Millisecond*200*time.Duration(warrior.Talents.ImprovedMortalStrike),
			},
		},

		BonusCritRating: core.TernaryFloat64(warrior.HasSetBonus(ItemSetSiegebreakerBattlegear, 4), 10, 0) * core.CritRatingPerCritChance,
		DamageMultiplier: 1 +
			[]float64{0, 0.03, 0.06, 0.1}[warrior.Talents.ImprovedMortalStrike] +
			core.TernaryFloat64(warrior.HasMajorGlyph(proto.WarriorMajorGlyph_GlyphOfMortalStrike), 0.1, 0) +
			core.TernaryFloat64(warrior.HasSetBonus(ItemSetOnslaughtBattlegear, 4), 0.05, 0),
		CritMultiplier:   warrior.critMultiplier(mh),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := 380 +
				spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower()) +
				spell.BonusWeaponDamage()

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)

			if !result.Landed() {
				spell.IssueRefund(sim)
			}
		},
	})
}

func (warrior *Warrior) CanMortalStrike(sim *core.Simulation) bool {
	if warrior.Talents.MortalStrike {
		return warrior.CurrentRage() >= warrior.MortalStrike.DefaultCast.Cost && warrior.MortalStrike.IsReady(sim)
	} else {
		return false
	}
}
