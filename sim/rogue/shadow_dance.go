package rogue

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
)

func (rogue *Rogue) registerShadowDanceCD() {
	if !rogue.Talents.ShadowDance {
		return
	}

	actionID := core.ActionID{SpellID: 51713}

	rogue.ShadowDanceAura = rogue.RegisterAura(core.Aura{
		Label:    "Shadow Dance",
		ActionID: actionID,
		Duration: time.Second * 6,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			// can now cast opening abilities outside of stealth
		},
	})

	rogue.ShadowDance = rogue.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: time.Minute,
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			rogue.ShadowDanceAura.Activate(sim)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell:    rogue.ShadowDance,
		Type:     core.CooldownTypeDPS,
		Priority: core.CooldownPriorityDefault,
		ShouldActivate: func(s *core.Simulation, c *core.Character) bool {
			return rogue.CurrentEnergy() > 90
		},
	})
}
