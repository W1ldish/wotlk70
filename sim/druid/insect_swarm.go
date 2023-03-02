package druid

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
)

func (druid *Druid) registerInsectSwarmSpell() {
	missAuras := druid.NewEnemyAuraArray(core.InsectSwarmAura)
	hasGlyph := druid.HasMajorGlyph(proto.DruidMajorGlyph_GlyphOfInsectSwarm)

	impISMultiplier := 1 + 0.01*float64(druid.Talents.ImprovedInsectSwarm)

	druid.InsectSwarm = druid.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 27013},
		SpellSchool: core.SpellSchoolNature,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagOmenTrigger,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.08,
			Multiplier: 1,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		DamageMultiplier: 1 +
			0.01*float64(druid.Talents.Genesis) +
			core.TernaryFloat64(druid.HasMajorGlyph(proto.DruidMajorGlyph_GlyphOfInsectSwarm), 0.3, 0),
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Insect Swarm",
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					druid.Wrath.DamageMultiplier *= impISMultiplier
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					druid.Wrath.DamageMultiplier /= impISMultiplier
				},
			},
			NumberOfTicks: 6 + core.TernaryInt32(druid.Talents.NaturesSplendor, 1, 0),
			TickLength:    time.Second * 2,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.SnapshotBaseDamage = 172 + 0.2*(dot.Spell.SpellPower())
				dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(dot.Spell.Unit.AttackTables[target.UnitIndex])
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)

				if druid.MoonkinT84PCAura != nil && sim.RandomFloat("Elune's Wrath proc") < 0.08 {
					druid.MoonkinT84PCAura.Activate(sim)
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHit)
			if result.Landed() {
				spell.Dot(target).Apply(sim)
				if !hasGlyph {
					missAuras.Get(target).Activate(sim)
				}
			}
			spell.DealOutcome(sim, result)
		},
	})
}
