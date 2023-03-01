import {
	Conjured,
	Consumes,
	Debuffs,
	EquipmentSpec,
	Flask,
	Food,
	Glyphs,
	IndividualBuffs,
	PartyBuffs,
	Potions,
	RaidBuffs,
	RaidTarget,
	TristateEffect,
	WeaponImbue
} from '../core/proto/common.js';
import {SavedTalents} from '../core/proto/ui.js';

import {
	BalanceDruid_Options as BalanceDruidOptions,
	BalanceDruid_Rotation as BalanceDruidRotation,
	BalanceDruid_Rotation_IsUsage,
	BalanceDruid_Rotation_MfUsage,
	BalanceDruid_Rotation_Type as RotationType,
	DruidMajorGlyph,
	DruidMinorGlyph,
} from '../core/proto/druid.js';

import * as Tooltips from '../core/constants/tooltips.js';
import {NO_TARGET} from "../core/proto_utils/utils";

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Phase 1',
	data: SavedTalents.create({
		talentsString: '5032003125331303213305311231--2',
		glyphs: Glyphs.create({
			major1: DruidMajorGlyph.GlyphOfStarfire,
			major2: DruidMajorGlyph.GlyphOfStarfall,
			major3: DruidMajorGlyph.DruidMajorGlyphNone,
			minor1: DruidMinorGlyph.GlyphOfTyphoon,
			minor2: DruidMinorGlyph.GlyphOfUnburdenedRebirth,
			minor3: DruidMinorGlyph.GlyphOfTheWild,
		}),
	}),
};

export const DefaultRotation = BalanceDruidRotation.create({
	type: RotationType.Adaptive,
	maintainFaerieFire: true,
	useSmartCooldowns: true,
	mfUsage: BalanceDruid_Rotation_MfUsage.NoMf,
	isUsage: BalanceDruid_Rotation_IsUsage.MaximizeIs,
	useStarfire:       true,
	useWrath:          true,
	useBattleRes: false,
	playerLatency: 200,
});

export const DefaultOptions = BalanceDruidOptions.create({
	innervateTarget: RaidTarget.create({
		targetIndex: NO_TARGET,
	}),
});

export const DefaultConsumes = Consumes.create({
    defaultPotion: Potions.HastePotion,
    flask: Flask.FlaskOfBlindingLight,
    food: Food.FoodBlackenedBasilisk,
    prepopPotion: Potions.HastePotion,
	defaultConjured: Conjured.ConjuredDarkRune,
    superSapper: true,
	weaponMain: WeaponImbue.ImbueBrilliantWizardOil,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	bloodlust: true,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	icyTalons: true,
	moonkinAura: TristateEffect.TristateEffectImproved,
	leaderOfThePack: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	sanctifiedRetribution: true,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	trueshotAura: true,
	wrathOfAirTotem: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	vampiricTouch: true,
});

export const DefaultPartyBuffs = PartyBuffs.create({
	heroicPresence: false,
});

export const DefaultDebuffs = Debuffs.create({
	bloodFrenzy: true,
	ebonPlaguebringer: true,
	faerieFire: TristateEffect.TristateEffectImproved,
	heartOfTheCrusader: true,
	judgementOfWisdom: true,
	shadowMastery: true,
	sunderArmor: true,
	totemOfWrath: true,
});

export const OtherDefaults = {
	distanceFromTarget: 18,
};

export const PRE_RAID_PRESET = {
	name: 'Pre-raid Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{ "items": [
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
			  24030,
			  24030,
			  0
			]
		  },
		  {
			"id": 24262,
			"enchant": 2748,
			"gems": [
			  24030,
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
			"id": 32387
		  }
	]}`),
};

