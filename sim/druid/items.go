package druid

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

// T4 Balance
var ItemSetMalorneRegalia = core.NewItemSet(core.ItemSet{
	Name: "Malorne Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			druid := agent.(DruidAgent).GetDruid()
			manaMetrics := druid.NewManaMetrics(core.ActionID{SpellID: 37295})

			druid.RegisterAura(core.Aura{
				Label:    "Malorne Regalia 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
						return
					}
					if !result.Landed() {
						return
					}
					if sim.RandomFloat("malorne 2p") > 0.05 {
						return
					}
					spell.Unit.AddMana(sim, 120, manaMetrics)
				},
			})
		},
		4: func(agent core.Agent) {
			// Currently this is handled in druid.go (reducing CD of innervate)
		},
	},
})

//T4 feral
var ItemSetMalorneHarness = core.NewItemSet(core.ItemSet{
	Name: "Malorne Harness",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			druid := agent.(DruidAgent).GetDruid()

			procChance := 0.04
			rageMetrics := druid.NewRageMetrics(core.ActionID{SpellID: 37306})
			energyMetrics := druid.NewEnergyMetrics(core.ActionID{SpellID: 37311})

			druid.RegisterAura(core.Aura{
				Label:    "Malorne 4pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
						if sim.RandomFloat("Malorne 2pc") < procChance {
							if druid.InForm(Bear) {
								druid.AddRage(sim, 10, rageMetrics)
							} else if druid.InForm(Cat) {
								druid.AddEnergy(sim, 20, energyMetrics)
							}
						}
					}
				},
			})
		},
		4: func(agent core.Agent) {
			druid := agent.(DruidAgent).GetDruid()
			if druid.InForm(Bear) {
				druid.AddStat(stats.Armor, 1400)
			} else if druid.InForm(Cat) {
				druid.AddStat(stats.Strength, 30)
			}
		},
	},
})

//T5 Balance
var ItemSetNordrassilRegalia = core.NewItemSet(core.ItemSet{
	Name: "Nordrassil Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			// Implemented in starfire.go.
		},
	},
})

//T5 Feral
var ItemSetNordrassilHarness = core.NewItemSet(core.ItemSet{
	Name: "Nordrassil Harness",
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			// Implemented in lacerate.go.
		},
	},
})

//T6 Balance
var ItemSetThunderheartRegalia = core.NewItemSet(core.ItemSet{
	Name: "Thunderheart Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Implemented in moonfire.go
		},
		4: func(agent core.Agent) {
			// Implemented in starfire.go
		},
	},
})

//T6 Feral
var ItemSetThunderheartHarness = core.NewItemSet(core.ItemSet{
	Name: "Thunderheart Harness",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Implemented in mangle.go.
		},
		4: func(agent core.Agent) {
			// Implemented in swipe.go.
		},
	},
})

var ItemSetGladiatorsWildhide = core.NewItemSet(core.ItemSet{
	Name: "Gladiator's Wildhide",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			druid := agent.(DruidAgent).GetDruid()
			druid.AddStat(stats.SpellPower, 29)
			druid.AddStat(stats.Resilience, 100)
		},
		4: func(agent core.Agent) {
			druid := agent.(DruidAgent).GetDruid()
			druid.AddStat(stats.SpellPower, 88)

			percentReduction := float64(time.Millisecond*1500) / float64(druid.starfireCastTime())
			swiftStarfireAura := druid.RegisterAura(core.Aura{
				Label:    "Swift Starfire",
				ActionID: core.ActionID{SpellID: 46832},
				Duration: time.Second * 15,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					druid.Starfire.CastTimeMultiplier -= percentReduction
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					druid.Starfire.CastTimeMultiplier += percentReduction
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if spell == druid.Starfire {
						aura.Deactivate(sim)
					}
				},
			})

			druid.RegisterAura(core.Aura{
				Label:    "Swift Starfire trigger",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if spell == druid.Wrath && sim.RandomFloat("Swift Starfire proc") > 0.85 {
						swiftStarfireAura.Activate(sim)
					}
				},
			})
		},
	},
})

