package wotlk

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

// Keep these in alphabetical order.

var ItemSetWrathOfSpellFire = core.NewItemSet(core.ItemSet{
	Name: "Wrath of Spellfire",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStatDependency(stats.Intellect, stats.SpellPower, 0.07)
		},
	},
})

var ItemSetEbonNetherscale = core.NewItemSet(core.ItemSet{
	Name: "Netherscale Armor",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 20)
			agent.GetCharacter().AddStat(stats.SpellHit, 20)
		},
	},
})

var ItemSetNetherstrike = core.NewItemSet(core.ItemSet{
	Name: "Netherstrike Armor",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 23)
		},
	},
})

var ItemSetRagesteel = core.NewItemSet(core.ItemSet{
	Name: "Burning Rage",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 20)
			agent.GetCharacter().AddStat(stats.SpellHit, 20)
		},
	},
})

var ItemSetTwinStars = core.NewItemSet(core.ItemSet{
	Name: "The Twin Stars",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 15)
		},
	},
})

var ItemSetSpellstrike = core.NewItemSet(core.ItemSet{
	Name: "Spellstrike Infusion",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			procAura := character.NewTemporaryStatsAura("Spellstrike Proc", core.ActionID{SpellID: 32106}, stats.Stats{stats.SpellPower: 92}, time.Second*10)

			character.RegisterAura(core.Aura{
				Label:    "Spellstrike",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if sim.RandomFloat("spellstrike") > 0.05 {
						return
					}
					procAura.Activate(sim)
				},
			})
		},
	},
})

var ItemSetManaEtched = core.NewItemSet(core.ItemSet{
	Name: "Mana-Etched Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			procAura := character.NewTemporaryStatsAura("Mana-Etched Insight Proc", core.ActionID{SpellID: 37619}, stats.Stats{stats.SpellPower: 110}, time.Second*15)

			character.RegisterAura(core.Aura{
				Label:    "Mana-Etched Insight",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if sim.RandomFloat("Mana-Etched Insight") > 0.02 {
						return
					}
					procAura.Activate(sim)
				},
			})
		},
	},
})
