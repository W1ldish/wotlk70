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

	newProcDamageEffect(ProcDamageEffect{
		ID: 12631,
		Trigger: core.ProcTrigger{
			Name:       "Fiery Plate Gauntlets",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMelee,
			Outcome:    core.OutcomeLanded,
			ProcChance: 1.0,
		},
		School: core.SpellSchoolFire,
		MinDmg: 4,
		MaxDmg: 4,
	})

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
}
