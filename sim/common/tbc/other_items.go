package tbc

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

func init() {
	core.NewItemEffect(30892, func(agent core.Agent) {
		for _, pet := range agent.GetCharacter().Pets {
			if pet.GetPet().IsGuardian() {
				continue // not sure if this applies to guardians.
			}
			pet.GetCharacter().PseudoStats.DamageDealtMultiplier *= 1.03
			pet.GetCharacter().AddStat(stats.MeleeCrit, core.CritRatingPerCritChance*2)
		}
	})

	core.NewItemEffect(19406, func(agent core.Agent) {
		character := agent.GetCharacter()

		character.AddStats(stats.Stats{
			stats.MeleeHit: 20,
			stats.SpellHit: 20,
		})

	})

	core.NewItemEffect(21709, func(agent core.Agent) {
		character := agent.GetCharacter()

		character.AddStats(stats.Stats{
			stats.MeleeHit: 8,
			stats.SpellHit: 8,
		})

	})

	core.NewItemEffect(33122, func(agent core.Agent) {
		character := agent.GetCharacter()

		character.AddStats(stats.Stats{
			stats.MeleeCrit: 24,
		})

	})

	core.NewItemEffect(34678, func(agent core.Agent) {
		character := agent.GetCharacter()
		const proc = 0.15

		var aldorAura *core.Aura
		var scryerSpell *core.Spell

		if character.ShattFaction == proto.ShattrathFaction_Aldor {
			aldorAura = character.NewTemporaryStatsAura("Light's Wrath", core.ActionID{SpellID: 45479}, stats.Stats{stats.SpellPower: 120}, time.Second*10)
		} else if character.ShattFaction == proto.ShattrathFaction_Scryer {
			scryerSpell = character.RegisterSpell(core.SpellConfig{
				ActionID:       core.ActionID{SpellID: 45429},
				SpellSchool:    core.SpellSchoolArcane,
				ProcMask:       core.ProcMaskEmpty,
				CritMultiplier: 1,
				ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
					spell.CalcAndDealDamage(sim, target, sim.Roll(333, 367), spell.OutcomeMagicHitAndCrit)
				},
			})
		}

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 45,
		}

		character.RegisterAura(core.Aura{
			Label:    "Shattered Sun Pendant of Acumen",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !spell.ProcMask.Matches(core.ProcMaskSpellDamage) {
					return
				}
				if !result.Landed() {
					return
				}
				if !icd.IsReady(sim) || sim.RandomFloat("Pendant of Acumen") > proc {
					return
				}
				icd.Use(sim)

				if character.ShattFaction == proto.ShattrathFaction_Aldor {
					aldorAura.Activate(sim)
				} else if character.ShattFaction == proto.ShattrathFaction_Scryer {
					scryerSpell.Cast(sim, result.Target)
				}

			},
		})
	})

	core.NewItemEffect(34679, func(agent core.Agent) {
		character := agent.GetCharacter()
		const proc = 0.15

		var aldorAura *core.Aura
		var scryerSpell *core.Spell

		if character.ShattFaction == proto.ShattrathFaction_Aldor {
			aldorAura = character.NewTemporaryStatsAura("Light's Strength", core.ActionID{SpellID: 45480}, stats.Stats{stats.AttackPower: 200}, time.Second*10)
		} else if character.ShattFaction == proto.ShattrathFaction_Scryer {
			scryerSpell = character.RegisterSpell(core.SpellConfig{
				ActionID:       core.ActionID{SpellID: 45428},
				SpellSchool:    core.SpellSchoolArcane,
				ProcMask:       core.ProcMaskEmpty,
				CritMultiplier: 1,
				ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
					spell.CalcAndDealDamage(sim, target, sim.Roll(333, 367), spell.OutcomeMeleeSpecialHitAndCrit)
				},
			})
		}

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 45,
		}

		character.RegisterAura(core.Aura{
			Label:    "Shattered Sun Pendant of Might",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !result.Landed() {
					return
				}
				if !icd.IsReady(sim) || sim.RandomFloat("Pendant of Might") > proc {
					return
				}
				icd.Use(sim)

				if character.ShattFaction == proto.ShattrathFaction_Aldor {
					aldorAura.Activate(sim)
				} else if character.ShattFaction == proto.ShattrathFaction_Scryer {
					scryerSpell.Cast(sim, result.Target)
				}

			},
		})
	})

}
