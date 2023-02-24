package tbc

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

func init() {
	core.AddEffectsToTest = false
	//// Battlemasters trinkets
	//sharedBattlemasterCooldownID := core.NewCooldownID()
	//addBattlemasterEffect := func(itemID int32) {
	//	core.NewItemEffect(itemID, core.MakeTemporaryStatsOnUseCDRegistration(
	//		"BattlemasterTrinket-"+strconv.Itoa(int(itemID)),
	//		stats.Stats{stats.Health: 1750},
	//		time.Second*15,
	//		core.MajorCooldown{
	//			ActionID:         core.ActionID{ItemID: itemID},
	//			CooldownID:       sharedBattlemasterCooldownID,
	//			Cooldown:         time.Minute * 3,
	//			SharedCooldownID: core.DefensiveTrinketSharedCooldownID,
	//		},
	//	))
	//}
	//addBattlemasterEffect(33832)
	//addBattlemasterEffect(34049)
	//addBattlemasterEffect(34050)
	//addBattlemasterEffect(34162)
	//addBattlemasterEffect(34163)

	// Offensive trinkets. Keep these in order by item ID.
	core.NewSimpleStatOffensiveTrinketEffect(22954, stats.Stats{stats.MeleeHaste: 200}, time.Second*15, time.Minute*2)                                 // Kiss of the Spider
	core.NewSimpleStatOffensiveTrinketEffect(23041, stats.Stats{stats.AttackPower: 260, stats.RangedAttackPower: 260}, time.Second*20, time.Minute*2)  // Slayer's Crest
	core.NewSimpleStatOffensiveTrinketEffect(24128, stats.Stats{stats.AttackPower: 320, stats.RangedAttackPower: 320}, time.Second*12, time.Minute*3)  // Figurine Nightseye Panther
	core.NewSimpleStatOffensiveTrinketEffect(28041, stats.Stats{stats.AttackPower: 200, stats.RangedAttackPower: 200}, time.Second*15, time.Minute*2)  // Bladefists Breadth
	core.NewSimpleStatOffensiveTrinketEffect(28121, stats.Stats{stats.ArmorPenetration: 600}, time.Second*20, time.Minute*2)                           // Icon of Unyielding Courage
	core.NewSimpleStatOffensiveTrinketEffect(28288, stats.Stats{stats.MeleeHaste: 260}, time.Second*10, time.Minute*2)                                 // Abacus of Violent Odds
	core.NewSimpleStatOffensiveTrinketEffect(29383, stats.Stats{stats.AttackPower: 278, stats.RangedAttackPower: 278}, time.Second*20, time.Minute*2)  // Bloodlust Brooch
	core.NewSimpleStatOffensiveTrinketEffect(29776, stats.Stats{stats.AttackPower: 200, stats.RangedAttackPower: 200}, time.Second*20, time.Minute*2)  // Core of Arkelos
	core.NewSimpleStatOffensiveTrinketEffect(32658, stats.Stats{stats.Agility: 150}, time.Second*20, time.Minute*2)                                    // Badge of Tenacity
	core.NewSimpleStatOffensiveTrinketEffect(33831, stats.Stats{stats.AttackPower: 360, stats.RangedAttackPower: 360}, time.Second*20, time.Minute*2)  // Berserkers Call
	core.NewSimpleStatOffensiveTrinketEffect(35702, stats.Stats{stats.AttackPower: 320, stats.RangedAttackPower: 320}, time.Second*15, time.Second*90) // Figurine Shadowsong Panther
	core.NewSimpleStatOffensiveTrinketEffect(38287, stats.Stats{stats.AttackPower: 278, stats.RangedAttackPower: 278}, time.Second*20, time.Minute*2)  // Empty Direbrew Mug

	// Defensive trinkets. Keep these in order by item ID.
	core.NewSimpleStatDefensiveTrinketEffect(27891, stats.Stats{stats.Armor: 1280}, time.Second*20, time.Minute*2)                                                          // Adamantine Figurine
	core.NewSimpleStatDefensiveTrinketEffect(28528, stats.Stats{stats.Dodge: 300}, time.Second*10, time.Minute*2)                                                           // Moroes Lucky Pocket Watch
	core.NewSimpleStatDefensiveTrinketEffect(29387, stats.Stats{stats.BlockValue: 200}, time.Second*20, time.Minute*2)                                                      // Gnomeregan Auto-Blocker 600
	core.NewSimpleStatDefensiveTrinketEffect(30300, stats.Stats{stats.Block: 125}, time.Second*15, time.Second*90)                                                          // Dabiris Enigma
	core.NewSimpleStatDefensiveTrinketEffect(30629, stats.Stats{stats.Defense: 165, stats.AttackPower: -330, stats.RangedAttackPower: -330}, time.Second*15, time.Minute*3) // Scarab of Displacement
	core.NewSimpleStatDefensiveTrinketEffect(32501, stats.Stats{stats.Health: 1750}, time.Second*20, time.Minute*3)                                                         // Shadowmoon Insignia
	core.NewSimpleStatDefensiveTrinketEffect(32534, stats.Stats{stats.Health: 1250}, time.Second*15, time.Minute*5)                                                         // Brooch of the Immortal King
	core.NewSimpleStatDefensiveTrinketEffect(33830, stats.Stats{stats.Armor: 2500}, time.Second*20, time.Minute*2)                                                          // Ancient Aqir Artifact
	core.NewSimpleStatDefensiveTrinketEffect(38289, stats.Stats{stats.BlockValue: 200}, time.Second*20, time.Minute*2)                                                      // Coren's Lucky Coin

	// Proc effects. Keep these in order by item ID.

	core.NewItemEffect(11815, func(agent core.Agent) {
		character := agent.GetCharacter()
		if !character.AutoAttacks.IsEnabled() {
			return
		}

		var handOfJusticeSpell *core.Spell
		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 2,
		}
		procChance := 0.013333

		character.RegisterAura(core.Aura{
			Label:    "Hand of Justice",
			Duration: core.NeverExpires,
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				handOfJusticeSpell = character.GetOrRegisterSpell(core.SpellConfig{
					ActionID:     core.ActionID{ItemID: 11815},
					SpellSchool:  core.SpellSchoolPhysical,
					ProcMask:     core.ProcMaskMeleeMHAuto,
					Flags:        core.SpellFlagMeleeMetrics | core.SpellFlagIncludeTargetBonusDamage | core.SpellFlagNoOnCastComplete,
					ApplyEffects: character.AutoAttacks.MHConfig.ApplyEffects,

					DamageMultiplier: 1,
					CritMultiplier:   character.DefaultMeleeCritMultiplier(),
					ThreatMultiplier: 1,
				})
			},
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				// https://wotlk.wowhead.com/spell=15600/hand-of-justice, proc mask = 20.
				if !result.Landed() || !spell.ProcMask.Matches(core.ProcMaskMelee) {
					return
				}

				if !icd.IsReady(sim) {
					return
				}

				if sim.RandomFloat("HandOfJustice") > procChance {
					return
				}
				icd.Use(sim)

				aura.Unit.AutoAttacks.MaybeReplaceMHSwing(sim, handOfJusticeSpell).Cast(sim, result.Target)
			},
		})
	})

	core.NewItemEffect(32654, func(agent core.Agent) {
		character := agent.GetCharacter()
		character.PseudoStats.BonusDamage += 7
		core.RegisterTemporaryStatsOnUseCD(
			character,
			"Crystalforged Trinket",
			stats.Stats{stats.AttackPower: 216, stats.RangedAttackPower: 216},
			time.Second*10,
			core.SpellConfig{
				ActionID: core.ActionID{ItemID: 32654},
				Cast: core.CastConfig{
					CD: core.Cooldown{
						Timer:    character.NewTimer(),
						Duration: time.Minute,
					},
					SharedCD: core.Cooldown{
						Timer:    character.GetOffensiveTrinketCD(),
						Duration: time.Second * 10,
					},
				},
			},
		)
	})

	core.NewItemEffect(21670, func(agent core.Agent) {
		character := agent.GetCharacter()

		procAura := character.RegisterAura(core.Aura{
			Label:     "Badge of the Swarmguard Proc",
			ActionID:  core.ActionID{SpellID: 26481},
			Duration:  core.NeverExpires,
			MaxStacks: 6,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
				character.AddStatDynamic(sim, stats.ArmorPenetration, 200*float64(newStacks-oldStacks))
			},
		})

		actionID := core.ActionID{ItemID: 21670}
		ppmm := character.AutoAttacks.NewPPMManager(10.0, core.ProcMaskMeleeOrRanged)
		activeAura := character.RegisterAura(core.Aura{
			Label:    "Badge of the Swarmguard",
			ActionID: actionID,
			Duration: time.Second * 30,
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				procAura.Deactivate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() {
					return
				}
				if !spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}

				if !ppmm.Proc(sim, spell.ProcMask, "Badge of the Swarmguard") {
					return
				}

				procAura.Activate(sim)
				procAura.AddStack(sim)
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 3,
				},
			},

			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
				activeAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Spell: spell,
			Type:  core.CooldownTypeDPS,
		})
	})

	core.NewItemEffect(23206, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.CurrentTarget.MobType == proto.MobType_MobTypeDemon || character.CurrentTarget.MobType == proto.MobType_MobTypeUndead {
			character.PseudoStats.MobTypeAttackPower += 150
		}
	})

	core.NewItemEffect(28034, func(agent core.Agent) {
		character := agent.GetCharacter()
		procAura := character.NewTemporaryStatsAura("Rage of the Unraveller", core.ActionID{ItemID: 28034}, stats.Stats{stats.AttackPower: 300, stats.RangedAttackPower: 300}, time.Second*10)
		const procChance = 0.1

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 50,
		}

		character.RegisterAura(core.Aura{
			Label:    "Hourglass of the Unraveller",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Outcome.Matches(core.OutcomeCrit) {
					return
				}
				if !spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !icd.IsReady(sim) {
					return
				}
				if sim.RandomFloat("Hourglass of the Unraveller") > procChance {
					return
				}

				icd.Use(sim)
				procAura.Activate(sim)
			},
		})
	})

	core.NewItemEffect(31857, func(agent core.Agent) {
		character := agent.GetCharacter()

		procAura := character.RegisterAura(core.Aura{
			Label:     "DMC Wrath Proc",
			ActionID:  core.ActionID{ItemID: 31857},
			Duration:  time.Second * 10,
			MaxStacks: 1000,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
				character.AddStatDynamic(sim, stats.MeleeCrit, 17*float64(newStacks-oldStacks))
				character.AddStatDynamic(sim, stats.SpellCrit, 17*float64(newStacks-oldStacks))
			},
		})

		character.RegisterAura(core.Aura{
			Label:    "DMC Wrath",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				// mask 340
				if !spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}

				if result.Outcome.Matches(core.OutcomeCrit) {
					procAura.Deactivate(sim)
				} else {
					procAura.Activate(sim)
					procAura.AddStack(sim)
				}
			},
		})
	})
	core.AddEffectsToTest = true
}
