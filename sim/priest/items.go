package priest

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

var ItemSetIncarnateRegalia = core.NewItemSet(core.ItemSet{
	Name: "Incarnate Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Your shadowfiend now has 75 more stamina and lasts 3 sec. longer.
			// Implemented in shadowfiend.go.
		},
		4: func(agent core.Agent) {
			// Your Mind Flay and Smite spells deal 5% more damage.
			// Implemented in mind_flay.go.
		},
	},
})

var ItemSetAvatarRegalia = core.NewItemSet(core.ItemSet{
	Name: "Avatar Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			priest := agent.(PriestAgent).GetPriest()
			manaMetrics := priest.NewManaMetrics(core.ActionID{SpellID: 37600})

			priest.RegisterAura(core.Aura{
				Label:    "Avatar Regalia 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if sim.RandomFloat("avatar 2p") > 0.06 {
						return
					}
					// This is a cheat...
					// easier than adding another aura the subtracts 150 mana from next cast.
					priest.AddMana(sim, 150, manaMetrics)
				},
			})
		},
		4: func(agent core.Agent) {
			priest := agent.(PriestAgent).GetPriest()

			procAura := priest.NewTemporaryStatsAura("Avatar Regalia 4pc Proc", core.ActionID{SpellID: 37604}, stats.Stats{stats.SpellPower: 100}, time.Second*15)
			procAura.OnSpellHitDealt = func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				aura.Deactivate(sim)
			}

			priest.RegisterAura(core.Aura{
				Label:    "Avatar Regalia 4pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if spell != priest.ShadowWordPain {
						return
					}

					if sim.RandomFloat("avatar 4p") > 0.4 { // 60% chance of not activating.
						return
					}

					procAura.Activate(sim)
				},
			})
		},
	},
})

var ItemSetAbsolutionRegalia = core.NewItemSet(core.ItemSet{
	Name: "Absolution Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Implemented in swp.go
		},
		4: func(agent core.Agent) {
			// Implemented in mindblast.go
		},
	},
})

var ItemSetVestmentsOfAbsolution = core.NewItemSet(core.ItemSet{
	Name: "Vestments of Absolution",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Implemented in prayer_of_healing.go
		},
		4: func(agent core.Agent) {
			// Implemented in greater_heal.go
		},
	},
})

func init() {

	core.NewSimpleStatDefensiveTrinketEffect(30665, stats.Stats{stats.Spirit: 300}, time.Second*20, time.Minute*2)

	core.NewItemEffect(32490, func(agent core.Agent) {
		priest := agent.(PriestAgent).GetPriest()
		// Not in the game yet so cant test; this logic assumes that:
		// - procrate is 10%
		// - no ICD on proc
		const procrate = 0.1
		procAura := priest.NewTemporaryStatsAura("Ashtongue Talisman Proc", core.ActionID{ItemID: 32490}, stats.Stats{stats.SpellPower: 220}, time.Second*10)

		priest.RegisterAura(core.Aura{
			Label:    "Ashtongue Talisman",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell != priest.ShadowWordPain {
					return
				}

				if sim.RandomFloat("Ashtongue Talisman of Acumen") > procrate {
					return
				}

				procAura.Activate(sim)
			},
		})
	})

}
