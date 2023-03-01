import { Consumes, WeaponEnchant } from '../core/proto/common.js';
import { EquipmentSpec } from '../core/proto/common.js';
import { Flask } from '../core/proto/common.js';
import { Food } from '../core/proto/common.js';
import { Glyphs } from '../core/proto/common.js';
import { ItemSpec } from '../core/proto/common.js';
import { Potions } from '../core/proto/common.js';
import { Faction } from '../core/proto/common.js';
import { RaidBuffs } from '../core/proto/common.js';
import { IndividualBuffs } from '../core/proto/common.js';
import { Debuffs } from '../core/proto/common.js';
import { TristateEffect } from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import { Player } from '../core/player.js';

import {
	ShadowPriest_Rotation as Rotation,
	ShadowPriest_Options as Options,
	ShadowPriest_Rotation_RotationType,
	PriestMajorGlyph as MajorGlyph,
	PriestMinorGlyph as MinorGlyph,
} from '../core/proto/priest.js';


import * as Tooltips from '../core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Standard',
	data: SavedTalents.create({
		talentsString: '--325323051223010323152301351',
		glyphs: Glyphs.create({
			major1: MajorGlyph.GlyphOfShadow,
			major2: MajorGlyph.GlyphOfMindFlay,
			major3: MajorGlyph.PriestMajorGlyphNone,
			minor1: MinorGlyph.GlyphOfFortitude,
			minor2: MinorGlyph.GlyphOfShadowProtection,
			minor3: MinorGlyph.GlyphOfShadowfiend,
		}),
	}),
};

export const DefaultRotation = Rotation.create({
	rotationType: ShadowPriest_Rotation_RotationType.Ideal,
});

export const DefaultOptions = Options.create({
	useShadowfiend: true,
	useMindBlast: true,
	useShadowWordDeath: true,
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskOfPureDeath,
	food: Food.FoodBlackenedBasilisk,
	defaultPotion: Potions.SuperManaPotion,
	prepopPotion: Potions.DestructionPotion,
	weaponMain: WeaponEnchant.EnchantBrilliantWizardOil,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	arcaneBrilliance: true,
	divineSpirit: true,
	trueshotAura: true,
	leaderOfThePack: TristateEffect.TristateEffectImproved,
	icyTalons: true,
	totemOfWrath: true,
	moonkinAura: TristateEffect.TristateEffectImproved,
	wrathOfAirTotem: true,
	sanctifiedRetribution: true,
	bloodlust: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	vampiricTouch: true,
});

export const DefaultDebuffs = Debuffs.create({
	sunderArmor: true,
	faerieFire: TristateEffect.TristateEffectImproved,
	bloodFrenzy: true,
	ebonPlaguebringer: true,
	heartOfTheCrusader: true,
	judgementOfWisdom: true,
});

export const PreBis_PRESET = {
	name: 'PreBis Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
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
			"id": 25777,
			"enchant": 2621
		  },
		  {
			"id": 21848,
			"enchant": 1144,
			"gems": [
			  31867,
			  31867
			]
		  },
		  {
			"id": 32655,
			"enchant": 2650,
			"gems": [
			  31867,
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
			  24030,
			  24030,
			  24030
			]
		  },
		  {
			"id": 28406,
			"gems": [
			  31867,
			  31867
			]
		  },
		  {
			"id": 28227
		  },
		  {
			"id": 32779
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
export const P1_PRESET = {
	name: 'P1 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(` {"items": [
		{
			"id": 40562,
			"enchant": 3820,
			"gems": [
				41285,
				39998
			]
		},
		{
			"id": 44661,
			"gems": [
				40026
			]
		},
		{
			"id": 40459,
			"enchant": 3810,
			"gems": [
				39998
			]
		},
		{
			"id": 44005,
			"enchant": 3722,
			"gems": [
				40026
			]
		},
		{
			"id": 44002,
			"enchant": 1144,
			"gems": [
				39998,
				39998
			]
		},
		{
			"id": 44008,
			"enchant": 2332,
			"gems": [
				39998,
				0
			]
		},
		{
			"id": 40454,
			"enchant": 3604,
			"gems": [
				40049,
				0
			]
		},
		{
			"id": 40561,
			"gems": [
				39998
			]
		},
		{
			"id": 40560,
			"enchant": 3719
		},
		{
			"id": 40558,
			"enchant": 3606
		},
		{
			"id": 40719
		},
		{
			"id": 40399
		},
		{
			"id": 40255
		},
		{
			"id": 40432
		},
		{
			"id": 40395,
			"enchant": 3834
		},
		{
			"id": 40273
		},
		{
			"id": 39712
		}
  ]}`),
};
export const P2_PRESET = {
	name: 'P2 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{ "items": [
        {
          "id": 46172,
          "enchant": 3820,
          "gems": [
            41285,
            45883
          ]
        },
        {
          "id": 45243,
          "gems": [
            39998
          ]
        },
        {
          "id": 46165,
          "enchant": 3810,
          "gems": [
            39998
          ]
        },
        {
          "id": 45242,
          "enchant": 3722,
          "gems": [
            40049
          ]
        },
        {
          "id": 46168,
          "enchant": 1144,
          "gems": [
            39998,
            39998
          ]
        },
        {
          "id": 45446,
          "enchant": 2332,
          "gems": [
            39998,
            0
          ]
        },
        {
          "id": 45665,
          "enchant": 3604,
          "gems": [
            39998,
            39998,
            0
          ]
        },
        {
          "id": 45619,
          "enchant": 3601,
          "gems": [
            39998,
            39998,
            39998
          ]
        },
        {
          "id": 46170,
          "enchant": 3719,
          "gems": [
            39998,
            40049
          ]
        },
        {
          "id": 45135,
          "enchant": 3606,
          "gems": [
            39998,
            40049
          ]
        },
        {
          "id": 45495,
          "gems": [
            40026
          ]
        },
        {
          "id": 46046,
          "gems": [
            39998
          ]
        },
        {
          "id": 45518
        },
        {
          "id": 45466
        },
        {
          "id": 45620,
          "enchant": 3834,
          "gems": [
            40026
          ]
        },
        {
          "id": 45617
        },
        {
          "id": 45294,
          "gems": [
            39998
          ]
        }
      ]
    }`),
};