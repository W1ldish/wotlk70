package dps

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/deathknight"
)

// end of fight oblit does not check diseases, it just presses it regardless, but will retry if fails to land.
func (dk *DpsDeathknight) RotationActionCallback_FrostSubUnh_EndOfFight_Obli(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	casted := false
	advance := true
	waitTime := time.Duration(-1)
	if dk.Obliterate.CanCast(sim, nil) {
		if dk.Deathchill != nil && dk.Deathchill.IsReady(sim) {
			dk.Deathchill.Cast(sim, target)
		}
		casted = dk.Obliterate.Cast(sim, target)
		advance = dk.LastOutcome.Matches(core.OutcomeLanded)
	}
	s.ConditionalAdvance(casted && advance)
	return core.TernaryDuration(casted, -1, waitTime)
}

func (dk *DpsDeathknight) RegularPrioPickSpell(sim *core.Simulation, target *core.Unit, untilTime time.Duration) *core.Spell {
	abGcd := 1500 * time.Millisecond
	spGcd := dk.SpellGCD()
	canCastAbility := sim.CurrentTime+abGcd <= untilTime
	canCastSpell := sim.CurrentTime+spGcd <= untilTime

	km := dk.KillingMachineAura.IsActive()
	rime := dk.RimeAura.IsActive()
	if canCastSpell && dk.RaiseDead.CanCast(sim, nil) && sim.GetRemainingDuration() >= time.Second*30 {
		return dk.RaiseDead
	} else if canCastSpell && dk.HowlingBlast.CanCast(sim, nil) && rime {
		return dk.HowlingBlast
	} else if canCastAbility && dk.FrostStrike.CanCast(sim, nil) && km {
		return dk.FrostStrike
	} else if canCastAbility && dk.FrostStrike.CanCast(sim, nil) && dk.CurrentRunicPower() >= 100.0 {
		return dk.FrostStrike
	} else if canCastAbility && dk.FrostStrike.CanCast(sim, nil) {
		return dk.FrostStrike
	} else if canCastSpell && dk.HornOfWinter.CanCast(sim, nil) {
		return dk.HornOfWinter
	} else {
		return nil
	}
}

//end of fight functions

func (dk *DpsDeathknight) RotationActionCallback_EndOfFightCheck(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	simDur := sim.CurrentTime + sim.GetRemainingDuration()

	if sim.CurrentTime+7000*time.Millisecond > simDur {
		s.Clear().NewAction(dk.RotationActionCallback_EndOfFightPrio)
	} else {
		s.Advance()
	}
	return sim.CurrentTime
}

