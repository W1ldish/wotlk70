package paladin

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
	"github.com/Tereneckla/wotlk/sim/core/proto"
	"github.com/Tereneckla/wotlk/sim/core/stats"
)

var ItemSetJusticarBattlegear = core.NewItemSet(core.ItemSet{
	Name: "Justicar Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// sim/debuffs.go handles this (and paladin/judgement.go)
		},
		4: func(agent core.Agent) {
			// TODO: if we ever implemented judgement of command, add bonus from 4p
		},
	},
})
var ItemSetCrystalforgeBattlegear = core.NewItemSet(core.ItemSet{
	Name: "Crystalforge Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// judgement.go
		},
		4: func(agent core.Agent) {
			// TODO: if we implement healing, this heals party.
		},
	},
})

// Tier 6 ret
var ItemSetLightbringerBattlegear = core.NewItemSet(core.ItemSet{
	Name: "Lightbringer Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()
			manaMetrics := paladin.NewManaMetrics(core.ActionID{SpellID: 38428})

			paladin.RegisterAura(core.Aura{
				Label:    "Lightbringer Battlegear 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if !spell.ProcMask.Matches(core.ProcMaskMelee) {
						return
					}
					if sim.RandomFloat("lightbringer 2pc") > 0.2 {
						return
					}
					paladin.AddMana(sim, 50, manaMetrics)
				},
			})
		},
		4: func(agent core.Agent) {
			// Implemented in hammer_of_wrath.go
		},
	},
})

func (paladin *Paladin) getItemSetLightbringerBattlegearBonus4() float64 {
	return core.TernaryFloat64(paladin.HasSetBonus(ItemSetLightbringerBattlegear, 4), .1, 0)
}

var ItemSetJusticarArmor = core.NewItemSet(core.ItemSet{
	Name: "Justicar Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the damage dealt by your Seal of Righteousness, Seal of
			// Vengeance, and Seal of Corruption by 10%.
			// Implemented in seals.go.
		},
		4: func(agent core.Agent) {
			// Increases the damage dealt by Holy Shield by 15.
			// Implemented in holy_shield.go.
		},
	},
})

var ItemSetCrystalforgeArmor = core.NewItemSet(core.ItemSet{
	Name: "Crystalforge Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the damage from your Retribution Aura by 15.
			// TODO
		},
		4: func(agent core.Agent) {
			// Each time you use your Holy Shield ability, you gain 100 Block Value
			// against a single attack in the next 6 seconds.
			paladin := agent.(PaladinAgent).GetPaladin()

			procAura := paladin.RegisterAura(core.Aura{
				Label:    "Crystalforge 4pc Proc",
				ActionID: core.ActionID{SpellID: 37191},
				Duration: time.Second * 6,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					paladin.AddStatDynamic(sim, stats.BlockValue, 100)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					paladin.AddStatDynamic(sim, stats.BlockValue, -100)
				},
				OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellResult) {
					if spellEffect.Outcome.Matches(core.OutcomeBlock) {
						aura.Deactivate(sim)
					}
				},
			})

			paladin.RegisterAura(core.Aura{
				Label:    "Crystalforge 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if spell == paladin.HolyShield {
						procAura.Activate(sim)
					}
				},
			})
		},
	},
})
var ItemSetLightbringerArmor = core.NewItemSet(core.ItemSet{
	Name: "Lightbringer Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the mana gained from your Spiritual Attunement ability by 10%.
		},
		4: func(agent core.Agent) {
			// Increases the damage dealt by Consecration by 10%.
		},
	},
})

func (paladin *Paladin) getItemSetGladiatorsVindicationBonusGloves() float64 {
	hasGloves := (paladin.Equip[proto.ItemSlot_ItemSlotHands].ID == 40798) || // S5a Hateful
		(paladin.Equip[proto.ItemSlot_ItemSlotHands].ID == 40802) || // S5b Hateful
		(paladin.Equip[proto.ItemSlot_ItemSlotHands].ID == 40805) || // S5c Deadly
		(paladin.Equip[proto.ItemSlot_ItemSlotHands].ID == 40808) || // S6 Furious
		(paladin.Equip[proto.ItemSlot_ItemSlotHands].ID == 40812) || // S7 Relentless
		(paladin.Equip[proto.ItemSlot_ItemSlotHands].ID == 51475) // S8 Wrathful
	return core.TernaryFloat64(hasGloves, .05, 0)
}

func init() {
	// Librams implemented in seals.go and judgement.go

	core.NewItemEffect(32368, func(agent core.Agent) {
		paladin := agent.(PaladinAgent).GetPaladin()
		procAura := paladin.NewTemporaryStatsAura("Tome of the Lightbringer Proc", core.ActionID{SpellID: 41042}, stats.Stats{stats.BlockValue: 186}, time.Second*10)

		paladin.RegisterAura(core.Aura{
			Label:    "Tome of the Lightbringer",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.Flags.Matches(SpellFlagPrimaryJudgement) {
					procAura.Activate(sim)
				}
			},
		})
	})

	core.NewItemEffect(27484, func(agent core.Agent) {
		paladin := agent.(PaladinAgent).GetPaladin()
		procAura := paladin.NewTemporaryStatsAura("Tome of the Avengement Proc", core.ActionID{SpellID: 34258}, stats.Stats{stats.MeleeCrit: 53, stats.SpellCrit: 53}, time.Second*5)

		paladin.RegisterAura(core.Aura{
			Label:    "Tome of the Avengement",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.Flags.Matches(SpellFlagPrimaryJudgement) {
					procAura.Activate(sim)
				}
			},
		})
	})

	core.NewItemEffect(32489, func(agent core.Agent) {
		paladin := agent.(PaladinAgent).GetPaladin()

		// The spell effect is https://www.wowhead.com/wotlk/spell=40472/enduring-judgement, most likely
		dotSpell := paladin.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{ItemID: 32489},
			SpellSchool:      core.SpellSchoolHoly,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "AshtongueTalismanOfZeal",
				},
				NumberOfTicks: 4,
				TickLength:    time.Second * 2,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.SnapshotBaseDamage = 480 / 4
					dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(dot.Spell.Unit.AttackTables[target.UnitIndex])
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})

		paladin.RegisterAura(core.Aura{
			Label:    "Ashtongue Talisman of Zeal",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.Flags.Matches(SpellFlagPrimaryJudgement) && sim.RandomFloat("AshtongueTalismanOfZeal") < 0.5 {
					dotSpell.Dot(result.Target).Apply(sim)
				}
			},
		})
	})

}
