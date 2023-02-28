package wotlk

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

type StackingStatBonusEffect struct {
	Name       string
	ID         int32
	Bonus      stats.Stats
	Duration   time.Duration
	MaxStacks  int32
	Callback   core.AuraCallback
	ProcMask   core.ProcMask
	SpellFlags core.SpellFlag
	Outcome    core.HitOutcome
	Harmful    bool
	ProcChance float64
}

func newStackingStatBonusEffect(config StackingStatBonusEffect) {
	core.NewItemEffect(config.ID, func(agent core.Agent) {
		character := agent.GetCharacter()

		procAura := core.MakeStackingAura(character, core.StackingStatAura{
			Aura: core.Aura{
				Label:     config.Name + " Proc",
				ActionID:  core.ActionID{ItemID: config.ID},
				Duration:  config.Duration,
				MaxStacks: config.MaxStacks,
			},
			BonusPerStack: config.Bonus,
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       config.Name,
			Callback:   config.Callback,
			ProcMask:   config.ProcMask,
			SpellFlags: config.SpellFlags,
			Outcome:    config.Outcome,
			Harmful:    config.Harmful,
			ProcChance: config.ProcChance,
			Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
				procAura.Activate(sim)
				procAura.AddStack(sim)
			},
		})
	})
}

type StackingStatBonusCD struct {
	Name        string
	ID          int32
	Bonus       stats.Stats
	Duration    time.Duration
	MaxStacks   int32
	CD          time.Duration
	Callback    core.AuraCallback
	ProcMask    core.ProcMask
	SpellFlags  core.SpellFlag
	Outcome     core.HitOutcome
	Harmful     bool
	ProcChance  float64
	IsDefensive bool
}

func newStackingStatBonusCD(config StackingStatBonusCD) {
	core.NewItemEffect(config.ID, func(agent core.Agent) {
		character := agent.GetCharacter()

		buffAura := core.MakeStackingAura(character, core.StackingStatAura{
			Aura: core.Aura{
				Label:     config.Name + " Aura",
				ActionID:  core.ActionID{ItemID: config.ID},
				Duration:  config.Duration,
				MaxStacks: config.MaxStacks,
			},
			BonusPerStack: config.Bonus,
		})

		core.ApplyProcTriggerCallback(&character.Unit, buffAura, core.ProcTrigger{
			Name:       config.Name,
			Callback:   config.Callback,
			ProcMask:   config.ProcMask,
			SpellFlags: config.SpellFlags,
			Outcome:    config.Outcome,
			Harmful:    config.Harmful,
			ProcChance: config.ProcChance,
			Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
				buffAura.AddStack(sim)
			},
		})

		var sharedTimer *core.Timer
		if config.IsDefensive {
			sharedTimer = character.GetDefensiveTrinketCD()
		} else {
			sharedTimer = character.GetOffensiveTrinketCD()
		}

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{ItemID: config.ID},
			Flags:    core.SpellFlagNoOnCastComplete,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: config.CD,
				},
				SharedCD: core.Cooldown{
					Timer:    sharedTimer,
					Duration: config.Duration,
				},
			},

			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Spell: spell,
			Type:  core.CooldownTypeDPS,
		})
	})
}

func init() {
	core.AddEffectsToTest = true
	core.NewItemEffect(34427, func(agent core.Agent) {
		character := agent.GetCharacter()

		procAura := core.MakeStackingAura(character, core.StackingStatAura{
			Aura: core.Aura{
				Label:     "Blackened Naaru Sliver Proc",
				ActionID:  core.ActionID{ItemID: 34427},
				Duration:  time.Second * 20,
				MaxStacks: 10,
				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
						aura.AddStack(sim)
					}
				},
			},
			BonusPerStack: stats.Stats{stats.AttackPower: 44, stats.RangedAttackPower: 44},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Blackened Naaru Sliver",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMeleeOrRanged,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.1,
			ICD:        time.Second * 45,
			Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
				procAura.Activate(sim)
			},
		})
	})

	core.AddEffectsToTest = false
}
