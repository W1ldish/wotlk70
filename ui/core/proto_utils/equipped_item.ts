import { GemColor, Suffix } from '../proto/common.js';
import { ItemSlot } from '../proto/common.js';
import { ItemSpec } from '../proto/common.js';
import { ItemType } from '../proto/common.js';
import { Profession } from '../proto/common.js';
import { Stat } from '../proto/common.js';
import {
	UIEnchant as Enchant,
	UIGem as Gem,
	UIItem as Item,
} from '../proto/ui.js';
import { distinct } from '../utils.js';

import { ActionId } from './action_id.js';
import { enchantAppliesToItem } from './utils.js';
import { gemEligibleForSocket, gemMatchesSocket } from './gems.js';
import { Stats } from './stats.js';

export function getWeaponDPS(item: Item): number {
	return ((item.weaponDamageMin + item.weaponDamageMax) / 2) / (item.weaponSpeed || 1);
}

/**
 * Represents an equipped item along with enchants/gems attached to it.
 *
 * This is an immutable type.
 */
export class EquippedItem {
	readonly _item: Item;
	readonly _enchant: Enchant | null;
	readonly _gems: Array<Gem | null>;

	readonly numPossibleSockets: number;

	constructor(item: Item, enchant?: Enchant | null, gems?: Array<Gem | null>) {
		this._item = item;
		this._enchant = enchant || null;
		this._gems = gems || [];
		
		this.numPossibleSockets = this.numSockets();

		// Fill gems with null so we always have the same number of gems as gem slots.
		if (this._gems.length < this.numPossibleSockets) {
			this._gems = this._gems.concat(new Array(this.numPossibleSockets - this._gems.length).fill(null));
		}
	}

	get item(): Item {
		// Make a defensive copy
		return Item.clone(this._item);
	}

	get enchant(): Enchant | null {
		// Make a defensive copy
		return this._enchant ? Enchant.clone(this._enchant) : null;
	}

	get gems(): Array<Gem | null> {
		// Make a defensive copy
		return this._gems.map(gem => gem == null ? null : Gem.clone(gem));
	}

	equals(other: EquippedItem) {
		if (!Item.equals(this._item, other.item))
			return false;

		if ((this._enchant == null) != (other.enchant == null))
			return false;

		if (this._enchant && other.enchant && !Enchant.equals(this._enchant, other.enchant))
			return false;

		if (this._gems.length != other.gems.length)
			return false;

		for (let i = 0; i < this._gems.length; i++) {
			if ((this._gems[i] == null) != (other.gems[i] == null))
				return false;

			if (this._gems[i] && other.gems[i] && !Gem.equals(this._gems[i]!, other.gems[i]!))
				return false;
		}

		return true;
	}

	/**
	 * Replaces the item and tries to keep the existing enchants/gems if possible.
	 */
	withItem(item: Item): EquippedItem {
		let newEnchant = null;
		if (this._enchant && enchantAppliesToItem(this._enchant, item))
			newEnchant = this._enchant;

		// Reorganize gems to match as many colors in the new item as possible.
		const newGems = new Array(item.gemSockets.length).fill(null);
		this._gems.slice(0, this._item.gemSockets.length).filter(gem => gem != null).forEach(gem => {
			const firstMatchingIndex = item.gemSockets.findIndex((socketColor, socketIdx) => !newGems[socketIdx] && gemMatchesSocket(gem!, socketColor));
			const firstEligibleIndex = item.gemSockets.findIndex((socketColor, socketIdx) => !newGems[socketIdx] && gemEligibleForSocket(gem!, socketColor));
			if (firstMatchingIndex != -1) {
				newGems[firstMatchingIndex] = gem;
			} else if (firstEligibleIndex != -1) {
				newGems[firstEligibleIndex] = gem;
			}
		});

		return new EquippedItem(item, newEnchant, newGems);
	}

	/**
	 * Returns a new EquippedItem with the given enchant applied.
	 */
	withEnchant(enchant: Enchant | null): EquippedItem {
		return new EquippedItem(this._item, enchant, this._gems);
	}

	/**
	 * Returns a new EquippedItem with the given gem socketed.
	 */
	private withGemHelper(gem: Gem | null, socketIdx: number): EquippedItem {
		if (this._gems.length <= socketIdx) {
			throw new Error('No gem socket with index ' + socketIdx);
		}

		const newGems = this._gems.slice();
		newGems[socketIdx] = gem;

		return new EquippedItem(this._item, this._enchant, newGems);
	}

