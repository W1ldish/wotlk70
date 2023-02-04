package core

import (
	"time"
)

const CharacterLevel = 70

const GCDMin = time.Second * 1
const GCDDefault = time.Millisecond * 1500

const MeleeAttackRatingPerDamage = 14.0
const ExpertisePerQuarterPercentReduction = 15.77 / 4 // TODO: Does it still cutoff at 1/4 percents?
const ArmorPenPerPercentArmor = 5.92

const HasteRatingPerHastePercent = 15.77
const CritRatingPerCritChance = 22.08

const SpellHitRatingPerHitChance = 12.62
const MeleeHitRatingPerHitChance = 15.77

const DefenseRatingPerDefense = 2.37
const DodgeRatingPerDodgeChance = 21.76
const ParryRatingPerParryChance = 21.76
const BlockRatingPerBlockChance = 7.88
const MissDodgeParryBlockCritChancePerDefense = 0.04

const DefenseRatingToChanceReduction = (1.0 / DefenseRatingPerDefense) * MissDodgeParryBlockCritChancePerDefense / 100

const ResilienceRatingPerCritReductionChance = 82.0
const ResilienceRatingPerCritDamageReductionPercent = 39.4231 / 2.2

// TODO: More log scraping to verify this value for WOTLK.
// Assuming 574 AP debuffs go to exactly zero and achieve -14.2%
const EnemyAutoAttackAPCoefficient = 0.0002883296

const AverageMagicPartialResistMultiplier = 0.94

// IDs for items used in core
const (
	ItemIDAtieshMage            = 22589
	ItemIDAtieshWarlock         = 22630
	ItemIDBraidedEterniumChain  = 24114
	ItemIDChainOfTheTwilightOwl = 24121
	ItemIDEyeOfTheNight         = 24116
	ItemIDJadePendantOfBlasting = 20966
	ItemIDTheLightningCapacitor = 28785
)

type Hand bool

const MainHand Hand = true
const OffHand Hand = false
