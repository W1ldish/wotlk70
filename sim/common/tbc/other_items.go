package tbc

import (
	"github.com/Tereneckla/wotlk/sim/core"
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
}
