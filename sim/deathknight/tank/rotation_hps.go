package tank

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
	"github.com/Tereneckla/wotlk70/sim/deathknight"
)

func (dk *TankDeathknight) TankRA_Hps(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	if !dk.GCD.IsReady(sim) {
		return dk.NextGCDAt()
	}

	if dk.DoDefensiveCds(sim, target, s) {
		return -1
	}

	if dk.DoDiseaseChecks(sim, target, s) {
		return -1
	}

	fd := dk.CurrentFrostRunes() + dk.CurrentDeathRunes()
	ud := dk.CurrentUnholyRunes() + dk.CurrentDeathRunes()

	if fd > 0 && ud > 0 && dk.FuSpell.CanCast(sim, target) && (dk.CurrentHealthPercent() < 0.75 || dk.FuSpell == dk.Obliterate) {
		dk.FuSpell.Cast(sim, target)
		return -1
	}

	if dk.Rotation.BloodTapPrio == proto.TankDeathknight_Rotation_Offensive {
		if dk.BloodTap.CanCast(sim, target) {
			dk.BloodTap.Cast(sim, target)
			dk.IcyTouch.Cast(sim, target)
			dk.CancelBloodTap(sim)
			return -1
		}
	}

	if dk.DoFrostCast(sim, target, s) {
		return -1
	}

	if dk.DoBloodCast(sim, target, s) {
		return -1
	}

	return -1
}
