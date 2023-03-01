package core

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core/proto"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

// Registers all consume-related effects to the Agent.
func applyConsumeEffects(agent Agent) {
	character := agent.GetCharacter()
	consumes := character.Consumes
	if consumes == nil {
		return
	}

	if consumes.Flask != proto.Flask_FlaskUnknown {
		switch consumes.Flask {
		case proto.Flask_LesserFlaskOfResistance:
			character.AddStats(stats.Stats{
				stats.ArcaneResistance: 50,
				stats.FireResistance:   50,
				stats.FrostResistance:  50,
				stats.NatureResistance: 50,
				stats.ShadowResistance: 50,
			})
		case proto.Flask_FlaskOfBlindingLight:
			character.OnSpellRegistered(func(spell *Spell) {
				if spell.SpellSchool.Matches(SpellSchoolArcane | SpellSchoolHoly | SpellSchoolNature) {
					spell.BonusSpellPower += 80
					if character.HasProfession(proto.Profession_Alchemy) {
						spell.BonusSpellPower += 24
					}
				}
			})
		case proto.Flask_FlaskOfMightyRestoration:
			character.AddStats(stats.Stats{
				stats.MP5: 25,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.MP5: 15,
				})
			}
		case proto.Flask_FlaskOfPureDeath:
			character.OnSpellRegistered(func(spell *Spell) {
				if spell.SpellSchool.Matches(SpellSchoolFire | SpellSchoolFrost | SpellSchoolShadow) {
					spell.BonusSpellPower += 80
					if character.HasProfession(proto.Profession_Alchemy) {
						spell.BonusSpellPower += 24
					}
				}
			})
		case proto.Flask_FlaskOfRelentlessAssault:
			character.AddStats(stats.Stats{
				stats.AttackPower:       120,
				stats.RangedAttackPower: 120,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.AttackPower:       36,
					stats.RangedAttackPower: 36,
				})
			}

		case proto.Flask_FlaskOfSupremePower:
			character.AddStats(stats.Stats{
				stats.SpellPower: 70,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.SpellPower: 21,
				})
			}
		case proto.Flask_FlaskOfFortification:
			character.AddStats(stats.Stats{
				stats.Health:  500,
				stats.Defense: 10,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Health:  150,
					stats.Defense: 3,
				})
			}
		case proto.Flask_FlaskOfChromaticWonder:
			character.AddStats(stats.Stats{
				stats.Stamina:          18,
				stats.Strength:         18,
				stats.Agility:          18,
				stats.Intellect:        18,
				stats.Spirit:           18,
				stats.ArcaneResistance: 35,
				stats.FireResistance:   35,
				stats.FrostResistance:  35,
				stats.NatureResistance: 35,
				stats.ShadowResistance: 35,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Stamina:          5,
					stats.Strength:         5,
					stats.Agility:          5,
					stats.Intellect:        5,
					stats.Spirit:           5,
					stats.ArcaneResistance: 10,
					stats.FireResistance:   10,
					stats.FrostResistance:  10,
					stats.NatureResistance: 10,
					stats.ShadowResistance: 10,
				})
			}
		case proto.Flask_FlaskOfDistilledWisdom:
			character.AddStats(stats.Stats{
				stats.Intellect: 65,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Intellect: 19,
				})
			}
		}
	} else {
		switch consumes.BattleElixir {
		case proto.BattleElixir_AdeptsElixir:
			character.AddStats(stats.Stats{
				stats.SpellCrit:  24,
				stats.SpellPower: 24,
				stats.MeleeCrit:  24,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.SpellCrit:  7,
					stats.SpellPower: 7,
					stats.MeleeCrit:  7,
				})
			}
		case proto.BattleElixir_ElixirOfHealingPower:
			character.AddStats(stats.Stats{
				stats.SpellPower: 24,
				stats.Spirit:     24,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.SpellPower: 7,
					stats.Spirit:     7,
				})
			}
		case proto.BattleElixir_ElixirOfDemonslaying:
			if character.CurrentTarget.MobType == proto.MobType_MobTypeDemon {
				character.PseudoStats.MobTypeAttackPower += 105
			}
		case proto.BattleElixir_ElixirOfMajorAgility:
			character.AddStats(stats.Stats{
				stats.Agility:   30,
				stats.MeleeCrit: 12,
				stats.SpellCrit: 12,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Agility:   9,
					stats.MeleeCrit: 3,
					stats.SpellCrit: 3,
				})
			}
		case proto.BattleElixir_ElixirOfMajorStrength:
			character.AddStats(stats.Stats{
				stats.Strength: 35,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Strength: 45,
				})
			}
		case proto.BattleElixir_ElixirOfMastery:
			character.AddStats(stats.Stats{
				stats.Stamina:   15,
				stats.Strength:  15,
				stats.Agility:   15,
				stats.Intellect: 15,
				stats.Spirit:    15,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Stamina:   4,
					stats.Strength:  4,
					stats.Agility:   4,
					stats.Intellect: 4,
					stats.Spirit:    4,
				})
			}
		case proto.BattleElixir_ElixirOfTheMongoose:
			character.AddStats(stats.Stats{
				stats.Agility:   25,
				stats.MeleeCrit: 10,
				stats.SpellCrit: 10,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Agility:   7,
					stats.MeleeCrit: 3,
					stats.SpellCrit: 3,
				})
			}
		case proto.BattleElixir_FelStrengthElixir:
			character.AddStats(stats.Stats{
				stats.AttackPower:       90,
				stats.RangedAttackPower: 90,
				stats.Stamina:           -10,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.AttackPower:       27,
					stats.RangedAttackPower: 27,
					stats.Stamina:           -3,
				})
			}
		case proto.BattleElixir_GreaterArcaneElixir:
			character.AddStats(stats.Stats{
				stats.SpellPower: 35,
				stats.SpellCrit:  10,
				stats.MeleeCrit:  10,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.SpellPower: 10,
					stats.SpellCrit:  3,
					stats.MeleeCrit:  3,
				})
			}
		}

		switch consumes.GuardianElixir {
		case proto.GuardianElixir_ElixirOfDraenicWisdom:
			character.AddStats(stats.Stats{
				stats.Intellect: 30,
				stats.Spirit:    30,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Intellect: 9,
					stats.Spirit:    9,
				})
			}
		case proto.GuardianElixir_ElixirOfIronskin:
			character.AddStats(stats.Stats{
				stats.Resilience: 30,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Resilience: 9,
				})
			}
		case proto.GuardianElixir_ElixirOfMajorDefense:
			character.AddStats(stats.Stats{
				stats.Armor: 550,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Armor: 165,
				})
			}
		case proto.GuardianElixir_ElixirOfMajorFortitude:
			character.AddStats(stats.Stats{
				stats.Health: 250,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.Health: 75,
				})
			}
		case proto.GuardianElixir_ElixirOfMajorMageblood:
			character.AddStats(stats.Stats{
				stats.MP5: 20,
			})
			if character.HasProfession(proto.Profession_Alchemy) {
				character.AddStats(stats.Stats{
					stats.MP5: 6,
				})
			}
		case proto.GuardianElixir_GiftOfArthas:
			character.AddStats(stats.Stats{
				stats.ShadowResistance: 10,
			})

			var debuffAuras []*Aura
			for _, target := range character.Env.Encounter.Targets {
				debuffAuras = append(debuffAuras, GiftOfArthasAura(&target.Unit))
			}

			actionID := ActionID{SpellID: 11374}
			goaProc := character.RegisterSpell(SpellConfig{
				ActionID:    actionID,
				SpellSchool: SpellSchoolNature,
				ProcMask:    ProcMaskEmpty,

				ThreatMultiplier: 1,
				FlatThreatBonus:  90,

				ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
					debuffAuras[target.Index].Activate(sim)
					spell.CalcAndDealOutcome(sim, target, spell.OutcomeAlwaysHit)
				},
			})

			character.RegisterAura(Aura{
				Label:    "Gift of Arthas",
				Duration: NeverExpires,
				OnReset: func(aura *Aura, sim *Simulation) {
					aura.Activate(sim)
				},
				OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
					if result.Landed() &&
						spell.SpellSchool == SpellSchoolPhysical &&
						sim.RandomFloat("Gift of Arthas") < 0.3 {
						goaProc.Cast(sim, spell.Unit)
					}
				},
			})
		}
	}

	switch consumes.Food {
	case proto.Food_FoodGrilledMudfish:
		character.AddStats(stats.Stats{
			stats.Agility: 20,
			stats.Spirit:  20,
		})
	case proto.Food_FoodRavagerDog:
		character.AddStats(stats.Stats{
			stats.AttackPower:       40,
			stats.RangedAttackPower: 40,
			stats.Spirit:            20,
		})
	case proto.Food_FoodRoastedClefthoof:
		character.AddStats(stats.Stats{
			stats.Strength: 20,
			stats.Spirit:   20,
		})
	case proto.Food_FoodSkullfishSoup:
		character.AddStats(stats.Stats{
			stats.SpellCrit: 20,
			stats.Spirit:    20,
		})
	case proto.Food_FoodSpicyHotTalbuk:
		character.AddStats(stats.Stats{
			stats.MeleeHit: 20,
			stats.Spirit:   20,
		})
	case proto.Food_FoodFishermansFeast:
		character.AddStats(stats.Stats{
			stats.Stamina: 30,
			stats.Spirit:  20,
		})
	}

	switch consumes.WeaponMain {
	case proto.WeaponEnchant_EnchantAdamantiteSharpeningStone:
		character.PseudoStats.BonusDamage += 12
		if character.Class != proto.Class_ClassHunter {
			character.AddStats(stats.Stats{
				stats.MeleeCrit: 14,
			})
		}

	case proto.WeaponEnchant_EnchantAdamantiteWeightStone:
		character.PseudoStats.BonusDamage += 12
		character.AddStats(stats.Stats{
			stats.MeleeCrit: 14,
		})

	case proto.WeaponEnchant_EnchantElementalSharpeningStone:
		character.AddStat(stats.MeleeCrit, 28)

	case proto.WeaponEnchant_EnchantBrilliantManaOil:
		character.AddStats(stats.Stats{
			stats.MP5:        12,
			stats.SpellPower: 13,
		})

	case proto.WeaponEnchant_EnchantBrilliantWizardOil:
		character.AddStats(stats.Stats{
			stats.SpellCrit:  14,
			stats.MeleeCrit:  14,
			stats.SpellPower: 36,
		})

	case proto.WeaponEnchant_EnchantSuperiorWizardOil:
		character.AddStat(stats.SpellPower, 42)

	case proto.WeaponEnchant_EnchantSuperiorManaOil:
		character.AddStat(stats.MP5, 14)
	}

	registerPotionCD(agent, consumes)
	registerConjuredCD(agent, consumes)
	registerExplosivesCD(agent, consumes)
}

