import { IndividualSimUI } from "../../individual_sim_ui";
import {
  BattleElixir,
  Class,
  Conjured,
  Flask,
  Food,
  GuardianElixir,
  Potions,
  Profession,
  Spec,
  Stat,
  WeaponImbue
} from "../../proto/common";
import { Component } from "../component";
import { IconEnumPicker } from "../icon_enum_picker";

import * as IconInputs from '../icon_inputs.js';
import { buildIconInput } from "../icon_inputs.js";
import { SettingsTab } from "./settings_tab";

export class ConsumesPicker extends Component {
  protected settingsTab: SettingsTab;
  protected simUI: IndividualSimUI<Spec>;

  constructor(parentElem: HTMLElement, settingsTab: SettingsTab, simUI: IndividualSimUI<Spec>) {
    super(parentElem, 'consumes-picker-root');
    this.settingsTab = settingsTab;
    this.simUI = simUI;

    this.buildPotionsPicker();
    this.buildElixirsPicker();
    this.buildFoodPicker();
    if (!((this.simUI.player.getClass() == Class.ClassRogue) || (this.simUI.player.getClass() == Class.ClassShaman))) {
      this.buildWeaponImbuePicker();
    }
    this.buildEngPicker();
    this.buildPetPicker();
  }

  private buildPotionsPicker() {
    let fragment = document.createElement('fragment');
    fragment.innerHTML = `
      <div class="consumes-row input-root input-inline">
        <label class="form-label">Potions</label>
        <div class="consumes-row-inputs">
          <div class="consumes-prepot"></div>
          <div class="consumes-potions"></div>
          <div class="consumes-conjured"></div>
        </div>
      </div>
    `;

    this.rootElem.appendChild(fragment.children[0] as HTMLElement);

    const prepopPotionOptions = this.simUI.splitRelevantOptions([
			// This list is smaller because some potions don't make sense to use as prepot.
			// E.g. healing/mana potions.
			{ item: Potions.DestructionPotion, stats: [Stat.StatSpellCrit,Stat.StatSpellPower] },
			{ item: Potions.HeroicPotion, stats: [Stat.StatStrength] },
			{ item: Potions.HastePotion, stats: [Stat.StatMeleeHaste, Stat.StatSpellHaste] },
      { item: Potions.IronshieldPotion, stats: [Stat.StatArmor] },
      { item: Potions.InsaneStrengthPotion, stats: [Stat.StatStrength] },
		]);
		if (prepopPotionOptions.length) {
			const elem = this.rootElem.querySelector('.consumes-prepot') as HTMLElement;
			new IconEnumPicker(
        elem,
        this.simUI.player,
        IconInputs.makePrepopPotionsInput(prepopPotionOptions, 'Prepop Potion (1s before combat)')
      );
    }

		const potionOptions = this.simUI.splitRelevantOptions([
			{ item: Potions.SuperManaPotion, stats: [Stat.StatIntellect] },
			{ item: Potions.IronshieldPotion, stats: [Stat.StatArmor] },
			{ item: Potions.HeroicPotion, stats: [Stat.StatStrength] },
			{ item: Potions.HastePotion, stats: [Stat.StatMeleeHaste, Stat.StatSpellHaste] },
			{ item: Potions.DestructionPotion, stats: [Stat.StatSpellCrit, Stat.StatSpellPower] },
      { item: Potions.InsaneStrengthPotion, stats: [Stat.StatStrength] },
      { item: Potions.FelManaPotion, stats: [Stat.StatIntellect]},
		]);
		if (potionOptions.length) {
			const elem = this.rootElem.querySelector('.consumes-potions') as HTMLElement;
			new IconEnumPicker(
        elem,
        this.simUI.player,
        IconInputs.makePotionsInput(potionOptions, 'Combat Potion')
      );
		}

		const conjuredOptions = this.simUI.splitRelevantOptions([
			this.simUI.player.getClass() == Class.ClassRogue ? { item: Conjured.ConjuredRogueThistleTea, stats: [] } : null,
			{ item: Conjured.ConjuredHealthstone, stats: [Stat.StatStamina] },
			{ item: Conjured.ConjuredDarkRune, stats: [Stat.StatIntellect] },
			{ item: Conjured.ConjuredFlameCap, stats: [] },
		]);
		if (conjuredOptions.length) {
			const elem = this.rootElem.querySelector('.consumes-conjured') as HTMLElement;
			new IconEnumPicker(elem, this.simUI.player, IconInputs.makeConjuredInput(conjuredOptions));
		}
  }

