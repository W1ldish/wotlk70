package tbc

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

func init() {
	core.AddEffectsToTest = false
	core.NewSimpleStatItemEffect(28484, stats.Stats{stats.Health: 1500, stats.Strength: 150}, time.Second*15, time.Minute*30) // Bulwark of Kings
	core.NewSimpleStatItemEffect(28485, stats.Stats{stats.Health: 1500, stats.Strength: 150}, time.Second*15, time.Minute*30) // Bulwark of Ancient Kings

	// Proc effects. Keep these in order by item ID.

	core.NewItemEffect(24114, func(agent core.Agent) {
		agent.GetCharacter().PseudoStats.BonusDamage += 5
	})

	core.NewItemEffect(29297, func(agent core.Agent) {
		character := agent.GetCharacter()

		const procChance = 0.03
		procAura := character.NewTemporaryStatsAura("Band of the Eternal Defender Proc", core.ActionID{ItemID: 29297}, stats.Stats{stats.Armor: 800}, time.Second*10)

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 60,
		}

		character.GetOrRegisterAura(core.Aura{
			Label:    "Band of the Eternal Defender",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || spell.SpellSchool != core.SpellSchoolPhysical {
					return
				}
				if !icd.IsReady(sim) {
					return
				}
				if sim.RandomFloat("Band of the Eternal Defender") < procChance {
					icd.Use(sim)
					procAura.Activate(sim)
				}
			},
		})
	})

	core.NewItemEffect(29962, func(agent core.Agent) {
		character := agent.GetCharacter()
		mh, oh := character.GetWeaponHands(29962)
		procMask := core.GetMeleeProcMaskForHands(mh, oh)

		procAura := character.NewTemporaryStatsAura("Heartrazor Proc", core.ActionID{ItemID: 29962}, stats.Stats{stats.AttackPower: 270, stats.RangedAttackPower: 270}, time.Second*10)
		ppmm := character.AutoAttacks.NewPPMManager(1.0, procMask)

		character.GetOrRegisterAura(core.Aura{
			Label:    "Heartrazor",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(procMask) {
					return
				}

				if !ppmm.Proc(sim, spell.ProcMask, "Heartrazor") {
					return
				}

				procAura.Activate(sim)
			},
		})
	})

	core.NewItemEffect(29996, func(agent core.Agent) {
		character := agent.GetCharacter()
		mh, oh := character.GetWeaponHands(29996)
		procMask := core.GetMeleeProcMaskForHands(mh, oh)

		const procChance = 2.7 / 60.0
		actionID := core.ActionID{ItemID: 29996}

		var resourceMetrics *core.ResourceMetrics
		if character.HasRageBar() {
			resourceMetrics = character.NewRageMetrics(actionID)
		} else if character.HasEnergyBar() {
			resourceMetrics = character.NewEnergyMetrics(actionID)
		}

		character.GetOrRegisterAura(core.Aura{
			Label:    "Rod of the Sun King",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(procMask) {
					return
				}

				if spell.Unit.HasRageBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					spell.Unit.AddRage(sim, 5, resourceMetrics)
				} else if spell.Unit.HasEnergyBar() {
					if sim.RandomFloat("Rod of the Sun King") > procChance {
						return
					}
					spell.Unit.AddEnergy(sim, 10, resourceMetrics)
				}
			},
		})
	})

	core.NewItemEffect(30090, func(agent core.Agent) {
		character := agent.GetCharacter()

		const procChance = 3.7 / 60.0
		procAura := character.NewTemporaryStatsAura("World Breaker Proc", core.ActionID{ItemID: 30090}, stats.Stats{stats.MeleeCrit: 900}, time.Second*4)

		character.RegisterAura(core.Aura{
			Label:    "World Breaker",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(core.ProcMaskMelee) {
					procAura.Deactivate(sim)
					return
				}
				if sim.RandomFloat("World Breaker") > procChance {
					procAura.Deactivate(sim)
					return
				}

				procAura.Activate(sim)
			},
		})
	})

	core.NewItemEffect(30311, func(agent core.Agent) {
		character := agent.GetCharacter()
		mh, oh := character.GetWeaponHands(30311)
		procMask := core.GetMeleeProcMaskForHands(mh, oh)

		const bonus = 1.2
		const inverseBonus = 1 / 1.2
		const procChance = 0.5

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Warp Slicer Proc",
			ActionID: core.ActionID{ItemID: 30311},
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyMeleeSpeed(sim, bonus)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyMeleeSpeed(sim, inverseBonus)
			},
		})

		character.GetOrRegisterAura(core.Aura{
			Label:    "Warp Slicer",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("WarpSlicer") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})

	core.NewItemEffect(30316, func(agent core.Agent) {
		character := agent.GetCharacter()

		const bonus = 1.2
		const inverseBonus = 1 / 1.2
		const procChance = 0.5

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Devastation Proc",
			ActionID: core.ActionID{ItemID: 30316},
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyMeleeSpeed(sim, bonus)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyMeleeSpeed(sim, inverseBonus)
			},
		})

		character.GetOrRegisterAura(core.Aura{
			Label:    "Devastation",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}
				if sim.RandomFloat("Devastation") > procChance {
					return
				}

				procAura.Activate(sim)
			},
		})
	})

	core.NewItemEffect(31332, func(agent core.Agent) {
		character := agent.GetCharacter()
		mh, oh := character.GetWeaponHands(31193)
		procMask := core.GetMeleeProcMaskForHands(mh, oh)

		ppmm := character.AutoAttacks.NewPPMManager(1.0, procMask)
		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Millisecond,
		}

		blinkStrikeSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 38308},
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    procMask,
			Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete,

			DamageMultiplier: character.AutoAttacks.MHConfig.DamageMultiplier,
			CritMultiplier:   character.DefaultMeleeCritMultiplier(),
			ThreatMultiplier: character.AutoAttacks.MHConfig.ThreatMultiplier,

			ApplyEffects: character.AutoAttacks.MHConfig.ApplyEffects,
		})

		character.RegisterAura(core.Aura{
			Label:    "Blinkstrike",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(procMask) {
					return
				}
				if !icd.IsReady(sim) {
					return
				}
				if !ppmm.Proc(sim, spell.ProcMask, "Blinkstrike") {
					return
				}
				icd.Use(sim)
				blinkStrikeSpell.Cast(sim, result.Target)
			},
		})

	})

	core.NewItemEffect(31193, func(agent core.Agent) {
		character := agent.GetCharacter()
		mh, oh := character.GetWeaponHands(31193)
		procMask := core.GetMeleeProcMaskForHands(mh, oh)

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 24585},
			SpellSchool: core.SpellSchoolShadow,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			CritMultiplier:   character.DefaultSpellCritMultiplier(),
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				baseDamage := sim.Roll(48, 54) + spell.SpellPower()
				spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Blade of Unquenched Thirst Trigger",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   procMask,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.02,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, result.Target)
			},
		})
	})

	core.NewItemEffect(31331, func(agent core.Agent) {
		character := agent.GetCharacter()
		mh, oh := character.GetWeaponHands(31331)
		procMask := core.GetMeleeProcMaskForHands(mh, oh)

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:     "The Night Blade Proc",
			ActionID:  core.ActionID{ItemID: 31331},
			Duration:  time.Second * 10,
			MaxStacks: 3,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
				character.AddStatDynamic(sim, stats.ArmorPenetration, 435*float64(newStacks-oldStacks))
			},
		})

		const procChance = 2 * 1.8 / 60.0
		character.GetOrRegisterAura(core.Aura{
			Label:    "The Night Blade",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(procMask) {
					return
				}
				if sim.RandomFloat("The Night Blade") > procChance {
					return
				}

				procAura.Activate(sim)
				procAura.AddStack(sim)
			},
		})
	})

	core.NewItemEffect(32375, func(agent core.Agent) {
		character := agent.GetCharacter()

		const procChance = 0.02
		procAura := character.NewTemporaryStatsAura("Bulwark Of Azzinoth Proc", core.ActionID{ItemID: 32375}, stats.Stats{stats.Armor: 2000}, time.Second*10)

		character.GetOrRegisterAura(core.Aura{
			Label:    "Bulwark Of Azzinoth",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.SpellSchool == core.SpellSchoolPhysical && sim.RandomFloat("Bulwark of Azzinoth") < procChance {
					procAura.Activate(sim)
				}
			},
		})
	})

	core.NewItemEffect(34473, func(agent core.Agent) {
		character := agent.GetCharacter()

		procAura := character.NewTemporaryStatsAura("Commendation of Kael'Thas Proc", core.ActionID{ItemID: 34473}, stats.Stats{stats.Dodge: 152}, time.Second*10)

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 30,
		}

		character.GetOrRegisterAura(core.Aura{
			Label:    "Commendation of Kael'Thas",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if aura.Unit.CurrentHealthPercent() >= 0.35 {
					return
				}

				if !icd.IsReady(sim) {
					return
				}

				icd.Use(sim)
				procAura.Activate(sim)
			},
		})
	})
	core.AddEffectsToTest = true
}