	/**
	 * Returns a new EquippedItem with the given gem socketed.
	 *
	 * Also ensures validity of the item on its own. Currently this just means enforcing unique gems.
	 */
	withGem(gem: Gem | null, socketIdx: number): EquippedItem {
		let curItem: EquippedItem | null = this;

		if (gem && gem.unique) {
			curItem = curItem.removeGemsWithId(gem.id);
		}

		return curItem.withGemHelper(gem, socketIdx);
	}

	removeGemsWithId(gemId: number): EquippedItem {
		let curItem: EquippedItem | null = this;
		// Remove any currently socketed identical gems.
		for (let i = 0; i < curItem._gems.length; i++) {
			if (curItem._gems[i]?.id == gemId) {
				curItem = curItem.withGemHelper(null, i);
			}
		}
		return curItem;
	}

	asActionId(): ActionId {
		return ActionId.fromItemId(this._item.id, this.item.suffix);
	}

	asSpec(): ItemSpec {
		return ItemSpec.create({
			id: this._item.id,
			enchant: this._enchant?.effectId,
			gems: this._gems.map(gem => gem?.id || 0),
			suffix: this._item.suffix,
			ivl: this._item.ilvl,
			quality: this._item.quality
		});
	}

	meetsSocketBonus(): boolean {
		return this._item.gemSockets.every((socketColor, i) => this._gems[i] && gemMatchesSocket(this._gems[i]!, socketColor));
	}

	socketBonusStats(): Stats {
		if (this.meetsSocketBonus()) {
			return new Stats(this._item.socketBonus);
		} else {
			return new Stats();
		}
	}

	getPossibleSuffixes(): Suffix[] {
		return this._item.suffixes
	}


	numSockets(): number {
		return this._item.gemSockets.length;
	}

	hasExtraGem(): boolean {
		return this._gems.length > this.item.gemSockets.length;
	}

	allSocketColors(): Array<GemColor> {
		return this._item.gemSockets;
	}
	curSocketColors(): Array<GemColor> {
		return this._item.gemSockets;
	}

	curGems(): Array<Gem> {
		return (this._gems.filter(g => g != null) as Array<Gem>).slice(0, this.numSockets());
	}

	getProfessionRequirements(): Array<Profession> {
		let profs: Array<Profession> = [];
		if (this._item.requiredProfession != Profession.ProfessionUnknown) {
			profs.push(this._item.requiredProfession);
		}
		if (this._enchant != null && this._enchant.requiredProfession != Profession.ProfessionUnknown) {
			profs.push(this._enchant.requiredProfession);
		}
		this._gems.forEach(gem => {
			if (gem != null && gem.requiredProfession != Profession.ProfessionUnknown) {
				profs.push(gem.requiredProfession);
			}
		});
		return distinct(profs);
	}
	getFailedProfessionRequirements(professions: Array<Profession>): Array<Item | Gem | Enchant> {
		let failed: Array<Item | Gem | Enchant> = [];
		if (this._item.requiredProfession != Profession.ProfessionUnknown && !professions.includes(this._item.requiredProfession)) {
			failed.push(this._item);
		}
		if (this._enchant != null && this._enchant.requiredProfession != Profession.ProfessionUnknown && !professions.includes(this._enchant.requiredProfession)) {
			failed.push(this._enchant);
		}
		this._gems.forEach(gem => {
			if (gem != null && gem.requiredProfession != Profession.ProfessionUnknown && !professions.includes(gem.requiredProfession)) {
				failed.push(gem);
			}
		});
		return failed;
	}