func ApplyPetConsumeEffects(pet *Character, ownerConsumes *proto.Consumes) {
	switch ownerConsumes.PetFood {
	case proto.PetFood_PetFoodKiblersBits:
		pet.AddStats(stats.Stats{
			stats.Strength: 20,
			stats.Stamina:  20,
		})
	}

	pet.AddStat(stats.Agility, []float64{0, 5, 9, 13, 17, 20}[ownerConsumes.PetScrollOfAgility])
	pet.AddStat(stats.Strength, []float64{0, 5, 9, 13, 17, 20}[ownerConsumes.PetScrollOfStrength])
}

var PotionAuraTag = "Potion"

func registerPotionCD(agent Agent, consumes *proto.Consumes) {
	character := agent.GetCharacter()
	defaultPotion := consumes.DefaultPotion
	startingPotion := consumes.PrepopPotion

	if defaultPotion == proto.Potions_UnknownPotion && startingPotion == proto.Potions_UnknownPotion {
		return
	}

	potionCD := character.NewTimer()

	startingMCD := makePotionActivation(startingPotion, character, potionCD)
	if startingMCD.Spell != nil {
		character.RegisterPrepullAction(-1*time.Second, func(sim *Simulation) {
			startingMCD.Spell.Cast(sim, nil)
			if startingPotion == proto.Potions_IndestructiblePotion {
				potionCD.Set(sim.CurrentTime + 2*time.Minute)
			} else {
				potionCD.Set(sim.CurrentTime + time.Minute)
			}
			character.UpdateMajorCooldowns()
		})
	}

	defaultMCD := makePotionActivation(defaultPotion, character, potionCD)
	if defaultMCD.Spell != nil {
		character.AddMajorCooldown(defaultMCD)
	}
}

