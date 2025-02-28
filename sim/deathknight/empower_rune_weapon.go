package deathknight

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
)

func (dk *Deathknight) registerEmpowerRuneWeaponSpell() {
	actionID := core.ActionID{SpellID: 47568}
	cdTimer := dk.NewTimer()
	cd := time.Minute * 5

	dk.EmpowerRuneWeapon = dk.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete,
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: cd,
			},
		},
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			dk.RegenAllRunes(sim)
			dk.AddRunicPower(sim, 25, dk.EmpowerRuneWeapon.RunicPowerMetrics())
		},
	})
}
