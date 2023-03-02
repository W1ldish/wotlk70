package shaman

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
)

// Totem IDs
const ()

func (shaman *Shaman) registerLavaLashSpell() {
	if !shaman.Talents.LavaLash {
		return
	}

	imbueMultiplier := 1.0
	if shaman.SelfBuffs.ImbueOH == proto.ShamanImbue_FlametongueWeapon || shaman.SelfBuffs.ImbueOH == proto.ShamanImbue_FlametongueWeaponDownrank {
		imbueMultiplier = 1.25
		if shaman.HasMajorGlyph(proto.ShamanMajorGlyph_GlyphOfLavaLash) {
			imbueMultiplier = 1.35
		}
	}

	var indomitabilityAura *core.Aura

	shaman.LavaLash = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 60103},
		SpellSchool: core.SpellSchoolFire,
		ProcMask:    core.ProcMaskMeleeOHSpecial,
		Flags:       core.SpellFlagMeleeMetrics,

		ManaCost: core.ManaCostOptions{
			BaseCost: 0.04,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    shaman.NewTimer(),
				Duration: time.Second * 6,
			},
		},

		DamageMultiplier: imbueMultiplier *
			core.TernaryFloat64(shaman.HasSetBonus(ItemSetWorldbreakerBattlegear, 2), 1.2, 1),
		CritMultiplier:   shaman.ElementalCritMultiplier(0),
		ThreatMultiplier: shaman.spellThreatMultiplier(),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := spell.Unit.OHWeaponDamage(sim, spell.MeleeAttackPower()) +
				spell.BonusWeaponDamage()

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if result.Landed() { //TODO: verify that it actually needs to hit
				if indomitabilityAura != nil {
					indomitabilityAura.Activate(sim)
				}
			}
		},
	})
}

func (shaman *Shaman) IsLavaLashCastable(sim *core.Simulation) bool {
	return shaman.LavaLash.IsReady(sim)
}