var AlchStoneItemIDs = []int32{44322, 44323, 44324}

func (character *Character) HasAlchStone() bool {
	alchStoneEquipped := false
	for _, itemID := range AlchStoneItemIDs {
		alchStoneEquipped = alchStoneEquipped || character.HasTrinketEquipped(itemID)
	}
	return character.HasProfession(proto.Profession_Alchemy) && alchStoneEquipped
}

func makePotionActivation(potionType proto.Potions, character *Character, potionCD *Timer) MajorCooldown {
	alchStoneEquipped := character.HasAlchStone()

	potionCast := CastConfig{
		CD: Cooldown{
			Timer:    potionCD,
			Duration: time.Minute * 60, // Infinite CD
		},
	}

	if potionType == proto.Potions_SuperManaPotion {
		actionID := ActionID{ItemID: 22832}
		manaMetrics := character.NewManaMetrics(actionID)
		return MajorCooldown{
			Type: CooldownTypeMana,
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
				totalRegen := character.ManaRegenPerSecondWhileCasting() * 5
				manaGain := 3000.0
				if alchStoneEquipped && potionType == proto.Potions_RunicManaPotion {
					manaGain *= 1.4
				}
				return character.MaxMana()-(character.CurrentMana()+totalRegen) >= manaGain
			},
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					manaGain := sim.RollWithLabel(1800, 3000, "RunicManaPotion")
					if alchStoneEquipped && potionType == proto.Potions_RunicManaPotion {
						manaGain *= 1.4
					}
					character.AddMana(sim, manaGain, manaMetrics)
				},
			}),
		}
	} else if potionType == proto.Potions_DestructionPotion {
		actionID := ActionID{ItemID: 22839}
		aura := character.NewTemporaryStatsAura("Destruction Potion", actionID, stats.Stats{stats.SpellPower: 120, stats.SpellCrit: 2 * CritRatingPerCritChance}, time.Second*15)
		return MajorCooldown{
			Type: CooldownTypeDPS,
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					aura.Activate(sim)
				},
			}),
		}
	} else if potionType == proto.Potions_SuperManaPotion {
		alchStoneEquipped := character.HasAlchStone()
		actionID := ActionID{ItemID: 22832}
		manaMetrics := character.NewManaMetrics(actionID)
		return MajorCooldown{
			Type: CooldownTypeMana,
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
				totalRegen := character.ManaRegenPerSecondWhileCasting() * 5
				return character.MaxMana()-(character.CurrentMana()+totalRegen) >= 3000
			},
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					// Restores 1800 to 3000 mana. (2 Min Cooldown)
					manaGain := sim.RollWithLabel(1800, 3000, "super mana")
					if alchStoneEquipped {
						manaGain *= 1.4
					}
					character.AddMana(sim, manaGain, manaMetrics)
				},
			}),
		}
	} else if potionType == proto.Potions_HastePotion {
		actionID := ActionID{ItemID: 22838}
		aura := character.NewTemporaryStatsAura("Haste Potion", actionID, stats.Stats{stats.MeleeHaste: 400}, time.Second*15)
		return MajorCooldown{
			Type: CooldownTypeDPS,
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					aura.Activate(sim)
				},
			}),
		}
	} else if potionType == proto.Potions_MightyRagePotion {
		actionID := ActionID{ItemID: 13442}
		aura := character.NewTemporaryStatsAura("Mighty Rage Potion", actionID, stats.Stats{stats.Strength: 60}, time.Second*15)
		rageMetrics := character.NewRageMetrics(actionID)
		return MajorCooldown{
			Type: CooldownTypeDPS,
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				if character.Class == proto.Class_ClassWarrior {
					return character.CurrentRage() < 25
				}
				return true
			},
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					aura.Activate(sim)
					if character.Class == proto.Class_ClassWarrior {
						bonusRage := sim.RollWithLabel(45, 75, "Mighty Rage Potion")
						character.AddRage(sim, bonusRage, rageMetrics)
					}
				},
			}),
		}
	} else if potionType == proto.Potions_FelManaPotion {
		actionID := ActionID{ItemID: 31677}

		// Restores 3200 mana over 24 seconds.
		manaGain := 3200.0
		if alchStoneEquipped {
			manaGain *= 1.4
		}
		mp5 := manaGain / 24 * 5

		buffAura := character.NewTemporaryStatsAura("Fel Mana Potion", actionID, stats.Stats{stats.MP5: mp5}, time.Second*24)
		debuffAura := character.NewTemporaryStatsAura("Fel Mana Potion Debuff", ActionID{SpellID: 38927}, stats.Stats{stats.SpellPower: -25}, time.Minute*15)

		return MajorCooldown{
			Type: CooldownTypeMana,
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Only pop if we have low enough mana. The potion takes effect over 24
				// seconds so we can pop it a little earlier than the full value.
				return character.MaxMana()-character.CurrentMana() >= 2000
			},
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					buffAura.Activate(sim)
					debuffAura.Activate(sim)
					debuffAura.Refresh(sim)
				},
			}),
		}
	} else if potionType == proto.Potions_InsaneStrengthPotion {
		actionID := ActionID{ItemID: 22828}
		aura := character.NewTemporaryStatsAura("Insane Strength Potion", actionID, stats.Stats{stats.Strength: 120, stats.Defense: -75}, time.Second*15)
		return MajorCooldown{
			Type: CooldownTypeDPS,
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					aura.Activate(sim)
				},
			}),
		}
	} else if potionType == proto.Potions_IronshieldPotion {
		actionID := ActionID{ItemID: 22849}
		aura := character.NewTemporaryStatsAura("Ironshield Potion", actionID, stats.Stats{stats.Armor: 2500}, time.Minute*2)
		return MajorCooldown{
			Type: CooldownTypeDPS,
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					aura.Activate(sim)
				},
			}),
		}
	} else if potionType == proto.Potions_HeroicPotion {
		actionID := ActionID{ItemID: 22837}
		aura := character.NewTemporaryStatsAura("Heroic Potion", actionID, stats.Stats{stats.Strength: 70, stats.Health: 700}, time.Second*15)
		return MajorCooldown{
			Type: CooldownTypeDPS,
			Spell: character.RegisterSpell(SpellConfig{
				ActionID: actionID,
				Flags:    SpellFlagNoOnCastComplete,
				Cast:     potionCast,
				ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
					aura.Activate(sim)
				},
			}),
		}
	} else {
		return MajorCooldown{}
	}
}

