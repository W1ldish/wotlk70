package paladin

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
)

func (paladin *Paladin) registerHolyWrathSpell() {
	results := make([]*core.SpellResult, len(paladin.Env.Encounter.Targets))
	bonusSpellPower := 0 +
		core.TernaryFloat64(paladin.Equip[proto.ItemSlot_ItemSlotRanged].ID == 28065, 120, 0)
	paladin.HolyWrath = paladin.RegisterSpell(core.SpellConfig{
		ActionID:        core.ActionID{SpellID: 27139},
		SpellSchool:     core.SpellSchoolHoly,
		ProcMask:        core.ProcMaskSpellDamage,
		Flags:           core.SpellFlagMeleeMetrics,
		BonusSpellPower: bonusSpellPower,
		ManaCost: core.ManaCostOptions{
			BaseCost:   0.20,
			Multiplier: 1 - 0.02*float64(paladin.Talents.Benediction),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: time.Second*30 - core.TernaryDuration(paladin.HasMajorGlyph(proto.PaladinMajorGlyph_GlyphOfHolyWrath), time.Second*15, 0),
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		CritMultiplier:   paladin.SpellCritMultiplier(),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			constBaseDamage := .07*spell.SpellPower() + .07*spell.MeleeAttackPower()

			for i, aoeTarget := range sim.Encounter.Targets {
				aoeUnit := &aoeTarget.Unit
				baseDamage := constBaseDamage + sim.Roll(779.2, 915.2)

				if aoeUnit.MobType == proto.MobType_MobTypeDemon || aoeUnit.MobType == proto.MobType_MobTypeUndead {
					results[i] = spell.CalcDamage(sim, aoeUnit, baseDamage, spell.OutcomeMagicHitAndCrit)
				} else {
					results[i] = spell.CalcDamage(sim, aoeUnit, baseDamage, spell.OutcomeAlwaysMiss)
				}
			}

			for i := range sim.Encounter.Targets {
				spell.DealDamage(sim, results[i])
			}
		},
	})
}
