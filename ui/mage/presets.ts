import { Conjured, WeaponEnchant } from '../core/proto/common.js';
import { Consumes } from '../core/proto/common.js';
import { EquipmentSpec } from '../core/proto/common.js';
import { Flask } from '../core/proto/common.js';
import { Food } from '../core/proto/common.js';
import { Glyphs } from '../core/proto/common.js';
import { ItemSpec } from '../core/proto/common.js';
import { Potions } from '../core/proto/common.js';
import { Spec } from '../core/proto/common.js';
import { Faction } from '../core/proto/common.js';
import { RaidTarget } from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import { Player } from '../core/player.js';
import { NO_TARGET } from '../core/proto_utils/utils';

import {
	Mage,
	MageTalents as MageTalents,
	Mage_Rotation as MageRotation,
	Mage_Rotation_Type as RotationType,
	Mage_Rotation_PrimaryFireSpell as PrimaryFireSpell,
	Mage_Rotation_AoeRotation as AoeRotationSpells,
	Mage_Options as MageOptions,
	Mage_Options_ArmorType as ArmorType,
	MageMajorGlyph,
	MageMinorGlyph,
} from '../core/proto/mage.js';

import * as Tooltips from '../core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.
export const ArcaneTalents = {
	name: 'Arcane',
	data: SavedTalents.create({
		talentsString: '230005233100330150323102505321-03',
		glyphs: Glyphs.create({
			major1: MageMajorGlyph.GlyphOfArcaneBlast,
			major2: MageMajorGlyph.GlyphOfArcaneMissiles,
			major3: MageMajorGlyph.MageMajorGlyphNone,
			minor1: MageMinorGlyph.GlyphOfSlowFall,
			minor2: MageMinorGlyph.GlyphOfFrostWard,
			minor3: MageMinorGlyph.GlyphOfBlastWave,
		}),
	}),
};
export const FireTalents = {
	name: 'Fire',
	data: SavedTalents.create({
		talentsString: '23000503110003-00550300123013300531203003',
		glyphs: Glyphs.create({
			major1: MageMajorGlyph.GlyphOfFireball,
			major2: MageMajorGlyph.GlyphOfMoltenArmor,
			major3: MageMajorGlyph.MageMajorGlyphNone,
			minor1: MageMinorGlyph.GlyphOfSlowFall,
			minor2: MageMinorGlyph.GlyphOfFrostWard,
			minor3: MageMinorGlyph.GlyphOfBlastWave,
		}),
	}),
};
export const FrostTalents = {
	name: 'Frost',
	data: SavedTalents.create({
		talentsString: '--0533033313233100230152231351',
		glyphs: Glyphs.create({
			major1: MageMajorGlyph.GlyphOfFrostbolt,
			major2: MageMajorGlyph.GlyphOfEternalWater,
			major3: MageMajorGlyph.MageMajorGlyphNone,
			minor1: MageMinorGlyph.GlyphOfSlowFall,
			minor2: MageMinorGlyph.GlyphOfFrostWard,
			minor3: MageMinorGlyph.GlyphOfBlastWave,
		}),
	}),
};

export const DefaultFireRotation = MageRotation.create({
	type: RotationType.Fire,
	primaryFireSpell: PrimaryFireSpell.Fireball,
	maintainImprovedScorch: false,
	pyroblastDelayMs: 50,
});

export const DefaultFireOptions = MageOptions.create({
	armor: ArmorType.MoltenArmor,
	focusMagicPercentUptime: 99,
	focusMagicTarget: RaidTarget.create({
		targetIndex: NO_TARGET,
	}),
	reactionTimeMs: 300,
	igniteMunching: true,
});

export const DefaultFireConsumes = Consumes.create({
	flask: Flask.FlaskOfPureDeath,
	food: Food.FoodBlackenedBasilisk,
	defaultPotion: Potions.HastePotion,
	defaultConjured: Conjured.ConjuredFlameCap,
	weaponMain: WeaponEnchant.EnchantBrilliantWizardOil,
});

export const DefaultFrostRotation = MageRotation.create({
	type: RotationType.Frost,
	waterElementalDisobeyChance: 0.1,
});

export const DefaultFrostOptions = MageOptions.create({
	armor: ArmorType.MoltenArmor,
	focusMagicTarget: RaidTarget.create({
		targetIndex: NO_TARGET,
	}),
	reactionTimeMs: 300,
});

export const DefaultFrostConsumes = Consumes.create({
	defaultPotion: Potions.PotionOfSpeed,
	flask: Flask.FlaskOfPureDeath,
	food: Food.FoodBlackenedBasilisk,
	weaponMain: WeaponEnchant.EnchantBrilliantWizardOil,
});

export const DefaultArcaneRotation = MageRotation.create({
	type: RotationType.Arcane,
	only3ArcaneBlastStacksBelowManaPercent: 0.15,
	blastWithoutMissileBarrageAboveManaPercent: 0.2,
	extraBlastsDuringFirstAp: 0,
	missileBarrageBelowArcaneBlastStacks: 0,
	missileBarrageBelowManaPercent: 0,
});

export const DefaultArcaneOptions = MageOptions.create({
	armor: ArmorType.MoltenArmor,
	focusMagicPercentUptime: 99,
	focusMagicTarget: RaidTarget.create({
		targetIndex: NO_TARGET,
	}),
	reactionTimeMs: 300,
});

export const DefaultArcaneConsumes = Consumes.create({
	defaultPotion: Potions.HastePotion,
	defaultConjured: Conjured.ConjuredDarkRune,
	flask: Flask.FlaskOfBlindingLight,
	food: Food.FoodBlackenedBasilisk,
	weaponMain: WeaponEnchant.EnchantBrilliantWizardOil,
});

export const OtherDefaults = {
	distanceFromTarget: 25,
};

export const ARCANE_PRERAID_PRESET = {
	name: "Arcane Preraid Preset",
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	enableWhen: (player: Player<Spec.SpecMage>) => player.getRotation().type == RotationType.Arcane,
	gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 24266,
			"enchant": 3002,
			"gems": [
			  31867,
			  31867,
			  31867
			]
		  },
		  {
			"id": 28134
		  },
		  {
			"id": 27796,
			"enchant": 2982,
			"gems": [
			  31867,
			  31867
			]
		  },
		  {
			"id": 29369,
			"enchant": 2621
		  },
		  {
			"id": 21848,
			"enchant": 2661,
			"gems": [
			  31867,
			  31867
			]
		  },
		  {
			"id": 32655,
			"enchant": 2650,
			"gems": [
			  24030,
			  0
			]
		  },
		  {
			"id": 21847,
			"enchant": 2937,
			"gems": [
			  31867,
			  31867,
			  0
			]
		  },
		  {
			"id": 21846,
			"gems": [
			  31867,
			  31867,
			  0
			]
		  },
		  {
			"id": 24262,
			"enchant": 2748,
			"gems": [
			  31867,
			  31867,
			  24030
			]
		  },
		  {
			"id": 28406,
			"gems": [
			  24030,
			  31867
			]
		  },
		  {
			"id": 32779
		  },
		  {
			"id": 28227
		  },
		  {
			"id": 31856
		  },
		  {
			"id": 29370
		  },
		  {
			"id": 23554,
			"enchant": 2669
		  },
		  {
			"id": 28412
		  },
		  {
			"id": 29350
		  }
	]}`),
};
