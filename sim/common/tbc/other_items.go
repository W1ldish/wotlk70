package tbc

import (
	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
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
}