  private buildElixirsPicker() {
    let fragment = document.createElement('fragment');
    fragment.innerHTML = `
      <div class="consumes-row input-root input-inline">
        <label class="form-label">Elixirs</label>
        <div class="consumes-row-inputs">
          <div class="consumes-flasks"></div>
          <span class="elixir-space">or</span>
          <div class="consumes-battle-elixirs"></div>
          <div class="consumes-guardian-elixirs"></div>
        </div>
      </div>
    `;

    this.rootElem.appendChild(fragment.children[0] as HTMLElement);

    const flaskOptions = this.simUI.splitRelevantOptions([
			{ item: Flask.FlaskOfSupremePower, stats: [Stat.StatSpellPower] },
			{ item: Flask.FlaskOfRelentlessAssault, stats: [Stat.StatAttackPower, Stat.StatRangedAttackPower] },
			{ item: Flask.FlaskOfMightyRestoration, stats: [Stat.StatMP5] },
			{ item: Flask.FlaskOfFortification, stats: [Stat.StatStamina, Stat.StatDefense] },
			{ item: Flask.FlaskOfPureDeath, stats: [Stat.StatSpellPower] },
      { item: Flask.FlaskOfBlindingLight, stats: [Stat.StatSpellPower] },
      { item: Flask.FlaskOfDistilledWisdom, stats: [Stat.StatIntellect] },
			{ item: Flask.FlaskOfChromaticWonder, stats: [Stat.StatAgility,Stat.StatStrength,Stat.StatIntellect,Stat.StatStamina,Stat.StatSpirit,Stat.StatArcaneResistance, Stat.StatFireResistance, Stat.StatFrostResistance, Stat.StatNatureResistance, Stat.StatShadowResistance] },
		]);
		if (flaskOptions.length) {
			const elem = this.rootElem.querySelector('.consumes-flasks') as HTMLElement;
			new IconEnumPicker(
        elem,
        this.simUI.player,
        IconInputs.makeFlasksInput(flaskOptions, 'Flask')
      );
		}

		const battleElixirOptions = this.simUI.splitRelevantOptions([
			{ item: BattleElixir.ElixirOfMajorShadowPower, stats: [Stat.StatSpellPower] },
			{ item: BattleElixir.ElixirOfMajorFirePower, stats: [Stat.StatSpellPower] },
			{ item: BattleElixir.FelStrengthElixir, stats: [Stat.StatAttackPower,Stat.StatRangedAttackPower] },
			{ item: BattleElixir.ElixirOfMajorAgility, stats: [Stat.StatAgility,Stat.StatMeleeCrit] },
			{ item: BattleElixir.ElixirOfMastery, stats: [Stat.StatAgility,Stat.StatStrength,Stat.StatIntellect,Stat.StatStamina,Stat.StatSpirit] },
      { item: BattleElixir.ElixirOfHealingPower, stats: [Stat.StatSpirit,Stat.StatSpellPower] },
			{ item: BattleElixir.AdeptsElixir, stats: [Stat.StatSpellCrit,Stat.StatSpellPower] },
			{ item: BattleElixir.ElixirOfMajorStrength, stats: [Stat.StatStrength] },
			{ item: BattleElixir.ElixirOfDemonslaying, stats: [Stat.StatAttackPower,Stat.StatRangedAttackPower] },
		]);

    const battleElixirsContainer = this.rootElem.querySelector('.consumes-battle-elixirs') as HTMLElement;
		if (battleElixirOptions.length) {
			new IconEnumPicker(
        battleElixirsContainer,
        this.simUI.player,
        IconInputs.makeBattleElixirsInput(battleElixirOptions, 'Battle Elixir')
      );
		} else {
      battleElixirsContainer.remove();
    }

		const guardianElixirOptions = this.simUI.splitRelevantOptions([
			{ item: GuardianElixir.ElixirOfDraenicWisdom, stats: [Stat.StatIntellect,Stat.StatSpirit] },
			{ item: GuardianElixir.ElixirOfMajorFortitude, stats: [Stat.StatStamina] },
			{ item: GuardianElixir.ElixirOfMajorMageblood, stats: [Stat.StatMP5] },
			{ item: GuardianElixir.ElixirOfMajorDefense, stats: [Stat.StatArmor] },
		]);

    const guardianElixirsContainer = this.rootElem.querySelector('.consumes-guardian-elixirs') as HTMLElement;
		if (guardianElixirOptions.length) {
			const guardianElixirsContainer = this.rootElem.querySelector('.consumes-guardian-elixirs') as HTMLElement;
			new IconEnumPicker(
        guardianElixirsContainer,
        this.simUI.player,
        IconInputs.makeGuardianElixirsInput(guardianElixirOptions, 'Guardian Elixir')
      );
		} else {
      guardianElixirsContainer.remove();
    }
  }

