package dps

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
	"github.com/Tereneckla/wotlk70/sim/deathknight"
)

type SharedRotation struct {
	dk         *DpsDeathknight
	recastedFF bool
	recastedBP bool

	ffFirst bool
	hasGod  bool
}

func (sr *SharedRotation) Reset(sim *core.Simulation) {
	sr.recastedFF = false
	sr.recastedBP = false
}

func (sr *SharedRotation) Initialize(dk *DpsDeathknight) {
	dk.sr.ffFirst = dk.Rotation.FirstDisease == proto.Deathknight_Rotation_FrostFever
	dk.sr.hasGod = dk.HasMajorGlyph(proto.DeathknightMajorGlyph_GlyphOfDisease)
}

func (dk *DpsDeathknight) shDiseaseCheck(sim *core.Simulation, target *core.Unit, spell *core.Spell, costRunes bool, casts int, ffSyncTime time.Duration) bool {
	ffRemaining := dk.FrostFeverSpell.Dot(target).RemainingDuration(sim)
	bpRemaining := dk.BloodPlagueSpell.Dot(target).RemainingDuration(sim)
	castGcd := dk.SpellGCD() * time.Duration(casts)

	// FF is not active or will drop before Gcd is ready after this cast
	if !dk.FrostFeverSpell.Dot(target).IsActive() || ffRemaining < castGcd {
		return false
	}
	// BP is not active or will drop before Gcd is ready after this cast
	if !dk.BloodPlagueSpell.Dot(target).IsActive() || bpRemaining < castGcd {
		return false
	}

	// If the ability we want to cast spends runes we check for possible disease drops
	// in the time we won't have runes to recast the disease
	if spell.CanCast(sim, nil) && costRunes {
		ffExpiresAt := ffRemaining + sim.CurrentTime
		bpExpiresAt := bpRemaining + sim.CurrentTime

		crpb := dk.CopyRunicPowerBar()
		spellCost := crpb.OptimalRuneCost(core.RuneCost(spell.DefaultCast.Cost))

		crpb.SpendRuneCost(sim, spell, spellCost)

		afterCastTime := sim.CurrentTime + castGcd
		currentFrostRunes := crpb.CurrentFrostRunes()
		currentUnholyRunes := crpb.CurrentUnholyRunes()
		nextFrostRuneAt := crpb.FrostRuneReadyAt(sim)
		nextUnholyRuneAt := crpb.UnholyRuneReadyAt(sim)

		// If FF is gonna drop while our runes are on CD
		if dk.shRecastAvailableCheck(ffExpiresAt-ffSyncTime, afterCastTime, int(spellCost.Frost()), int32(currentFrostRunes), nextFrostRuneAt) {
			return false
		}

		// If BP is gonna drop while our runes are on CD
		if dk.shRecastAvailableCheck(bpExpiresAt, afterCastTime, int(spellCost.Unholy()), int32(currentUnholyRunes), nextUnholyRuneAt) {
			return false
		}
	}

	return true
}

func (dk *DpsDeathknight) shRecastAvailableCheck(expiresAt time.Duration, afterCastTime time.Duration,
	spellCost int, currentRunes int32, nextRuneAt time.Duration) bool {
	if spellCost > 0 && currentRunes == 0 {
		if expiresAt <= nextRuneAt {
			return true
		}
	} else if afterCastTime >= expiresAt {
		return true
	}
	return false
}

func (dk *DpsDeathknight) shShouldSpreadDisease(sim *core.Simulation) bool {
	return dk.sr.recastedFF && dk.sr.recastedBP && dk.Env.GetNumTargets() > 1
}

func (dk *DpsDeathknight) RotationAction_CancelBT(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	dk.BloodTapAura.Deactivate(sim)
	s.Advance()
	return sim.CurrentTime
}
