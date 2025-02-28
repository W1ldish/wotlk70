package hunter

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
)

func (hp *HunterPet) ApplyTalents() {
	talents := hp.Talents()
	// TODO:
	// Thunderstomp

	hp.AddStat(stats.MeleeCrit, 3*core.CritRatingPerCritChance*float64(talents.SpidersBite))
	hp.AddStat(stats.SpellCrit, 3*core.CritRatingPerCritChance*float64(talents.SpidersBite))
	hp.PseudoStats.DamageDealtMultiplier *= 1 + 0.03*float64(talents.SpikedCollar)
	hp.PseudoStats.DamageDealtMultiplier *= 1 + 0.03*float64(talents.SharkAttack)

	hp.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexArcane] *= 1 - 0.05*float64(talents.GreatResistance)
	hp.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexFire] *= 1 - 0.05*float64(talents.GreatResistance)
	hp.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexFrost] *= 1 - 0.05*float64(talents.GreatResistance)
	hp.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexNature] *= 1 - 0.05*float64(talents.GreatResistance)
	hp.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow] *= 1 - 0.05*float64(talents.GreatResistance)

	if talents.GreatStamina != 0 {
		hp.MultiplyStat(stats.Stamina, 1.0+0.04*float64(talents.GreatStamina))
	}

	if talents.NaturalArmor != 0 {
		hp.MultiplyStat(stats.Armor, 1.0+0.05*float64(talents.NaturalArmor))
	}

	if talents.BloodOfTheRhino != 0 {
		hp.PseudoStats.HealingTakenMultiplier *= 1 + 0.2*float64(talents.BloodOfTheRhino)

		hp.MultiplyStat(stats.Stamina, 1.0+0.02*float64(talents.BloodOfTheRhino))
	}

	if talents.PetBarding != 0 {
		hp.AddStat(stats.Dodge, 1*core.DodgeRatingPerDodgeChance*float64(talents.PetBarding))
		hp.MultiplyStat(stats.Armor, 1.0+0.05*float64(talents.PetBarding))
	}

	hp.applyOwlsFocus()
	hp.applyCullingTheHerd()
	hp.applyFeedingFrenzy()

	hp.registerRoarOfRecoveryCD()
	hp.registerRabidCD()
	hp.registerCallOfTheWildCD()
	hp.registerWolverineBite()
}

func (hp *HunterPet) applyOwlsFocus() {
	if hp.Talents().OwlsFocus == 0 {
		return
	}

	procChance := 0.15 * float64(hp.Talents().OwlsFocus)

	procAura := hp.RegisterAura(core.Aura{
		Label:    "Owl's Focus Proc",
		ActionID: core.ActionID{SpellID: 53515},
		Duration: time.Second * 8,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.CostMultiplier -= 1
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.CostMultiplier += 1
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskSpecial) {
				aura.Deactivate(sim)
			}
		},
	})

	hp.RegisterAura(core.Aura{
		Label:    "Owls Focus",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskSpecial) && sim.RandomFloat("Owls Focus") < procChance {
				procAura.Activate(sim)
			}
		},
	})
}

func (hp *HunterPet) applyCullingTheHerd() {
	if hp.Talents().CullingTheHerd == 0 {
		return
	}

	damageMult := 1 + 0.01*float64(hp.Talents().CullingTheHerd)

	makeProcAura := func(unit *core.Unit) *core.Aura {
		return unit.RegisterAura(core.Aura{
			Label:    "Culling the Herd Proc",
			ActionID: core.ActionID{SpellID: 52858},
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.DamageDealtMultiplier *= damageMult
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.DamageDealtMultiplier /= damageMult
			},
		})
	}
	petAura := makeProcAura(&hp.Unit)
	ownerAura := makeProcAura(&hp.hunterOwner.Unit)

	hp.RegisterAura(core.Aura{
		Label:    "Culling the Herd",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Outcome.Matches(core.OutcomeCrit) && (spell.IsSpellAction(BiteSpellID) || spell.IsSpellAction(ClawSpellID) || spell.IsSpellAction(SmackSpellID)) {
				petAura.Activate(sim)
				ownerAura.Activate(sim)
			}
		},
	})
}

func (hp *HunterPet) applyFeedingFrenzy() {
	if hp.Talents().FeedingFrenzy == 0 {
		return
	}

	multiplier := 1.0 + 0.08*float64(hp.Talents().FeedingFrenzy)

	hp.RegisterResetEffect(func(sim *core.Simulation) {
		sim.RegisterExecutePhaseCallback(func(sim *core.Simulation, isExecute int) {
			if isExecute == 35 {
				hp.PseudoStats.DamageDealtMultiplier *= multiplier
			}
		})
	})
}

