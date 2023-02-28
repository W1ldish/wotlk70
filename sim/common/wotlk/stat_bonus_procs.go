package wotlk

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

type ProcStatBonusEffect struct {
	Name       string
	ID         int32
	Bonus      stats.Stats
	Duration   time.Duration
	Callback   core.AuraCallback
	ProcMask   core.ProcMask
	Outcome    core.HitOutcome
	Harmful    bool
	ProcChance float64
	PPM        float64
	ICD        time.Duration

	// For ignoring a hardcoded spell.
	IgnoreSpellID int32
}

func newProcStatBonusEffect(config ProcStatBonusEffect) {
	core.NewItemEffect(config.ID, func(agent core.Agent) {
		character := agent.GetCharacter()
		procAura := character.NewTemporaryStatsAura(config.Name+" Proc", core.ActionID{ItemID: config.ID}, config.Bonus, config.Duration)

		handler := func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
			procAura.Activate(sim)
		}
		if config.IgnoreSpellID != 0 {
			ignoreSpellID := config.IgnoreSpellID
			handler = func(sim *core.Simulation, spell *core.Spell, _ *core.SpellResult) {
				if !spell.IsSpellAction(ignoreSpellID) {
					procAura.Activate(sim)
				}
			}
		}

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       config.Name,
			Callback:   config.Callback,
			ProcMask:   config.ProcMask,
			Outcome:    config.Outcome,
			Harmful:    config.Harmful,
			ProcChance: config.ProcChance,
			PPM:        config.PPM,
			ICD:        config.ICD,
			Handler:    handler,
		})
	})
}

func init() {
	// Keep these separated by stat, ordered by item ID within each group.
	core.AddEffectsToTest = true
	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Empyrean Demolisher",
		ID:       17112,
		Bonus:    stats.Stats{stats.MeleeHaste: 212, stats.SpellHaste: 212},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      2.8,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Khorium Champion",
		ID:       23541,
		Bonus:    stats.Stats{stats.Strength: 120},
		Duration: time.Second * 30,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      1.65,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Blackout Truncheon",
		ID:       27901,
		Bonus:    stats.Stats{stats.MeleeHaste: 132, stats.SpellHaste: 132},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      1.5 * 0.8,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Quagmirran's Eye",
		ID:         27683,
		Bonus:      stats.Stats{stats.SpellHaste: 320},
		Duration:   time.Second * 6,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskSpellDamage,
		Outcome:    core.OutcomeHit,
		ProcChance: 0.1,
		ICD:        time.Second * 45,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Lionheart Champion",
		ID:       28429,
		Bonus:    stats.Stats{stats.Strength: 100},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      3.6,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Lionheart Executioner",
		ID:       28430,
		Bonus:    stats.Stats{stats.Strength: 100},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      3.6,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Drakefist Hammer",
		ID:       28437,
		Bonus:    stats.Stats{stats.MeleeHaste: 212, stats.SpellHaste: 212},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      2.7,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Dragonmaw",
		ID:       28438,
		Bonus:    stats.Stats{stats.MeleeHaste: 212, stats.SpellHaste: 212},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      2.7,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Dragonstrike",
		ID:       28439,
		Bonus:    stats.Stats{stats.MeleeHaste: 212, stats.SpellHaste: 212},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeHit,
		PPM:      2.7,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Robe of the Elder Scribes",
		ID:         28602,
		Bonus:      stats.Stats{stats.SpellPower: 130},
		Duration:   time.Second * 10,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskSpellDamage,
		Outcome:    core.OutcomeHit,
		ProcChance: 0.2,
		ICD:        time.Second * 50,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Shiffar's Nexus Horn",
		ID:         28418,
		Bonus:      stats.Stats{stats.SpellPower: 225},
		Duration:   time.Second * 10,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskSpellDamage,
		Outcome:    core.OutcomeCrit,
		ProcChance: 0.2,
		ICD:        time.Second * 45,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Eye of Magtheridon",
		ID:       28789,
		Bonus:    stats.Stats{stats.SpellPower: 170},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt | core.CallbackOnPeriodicDamageDealt,
		ProcMask: core.ProcMaskSpellDamage,
		Outcome:  core.OutcomeMiss,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Dragonspine Trophy",
		ID:       28830,
		Bonus:    stats.Stats{stats.MeleeHaste: 325, stats.SpellHaste: 325},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMeleeOrRanged,
		Outcome:  core.OutcomeLanded,
		PPM:      1.0,
		ICD:      time.Second * 20,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Band of the Eternal Champion",
		ID:         29301,
		Bonus:      stats.Stats{stats.AttackPower: 160, stats.RangedAttackPower: 160},
		Duration:   time.Second * 10,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskMeleeOrRanged,
		Outcome:    core.OutcomeLanded,
		ProcChance: 1,
		ICD:        time.Second * 60,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Band of the Eternal Sage",
		ID:         29305,
		Bonus:      stats.Stats{stats.SpellPower: 95},
		Duration:   time.Second * 10,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskSpellDamage,
		Outcome:    core.OutcomeLanded,
		ProcChance: 0.1,
		ICD:        time.Second * 60,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Bladefirst Proc",
		ID:       29348,
		Bonus:    stats.Stats{stats.MeleeHaste: 180},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMeleeMH,
		Outcome:  core.OutcomeLanded,
		PPM:      2.7,
		ICD:      time.Second * 45,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Warp-Spring Coil",
		ID:         30450,
		Bonus:      stats.Stats{stats.ArmorPenetration: 142},
		Duration:   time.Second * 15,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskMelee,
		Outcome:    core.OutcomeHit,
		ProcChance: 0.25,
		ICD:        time.Second * 30,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Sextant of Unstable Currents",
		ID:         30626,
		Bonus:      stats.Stats{stats.SpellPower: 190},
		Duration:   time.Second * 15,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskSpellDamage,
		Outcome:    core.OutcomeCrit,
		ProcChance: 0.2,
		ICD:        time.Second * 45,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Tsunami Talisman Proc",
		ID:         30627,
		Bonus:      stats.Stats{stats.AttackPower: 340, stats.RangedAttackPower: 340},
		Duration:   time.Second * 10,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskMeleeOrRanged,
		Outcome:    core.OutcomeLanded,
		ProcChance: 0.1,
		ICD:        time.Second * 45,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Singing Crystal Axe",
		ID:       31318,
		Bonus:    stats.Stats{stats.MeleeHaste: 400},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMelee,
		Outcome:  core.OutcomeLanded,
		PPM:      3.5,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:       "Shard of Contempt Proc",
		ID:         34472,
		Bonus:      stats.Stats{stats.AttackPower: 42, stats.RangedAttackPower: 42},
		Duration:   time.Second * 20,
		Callback:   core.CallbackOnSpellHitDealt,
		ProcMask:   core.ProcMaskMeleeOrRanged,
		Outcome:    core.OutcomeLanded,
		ProcChance: 0.1,
		ICD:        time.Second * 45,
	})

	newProcStatBonusEffect(ProcStatBonusEffect{
		Name:     "Madness of the Betrayer Proc",
		ID:       32505,
		Bonus:    stats.Stats{stats.ArmorPenetration: 42},
		Duration: time.Second * 10,
		Callback: core.CallbackOnSpellHitDealt,
		ProcMask: core.ProcMaskMeleeOrRanged,
		Outcome:  core.OutcomeLanded,
		PPM:      1.0,
	})
	core.AddEffectsToTest = false
}
