package warlock

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
)

func (warlock *Warlock) registerHauntSpell() {
	actionID := core.ActionID{SpellID: 59164}

	debuffMultiplier := 1.2
	if warlock.HasMajorGlyph(proto.WarlockMajorGlyph_GlyphOfHaunt) {
		debuffMultiplier += 0.03
	}

	warlock.HauntDebuffAura = warlock.CurrentTarget.GetOrRegisterAura(core.Aura{
		Label:    "Haunt-" + warlock.Label,
		ActionID: core.ActionID{SpellID: 59164},
		Duration: time.Second * 12,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			warlock.AttackTables[aura.Unit.UnitIndex].PeriodicShadowDamageTakenMultiplier *= debuffMultiplier
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			warlock.AttackTables[aura.Unit.UnitIndex].PeriodicShadowDamageTakenMultiplier /= debuffMultiplier
		},
	})

	warlock.Haunt = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     actionID,
		SpellSchool:  core.SpellSchoolShadow,
		ProcMask:     core.ProcMaskSpellDamage,
		MissileSpeed: 20,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.12,
			Multiplier: 1 - 0.02*float64(warlock.Talents.Suppression),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
			CD: core.Cooldown{
				Timer:    warlock.NewTimer(),
				Duration: time.Second * 8,
			},
		},

		BonusCritRating: 0 +
			warlock.masterDemonologistShadowCrit,
		DamageMultiplierAdditive: 1 +
			warlock.GrandFirestoneBonus() +
			0.03*float64(warlock.Talents.ShadowMastery),
		CritMultiplier:   warlock.SpellCritMultiplier(1, core.TernaryFloat64(warlock.Talents.Pandemic, 1, 0)),
		ThreatMultiplier: 1 - 0.1*float64(warlock.Talents.ImprovedDrainSoul),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(487, 569) + 0.4286*spell.SpellPower()
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
				if result.Landed() {
					warlock.HauntDebuffAura.Activate(sim)
				}
			})
		},
	})
}
