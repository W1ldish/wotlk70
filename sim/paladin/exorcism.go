package paladin

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
)

func (paladin *Paladin) registerExorcismSpell() {
	bonusSpellPower := 0 +
		core.TernaryFloat64(paladin.Equip[proto.ItemSlot_ItemSlotRanged].ID == 28065, 120, 0)

	paladin.Exorcism = paladin.RegisterSpell(core.SpellConfig{
		ActionID:        core.ActionID{SpellID: 27138},
		SpellSchool:     core.SpellSchoolHoly,
		ProcMask:        core.ProcMaskSpellDamage,
		Flags:           core.SpellFlagMeleeMetrics,
		BonusSpellPower: bonusSpellPower,
		ManaCost: core.ManaCostOptions{
			BaseCost:   0.08,
			Multiplier: 1 - 0.02*float64(paladin.Talents.Benediction),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: time.Second * 15,
			},
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				if paladin.CurrentMana() >= cast.Cost {
					castTime := time.Duration(float64(cast.CastTime) * spell.CastTimeMultiplier)
					if castTime > 0 {
						paladin.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime+castTime, false)
					}
				}
			},
		},

		DamageMultiplierAdditive: 1 +
			paladin.getTalentSanctityOfBattleBonus() +
			paladin.getMajorGlyphOfExorcismBonus(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		CritMultiplier:   paladin.SpellCritMultiplier(),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(694, 772) +
				.15*spell.SpellPower() +
				.15*spell.MeleeAttackPower()

			bonusCrit := core.TernaryFloat64(
				target.MobType == proto.MobType_MobTypeDemon || target.MobType == proto.MobType_MobTypeUndead,
				100*core.CritRatingPerCritChance,
				0)

			spell.BonusCritRating += bonusCrit
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			spell.BonusCritRating -= bonusCrit
		},
	})
}
