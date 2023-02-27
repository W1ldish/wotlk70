package database

import (
	"github.com/Tereneckla/wotlk/sim/core/proto"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

// Note: EffectId AND SpellId are required for all enchants, because they are
// used by various importers/exporters. ItemId is optional.

var EnchantOverrides = []*proto.UIEnchant{
	// Head
	{EffectId: 2999, ItemId: 29186, SpellId: 35443, Name: "Arcanum of the Defender", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Defense: 16, stats.Dodge: 17}.ToFloatArray(), Type: proto.ItemType_ItemTypeHead},
	{EffectId: 3002, ItemId: 29191, SpellId: 35447, Name: "Arcanum of Power", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.SpellPower: 22, stats.SpellHit: 14}.ToFloatArray(), Type: proto.ItemType_ItemTypeHead},
	{EffectId: 3003, ItemId: 29192, SpellId: 35452, Name: "Arcanum of Ferocity", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.AttackPower: 34, stats.RangedAttackPower: 34, stats.MeleeHit: 16, stats.SpellHit: 16}.ToFloatArray(), Type: proto.ItemType_ItemTypeHead},
	{EffectId: 3096, ItemId: 30846, SpellId: 37891, Name: "Arcanum of the Outcast", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Strength: 17, stats.Intellect: 16}.ToFloatArray(), Type: proto.ItemType_ItemTypeHead},
	{EffectId: 3004, ItemId: 29193, SpellId: 35453, Name: "Arcanum of the Gladiator", Phase: 6, Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Stamina: 18, stats.Resilience: 20}.ToFloatArray(), Type: proto.ItemType_ItemTypeHead},

	// ZG Head Enchants
	{EffectId: 2583, ItemId: 19782, SpellId: 24149, Name: "Presence of Might", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Stamina: 10, stats.Defense: 10, stats.BlockValue: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeHead, ClassAllowlist: []proto.Class{proto.Class_ClassWarrior}},

	// Shoulder
	{EffectId: 2982, ItemId: 28886, SpellId: 35406, Name: "Greater Inscription of Discipline", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.SpellPower: 18, stats.SpellCrit: 10}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2986, ItemId: 28888, SpellId: 35417, Name: "Greater Inscription of Vengeance", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.AttackPower: 30, stats.RangedAttackPower: 30, stats.MeleeCrit: 10}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2978, ItemId: 28889, SpellId: 35402, Name: "Greater Inscription of Warding", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Defense: 10, stats.Dodge: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2995, ItemId: 28909, SpellId: 35437, Name: "Greater Inscription of the Orb", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.SpellPower: 12, stats.SpellCrit: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2997, ItemId: 28910, SpellId: 35437, Name: "Greater Inscription of the Blade", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.AttackPower: 20, stats.RangedAttackPower: 20, stats.MeleeCrit: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2991, ItemId: 28911, SpellId: 35439, Name: "Greater Inscription of the Knight", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Defense: 15, stats.Dodge: 10}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2605, ItemId: 20076, SpellId: 24421, Name: "Zandalar Signet of Mojo", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.SpellPower: 18}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2721, ItemId: 23545, SpellId: 29467, Name: "Power of the Scourge", Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.SpellPower: 15, stats.SpellCrit: 14}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2717, ItemId: 23548, SpellId: 29483, Name: "Might of the Scourge", Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.AttackPower: 26, stats.RangedAttackPower: 26, stats.MeleeCrit: 14}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},
	{EffectId: 2716, ItemId: 23549, SpellId: 29480, Name: "Fortitude of the Scourge", Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Stamina: 16, stats.Armor: 100}.ToFloatArray(), Type: proto.ItemType_ItemTypeShoulder},

	// Back
	{EffectId: 2622, ItemId: 33148, SpellId: 25086, Name: "Enchant Cloak - Dodge", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Dodge: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeBack},
	{EffectId: 2621, ItemId: 33150, SpellId: 25084, Name: "Enchant Cloak - Subtlety", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeBack},
	{EffectId: 849, ItemId: 11206, SpellId: 13882, Name: "Enchant Cloak - Lesser Agility", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Agility: 3}.ToFloatArray(), Type: proto.ItemType_ItemTypeBack},
	{EffectId: 368, SpellId: 34004, Name: "Enchant Cloak - Greater Agility", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Agility: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeBack},
	{EffectId: 2938, ItemId: 28274, SpellId: 34003, Name: "Enchant Cloak - Spell Penetration", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.SpellPenetration: 20}.ToFloatArray(), Type: proto.ItemType_ItemTypeBack},
	{EffectId: 1441, ItemId: 28277, SpellId: 34006, Name: "Enchant Cloak - Greater Shadow Resistance", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.ShadowResistance: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeBack},
	{EffectId: 2648, ItemId: 35756, SpellId: 47051, Name: "Enchant Cloak - Steelweave", Phase: 6, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Defense: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeBack},

	// Chest
	{EffectId: 2659, SpellId: 27957, Name: "Chest - Exceptional Health", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Health: 150}.ToFloatArray(), Type: proto.ItemType_ItemTypeChest},
	{EffectId: 2661, ItemId: 24003, SpellId: 27960, Name: "Chest - Exceptional Stats", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Stamina: 6, stats.Intellect: 6, stats.Spirit: 6, stats.Strength: 6, stats.Agility: 6}.ToFloatArray(), Type: proto.ItemType_ItemTypeChest},
	{EffectId: 2933, ItemId: 28270, SpellId: 33992, Name: "Chest - Major Resilience", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Resilience: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeChest},
	{EffectId: 1144, SpellId: 33990, Name: "Chest - Major Spirit", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Spirit: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeChest},
	{EffectId: 3150, SpellId: 33991, Name: "Chest - Restore Mana Prime", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.MP5: 6}.ToFloatArray(), Type: proto.ItemType_ItemTypeChest},
	{EffectId: 1950, ItemId: 35500, SpellId: 46594, Name: "Chest - Defense", Phase: 6, Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Defense: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeChest},

	// Wrist
	{EffectId: 2649, ItemId: 22533, SpellId: 27914, Name: "Bracer - Fortitude", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Stamina: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeWrist},
	{EffectId: 2650, ItemId: 22534, SpellId: 27917, Name: "Bracer - Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.SpellPower: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeWrist},
	{EffectId: 369, SpellId: 34001, Name: "Bracer - Major Intellect", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Intellect: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeWrist},
	{EffectId: 2647, SpellId: 27899, Name: "Bracer - Brawn", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Strength: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeWrist},
	{EffectId: 1593, SpellId: 34002, Name: "Bracer - Assault", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.AttackPower: 24, stats.RangedAttackPower: 24}.ToFloatArray(), Type: proto.ItemType_ItemTypeWrist},
	{EffectId: 1891, SpellId: 27905, Name: "Bracer - Stats", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Stamina: 4, stats.Intellect: 4, stats.Spirit: 4, stats.Strength: 4, stats.Agility: 4}.ToFloatArray(), Type: proto.ItemType_ItemTypeWrist},

	// Hands
	{EffectId: 2935, ItemId: 28271, SpellId: 33994, Name: "Gloves - Spell Strike", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.SpellHit: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeHands},
	{EffectId: 2937, ItemId: 28272, SpellId: 33997, Name: "Gloves - Major Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.SpellPower: 20}.ToFloatArray(), Type: proto.ItemType_ItemTypeHands},
	{EffectId: 684, SpellId: 33995, Name: "Gloves - Major Strength", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Strength: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeHands},
	{EffectId: 2564, ItemId: 33152, SpellId: 25080, Name: "Gloves - Major Agility", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Agility: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeHands},
	{EffectId: 2613, ItemId: 33153, SpellId: 25072, Name: "Gloves - Threat", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeHands},

	// Legs
	{EffectId: 2748, ItemId: 24274, SpellId: 31372, Name: "Runic Spellthread", Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.SpellPower: 35, stats.Stamina: 20}.ToFloatArray(), Type: proto.ItemType_ItemTypeLegs},
	{EffectId: 2747, ItemId: 24273, SpellId: 31371, Name: "Mystic Spellthread", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.SpellPower: 25, stats.Stamina: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeLegs},
	{EffectId: 3010, ItemId: 29533, SpellId: 35488, Name: "Cobrahide Leg Armor", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.AttackPower: 40, stats.RangedAttackPower: 40, stats.MeleeCrit: 10}.ToFloatArray(), Type: proto.ItemType_ItemTypeLegs},
	{EffectId: 3012, ItemId: 29535, SpellId: 35490, Name: "Nethercobra Leg Armor", Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.AttackPower: 50, stats.RangedAttackPower: 50, stats.MeleeCrit: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeLegs},
	{EffectId: 3013, ItemId: 29536, SpellId: 35495, Name: "Nethercleft Leg Armor", Quality: proto.ItemQuality_ItemQualityEpic, Stats: stats.Stats{stats.Stamina: 40, stats.Agility: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeLegs},

	// Feet
	{EffectId: 851, ItemId: 16220, SpellId: 20024, Name: "Enchant Boots - Spirit", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Spirit: 5}.ToFloatArray(), Type: proto.ItemType_ItemTypeFeet},
	{EffectId: 2940, ItemId: 35297, SpellId: 34008, Name: "Enchant Boots - Boar's Speed", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Stamina: 9}.ToFloatArray(), Type: proto.ItemType_ItemTypeFeet},
	{EffectId: 2656, ItemId: 35298, SpellId: 27948, Name: "Enchant Boots - Vitality", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.MP5: 4}.ToFloatArray(), Type: proto.ItemType_ItemTypeFeet},
	{EffectId: 2649, ItemId: 22543, SpellId: 27950, Name: "Enchant Boots - Fortitude", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Stamina: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeFeet},
	{EffectId: 2657, ItemId: 22544, SpellId: 27951, Name: "Enchant Boots - Dexterity", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Agility: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeFeet},
	{EffectId: 2939, ItemId: 28279, SpellId: 34007, Name: "Enchant Boots - Cat's Swiftness", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.Agility: 6}.ToFloatArray(), Type: proto.ItemType_ItemTypeFeet},
	{EffectId: 2658, ItemId: 22545, SpellId: 27954, Name: "Enchant Boots - Surefooted", Phase: 2, Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.MeleeHit: 10}.ToFloatArray(), Type: proto.ItemType_ItemTypeFeet},

	// Weapon
	{EffectId: 1897, ItemId: 16250, SpellId: 20031, Name: "Superior Striking", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 963, ItemId: 22552, SpellId: 27967, Name: "Major Striking", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 1900, ItemId: 16252, SpellId: 20034, Name: "Crusader", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 2666, ItemId: 22551, SpellId: 27968, Name: "Major Intellect", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Intellect: 30}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 2667, ItemId: 22554, SpellId: 27971, Name: "Savagery", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.AttackPower: 70, stats.RangedAttackPower: 70}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon, EnchantType: proto.EnchantType_EnchantTypeTwoHand},
	{EffectId: 2669, ItemId: 22555, SpellId: 27975, Name: "Major Spellpower", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.SpellPower: 40}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 2671, ItemId: 22560, SpellId: 27981, Name: "Sunfire", Phase: 2, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 2672, ItemId: 22561, SpellId: 27982, Name: "Soulfrost", Phase: 2, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 2673, ItemId: 22559, SpellId: 27984, Name: "Mongoose", Phase: 2, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 2564, ItemId: 19445, SpellId: 23800, Name: "Agility", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Agility: 15}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 3222, ItemId: 33165, SpellId: 42620, Name: "Greater Agility", Phase: 2, Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Agility: 20}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 2670, ItemId: 22556, SpellId: 27977, Name: "2H Weapon - Major Agility", Quality: proto.ItemQuality_ItemQualityUncommon, Stats: stats.Stats{stats.Agility: 35}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon, EnchantType: proto.EnchantType_EnchantTypeTwoHand},
	{EffectId: 3225, ItemId: 33307, SpellId: 42974, Name: "Executioner", Phase: 5, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 3273, ItemId: 35498, SpellId: 46578, Name: "Deathfrost", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon},
	{EffectId: 3855, ItemId: 45060, SpellId: 62959, Name: "Staff - Spellpower", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{stats.SpellPower: 69}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon, EnchantType: proto.EnchantType_EnchantTypeStaff},

	// Shield
	{EffectId: 2654, ItemId: 22539, SpellId: 27945, Name: "Intellect", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Intellect: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon, EnchantType: proto.EnchantType_EnchantTypeShield},
	{EffectId: 1071, ItemId: 28282, SpellId: 34009, Name: "Major Stamina", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Stamina: 18}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon, EnchantType: proto.EnchantType_EnchantTypeShield},
	{EffectId: 3229, SpellId: 44383, Name: "Resilience", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Resilience: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeWeapon, EnchantType: proto.EnchantType_EnchantTypeShield},

	// Ring
	{EffectId: 2929, ItemId: 22535, SpellId: 27920, Name: "Striking", Phase: 2, Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeFinger, RequiredProfession: proto.Profession_Enchanting},
	{EffectId: 2928, ItemId: 22536, SpellId: 27924, Name: "Spellpower", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.SpellPower: 12}.ToFloatArray(), Type: proto.ItemType_ItemTypeFinger, RequiredProfession: proto.Profession_Enchanting},
	{EffectId: 2931, ItemId: 22538, SpellId: 27927, Name: "Stats", Quality: proto.ItemQuality_ItemQualityCommon, Stats: stats.Stats{stats.Stamina: 4, stats.Intellect: 4, stats.Spirit: 4, stats.Strength: 4, stats.Agility: 4}.ToFloatArray(), Type: proto.ItemType_ItemTypeFinger, RequiredProfession: proto.Profession_Enchanting},

	// Ranged
	{EffectId: 2523, ItemId: 18283, SpellId: 22779, Name: "Biznicks 247x128 Accurascope", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeRanged},
	{EffectId: 2723, ItemId: 23765, SpellId: 30252, Name: "Khorium Scope", Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeRanged},
	{EffectId: 2724, ItemId: 23766, SpellId: 30260, Name: "Stabilized Eternium Scope", Phase: 2, Quality: proto.ItemQuality_ItemQualityRare, Stats: stats.Stats{}.ToFloatArray(), Type: proto.ItemType_ItemTypeRanged},
}
