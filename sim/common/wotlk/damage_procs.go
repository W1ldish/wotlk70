package wotlk

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
)

type ProcDamageEffect struct {
	ID      int32
	Trigger core.ProcTrigger

	School core.SpellSchool
	MinDmg float64
	MaxDmg float64
}

func newProcDamageEffect(config ProcDamageEffect) {
	core.NewItemEffect(config.ID, func(agent core.Agent) {
		character := agent.GetCharacter()

		minDmg := config.MinDmg
		maxDmg := config.MaxDmg
		damageSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{ItemID: config.ID},
			SpellSchool: config.School,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			CritMultiplier:   character.DefaultSpellCritMultiplier(),
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(minDmg, maxDmg), spell.OutcomeMagicHitAndCrit)
			},
		})

		triggerConfig := config.Trigger
		triggerConfig.Handler = func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
			damageSpell.Cast(sim, character.CurrentTarget)
		}
		core.MakeProcTriggerAura(&character.Unit, triggerConfig)
	})
}

func init() {
	core.AddEffectsToTest = false

	newProcDamageEffect(ProcDamageEffect{
		ID: 12632,
		Trigger: core.ProcTrigger{
			Name:       "Storm Gauntlets",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMelee,
			Outcome:    core.OutcomeLanded,
			ProcChance: 1.0,
		},
		School: core.SpellSchoolNature,
		MinDmg: 3,
		MaxDmg: 3,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 17111,
		Trigger: core.ProcTrigger{
			Name:       "Blazefury Medallion",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMelee,
			Outcome:    core.OutcomeLanded,
			ProcChance: 1.0,
		},
		School: core.SpellSchoolFire,
		MinDmg: 2,
		MaxDmg: 2,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 28573,
		Trigger: core.ProcTrigger{
			Name:     "Despair",
			Callback: core.CallbackOnSpellHitDealt,
			ProcMask: core.ProcMaskMelee,
			Outcome:  core.OutcomeLanded,
			PPM:      0.5 * 3.5,
		},
		School: core.SpellSchoolPhysical,
		MinDmg: 600,
		MaxDmg: 600,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 28579,
		Trigger: core.ProcTrigger{
			Name:     "Romulo's Poison Vial",
			Callback: core.CallbackOnSpellHitDealt,
			ProcMask: core.ProcMaskMeleeOrRanged,
			Outcome:  core.OutcomeLanded,
			PPM:      1.0,
		},
		School: core.SpellSchoolNature,
		MinDmg: 222,
		MaxDmg: 332,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 28767,
		Trigger: core.ProcTrigger{
			Name:     "The Decapitator",
			Callback: core.CallbackOnSpellHitDealt,
			ProcMask: core.ProcMaskMelee,
			Outcome:  core.OutcomeLanded,
			ICD:      time.Minute * 3,
		},
		School: core.SpellSchoolPhysical,
		MinDmg: 513,
		MaxDmg: 567,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 28774,
		Trigger: core.ProcTrigger{
			Name: "Glaive of the Pit	",
			Callback: core.CallbackOnSpellHitDealt,
			ProcMask: core.ProcMaskMelee,
			Outcome:  core.OutcomeLanded,
			PPM:      3.7,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 285,
		MaxDmg: 315,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 34470,
		Trigger: core.ProcTrigger{
			Name:       "Timbal's Focusing Crystal",
			Callback:   core.CallbackOnPeriodicDamageDealt,
			Harmful:    true,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.1,
			ICD:        time.Second * 15,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 285,
		MaxDmg: 475,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 37064,
		Trigger: core.ProcTrigger{
			Name:       "Vestige of Haldor",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMeleeOrRanged,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolFire,
		MinDmg: 1024,
		MaxDmg: 1536,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 37264,
		Trigger: core.ProcTrigger{
			Name:       "Pendulum of Telluric Currents",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskSpellDamage,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 1168,
		MaxDmg: 1752,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 39889,
		Trigger: core.ProcTrigger{
			Name:       "Horn of Agent Fury",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMeleeOrRanged,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolHoly,
		MinDmg: 1024,
		MaxDmg: 1536,
	})

	core.AddEffectsToTest = true

	newProcDamageEffect(ProcDamageEffect{
		ID: 40371,
		Trigger: core.ProcTrigger{
			Name:       "Bandit's Insignia",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMeleeOrRanged,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolArcane,
		MinDmg: 1504,
		MaxDmg: 2256,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 40373,
		Trigger: core.ProcTrigger{
			Name:       "Extract of Necromantic Power",
			Callback:   core.CallbackOnPeriodicDamageDealt,
			Harmful:    true,
			ProcChance: 0.10,
			ICD:        time.Second * 15,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 788,
		MaxDmg: 1312,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 42990,
		Trigger: core.ProcTrigger{
			Name:       "DMC Death",
			Callback:   core.CallbackOnSpellHitDealt | core.CallbackOnPeriodicDamageDealt,
			Harmful:    true,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 1750,
		MaxDmg: 2250,
	})
}