var ConjuredAuraTag = "Conjured"

func registerConjuredCD(agent Agent, consumes *proto.Consumes) {
	character := agent.GetCharacter()
	conjuredType := consumes.DefaultConjured

	if conjuredType == proto.Conjured_ConjuredDarkRune {
		actionID := ActionID{ItemID: 20520}
		manaMetrics := character.NewManaMetrics(actionID)
		// damageTakenManaMetrics := character.NewManaMetrics(ActionID{SpellID: 33776})
		spell := character.RegisterSpell(SpellConfig{
			ActionID: actionID,
			Flags:    SpellFlagNoOnCastComplete,
			Cast: CastConfig{
				CD: Cooldown{
					Timer:    character.GetConjuredCD(),
					Duration: time.Minute * 15,
				},
			},
			ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
				// Restores 900 to 1500 mana. (2 Min Cooldown)
				manaGain := sim.RollWithLabel(900, 1500, "dark rune")
				character.AddMana(sim, manaGain, manaMetrics)

				// if character.Class == proto.Class_ClassPaladin {
				// 	// Paladins gain extra mana from self-inflicted damage
				// 	// TO-DO: It is possible for damage to be resisted or to crit
				// 	// This would affect mana returns for Paladins
				// 	manaFromDamage := manaGain * 2.0 / 3.0 * 0.1
				// 	character.AddMana(sim, manaFromDamage, damageTakenManaMetrics, false)
				// }
			},
		})
		character.AddMajorCooldown(MajorCooldown{
			Spell: spell,
			Type:  CooldownTypeMana,
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
				totalRegen := character.ManaRegenPerSecondWhileCasting() * 5
				return character.MaxMana()-(character.CurrentMana()+totalRegen) >= 1500
			},
		})
	} else if conjuredType == proto.Conjured_ConjuredFlameCap {
		actionID := ActionID{ItemID: 22788}

		flameCapProc := character.RegisterSpell(SpellConfig{
			ActionID:    actionID,
			ProcMask:    ProcMaskEmpty,
			SpellSchool: SpellSchoolFire,

			DamageMultiplier: 1,
			CritMultiplier:   character.DefaultSpellCritMultiplier(),
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
				spell.CalcAndDealDamage(sim, target, 40, spell.OutcomeMagicHitAndCrit)
			},
		})

		const procChance = 0.185
		var fireSpells []*Spell
		character.OnSpellRegistered(func(spell *Spell) {
			if spell.SpellSchool.Matches(SpellSchoolFire) {
				fireSpells = append(fireSpells, spell)
			}
		})

		flameCapAura := character.RegisterAura(Aura{
			Label:    "Flame Cap",
			ActionID: actionID,
			Duration: time.Minute,
			OnGain: func(aura *Aura, sim *Simulation) {
				for _, spell := range fireSpells {
					spell.BonusSpellPower += 80
				}
			},
			OnExpire: func(aura *Aura, sim *Simulation) {
				for _, spell := range fireSpells {
					spell.BonusSpellPower -= 80
				}
			},
			OnSpellHitDealt: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(ProcMaskMeleeOrRanged) {
					return
				}
				if sim.RandomFloat("Flame Cap Melee") > procChance {
					return
				}

				flameCapProc.Cast(sim, result.Target)
			},
		})

		spell := character.RegisterSpell(SpellConfig{
			ActionID: actionID,
			Flags:    SpellFlagNoOnCastComplete,
			Cast: CastConfig{
				CD: Cooldown{
					Timer:    character.GetConjuredCD(),
					Duration: time.Minute * 3,
				},
			},
			ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
				flameCapAura.Activate(sim)
			},
		})
		character.AddMajorCooldown(MajorCooldown{
			Spell: spell,
			Type:  CooldownTypeDPS,
		})
	} else if conjuredType == proto.Conjured_ConjuredHealthstone {
		actionID := ActionID{ItemID: 36892}
		healthMetrics := character.NewHealthMetrics(actionID)

		spell := character.RegisterSpell(SpellConfig{
			ActionID: actionID,
			Flags:    SpellFlagNoOnCastComplete,
			Cast: CastConfig{
				CD: Cooldown{
					Timer:    character.GetConjuredCD(),
					Duration: time.Minute * 2,
				},
			},
			ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
				character.GainHealth(sim, 2080*character.PseudoStats.HealingTakenMultiplier, healthMetrics)
			},
		})
		character.AddMajorCooldown(MajorCooldown{
			Spell: spell,
			Type:  CooldownTypeSurvival,
		})
	}
}