func init() {

	core.NewItemEffect(32486, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()

		// Not in the game yet so cant test; this logic assumes that:
		// - does not affect the starfire which procs it
		// - can proc off of any completed cast, not just hits
		actionID := core.ActionID{ItemID: 32486}

		var procAura *core.Aura
		if druid.InForm(Moonkin) {
			procAura = druid.NewTemporaryStatsAura("Ashtongue Talisman Proc", actionID, stats.Stats{stats.SpellPower: 150}, time.Second*8)
		} else if druid.InForm(Bear | Cat) {
			procAura = druid.NewTemporaryStatsAura("Ashtongue Talisman Proc", actionID, stats.Stats{stats.Strength: 140}, time.Second*8)
		} else {
			return
		}

		core.MakePermanent(druid.RegisterAura(core.Aura{
			Label: "Ashtongue Talisman",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() {
					return
				}
				if spell == druid.Starfire {
					if sim.RandomFloat("Ashtongue Talisman") < 0.25 {
						procAura.Activate(sim)
					}
				} else if druid.IsMangle(spell) {
					if sim.RandomFloat("Ashtongue Talisman") < 0.4 {
						procAura.Activate(sim)
					}
				}
			},
		}))
	})

	core.NewItemEffect(32257, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()

		procAura := druid.NewTemporaryStatsAura("Idol of the White Stag Proc", core.ActionID{ItemID: 32257}, stats.Stats{stats.AttackPower: 94}, time.Second*20)

		core.MakePermanent(druid.RegisterAura(core.Aura{
			Label: "Idol of the White Stag",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if druid.IsMangle(spell) {
					procAura.Activate(sim)
				}
			},
		}))
	})

	core.NewItemEffect(33509, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()
		procAura := druid.NewTemporaryStatsAura("Idol of Terror Proc", core.ActionID{SpellID: 43737}, stats.Stats{stats.Agility: 65}, time.Second*10)
		icd := core.Cooldown{
			Timer:    druid.NewTimer(),
			Duration: time.Second * 10,
		}
		druid.RegisterAura(core.Aura{
			Label:    "Idol of Terror",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !druid.IsMangle(spell) || !icd.IsReady(sim) {
					return
				}
				if sim.RandomFloat("Idol of Terror") > 0.85 {
					return
				}
				icd.Use(sim)
				procAura.Activate(sim)

			},
		})
	})

	core.NewItemEffect(33510, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()

		procAura := druid.NewTemporaryStatsAura("Idol of the Unseen Moon Proc", core.ActionID{ItemID: 33510}, stats.Stats{stats.SpellPower: 140}, time.Second*10)

		druid.RegisterAura(core.Aura{
			Label:    "Idol of the Unseen Moon",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() {
					return
				}

				if spell == druid.Moonfire {
					if sim.RandomFloat("Idol of the Unseen Moon") > 0.5 {
						return
					}
					procAura.Activate(sim)
				}
			},
		})
	})

	core.NewItemEffect(30664, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()

		var procAura *core.Aura
		if druid.InForm(Moonkin) {
			procAura = druid.NewTemporaryStatsAura("Living Root Moonkin Proc", core.ActionID{SpellID: 37343}, stats.Stats{stats.SpellPower: 209}, time.Second*15)
		} else if druid.InForm(Bear) {
			procAura = druid.NewTemporaryStatsAura("Living Root Bear Proc", core.ActionID{SpellID: 37340}, stats.Stats{stats.Armor: 4070}, time.Second*15)
		} else if druid.InForm(Cat) {
			procAura = druid.NewTemporaryStatsAura("Living Root Cat Proc", core.ActionID{SpellID: 37341}, stats.Stats{stats.Strength: 64}, time.Second*15)
		} else {
			return
		}

		druid.RegisterAura(core.Aura{
			Label:    "Living Root of the Wildheart",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if druid.InForm(Moonkin) && sim.RandomFloat("Living Root of the Wildheart") < 0.03 {
					procAura.Activate(sim)
				}
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !spell.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("Living Root of the Wildheart") > 0.03 {
					return
				}

				procAura.Activate(sim)
			},
		})
	})

	core.NewItemEffect(32387, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()
		core.MakePermanent(druid.RegisterAura(core.Aura{
			Label:      "Idol of the Raven Goddess",
			BuildPhase: core.CharacterBuildPhaseGear,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				// For now this assume we'll never leave main form
				if druid.StartingForm.Matches(Bear | Cat) {
					druid.AddStatDynamic(sim, stats.MeleeCrit, 40.0)
				} else if druid.StartingForm.Matches(Moonkin) {
					druid.AddStatDynamic(sim, stats.SpellCrit, 40.0)
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				if druid.StartingForm.Matches(Bear | Cat) {
					druid.AddStatDynamic(sim, stats.MeleeCrit, -40.0)
				} else if druid.StartingForm.Matches(Moonkin) {
					druid.AddStatDynamic(sim, stats.SpellCrit, -40.0)
				}
			},
		}))
	})

	core.NewItemEffect(33947, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()
		procAura := druid.NewTemporaryStatsAura("Vengeful Gladiator's Idol of Resolve Proc", core.ActionID{ItemID: 33947}, stats.Stats{stats.Resilience: 34}, time.Second*6)

		core.MakePermanent(druid.RegisterAura(core.Aura{
			Label: "Vengeful Gladiator's Idol of Resolve",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !druid.IsMangle(spell) {
					return
				}
				procAura.Activate(sim)
			},
		}))
	})

	core.NewItemEffect(35019, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()
		procAura := druid.NewTemporaryStatsAura("Brutal Gladiator's Idol of Resolve Proc", core.ActionID{ItemID: 35019}, stats.Stats{stats.Resilience: 39}, time.Second*6)

		core.MakePermanent(druid.RegisterAura(core.Aura{
			Label: "Brutal Gladiator's Idol of Resolve",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !druid.IsMangle(spell) {
					return
				}
				procAura.Activate(sim)
			},
		}))
	})
}
