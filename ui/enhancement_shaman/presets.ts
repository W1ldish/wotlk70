import {
	Consumes,
	Flask,
	Food,
	Glyphs,
	EquipmentSpec,
	Potions,
	RaidBuffs,
	TristateEffect,
	Debuffs,
  CustomRotation,
  CustomSpell,
  ItemSwap,
  ItemSpec,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';

import { EnhancementShaman_Rotation as EnhancementShamanRotation, EnhancementShaman_Options as EnhancementShamanOptions, ShamanShield } from '../core/proto/shaman.js';
import {
	AirTotem,
	EarthTotem,
	FireTotem,
	WaterTotem,
	ShamanTotems,
	ShamanImbue,
	ShamanSyncType,
	ShamanMajorGlyph,
	EnhancementShaman_Rotation_PrimaryShock as PrimaryShock,
  EnhancementShaman_Rotation_RotationType as RotationType,
  EnhancementShaman_Rotation_CustomRotationSpell as CustomRotationSpell
} from '../core/proto/shaman.js';

import * as Tooltips from '../core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Standard',
	data: SavedTalents.create({
		talentsString: '053030152-3040500310984633031131031051',
		glyphs: Glyphs.create({
			major1: ShamanMajorGlyph.GlyphOfStormstrike,
			major2: ShamanMajorGlyph.GlyphOfFlametongueWeapon,
			major3: ShamanMajorGlyph.GlyphOfFeralSpirit,
			//minor glyphs dont affect damage done, all convenience/QoL
		})
	}),
};

export const DefaultRotation = EnhancementShamanRotation.create({
	totems: ShamanTotems.create({
		earth: EarthTotem.StrengthOfEarthTotem,
		air: AirTotem.WindfuryTotem,
		fire: FireTotem.MagmaTotem,
		water: WaterTotem.ManaSpringTotem,
		useFireElemental: true,
	}),
	maelstromweaponMinStack: 3,
	lightningboltWeave: true,
	autoWeaveDelay: 500,
	delayGcdWeave: 750,
	lavaburstWeave: false,
	firenovaManaThreshold: 3000,
	shamanisticRageManaThreshold: 25,
	primaryShock: PrimaryShock.Earth,
	weaveFlameShock: true,
  	rotationType: RotationType.Priority,
  	customRotation: CustomRotation.create({
			spells: [
			CustomSpell.create({ spell: CustomRotationSpell.LightningBolt }),
			CustomSpell.create({ spell: CustomRotationSpell.StormstrikeDebuffMissing }),
			CustomSpell.create({ spell: CustomRotationSpell.LightningBoltWeave }),
			CustomSpell.create({ spell: CustomRotationSpell.Stormstrike }),
			CustomSpell.create({ spell: CustomRotationSpell.FlameShock }),
			CustomSpell.create({ spell: CustomRotationSpell.EarthShock }),
			CustomSpell.create({ spell: CustomRotationSpell.MagmaTotem}),
			CustomSpell.create({ spell: CustomRotationSpell.LightningShield }),
			CustomSpell.create({ spell: CustomRotationSpell.FireNova }),
			CustomSpell.create({ spell: CustomRotationSpell.LightningBoltDelayedWeave }),
			CustomSpell.create({ spell: CustomRotationSpell.LavaLash }),
		],
	}),
});

export const DefaultOptions = EnhancementShamanOptions.create({
	shield: ShamanShield.LightningShield,
	bloodlust: true,
	imbueMh: ShamanImbue.WindfuryWeapon,
	imbueOh: ShamanImbue.FlametongueWeapon,
	syncType: ShamanSyncType.SyncMainhandOffhandSwings,
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.PotionOfSpeed,
	flask: Flask.FlaskOfEndlessRage,
	food: Food.FoodFishFeast,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	arcaneBrilliance: true,
	leaderOfThePack: TristateEffect.TristateEffectImproved,
	totemOfWrath: true,
	wrathOfAirTotem: true,
	moonkinAura: TristateEffect.TristateEffectImproved,
	sanctifiedRetribution: true,
	divineSpirit: true,
	battleShout: TristateEffect.TristateEffectImproved,
});

export const DefaultDebuffs = Debuffs.create({
	bloodFrenzy: true,
	sunderArmor: true,
	curseOfWeakness: TristateEffect.TristateEffectRegular,
	curseOfElements: true,
	faerieFire: TristateEffect.TristateEffectImproved,
	judgementOfWisdom: true,
	misery: true,
	totemOfWrath: true,
	shadowMastery: true,
});


export const PreRaid_PRESET = {
	name: 'Preraid Preset',
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
			"id": 32078,
			"enchant": 2982
		  },
		  {
			"id": 27981,
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
			  31867
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
			"id": 29268,
			"enchant": 2654
		  },
		  {
			"id": 28248
		  }
	]}`),
}