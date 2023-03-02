package warlock

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
)

type ProcTracker struct {
	id          int32
	aura        *core.Aura
	didActivate bool
	isActive    bool
	expiresAt   time.Duration
}

func (warlock *Warlock) addProc(id int32, label string, isActive bool) bool {
	if !warlock.HasAura(label) {
		return false
	}
	warlock.procTrackers = append(warlock.procTrackers, &ProcTracker{
		id:          id,
		didActivate: false,
		isActive:    isActive,
		expiresAt:   -1,
		aura:        warlock.GetAura(label),
	})
	return true
}

func (warlock *Warlock) resetProcTrackers() {
	for _, procTracker := range warlock.procTrackers {
		procTracker.didActivate = false
		procTracker.expiresAt = -1
	}
}

func (warlock *Warlock) initProcTrackers() {
	warlock.procTrackers = make([]*ProcTracker, 0)

	warlock.addProc(40211, "Potion of Speed", true)
	warlock.addProc(47197, "Eradication", true)
	warlock.addProc(10060, "Power Infusion", true)
	warlock.addProc(54999, "Hyperspeed Acceleration", true)
	warlock.addProc(26297, "Berserking (Troll)", true)
	warlock.addProc(33697, "Blood Fury", true)
}

func (warlock *Warlock) setupDSCooldowns() {
	warlock.majorCds = make([]*core.MajorCooldown, 0)

	// berserking (troll)
	warlock.DSCooldownSync(core.ActionID{SpellID: 26297}, false)

	// blood fury (orc)
	warlock.DSCooldownSync(core.ActionID{SpellID: 33697}, false)

	// Power Infusion
	warlock.DSCooldownSync(core.ActionID{SpellID: 10060}, false)
}

func (warlock *Warlock) DSCooldownSync(actionID core.ActionID, isPotion bool) {
	if majorCd := warlock.Character.GetMajorCooldown(actionID); majorCd != nil {
		warlock.majorCds = append(warlock.majorCds, majorCd)
	}
}

func logMessage(sim *core.Simulation, message string) {
	if sim.Log != nil {
		sim.Log(message)
	}
}

func (warlock *Warlock) DSProcCheck(sim *core.Simulation, castTime time.Duration) bool {
	for _, procTracker := range warlock.procTrackers {
		if !procTracker.didActivate && procTracker.aura.IsActive() {
			procTracker.didActivate = true
			procTracker.expiresAt = procTracker.aura.ExpiresAt()
		}

		// A proc is about to drop
		if procTracker.didActivate && procTracker.expiresAt <= sim.CurrentTime+castTime {
			logMessage(sim, "Proc dropping "+procTracker.aura.Label)
			return false
		}
	}

	for _, procTracker := range warlock.procTrackers {
		if !procTracker.didActivate && !procTracker.isActive {
			logMessage(sim, "Waiting on procs..")
			return true
		}
	}

	return false
}
