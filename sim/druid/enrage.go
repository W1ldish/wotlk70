package druid

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
)

func (druid *Druid) registerEnrageSpell() {
	actionID := core.ActionID{SpellID: 5229}
	rageMetrics := druid.NewRageMetrics(actionID)

	instantRage := []float64{20, 24, 27, 30}[druid.Talents.Intensity]

	dmgBonus := 0.05 * float64(druid.Talents.KingOfTheJungle)
	armorLoss := druid.ScaleBaseArmor(0.16 * druid.TotalBearArmorMultiplier())

	druid.EnrageAura = druid.RegisterAura(core.Aura{
		Label:    "Enrage Aura",
		ActionID: actionID,
		Duration: 10 * time.Second,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			druid.PseudoStats.DamageDealtMultiplier *= 1.0 + dmgBonus
			druid.AddStatDynamic(sim, stats.Armor, -armorLoss)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			druid.PseudoStats.DamageDealtMultiplier /= 1.0 + dmgBonus
			druid.AddStatDynamic(sim, stats.Armor, armorLoss)
		},
	})

	spell := druid.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    druid.NewTimer(),
				Duration: time.Minute,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			druid.AddRage(sim, instantRage, rageMetrics)

			core.StartPeriodicAction(sim, core.PeriodicActionOptions{
				NumTicks: 10,
				Period:   time.Second * 1,
				OnAction: func(sim *core.Simulation) {
					if druid.EnrageAura.IsActive() {
						druid.AddRage(sim, 1, rageMetrics)
					}
				},
			})

			druid.EnrageAura.Activate(sim)
		},
	})

	druid.Enrage = spell
}