func (dk *DpsDeathknight) RotationActionCallback_EndOfFightPrio(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {

	simDur := sim.CurrentTime + sim.GetRemainingDuration()
	simTimeLeft := sim.GetRemainingDuration()
	ffExpiresAt := dk.FrostFeverSpell.Dot(target).ExpiresAt()
	bpExpiresAt := dk.BloodPlagueSpell.Dot(target).ExpiresAt()
	diseaseExpiresAt := core.MinDuration(ffExpiresAt, bpExpiresAt)
	abGcd := 1500 * time.Millisecond
	spGcd := dk.SpellGCD()
	frAt := dk.NormalFrostRuneReadyAt(sim)
	uhAt := dk.NormalUnholyRuneReadyAt(sim)
	obAt := core.MaxDuration(frAt, uhAt)
	fsCost := float64(core.RuneCost(dk.FrostStrike.CurCast.Cost).RunicPower())
	delayAmount := core.MinDuration(time.Duration(dk.Rotation.OblitDelayDuration)*time.Millisecond, 2501*time.Millisecond)
	bothblAt := dk.BloodDeathRuneBothReadyAt()

	if bothblAt == core.NeverExpires {
		bothblAt = 1
	}
	if bothblAt == -1 {
		bothblAt = core.NeverExpires
	}

	if dk.Talents.Epidemic == 2 || diseaseExpiresAt >= simDur {
		obAt = core.MinDuration(obAt, bothblAt)
	}

	if diseaseExpiresAt >= simDur { //diseases last until end of fight
		if sim.CurrentTime >= obAt { //have runes to oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if spGcd > simTimeLeft && (obAt <= simDur || bothblAt <= simDur) { //cant fit a spell GCD before end of fight, and will be able to oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if spGcd < simTimeLeft && abGcd > simTimeLeft && obAt < simDur && dk.RimeAura.IsActive() { //can fit a spell but not melee ability before last GCD of fight and have rime
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (abGcd > simTimeLeft || sim.CurrentTime+abGcd > obAt+delayAmount) && obAt < simDur { //oblit can be used, it's the last GCD or it goes over oblit delay
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentRunicPower() >= fsCost && sim.CurrentTime+abGcd < obAt+delayAmount { //can FS and wont cross oblit delay
			s.Clear().
				NewAction(dk.RotationActionCallback_FS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.RimeAura.IsActive() && sim.CurrentTime+spGcd < obAt+delayAmount { //can rime and wont cross oblit delay
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() == 1) && (sim.CurrentTime+abGcd < obAt+delayAmount) { //have runes for BS, and it cant be used for oblit instead
			s.Clear().
				NewAction(dk.RotationActionCallback_BS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) && sim.CurrentTime+spGcd < obAt+delayAmount && simTimeLeft > spGcd { //can horn and wont cross oblit delay
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else { //cant do anything, wait to oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if diseaseExpiresAt >= simDur-abGcd { //disease expires less than 1 gcd before end of fight
		if sim.CurrentTime >= obAt || (sim.CurrentTime >= bothblAt && dk.CurrentDeathRunes() >= 2) { //have runes to oblit, will not refresh diseases
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if spGcd > simTimeLeft && obAt < simDur { //cant fit a spell GCD before end of fight, and will be able to oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if spGcd < simTimeLeft && abGcd > simTimeLeft && obAt < simDur && dk.RimeAura.IsActive() { //can fit a spell but not melee ability before last GCD of fight and have rime
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if abGcd > simTimeLeft && obAt < simDur { //oblit can be used, it's the last GCD or it goes over oblit delay
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentRunicPower() >= fsCost && sim.CurrentTime+abGcd < obAt+delayAmount { //can FS and wont cross oblit delay
			s.Clear().
				NewAction(dk.RotationActionCallback_FS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.RimeAura.IsActive() && sim.CurrentTime+spGcd < obAt+delayAmount { //can rime and wont cross oblit delay
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentBloodRunes() >= 1 || sim.CurrentTime < diseaseExpiresAt { //can pesti and diseases are still up, this serves to pesti if there's nothing more valuable to press
			s.Clear().
				NewAction(dk.RotationActionCallback_Pesti).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) && simTimeLeft > spGcd { //can't do anything but horn
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if sim.CurrentTime+spGcd < diseaseExpiresAt && sim.CurrentTime+abGcd > diseaseExpiresAt && dk.CurrentRunicPower() < 100 && (dk.RimeAura.IsActive() || dk.HornOfWinter.IsReady(sim)) { //if you can fit a spellgcd before disease dropping, and only if spell hit cap is reached
		if dk.RimeAura.IsActive() && dk.CurrentRunicPower() < 100 { //rime prio
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) && dk.CurrentRunicPower() < 100 { //fit a horn in if you can
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if sim.CurrentTime+2*abGcd < diseaseExpiresAt { //there's at least 2 physical GCDs until disease fall so press normal prio, but do not double death oblit as you need it for pesti
		if sim.CurrentTime >= obAt && (dk.Talents.Epidemic == 2 || (dk.CurrentFrostRunes() >= 1 && dk.CurrentUnholyRunes() >= 1)) { //can oblit, either be unh sub or have frost/unh runes ready
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd > obAt+delayAmount && (dk.Talents.Epidemic == 2 || (dk.CurrentFrostRunes() >= 1 && dk.CurrentUnholyRunes() >= 1)) { //same as above, no time to spellGCD before oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd < obAt+delayAmount && sim.CurrentTime+abGcd > obAt && (dk.RimeAura.IsActive() || dk.CurrentRunicPower() < fsCost*4-2*dk.fr.oblitRPRegen && dk.HornOfWinter.IsReady(sim)) { //if you can fit a spGcd before oblit and won't overcap RP with horn
			if dk.RimeAura.IsActive() {
				s.Clear().
					NewAction(dk.RotationActionCallback_HB).
					NewAction(dk.RotationActionCallback_EndOfFightCheck)
			} else if dk.CurrentRunicPower() < fsCost*4-2*dk.fr.oblitRPRegen && dk.HornOfWinter.IsReady(sim) { //if u wont overflow RP
				s.Clear().
					NewAction(dk.RotationActionCallback_HW).
					NewAction(dk.RotationActionCallback_EndOfFightCheck)
			}
		} else if sim.CurrentTime+abGcd > obAt+delayAmount { //if no time to abGcd before oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentRunicPower() >= fsCost { //fs if can
			s.Clear().
				NewAction(dk.RotationActionCallback_FS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.RimeAura.IsActive() { //rime if can
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() == 1) && diseaseExpiresAt > simDur-abGcd { //if can BS and there's only 1 gcd left
			s.Clear().
				NewAction(dk.RotationActionCallback_BS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() == 1) && diseaseExpiresAt < simDur-abGcd { //if can pesti and there's more than 1 gcd left
			s.Clear().
				NewAction(dk.RotationActionCallback_Pesti).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) && simTimeLeft > spGcd { //if can horn
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else { //wait for oblit
			dk.WaitUntil(sim, obAt)
			s.NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if sim.CurrentTime+2*abGcd > diseaseExpiresAt { //if u can only fit 1 spGcd + 1 abGcd before disease falls, do pesti first as it might miss
		s.Clear().
			NewAction(dk.RotationActionCallback_Pesti).
			NewAction(dk.RotationActionCallback_EndOfFightCheck)
	} else {
		return -1
	}
	return sim.CurrentTime
}
