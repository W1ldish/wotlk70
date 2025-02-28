package tank

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
	"github.com/Tereneckla/wotlk70/sim/deathknight"
)

func (dk *TankDeathknight) TankRA_Tps(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	if !dk.GCD.IsReady(sim) {
		return dk.NextGCDAt()
	}

	if dk.DoDefensiveCds(sim, target, s) {
		return -1
	}

	t := sim.CurrentTime
	ff := dk.FrostFeverSpell.Dot(target).ExpiresAt() - t
	bp := dk.BloodPlagueSpell.Dot(target).ExpiresAt() - t

	if ff <= 0 && dk.IcyTouch.CanCast(sim, target) {
		dk.IcyTouch.Cast(sim, target)
		return -1
	}

	if bp <= 0 && dk.PlagueStrike.CanCast(sim, target) {
		dk.PlagueStrike.Cast(sim, target)
		return -1
	}

	if ff <= 2*time.Second || bp <= 2*time.Second && dk.Pestilence.CanCast(sim, target) {
		dk.Pestilence.Cast(sim, target)
		return -1
	}

	if dk.switchIT && dk.IcyTouch.CanCast(sim, target) {
		dk.IcyTouch.Cast(sim, target)

		if dk.DeathRunesInFU() == 0 {
			dk.switchIT = false
		}

		return -1
	}

	if !dk.switchIT && dk.FuSpell.CanCast(sim, target) {
		dk.FuSpell.Cast(sim, target)

		if dk.DeathRunesInFU() == 4 {
			dk.switchIT = true
		}

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