var SuperSapperActionID = ActionID{ItemID: 23827}
var BiggerOneActionID = ActionID{ItemID: 23826}
var AdamantiteGrenadeID = ActionID{ItemID: 23737}
var FelIronBombID = ActionID{ItemID: 23736}

func registerExplosivesCD(agent Agent, consumes *proto.Consumes) {
	character := agent.GetCharacter()
	hasFiller := consumes.FillerExplosive != proto.Explosive_ExplosiveUnknown
	if !character.HasProfession(proto.Profession_Engineering) {
		return
	}
	if !consumes.SuperSapper && !hasFiller {
		return
	}
	sharedTimer := character.NewTimer()

	if consumes.SuperSapper {
		character.AddMajorCooldown(MajorCooldown{
			Spell:    character.newSuperSapperSpell(sharedTimer),
			Type:     CooldownTypeDPS | CooldownTypeExplosive,
			Priority: CooldownPriorityLow + 0.03,
		})
	}

	if hasFiller {
		var filler *Spell
		switch consumes.FillerExplosive {
		case proto.Explosive_ExplosiveBiggerOne:
			filler = character.newBiggerOneSpell(sharedTimer)
		case proto.Explosive_ExplosiveAdamantiteGrenade:
			filler = character.newAdamantiteGrenadeSpell(sharedTimer)
		case proto.Explosive_ExplosiveFelIronBomb:
			filler = character.newFelIronBombSpell(sharedTimer)
		}

		character.AddMajorCooldown(MajorCooldown{
			Spell:    filler,
			Type:     CooldownTypeDPS | CooldownTypeExplosive,
			Priority: CooldownPriorityLow + 0.01,
		})
	}
}