func (hp *HunterPet) registerRoarOfRecoveryCD() {
	if !hp.Talents().RoarOfRecovery {
		return
	}

	hunter := hp.hunterOwner
	actionID := core.ActionID{SpellID: 53517}
	manaMetrics := hunter.NewManaMetrics(actionID)

	rorSpell := hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: hunter.applyLongevity(time.Minute * 3),
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			core.StartPeriodicAction(sim, core.PeriodicActionOptions{
				Period:   time.Second * 3,
				NumTicks: 3,
				OnAction: func(sim *core.Simulation) {
					hunter.AddMana(sim, hunter.MaxMana()*0.1, manaMetrics, false)
				},
			})
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		Spell: rorSpell,
		Type:  core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return hp.IsEnabled() && hunter.CurrentManaPercent() < 0.6
		},
	})
}

func (hp *HunterPet) registerRabidCD() {
	if !hp.Talents().Rabid {
		return
	}

	hunter := hp.hunterOwner
	actionID := core.ActionID{SpellID: 53401}
	procChance := 0.2

	statDeps := []*stats.StatDependency{nil}
	for i := 1; i <= 5; i++ {
		statDeps = append(statDeps, hp.NewDynamicMultiplyStat(stats.AttackPower, 1+0.05*float64(i)))
	}

	procAura := hp.RegisterAura(core.Aura{
		Label:     "Rabid Power",
		ActionID:  core.ActionID{SpellID: 53403},
		Duration:  core.NeverExpires,
		MaxStacks: 5,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			if oldStacks != 0 {
				aura.Unit.DisableDynamicStatDep(sim, statDeps[oldStacks])
			}
			if newStacks != 0 {
				aura.Unit.EnableDynamicStatDep(sim, statDeps[newStacks])
			}
		},
	})

	buffAura := hp.RegisterAura(core.Aura{
		Label:    "Rabid",
		ActionID: actionID,
		Duration: time.Second * 20,
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			procAura.Deactivate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() {
				return
			}
			if !spell.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}
			if sim.RandomFloat("Rabid") > procChance {
				return
			}

			procAura.Activate(sim)
			procAura.AddStack(sim)
		},
	})

	rabidSpell := hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: hunter.applyLongevity(time.Second * 45),
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			buffAura.Activate(sim)
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		Spell: rabidSpell,
		Type:  core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return hp.IsEnabled()
		},
	})
}

func (hp *HunterPet) registerCallOfTheWildCD() {
	if !hp.Talents().CallOfTheWild {
		return
	}

	hunter := hp.hunterOwner
	actionID := core.ActionID{SpellID: 53434}

	ownerMAPDep := hunter.NewDynamicMultiplyStat(stats.AttackPower, 1.1)
	ownerRAPDep := hunter.NewDynamicMultiplyStat(stats.RangedAttackPower, 1.1)
	petMAPDep := hp.NewDynamicMultiplyStat(stats.AttackPower, 1.1)
	makeProcAura := func(unit *core.Unit, mapDep *stats.StatDependency, rapDep *stats.StatDependency) *core.Aura {
		return unit.RegisterAura(core.Aura{
			Label:    "Call of the Wild",
			ActionID: actionID,
			Duration: time.Second * 20,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				unit.EnableDynamicStatDep(sim, mapDep)
				if rapDep != nil {
					unit.EnableDynamicStatDep(sim, rapDep)
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				unit.DisableDynamicStatDep(sim, mapDep)
				if rapDep != nil {
					unit.DisableDynamicStatDep(sim, rapDep)
				}
			},
		})
	}
	petAura := makeProcAura(&hp.Unit, petMAPDep, nil)
	ownerAura := makeProcAura(&hp.hunterOwner.Unit, ownerMAPDep, ownerRAPDep)

	cotwSpell := hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: hunter.applyLongevity(time.Minute * 5),
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			petAura.Activate(sim)
			ownerAura.Activate(sim)
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		Spell: cotwSpell,
		Type:  core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return hp.IsEnabled()
		},
	})
}

func (hp *HunterPet) registerWolverineBite() {
	if !hp.Talents().WolverineBite {
		return
	}

	hunter := hp.hunterOwner
	actionID := core.ActionID{SpellID: 53508}

	var wbValidUntil time.Duration
	hp.RegisterAura(core.Aura{
		Label:    "Wolverine Bite Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			wbValidUntil = 0
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.DidCrit() {
				wbValidUntil = sim.CurrentTime + time.Second*5
			}
		},
	})

	wbSpell := hp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagIncludeTargetBonusDamage,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    hp.NewTimer(),
				Duration: hunter.applyLongevity(time.Second * 10),
			},
		},

		DamageMultiplier: 1 * hp.hunterOwner.markedForDeathMultiplier(),
		CritMultiplier:   2,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := 5*80 + 0.07*spell.MeleeAttackPower()
			baseDamage *= hp.killCommandMult()

			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialNoBlockDodgeParry)
			wbValidUntil = 0
		},
	})

	hp.AddMajorCooldown(core.MajorCooldown{
		Spell: wbSpell,
		Type:  core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return hp.IsEnabled() && wbSpell.IsReady(sim) && wbValidUntil > sim.CurrentTime
		},
	})
}
