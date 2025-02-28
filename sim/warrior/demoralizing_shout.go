package warrior

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
)

func (warrior *Warrior) registerDemoralizingShoutSpell() {
	warrior.DemoralizingShoutAuras = make([]*core.Aura, warrior.Env.GetNumTargets())
	for _, target := range warrior.Env.Encounter.Targets {
		warrior.DemoralizingShoutAuras[target.Index] = core.DemoralizingShoutAura(&target.Unit, warrior.Talents.BoomingVoice, warrior.Talents.ImprovedDemoralizingShout)
	}

	warrior.DemoralizingShout = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 25203},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskEmpty,

		RageCost: core.RageCostOptions{
			Cost: 10 - float64(warrior.Talents.FocusedRage),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ThreatMultiplier: 1,
		FlatThreatBonus:  63.2,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for _, aoeTarget := range sim.Encounter.Targets {
				result := spell.CalcAndDealOutcome(sim, &aoeTarget.Unit, spell.OutcomeMagicHit)
				if result.Landed() {
					warrior.DemoralizingShoutAuras[aoeTarget.Index].Activate(sim)
				}
			}
		},
	})
}

func (warrior *Warrior) CanDemoralizingShout(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.DemoralizingShout.DefaultCast.Cost
}

func (warrior *Warrior) ShouldDemoralizingShout(sim *core.Simulation, filler bool, maintainOnly bool) bool {
	if !warrior.CanDemoralizingShout(sim) {
		return false
	}

	if filler {
		return true
	}

	return maintainOnly &&
		warrior.DemoralizingShoutAuras[warrior.CurrentTarget.Index].ShouldRefreshExclusiveEffects(sim, time.Second*2)
}
