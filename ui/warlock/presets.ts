import {
	Consumes,
	Flask,
	Food,
	PetFood,
	Glyphs,
	EquipmentSpec,
	Potions,
	RaidBuffs,
	IndividualBuffs,
	Debuffs,
	TristateEffect,
	Spec,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import { Player } from '../core/player.js';

import {
	Warlock_Rotation as WarlockRotation,
	Warlock_Options as WarlockOptions,
	Warlock_Rotation_PrimarySpell as PrimarySpell,
	Warlock_Rotation_SecondaryDot as SecondaryDot,
	Warlock_Rotation_SpecSpell as SpecSpell,
	Warlock_Rotation_Curse as Curse,
	Warlock_Rotation_Type as RotationType,
	Warlock_Options_WeaponImbue as WeaponImbue,
	Warlock_Options_Armor as Armor,
	Warlock_Options_Summon as Summon,
	WarlockMajorGlyph as MajorGlyph,
	WarlockMinorGlyph as MinorGlyph,
} from '../core/proto/warlock.js';

import * as WarlockTooltips from './tooltips.js';

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.

export const AfflictionTalents = {
	name: 'Affliction',
	data: SavedTalents.create({
		talentsString: '2350222031123510253500331151--',
		glyphs: Glyphs.create({
			major1: MajorGlyph.GlyphOfHaunt,
			major2: MajorGlyph.GlyphOfLifeTap,
			major3: MajorGlyph.WarlockMajorGlyphNone,
			minor1: MinorGlyph.GlyphOfSouls,
			minor2: MinorGlyph.GlyphOfDrainSoul,
			minor3: MinorGlyph.GlyphOfSubjugateDemon,
		}),
	}),
};

export const DemonologyTalents = {
	name: 'Demonology',
	data: SavedTalents.create({
		talentsString: '03-003203301135112530135201251-05',
		glyphs: Glyphs.create({
			major1: MajorGlyph.GlyphOfLifeTap,
			major2: MajorGlyph.GlyphOfFelguard,
			major3: MajorGlyph.WarlockMajorGlyphNone,
			minor1: MinorGlyph.GlyphOfSouls,
			minor2: MinorGlyph.GlyphOfDrainSoul,
			minor3: MinorGlyph.GlyphOfSubjugateDemon,
		}),
	}),
};

export const DestructionTalents = {
	name: 'Destruction',
	data: SavedTalents.create({
		talentsString: '03-03-05203205220331051035031351',
		glyphs: Glyphs.create({
			major1: MajorGlyph.GlyphOfConflagrate,
			major2: MajorGlyph.GlyphOfLifeTap,
			major3: MajorGlyph.WarlockMajorGlyphNone,
			minor1: MinorGlyph.GlyphOfSouls,
			minor2: MinorGlyph.GlyphOfDrainSoul,
			minor3: MinorGlyph.GlyphOfSubjugateDemon,
		}),
	}),
};

export const AfflictionRotation = WarlockRotation.create({
	primarySpell: PrimarySpell.ShadowBolt,
	secondaryDot: SecondaryDot.UnstableAffliction,
	specSpell: SpecSpell.Haunt,
	curse: Curse.Agony,
	corruption: true,
	useInfernal: false,
	detonateSeed: true,
});

export const DemonologyRotation = WarlockRotation.create({
	primarySpell: PrimarySpell.ShadowBolt,
	secondaryDot: SecondaryDot.Immolate,
	specSpell: SpecSpell.NoSpecSpell,
	curse: Curse.Doom,
	corruption: true,
	useInfernal: false,
	detonateSeed: true,
});

export const DestructionRotation = WarlockRotation.create({
	primarySpell: PrimarySpell.Incinerate,
	secondaryDot: SecondaryDot.Immolate,
	specSpell: SpecSpell.ChaosBolt,
	curse: Curse.Doom,
	corruption: false,
	useInfernal: false,
	detonateSeed: true,
});

export const AfflictionOptions = WarlockOptions.create({
	armor: Armor.FelArmor,
	summon: Summon.Felhunter,
	weaponImbue: WeaponImbue.GrandSpellstone,
});

export const DemonologyOptions = WarlockOptions.create({
	armor: Armor.FelArmor,
	summon: Summon.Felguard,
	weaponImbue: WeaponImbue.GrandSpellstone,
});

export const DestructionOptions = WarlockOptions.create({
	armor: Armor.FelArmor,
	summon: Summon.Imp,
	weaponImbue: WeaponImbue.GrandFirestone,
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskOfPureDeath,
	food: Food.FoodBlackenedBasilisk,
	petFood: PetFood.PetFoodKiblersBits,
	defaultPotion: Potions.DestructionPotion,
	prepopPotion: Potions.DestructionPotion,
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

export const DestroIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	blessingOfMight: TristateEffect.TristateEffectImproved,
});

export const DefaultDebuffs = Debuffs.create({
	sunderArmor: true,
	faerieFire: TristateEffect.TristateEffectImproved,
	bloodFrenzy: true,
	ebonPlaguebringer: true,
	heartOfTheCrusader: true,
	judgementOfWisdom: true,
});

export const DestroDebuffs = Debuffs.create({
	sunderArmor: true,
	faerieFire: TristateEffect.TristateEffectImproved,
	bloodFrenzy: true,
	ebonPlaguebringer: true,
	heartOfTheCrusader: true,
	judgementOfWisdom: true,
	shadowMastery: true,
});

export const OtherDefaults = {
	distanceFromTarget: 25,
};

export const SWP_BIS = {
	name: 'Straight Outa SWP',
	gear: EquipmentSpec.fromJsonString(`{"items": [
		{
			"id": 34340,
			"enchant": 3002,
			"gems": [
				34220,
				32215
			]
		},
		{
			"id": 34204
		},
		{
			"id": 31054,
			"enchant": 2982,
			"gems": [
				32215,
				35760
			]
		},
		{
			"id": 34242,
			"enchant": 2621,
			"gems": [
				32196
			]
		},
		{
			"id": 34364,
			"enchant": 2661,
			"gems": [
				32196,
				35488,
				32196
			]
		},
		{
			"id": 34436,
			"enchant": 2650,
			"gems": [
				35760,
				0
			]
		},
		{
			"id": 34344,
			"enchant": 2937,
			"gems": [
				35760,
				32196,
				0
			]
		},
		{
			"id": 34541,
			"gems": [
				35760,
				0
			]
		},
		{
			"id": 34181,
			"enchant": 2748,
			"gems": [
				32196,
				32196,
				35760
			]
		},
		{
			"id": 34564,
			"enchant": 2940,
			"gems": [
				35760
			]
		},
		{
			"id": 34362,
			"enchant": 2928
		},
		{
			"id": 34230,
			"enchant": 2928
		},
		{
			"id": 32483
		},
		{
			"id": 34429
		},
		{
			"id": 34336,
			"enchant": 2672
		},
		{
			"id": 34179
		},
		{
			"id": 34347,
			"gems": [
				35760
			]
		}
  ]}`),
};
export const P1_PreBiS_11 = {
	name: 'Pre-Raid Affliction',
	tooltip: WarlockTooltips.BIS_TOOLTIP,
	enableWhen: (player: Player<Spec.SpecWarlock>) => player.getRotation().type == RotationType.Affliction,
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
			"id": 31338
		  },
		  {
			"id": 27775,
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
			"id": 28179,
			"gems": [
			  31867,
			  31867
			]
		  },
		  {
			"id": 28227
		  },
		  {
			"id": 31339
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
			"id": 29269
		  },
		  {
			"id": 29350
		  }
  ]}`),
}