// Creates a spell object for the common explosive case.
func (character *Character) newBasicExplosiveSpellConfig(sharedTimer *Timer, actionID ActionID, school SpellSchool, minDamage float64, maxDamage float64, cooldown Cooldown, minSelfDamage float64, maxSelfDamage float64) SpellConfig {
	dealSelfDamage := actionID.SameAction(SuperSapperActionID)

	return SpellConfig{
		ActionID:    actionID,
		SpellSchool: school,
		ProcMask:    ProcMaskEmpty,

		Cast: CastConfig{
			CD: cooldown,
			SharedCD: Cooldown{
				Timer:    sharedTimer,
				Duration: time.Minute,
			},
		},

		// Explosives always have 1% resist chance, so just give them hit cap.
		BonusHitRating:   100 * SpellHitRatingPerHitChance,
		DamageMultiplier: 1,
		CritMultiplier:   2,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			for _, aoeTarget := range sim.Encounter.Targets {
				baseDamage := sim.Roll(minDamage, maxDamage) * sim.Encounter.AOECapMultiplier()
				spell.CalcAndDealDamage(sim, &aoeTarget.Unit, baseDamage, spell.OutcomeMagicHitAndCrit)
			}

			if dealSelfDamage {
				baseDamage := sim.Roll(minDamage, maxDamage)
				spell.CalcAndDealDamage(sim, &character.Unit, baseDamage, spell.OutcomeMagicHitAndCrit)
			}
		},
	}
}
func (character *Character) newSuperSapperSpell(sharedTimer *Timer) *Spell {
	return character.GetOrRegisterSpell(character.newBasicExplosiveSpellConfig(sharedTimer, SuperSapperActionID, SpellSchoolFire, 900, 1500, Cooldown{Timer: character.NewTimer(), Duration: time.Minute * 5}, 675, 1125))
}
func (character *Character) newBiggerOneSpell(sharedTimer *Timer) *Spell {
	return character.GetOrRegisterSpell(character.newBasicExplosiveSpellConfig(sharedTimer, BiggerOneActionID, SpellSchoolFire, 600, 1000, Cooldown{}, 0, 0))
}
func (character *Character) newAdamantiteGrenadeSpell(sharedTimer *Timer) *Spell {
	return character.GetOrRegisterSpell(character.newBasicExplosiveSpellConfig(sharedTimer, AdamantiteGrenadeID, SpellSchoolFire, 450, 750, Cooldown{}, 0, 0))
}
func (character *Character) newFelIronBombSpell(sharedTimer *Timer) *Spell {
	return character.GetOrRegisterSpell(character.newBasicExplosiveSpellConfig(sharedTimer, FelIronBombID, SpellSchoolFire, 330, 770, Cooldown{}, 0, 0))
}
