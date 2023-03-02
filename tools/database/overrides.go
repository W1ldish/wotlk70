package database

import (
	"regexp"

	"github.com/Tereneckla/wotlk/sim/core/proto"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

var ItemOverrides = []*proto.UIItem{
	{ /** Destruction Holo-gogs */ Id: 32494, ClassAllowlist: []proto.Class{proto.Class_ClassMage, proto.Class_ClassPriest, proto.Class_ClassWarlock}},
	{ /** Gadgetstorm Goggles */ Id: 32476, ClassAllowlist: []proto.Class{proto.Class_ClassShaman}},
	{ /** Magnified Moon Specs */ Id: 32480, ClassAllowlist: []proto.Class{proto.Class_ClassDruid}},
	{ /** Quad Deathblow X44 Goggles */ Id: 34353, ClassAllowlist: []proto.Class{proto.Class_ClassDruid, proto.Class_ClassRogue}},
	{ /** Hyper-Magnified Moon Specs */ Id: 35182, ClassAllowlist: []proto.Class{proto.Class_ClassDruid}},
	{ /** Lightning Etched Specs */ Id: 34355, ClassAllowlist: []proto.Class{proto.Class_ClassShaman}},
	{ /** Annihilator Holo-Gogs */ Id: 34847, ClassAllowlist: []proto.Class{proto.Class_ClassMage, proto.Class_ClassPriest, proto.Class_ClassWarlock}},

	{Id: 29276, Phase: 2}, //Violet Signet
	{Id: 29277, Phase: 2}, //Violet Signet
	{Id: 29278, Phase: 2}, //Violet Signet
	{Id: 29279, Phase: 2}, //Violet Signet of the Great Protector
	{Id: 29280, Phase: 2}, //Violet Signet
	{Id: 29281, Phase: 2}, //Violet Signet
	{Id: 29282, Phase: 2}, //Violet Signet
	{Id: 29283, Phase: 2}, //Violet Signet of the Master Assassin
	{Id: 29284, Phase: 2}, //Violet Signet
	{Id: 29285, Phase: 2}, //Violet Signet
	{Id: 29286, Phase: 2}, //Violet Signet
	{Id: 29287, Phase: 2}, //Violet Signet of the Archmage
	{Id: 29288, Phase: 2}, //Violet Signet
	{Id: 29289, Phase: 2}, //Violet Signet
	{Id: 29290, Phase: 2}, //Violet Signet of the Grand Restorer
	{Id: 29291, Phase: 2}, //Violet Signet
	{Id: 33204, Phase: 2}, //Shadowprowler's Chestguard
	{Id: 33122, Phase: 2}, //Cloak of Darkness

	{Id: 29294, Phase: 4}, //Band of Eternity
	{Id: 29295, Phase: 4}, //Band of Eternity
	{Id: 29296, Phase: 4}, //Band of Eternity
	{Id: 29297, Phase: 4}, //Band of the Eternal Defender
	{Id: 29298, Phase: 4}, //Band of Eternity
	{Id: 29299, Phase: 4}, //Band of Eternity
	{Id: 29300, Phase: 4}, //Band of Eternity
	{Id: 29301, Phase: 4}, //Band of the Eternal Champion
	{Id: 29302, Phase: 4}, //Band of Eternity
	{Id: 29303, Phase: 4}, //Band of Eternity
	{Id: 29304, Phase: 4}, //Band of Eternity
	{Id: 29305, Phase: 4}, //Band of the Eternal Sage
	{Id: 29306, Phase: 4}, //Band of Eternity
	{Id: 29307, Phase: 4}, //Band of Eternity
	{Id: 29308, Phase: 4}, //Band of Eternity
	{Id: 29309, Phase: 4}, //Band of the Eternal Restorer

	{Id: 33478, Phase: 5}, //Jin'rohk, The Great Apocalypse
	{Id: 33640, Phase: 5}, //Fury
	{Id: 33492, Phase: 5}, //Trollbane
	{Id: 33465, Phase: 5}, //Staff of Primal Fury
	{Id: 33468, Phase: 5}, //Dark Blessing
	{Id: 33490, Phase: 5}, //Staff of Dark Mending
	{Id: 33495, Phase: 5}, //Rage
	{Id: 33298, Phase: 5}, //Prowler's Strikeblade
	{Id: 33467, Phase: 5}, //Blade of Twisted Visions
	{Id: 33474, Phase: 5}, //Ancient Amani Longbow
	{Id: 33354, Phase: 5}, //Wub's Cursed Hexblade
	{Id: 33388, Phase: 5}, //Heartless
	{Id: 33476, Phase: 5}, //Cleaver of the Unforgiving
	{Id: 33389, Phase: 5}, //Dagger of Bad Mojo
	{Id: 33493, Phase: 5}, //Umbral Shiv
	{Id: 33494, Phase: 5}, //Amani Divining Staff
	{Id: 33214, Phase: 5}, //Akil'zon's Talonblade
	{Id: 33283, Phase: 5}, //Amani Punisher
	{Id: 33491, Phase: 5}, //Tuskbreaker
	{Id: 33191, Phase: 5}, //Jungle Stompers
	{Id: 33499, Phase: 5}, //Signet of the Last Defender
	{Id: 33481, Phase: 5}, //Pauldrons of Stone Resolve
	{Id: 33831, Phase: 5}, //Berserker's Call
	{Id: 34029, Phase: 5}, //Tiny Voodoo Mask
	{Id: 33432, Phase: 5}, //Coif of the Jungle Stalker
	{Id: 33471, Phase: 5}, //Two-toed Sandals
	{Id: 33497, Phase: 5}, //Mana Attuned Band
	{Id: 33215, Phase: 5}, //Bloodstained Elven Battlevest
	{Id: 33281, Phase: 5}, //Brooch of Nature's Mercy
	{Id: 33464, Phase: 5}, //Hex Lord's Voodoo Pauldrons
	{Id: 33469, Phase: 5}, //Hauberk of the Empire's Champion
	{Id: 33206, Phase: 5}, //Pauldrons of Primal Fury
	{Id: 33328, Phase: 5}, //Arrow-fall Chestguard
	{Id: 33500, Phase: 5}, //Signet of Eternal Life
	{Id: 33286, Phase: 5}, //Mojo-mender's Mask
	{Id: 33327, Phase: 5}, //Mask of Introspection
	{Id: 33332, Phase: 5}, //Enamelled Disc of Mojo
	{Id: 33421, Phase: 5}, //Battleworn Tuskguard
	{Id: 33805, Phase: 5}, //Shadowhunter's Treads
	{Id: 33829, Phase: 5}, //Hex Shrunken Head
	{Id: 33326, Phase: 5}, //Bulwark of the Amani Empire
	{Id: 33356, Phase: 5}, //Helm of Natural Regeneration
	{Id: 33480, Phase: 5}, //Cord of Braided Troll Hair
	{Id: 33489, Phase: 5}, //Mantle of Ill Intent
	{Id: 33533, Phase: 5}, //Avalanche Leggings
	{Id: 33203, Phase: 5}, //Robes of Heavenly Purpose
	{Id: 33211, Phase: 5}, //Bladeangel's Money Belt
	{Id: 33285, Phase: 5}, //Fury of the Ursine
	{Id: 33303, Phase: 5}, //Skullshatter Warboots
	{Id: 33446, Phase: 5}, //Girdle of Stromgarde's Hope
	{Id: 33453, Phase: 5}, //Hood of Hexing
	{Id: 33498, Phase: 5}, //Signet of the Quiet Forest
	{Id: 33830, Phase: 5}, //Ancient Aqir Artifact
	{Id: 33216, Phase: 5}, //Chestguard of Hidden Purpose
	{Id: 33293, Phase: 5}, //Signet of Ancient Magics
	{Id: 33297, Phase: 5}, //The Savage's Choker
	{Id: 33299, Phase: 5}, //Spaulders of the Advocate
	{Id: 33300, Phase: 5}, //Shoulderpads of Dancing Blades
	{Id: 33317, Phase: 5}, //Robe of Departed Spirits
	{Id: 33322, Phase: 5}, //Shimmer-pelt Vest
	{Id: 33329, Phase: 5}, //Shadowtooth Trollskin Cuirass
	{Id: 33357, Phase: 5}, //Footpads of Madness
	{Id: 33463, Phase: 5}, //Hood of the Third Eye
	{Id: 33466, Phase: 5}, //Loop of Cursed Bones
	{Id: 33473, Phase: 5}, //Chestguard of the Warlord
	{Id: 33479, Phase: 5}, //Grimgrin Faceguard
	{Id: 33483, Phase: 5}, //Life-step Belt
	{Id: 33496, Phase: 5}, //Signet of Primal Wrath
	{Id: 33590, Phase: 5}, //Cloak of Fiends
	{Id: 33591, Phase: 5}, //Shadowcaster's Drape
	{Id: 33592, Phase: 5}, //Cloak of Ancient Rituals
	{Id: 33828, Phase: 5}, //Tome of Diabolic Remedy
	{Id: 33971, Phase: 5}, //Elunite Imbued Leggings
	{Id: 33207, Phase: 5}, //Implacable Guardian Sabatons
	{Id: 33222, Phase: 5}, //Nyn'jah's Tabi Boots
	{Id: 33279, Phase: 5}, //Iron-tusk Girdle
	{Id: 33280, Phase: 5}, //War-Feathered Loop
	{Id: 33287, Phase: 5}, //Gnarled Ironwood Pauldrons
	{Id: 33291, Phase: 5}, //Voodoo-woven Belt
	{Id: 33296, Phase: 5}, //Brooch of Deftness
	{Id: 33324, Phase: 5}, //Treads of the Life Path
	{Id: 33325, Phase: 5}, //Voodoo Shaker
	{Id: 33331, Phase: 5}, //Chain of Unleashed Rage
	{Id: 33334, Phase: 5}, //Fetish of the Primal Gods
	{Id: 33386, Phase: 5}, //Man'kin'do's Belt
	{Id: 33501, Phase: 5}, //Bloodthirster's Wargreaves
	{Id: 33502, Phase: 5}, //Libram of Mending
	{Id: 33503, Phase: 5}, //Libram of Divine Judgement
	{Id: 33504, Phase: 5}, //Libram of Divine Purpose
	{Id: 33505, Phase: 5}, //Totem of Living Water
	{Id: 33506, Phase: 5}, //Skycall Totem
	{Id: 33507, Phase: 5}, //Stonebreaker's Totem
	{Id: 33508, Phase: 5}, //Idol of Budding Life
	{Id: 33509, Phase: 5}, //Idol of Terror
	{Id: 33510, Phase: 5}, //Idol of the Unseen Moon
	{Id: 33512, Phase: 5}, //Furious Deathgrips
	{Id: 33513, Phase: 5}, //Eternium Rage-shackles
	{Id: 33514, Phase: 5}, //Pauldrons of Gruesome Fate
	{Id: 33515, Phase: 5}, //Unwavering Legguards
	{Id: 33516, Phase: 5}, //Bracers of the Ancient Phalanx
	{Id: 33517, Phase: 5}, //Bonefist Gauntlets
	{Id: 33518, Phase: 5}, //High Justicar's Legplates
	{Id: 33519, Phase: 5}, //Handguards of the Templar
	{Id: 33520, Phase: 5}, //Vambraces of the Naaru
	{Id: 33522, Phase: 5}, //Chestguard of the Stoic Guardian
	{Id: 33523, Phase: 5}, //Sabatons of the Righteous Defender
	{Id: 33524, Phase: 5}, //Girdle of the Protector
	{Id: 33527, Phase: 5}, //Shifting Camouflage Pants
	{Id: 33528, Phase: 5}, //Gauntlets of Sniping
	{Id: 33529, Phase: 5}, //Steadying Bracers
	{Id: 33530, Phase: 5}, //Natural Life Leggings
	{Id: 33531, Phase: 5}, //Polished Waterscale Gloves
	{Id: 33532, Phase: 5}, //Gleaming Earthen Bracers
	{Id: 33534, Phase: 5}, //Grips of Nature's Wrath
	{Id: 33535, Phase: 5}, //Earthquake Bracers
	{Id: 33536, Phase: 5}, //Stormwrap
	{Id: 33537, Phase: 5}, //Treads of Booming Thunder
	{Id: 33538, Phase: 5}, //Shallow-grave Trousers
	{Id: 33539, Phase: 5}, //Trickster's Stickyfingers
	{Id: 33540, Phase: 5}, //Master Assassin Wristguards
	{Id: 33552, Phase: 5}, //Pants of Splendid Recovery
	{Id: 33557, Phase: 5}, //Gargon's Bracers of Peaceful Slumber
	{Id: 33559, Phase: 5}, //Starfire Waistband
	{Id: 33566, Phase: 5}, //Blessed Elunite Coverings
	{Id: 33577, Phase: 5}, //Moon-walkers
	{Id: 33578, Phase: 5}, //Armwraps of the Kaldorei Protector
	{Id: 33579, Phase: 5}, //Vestments of Hibernation
	{Id: 33580, Phase: 5}, //Band of the Swift Paw
	{Id: 33582, Phase: 5}, //Footwraps of Wild Encroachment
	{Id: 33583, Phase: 5}, //Waistguard of the Great Beast
	{Id: 33584, Phase: 5}, //Pantaloons of Arcane Annihilation
	{Id: 33585, Phase: 5}, //Achromic Trousers of the Naaru
	{Id: 33586, Phase: 5}, //Studious Wraps
	{Id: 33587, Phase: 5}, //Light-Blessed Bonds
	{Id: 33588, Phase: 5}, //Runed Spell-cuffs
	{Id: 33589, Phase: 5}, //Wristguards of Tranquil Thought
	{Id: 33593, Phase: 5}, //Slikk's Cloak of Placation
	{Id: 33810, Phase: 5}, //Amani Mask of Death
	{Id: 33965, Phase: 5}, //Hauberk of the Furious Elements
	{Id: 33970, Phase: 5}, //Pauldrons of the Furious Elements
	{Id: 33972, Phase: 5}, //Mask of Primal Power
	{Id: 33973, Phase: 5}, //Pauldrons of Tribal Fury
	{Id: 33974, Phase: 5}, //Grasp of the Moonkin
	{Id: 33192, Phase: 5}, //Carved Witch Doctor's Stick
	{Id: 33832, Phase: 5}, //Battlemaster's Determination
	{Id: 34049, Phase: 5}, //Battlemaster's Audacity
	{Id: 34050, Phase: 5}, //Battlemaster's Perserverance
	{Id: 34162, Phase: 5}, //Battlemaster's Depravity
	{Id: 34163, Phase: 5}, //Battlemaster's Cruelty
	{Id: 35326, Phase: 5}, //Battlemaster's Alacrity
	{Id: 33304, Phase: 5}, //Cloak of Subjugated Power
	{Id: 33333, Phase: 5}, //Kharmaa's Shroud of Hope
	{Id: 33484, Phase: 5}, //Dory's Embrace
	{Id: 35321, Phase: 5}, //Cloak of Arcane Alacrity
	{Id: 35324, Phase: 5}, //Cloak of Swift Reprieve

	{Id: 34678, Phase: 6}, //Shattered Sun Pendant of Acumen
	{Id: 34679, Phase: 6}, //Shattered Sun Pendant of Might
	{Id: 34680, Phase: 6}, //Shattered Sun Pendant of Resolve
	{Id: 34677, Phase: 6}, //Shattered Sun Pendant of Restoration
	{Id: 34675, Phase: 6}, //Sunward Crest
	{Id: 34676, Phase: 6}, //Dawnforged Defender
	{Id: 34665, Phase: 6}, //Bombardier's Blade
	{Id: 34666, Phase: 6}, //The Sunbreaker
	{Id: 34667, Phase: 6}, //Archmage's Guile
	{Id: 34670, Phase: 6}, //Seeker's Gavel
	{Id: 34671, Phase: 6}, //K'iru's Presage
	{Id: 34672, Phase: 6}, //Inuuro's Blade
	{Id: 34673, Phase: 6}, //Legionfoe
	{Id: 34674, Phase: 6}, //Truestrike Crossbow
	{Id: 35693, Phase: 6}, //Figurine - Empyrean Tortoise
	{Id: 35694, Phase: 6}, //Figurine - Khorium Boar
	{Id: 35700, Phase: 6}, //Figurine - Crimson Serpent
	{Id: 35702, Phase: 6}, //Figurine - Shadowsong Panther
	{Id: 35703, Phase: 6}, //Figurine - Seaspray Albatross
	{Id: 35748, Phase: 6}, //Guardian's Alchemist Stone
	{Id: 35749, Phase: 6}, //Sorcerer's Alchemist Stone
	{Id: 35750, Phase: 6}, //Redeemer's Alchemist Stone
	{Id: 35751, Phase: 6}, //Assassin's Alchemist Stone
	{Id: 34470, Phase: 6}, //Timbal's Focusing Crystal
	{Id: 34471, Phase: 6}, //Vial of the Sunwell
	{Id: 34472, Phase: 6}, //Shard of Contempt
	{Id: 34473, Phase: 6}, //Commendation of Kael'thas
	{Id: 34601, Phase: 6}, //Shoulderplates of Everlasting Pain
	{Id: 34602, Phase: 6}, //Eversong Cuffs
	{Id: 34603, Phase: 6}, //Distracting Blades
	{Id: 34604, Phase: 6}, //Jaded Crystal Dagger
	{Id: 34605, Phase: 6}, //Breastplate of Fierce Survival
	{Id: 34606, Phase: 6}, //Edge of Oppression
	{Id: 34607, Phase: 6}, //Fel-tinged Mantle
	{Id: 34608, Phase: 6}, //Rod of the Blazing Light
	{Id: 34609, Phase: 6}, //Quickening Blade of the Prince
	{Id: 34610, Phase: 6}, //Scarlet Sin'dorei Robes
	{Id: 34611, Phase: 6}, //Cudgel of Consecration
	{Id: 34612, Phase: 6}, //Greaves of the Penitent Knight
	{Id: 34613, Phase: 6}, //Shoulderpads of the Silvermoon Retainer
	{Id: 34614, Phase: 6}, //Tunic of the Ranger Lord
	{Id: 34615, Phase: 6}, //Netherforce Chestplate
	{Id: 34616, Phase: 6}, //Breeching comet
	{Id: 34625, Phase: 6}, //Kharmaa's Ring of Fate
	{Id: 34697, Phase: 6}, //Bindings of Raging Fire
	{Id: 34698, Phase: 6}, //Bracers of the Forest Stalker
	{Id: 34699, Phase: 6}, //Sun-forged Cleaver
	{Id: 34700, Phase: 6}, //Gauntlets of Divine Blessings
	{Id: 34701, Phase: 6}, //Leggings of the Betrayed
	{Id: 34702, Phase: 6}, //Cloak of Swift Mending
	{Id: 34703, Phase: 6}, //Lantro's Dancing Blade
	{Id: 34704, Phase: 6}, //Band of Arcane Alacrity
	{Id: 34705, Phase: 6}, //Bracers of Divine Infusion
	{Id: 34706, Phase: 6}, //Band of Determination
	{Id: 34707, Phase: 6}, //Boots of Resuscitation
	{Id: 34708, Phase: 6}, //Cloak of the Coming Night
	{Id: 34783, Phase: 6}, //Nightstrike
	{Id: 34788, Phase: 6}, //Duskhallow Mantle
	{Id: 34789, Phase: 6}, //Bracers of Slaughter
	{Id: 34790, Phase: 6}, //Battle-mace of the High Priestess
	{Id: 34791, Phase: 6}, //Gauntlets of the Tranquil Waves
	{Id: 34792, Phase: 6}, //Cloak of the Betrayed
	{Id: 34793, Phase: 6}, //Cord of Reconstruction
	{Id: 34794, Phase: 6}, //Axe of Shattered Dreams
	{Id: 34795, Phase: 6}, //Helm of Sanctification
	{Id: 34796, Phase: 6}, //Robes of Summer Flame
	{Id: 34797, Phase: 6}, //Sun-infused Focus Staff
	{Id: 34798, Phase: 6}, //Band of Celerity
	{Id: 34799, Phase: 6}, //Hauberk of the War Bringer
	{Id: 34807, Phase: 6}, //Sunstrider Warboots
	{Id: 34808, Phase: 6}, //Gloves of Arcane Acuity
	{Id: 34809, Phase: 6}, //Sunrage Treads
	{Id: 34810, Phase: 6}, //Cloak of Blade Turning

	{Id: 34891, Phase: 6}, //The Blade of Harbingers
	{Id: 34892, Phase: 6}, //Crossbow of Relentless Strikes
	{Id: 34893, Phase: 6}, //Vanir's Right Fist of Brutality
	{Id: 34894, Phase: 6}, //Blade of Serration
	{Id: 34895, Phase: 6}, //Scryer's Blade of Focus
	{Id: 34896, Phase: 6}, //Gavel of Naaru Blessings
	{Id: 34898, Phase: 6}, //Staff of the Forest Lord
	{Id: 34949, Phase: 6}, //Swift Blade of Uncertainty
	{Id: 34950, Phase: 6}, //Vanir's Left Fist of Savagery
	{Id: 34951, Phase: 6}, //Vanir's Left Fist of Brutality
	{Id: 34952, Phase: 6}, //The Mutilator
	{Id: 34887, Phase: 6}, //Angelista's Revenge
	{Id: 34888, Phase: 6}, //Ring of the Stalwart Protector
	{Id: 34889, Phase: 6}, //Fused Nethergon Band
	{Id: 34890, Phase: 6}, //Anveena's Touch
	{Id: 34900, Phase: 6}, //Shroud of Nature's Harmony
	{Id: 34901, Phase: 6}, //Grovewalker's Leggings
	{Id: 34902, Phase: 6}, //Oakleaf-Spun Handguards
	{Id: 34903, Phase: 6}, //Embrace of Starlight
	{Id: 34904, Phase: 6}, //Barbed Gloves of the Sage
	{Id: 34905, Phase: 6}, //Crystalwind Leggings
	{Id: 34906, Phase: 6}, //Embrace of Everlasting Prowess
	{Id: 34910, Phase: 6}, //Tameless Breeches
	{Id: 34911, Phase: 6}, //Handwraps of the Aggressor
	{Id: 34912, Phase: 6}, //Scaled Drakeskin Chestguard
	{Id: 34914, Phase: 6}, //Leggings of the Pursuit
	{Id: 34916, Phase: 6}, //Gauntlets of Rapidity
	{Id: 34917, Phase: 6}, //Shroud of the Lore`nial
	{Id: 34918, Phase: 6}, //Legwraps of Sweltering Flame
	{Id: 34919, Phase: 6}, //Boots of Incantations
	{Id: 34921, Phase: 6}, //Ecclesiastical Cuirass
	{Id: 34922, Phase: 6}, //Greaves of Pacification
	{Id: 34923, Phase: 6}, //Waistguard of Reparation
	{Id: 34924, Phase: 6}, //Gown of Spiritual Wonder
	{Id: 34925, Phase: 6}, //Adorned Supernal Legwraps
	{Id: 34926, Phase: 6}, //Slippers of Dutiful Mending
	{Id: 34927, Phase: 6}, //Tunic of the Dark Hour
	{Id: 34928, Phase: 6}, //Trousers of the Scryers' Retainer
	{Id: 34929, Phase: 6}, //Belt of the Silent Path
	{Id: 34930, Phase: 6}, //Wae of Life Chestguard
	{Id: 34931, Phase: 6}, //Runed Scales of Antiquity
	{Id: 34932, Phase: 6}, //Clutch of the Soothing Breeze
	{Id: 34933, Phase: 6}, //Hauberk of Whirling Fury
	{Id: 34934, Phase: 6}, //Rushing Storm Kilt
	{Id: 34935, Phase: 6}, //Aftershock Waistguard
	{Id: 34936, Phase: 6}, //Tormented Demonsoul Robes
	{Id: 34937, Phase: 6}, //Corrupted Soulcloth Pantaloons
	{Id: 34938, Phase: 6}, //Enslaved Doomguard Soulgrips
	{Id: 34939, Phase: 6}, //Chestplate of Stoicism
	{Id: 34940, Phase: 6}, //Sunguard Legplates
	{Id: 34941, Phase: 6}, //Girdle of the Fearless
	{Id: 34942, Phase: 6}, //Breastplate of Ire
	{Id: 34943, Phase: 6}, //Legplates of Unending Fury
	{Id: 34944, Phase: 6}, //Girdle of Seething Rage
	{Id: 34945, Phase: 6}, //Shattrath Protectorate's Breastplate
	{Id: 34946, Phase: 6}, //Inscribed Legplates of the Aldor
	{Id: 34947, Phase: 6}, //Blue's Greaves of the Righteous Guardian

}

// Keep these sorted by item ID.
var ItemAllowList = map[int32]struct{}{
	11815: {}, // Hand of Justice
	12590: {}, // Felstriker
	15808: {}, // Fine Light Crossbow (for hunter testing).
	18843: {},
	18844: {},
	18847: {},
	18848: {},
	19019: {}, // Thunderfury
	19808: {}, // Rockhide Strongfish
	20837: {}, // Sunstrider Axe
	20966: {}, // Jade Pendant of Blasting
	21625: {}, // Scarab Brooch
	24114: {}, // Braided Eternium Chain
	28572: {}, // Blade of the Unrequited
	28830: {}, // Dragonspine Trophy
	29383: {}, // Bloodlust Brooch
	29387: {}, // Gnomeregan Auto-Blocker 600
	29994: {}, // Thalassian Wildercloak
	29996: {}, // Rod of the Sun King
	30032: {}, // Red Belt of Battle
	30627: {}, // Tsunami Talisman
	30720: {}, // Serpent-Coil Braid
	31193: {}, // Blade of Unquenched Thirst
	32387: {}, // Idol of the Raven Goddess
	32658: {}, // Badge of Tenacity
	33135: {}, // Falling Star
	33140: {}, // Blood of Amber
	33143: {}, // Stone of Blades
	33144: {}, // Facet of Eternity
	33504: {}, // Libram of Divine Purpose
	33506: {}, // Skycall Totem
	33507: {}, // Stonebreaker's Totem
	33508: {}, // Idol of Budding Life
	33510: {}, // Unseen moon idol
	33829: {}, // Hex Shrunken Head
	33831: {}, // Berserkers Call
	34472: {}, // Shard of Contempt
	34473: {}, // Commendation of Kael'thas

	// Sets
	27510: {}, // Tidefury Gauntlets
	27802: {}, // Tidefury Shoulderguards
	27909: {}, // Tidefury Kilt
	28231: {}, // Tidefury Chestpiece
	28349: {}, // Tidefury Helm

	15056: {}, // Stormshroud Armor
	15057: {}, // Stormshroud Pants
	15058: {}, // Stormshroud Shoulders
	21278: {}, // Stormshroud Gloves
}

// Keep these sorted by item ID.
var ItemDenyList = map[int32]struct{}{
	17782: {}, // talisman of the binding shard
	17783: {}, // talisman of the binding fragment
	17802: {}, // Deprecated version of Thunderfury
	18582: {},
	18583: {},
	18584: {},
	24265: {},
	32384: {},
	32421: {},
	32422: {},
	33482: {},
	33350: {},
	34576: {}, // Battlemaster's Cruelty
	34577: {}, // Battlemaster's Depreavity
	34578: {}, // Battlemaster's Determination
	34579: {}, // Battlemaster's Audacity
	34580: {}, // Battlemaster's Perseverence
}

// Item icons to include in the DB, so they don't need to be separately loaded in the UI.
var ExtraItemIcons = []int32{
	// Pet foods
	33874,
	43005,

	// Spellstones
	41174,
	41196,

	// Demonic Rune
	12662,

	// Food IDs
	27655,
	27657,
	27658,
	27664,
	33052,
	33825,
	33872,
	34753,
	34754,
	34756,
	34758,
	34767,
	34769,
	42994,
	42995,
	42996,
	42998,
	42999,
	43000,
	43015,

	// Flask IDs
	13511,
	13512,
	22851,
	22853,
	22854,
	22861,
	22866,
	33208,
	40079,
	44939,
	46376,
	46377,
	46378,
	46379,

	// Elixer IDs
	40072,
	40078,
	40097,
	40109,
	44328,
	44332,

	// Elixer IDs
	13452,
	13454,
	22824,
	22827,
	22831,
	22833,
	22834,
	22835,
	22840,
	28103,
	28104,
	31679,
	32062,
	32067,
	32068,
	39666,
	40068,
	40070,
	40073,
	40076,
	44325,
	44327,
	44329,
	44330,
	44331,
	9088,
	9224,
	22825,

	// Potions / In Battle Consumes
	13442,
	20520,
	22105,
	22788,
	22828,
	22832,
	22837,
	22838,
	22839,
	22849,
	31677,
	33447,
	33448,
	36892,
	40093,
	40211,
	40212,
	40536,
	40771,
	41119,
	41166,
	42545,
	23827,

	// Poisons
	43231,
	43233,
	43235,

	// Thistle Tea
	7676,

	// Scrolls
	33461,
	33462,
	33457,
	33458,
	33460,
	33459,
}

// Raid buffs / debuffs
var SharedSpellsIcons = []int32{
	// Revitalize, Rejuv, WG
	48545,
	26982,
	53251,

	// Registered CD's
	49016,
	57933,
	64382,
	10060,
	16190,
	29166,
	53530,
	33206,
	2825,
	54758,

	// Raid Buffs
	27127,
	57566,
	54038,

	26991,
	17051,

	25898,
	25899,

	27149,
	20140,
	25509,
	16293,

	25389,
	14767,

	25528,
	52456,
	57330,

	25312,

	27141,
	20045,
	2048,

	53138,
	30809,
	19506,

	31869,
	31583,
	34460,

	57472,
	50720,

	53648,

	469,
	12861,
	27268,
	18696,

	27143,
	20245,
	25570,
	16206,

	17007,
	34300,
	29801,

	55610,
	65990,
	29193,

	48160,
	31878,
	53292,
	54118,
	44561,

	24907,
	48396,
	51470,

	3738,
	47240,
	57721,
	25557,

	27150,
	39374,
	31025,
	31035,
	6562,
	31033,
	26992,
	16840,
	54648,

	// Raid Debuffs
	8647,
	25225,
	55754,

	770,
	33602,
	30909,
	18180,
	56631,
	53598,

	26016,
	25203,
	12879,
	48560,
	16862,
	55487,

	33876,
	46855,
	57393,

	30706,
	20337,
	58410,

	25264,
	12666,
	55095,
	51456,
	53696,
	48485,

	3043,
	29859,
	58413,
	65855,

	17800,
	17803,
	12873,
	28593,

	33198,
	51161,
	48511,
	27228,

	20271,
	53408,

	11374,
}

// If any of these match the item name, don't include it.
var DenyListNameRegexes = []*regexp.Regexp{
	regexp.MustCompile(`30 Epic`),
	regexp.MustCompile(`63 Blue`),
	regexp.MustCompile(`63 Green`),
	regexp.MustCompile(`66 Epic`),
	regexp.MustCompile(`90 Epic`),
	regexp.MustCompile(`90 Green`),
	regexp.MustCompile(`Boots 1`),
	regexp.MustCompile(`Boots 2`),
	regexp.MustCompile(`Boots 3`),
	regexp.MustCompile(`Bracer 1`),
	regexp.MustCompile(`Bracer 2`),
	regexp.MustCompile(`Bracer 3`),
	regexp.MustCompile(`DB\d`),
	regexp.MustCompile(`DEPRECATED`),
	regexp.MustCompile(`Deprecated: Keanna`),
	regexp.MustCompile(`Indalamar`),
	regexp.MustCompile(`Monster -`),
	regexp.MustCompile(`NEW`),
	regexp.MustCompile(`PH`),
	regexp.MustCompile(`QR XXXX`),
	regexp.MustCompile(`TEST`),
	regexp.MustCompile(`Test`),
	regexp.MustCompile(`zOLD`),
}

// Allows manual overriding for Gem fields in case WowHead is wrong.
var GemOverrides = []*proto.UIGem{
	{Id: 33131, Stats: stats.Stats{stats.AttackPower: 32, stats.RangedAttackPower: 32}.ToFloatArray()},
	{Id: 35707, Phase: 6}, //Regal Nightseye
	{Id: 35503, Phase: 6}, //Ember Skyfire Diamond
	{Id: 35501, Phase: 6}, //Eternal Earthstorm Diamond
	{Id: 25896, Phase: 2},
	{Id: 25897, Phase: 2},
	{Id: 47055, Phase: 6}, //Reckless Pyrestone
	{Id: 35761, Phase: 6}, //Quick Lionseye
	{Id: 35315, Phase: 6}, //Quick Dawnstone
	{Id: 35759, Phase: 6}, //Forceful Seaspray Emerald
	{Id: 37503, Phase: 6}, //Purified Shadowsong Amethyst
	{Id: 35316, Phase: 6}, //Reckless Noble Topaz
	{Id: 35318, Phase: 6}, //Forceful Talasite
}
var GemDenyList = map[int32]struct{}{
	// pvp non-unique gems not in game currently.
	32735: {},
	35489: {},
	38545: {},
	38546: {},
	38547: {},
	38548: {},
	38549: {},
	38550: {},
}