  private buildFoodPicker() {
    let fragment = document.createElement('fragment');
    fragment.innerHTML = `
      <div class="consumes-row input-root input-inline">
        <label class="form-label">Food</label>
        <div class="consumes-row-inputs">
          <div class="consumes-food"></div>
        </div>
      </div>
    `;

    this.rootElem.appendChild(fragment.children[0] as HTMLElement);

    const foodOptions = this.simUI.splitRelevantOptions([
			{ item: Food.FoodBlackenedBasilisk, stats: [Stat.StatSpellPower] },
			{ item: Food.FoodRoastedClefthoof, stats: [Stat.StatStrength] },
      { item: Food.FoodGrilledMudfish, stats: [Stat.StatAgility]},
			{ item: Food.FoodRavagerDog, stats: [Stat.StatAttackPower, Stat.StatRangedAttackPower] },
			{ item: Food.FoodSpicyHotTalbuk, stats: [Stat.StatMeleeHit, Stat.StatSpellHit] },
      { item: Food.FoodFishermansFeast, stats: [Stat.StatStamina] },
      { item: Food.FoodSkullfishSoup, stats: [Stat.StatMeleeCrit, Stat.StatSpellCrit]},
		]);
		if (foodOptions.length) {
			const elem = this.rootElem.querySelector('.consumes-food') as HTMLElement;
			new IconEnumPicker(elem, this.simUI.player, IconInputs.makeFoodInput(foodOptions));
		}
  }

  private buildWeaponImbuePicker() {
    let fragment = document.createElement('fragment');
    fragment.innerHTML = `
      <div class="consumes-row input-root input-inline">
        <label class="form-label">Weapon Imbue</label>
        <div class="consumes-row-inputs">
          <div class="consumes-weapon-main"></div>
          <div class="consumes-weapon-off"></div>
        </div>
      </div>
    `;

    this.rootElem.appendChild(fragment.children[0] as HTMLElement);

    const weaponOptions = this.simUI.splitRelevantOptions([
      this.simUI.player.getClass() == Class.ClassWarlock ? { item: WeaponImbue.ImbueSpellStone, stats: [] } : null,
      this.simUI.player.getClass() == Class.ClassWarlock ? { item: WeaponImbue.ImbueFireStone, stats: [] } : null,
			{ item: WeaponImbue.ImbueBrilliantWizardOil, stats: [Stat.StatSpellPower,Stat.StatSpellCrit] },
      { item: WeaponImbue.ImbueBrilliantManaOil, stats: [Stat.StatSpellPower,Stat.StatMP5]},
      { item: WeaponImbue.ImbueSuperiorWizardOil, stats: [Stat.StatSpellPower]},
      { item: WeaponImbue.ImbueSuperiorManaOil, stats: [Stat.StatMP5]},
      { item: WeaponImbue.ImbueAdamantiteSharpeningStone, stats: [Stat.StatAttackPower, Stat.StatMeleeCrit]},
      { item: WeaponImbue.ImbueAdamantiteWeightStone, stats: [Stat.StatAttackPower, Stat.StatMeleeCrit]},
      { item: WeaponImbue.ImbueElementalSharpeningStone, stats: [Stat.StatMeleeCrit]},

		]);
    const main = this.rootElem.querySelector('.consumes-weapon-main') as HTMLElement;
    const off = this.rootElem.querySelector('.consumes-weapon-off') as HTMLElement;

		if (weaponOptions.length) {
			new IconEnumPicker(main, this.simUI.player, IconInputs.makeMainWeaponImbueInput(weaponOptions));
      new IconEnumPicker(off, this.simUI.player, IconInputs.makeOffWeaponImbueInput(weaponOptions));
		}
    
    const updateGear = () => {
      if (this.simUI.player.getGear().hasSharpOHWeapon() || this.simUI.player.getGear().hasBluntOHWeapon())
				off.classList.remove('hide');
			else
        off.classList.add('hide');
      
    }

    this.simUI.player.gearChangeEmitter.on(updateGear);
    updateGear();

  }

  private buildEngPicker() {
    let fragment = document.createElement('fragment');
    fragment.innerHTML = `
      <div class="consumes-row input-root input-inline">
        <label class="form-label">Engineering</label>
        <div class="consumes-row-inputs consumes-trade"></div>
      </div>
    `;

    this.rootElem.appendChild(fragment.children[0] as HTMLElement);

    const tradeConsumesElem = this.rootElem.querySelector('.consumes-trade') as HTMLElement;

		buildIconInput(tradeConsumesElem, this.simUI.player, IconInputs.SuperSapper);
		buildIconInput(tradeConsumesElem, this.simUI.player, IconInputs.FillerExplosiveInput);

		const updateProfession = () => {
			if (this.simUI.player.hasProfession(Profession.Engineering))
				tradeConsumesElem.parentElement!.classList.remove('hide');
			else
				tradeConsumesElem.parentElement!.classList.add('hide');
		};
		this.simUI.player.professionChangeEmitter.on(updateProfession);
		updateProfession();
  }

  private buildPetPicker() {
    if (this.simUI.individualConfig.petConsumeInputs?.length) {
      let fragment = document.createElement('fragment');
      fragment.innerHTML = `
        <div class="consumes-row input-root input-inline">
          <label class="form-label">Pet</label>
          <div class="consumes-row-inputs consumes-pet"></div>
        </div>
      `;

      this.rootElem.appendChild(fragment.children[0] as HTMLElement);

      const petConsumesElem = this.rootElem.querySelector('.consumes-pet') as HTMLElement;
			this.simUI.individualConfig.petConsumeInputs.map(iconInput => buildIconInput(petConsumesElem, this.simUI.player, iconInput));
		}
  }
}