	static getSuffixName(suffix: Suffix): string {
		switch(suffix) {
			case Suffix.SuffixMonkey: return "of the Monkey";
			case Suffix.SuffixEagle: return "of the Eagle";
			case Suffix.SuffixBear: return "of the Bear";
			case Suffix.SuffixWhale: return "of the Whale";
			case Suffix.SuffixOwl: return "of the Owl";
			case Suffix.SuffixGorilla: return "of the Gorilla";
			case Suffix.SuffixFalcon: return "of the Falcon";
			case Suffix.SuffixBoar: return "of the Boar";
			case Suffix.SuffixWolf: return "of the Wolf";
			case Suffix.SuffixTiger: return "of the Tiger";
			case Suffix.SuffixSpirit: return "of Spirit";
			case Suffix.SuffixStamina: return "of Stamina";
			case Suffix.SuffixStrength: return "of Strength";
			case Suffix.SuffixAgility: return "of Agility";
			case Suffix.SuffixIntellect: return "of Intellect";
			case Suffix.SuffixPower: return "of Power";
			case Suffix.SuffixArcaneWrath: return "of Arcane Wrath";
			case Suffix.SuffixFieryWrath: return "of Fiery Wrath";
			case Suffix.SuffixFrozenWrath: return "of Frozen Wrath";
			case Suffix.SuffixNaturesWrath: return "of Nature's Wrath";
			case Suffix.SuffixShadowWrath: return "of Shadow Wrath";
			case Suffix.SuffixSpellPower: return "of Spell Power";
			case Suffix.SuffixDefense: return "of Defense";
			case Suffix.SuffixRegeneration: return "of Regeneration";
			case Suffix.SuffixEluding: return "of Eluding";
			case Suffix.SuffixConcentration: return "of Concentration";
			case Suffix.SuffixArcaneProtection: return "of Arcane Protection";
			case Suffix.SuffixFireProtection: return "of Fire Protection";
			case Suffix.SuffixFrostProtection: return "of Frost Protection";
			case Suffix.SuffixNatureProtection: return "of Nature Protection";
			case Suffix.SuffixShadowProtection: return "of Shadow Protection";
			case Suffix.SuffixSorcerer: return "of the Sorcerer";
			case Suffix.SuffixPhysician: return "of the Physician";
			case Suffix.SuffixProphet: return "of the Prophet";
			case Suffix.SuffixInvoker: return "of the Invoker";
			case Suffix.SuffixBandit: return "of the Bandit";
			case Suffix.SuffixBeast: return "of the Beast";
			case Suffix.SuffixHierophant: return "of the Hierophant";
			case Suffix.SuffixSoldier: return "of the Soldier";
			case Suffix.SuffixElder: return "of the Elder";
			case Suffix.SuffixChampion: return "of the Champion";
			case Suffix.SuffixTest: return "of the Test";
			case Suffix.SuffixBlocking: return "of Blocking";
			case Suffix.SuffixPaladinTesting: return "of Paladin Testing";
			case Suffix.SuffixGrove: return "of the Grove";
			case Suffix.SuffixHunt: return "of the Hunt";
			case Suffix.SuffixMind: return "of the Mind";
			case Suffix.SuffixCrusade: return "of the Crusade";
			case Suffix.SuffixVision: return "of the Vision";
			case Suffix.SuffixAncestor: return "of the Ancestor";
			case Suffix.SuffixNightmare: return "of the Nightmare";
			case Suffix.SuffixBattle: return "of the Battle";
			case Suffix.SuffixShadow: return "of the Shadow";
			case Suffix.SuffixSun: return "of the Sun";
			case Suffix.SuffixMoon: return "of the Moon";
			case Suffix.SuffixWild: return "of the Wild";
			case Suffix.SuffixSpellPowerResistance: return "of Spell Power";
			case Suffix.SuffixStrengthResistance: return "of Strength";
			case Suffix.SuffixAgilityResistance: return "of Agility";
			case Suffix.SuffixPowerResistance: return "of Power";
			case Suffix.SuffixMagicResistance: return "of Magic";
			case Suffix.SuffixKnight: return "of the Knight";
			case Suffix.SuffixSeer: return "of the Seer";
			case Suffix.SuffixBear60: return "of the Bear";
			case Suffix.SuffixEagle60: return "of the Eagle";
			case Suffix.SuffixAncestor60: return "of the Ancestor";
			case Suffix.SuffixBandit60: return "of the Bandit";
			case Suffix.SuffixBattle60: return "of the Battle";
			case Suffix.SuffixElder60: return "of the Elder";
			case Suffix.SuffixBeast60: return "of the Beast";
			case Suffix.SuffixChampion60: return "of the Champion";
			case Suffix.SuffixGrove60: return "of the Grove";
			case Suffix.SuffixKnight60: return "of the Knight";
			case Suffix.SuffixMonkey60: return "of the Monkey";
			case Suffix.SuffixMoon60: return "of the Moon";
			case Suffix.SuffixWild60: return "of the Wild";
			case Suffix.SuffixWhale60: return "of the Whale";
			case Suffix.SuffixVision60: return "of the Vision";
			case Suffix.SuffixSun60: return "of the Sun";
			case Suffix.SuffixStamina60: return "of Stamina";
			case Suffix.SuffixSorcerer60: return "of the Sorcerer";
			case Suffix.SuffixSoldier60: return "of the Soldier";
			case Suffix.SuffixShadow60: return "of the Shadow";
			case Suffix.SuffixForeseer: return "of the Foreseer";
			case Suffix.SuffixThief: return "of the Thief";
			case Suffix.SuffixNecromancer: return "of the Necromancer";
			case Suffix.SuffixMarksman: return "of the Marksman";
			case Suffix.SuffixSquire: return "of the Squire";
			case Suffix.SuffixRestoration: return "of Restoration";
			default: return "";
		}
	}
};
