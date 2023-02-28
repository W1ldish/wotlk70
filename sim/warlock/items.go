package warlock

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

var ItemSetOblivionRaiment = core.NewItemSet(core.ItemSet{
	Name: "Oblivion Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// in pet.go constructor
		},
		4: func(agent core.Agent) {
			// in seed.go
		},
	},
})

var ItemSetVoidheartRaiment = core.NewItemSet(core.ItemSet{
	Name: "Voidheart Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()

			shadowBonus := warlock.RegisterAura(core.Aura{
				Label:    "Shadowflame",
				Duration: time.Second * 15,
				ActionID: core.ActionID{SpellID: 37377},
			})

			fireBonus := warlock.RegisterAura(core.Aura{
				Label:    "Shadowflame Hellfire",
				Duration: time.Second * 15,
				ActionID: core.ActionID{SpellID: 39437},
			})

			warlock.RegisterAura(core.Aura{
				Label:    "Voidheart Raiment 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if sim.RandomFloat("cycl4p") > 0.05 {
						return
					}
					if spell.SpellSchool.Matches(core.SpellSchoolShadow) {
						shadowBonus.Activate(sim)
					}
					if spell.SpellSchool.Matches(core.SpellSchoolFire) {
						fireBonus.Activate(sim)
					}
				},
			})
		},
		4: func(agent core.Agent) {
			// implemented in immolate.go and corruption.go
		},
	},
})

var ItemSetCorruptorRaiment = core.NewItemSet(core.ItemSet{
	Name: "Corruptor Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// heals pet
		},
		4: func(agent core.Agent) {
			// implemented in immolate.go and corruption.go
		},
	},
})

// T6
var ItemSetMaleficRaiment = core.NewItemSet(core.ItemSet{
	Name: "Malefic Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// heals... not implemented yet
		},
		4: func(agent core.Agent) {
			// Increases damage done by shadowbolt and incinerate by 6%.
			// Implemented in shadowbolt.go and incinerate.go
		},
	},
})

func init() {

	core.NewItemEffect(30449, func(agent core.Agent) {
		warlock := agent.(WarlockAgent).GetWarlock()
		warlock.AddStat(stats.SpellPower, 48)
		if warlock.Pet != nil {
			warlock.Pet.AddStats(stats.Stats{
				stats.ArcaneResistance: 130,
				stats.FireResistance:   130,
				stats.FrostResistance:  130,
				stats.NatureResistance: 130,
				stats.ShadowResistance: 130,
			})
		}
	})

	core.NewItemEffect(32493, func(agent core.Agent) {
		warlock := agent.(WarlockAgent).GetWarlock()
		procAura := warlock.NewTemporaryStatsAura("Asghtongue Talisman Proc", core.ActionID{SpellID: 40478}, stats.Stats{stats.SpellPower: 220}, time.Second*5)

		warlock.RegisterAura(core.Aura{
			Label:    "Ashtongue Talisman",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell == warlock.Corruption && sim.RandomFloat("Ashtongue Talisman of Insight") < 0.2 {
					procAura.Activate(sim)
				}
			},
		})
	})
}
