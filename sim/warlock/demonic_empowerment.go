package warlock

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/proto"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
)

func (warlock *Warlock) registerDemonicEmpowermentSpell() {
	if !warlock.Talents.DemonicEmpowerment {
		return
	}

	var petAura core.Aura
	switch warlock.Options.Summon {
	case proto.Warlock_Options_Imp:
		petAura = core.Aura{
			Label:    "Demonic Empowerment Aura",
			ActionID: core.ActionID{SpellID: 47193},
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				warlock.Pet.AddStatsDynamic(sim, stats.Stats{
					stats.SpellCrit: 20 * core.CritRatingPerCritChance,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				warlock.Pet.AddStatsDynamic(sim, stats.Stats{
					stats.SpellCrit: -20 * core.CritRatingPerCritChance,
				})
			},
		}
	case proto.Warlock_Options_Felguard:
		petAura = core.Aura{
			Label:    "Demonic Empowerment Aura",
			ActionID: core.ActionID{SpellID: 47193},
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				warlock.Pet.MultiplyAttackSpeed(sim, 1.2)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				warlock.Pet.MultiplyAttackSpeed(sim, 1/1.2)
			},
		}
	default:
		petAura = core.Aura{
			Label:    "Demonic Empowerment Aura",
			ActionID: core.ActionID{SpellID: 47193},
			Duration: time.Second * 15,
		}
	}

	if warlock.Options.Summon != proto.Warlock_Options_NoSummon {
		warlock.Pet.DemonicEmpowermentAura = warlock.Pet.RegisterAura(petAura)
	}

	warlock.DemonicEmpowerment = warlock.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: 47193},

		ManaCost: core.ManaCostOptions{
			BaseCost: 0.06,
		},
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    warlock.NewTimer(),
				Duration: time.Second * time.Duration(60.*(1.-0.1*float64(warlock.Talents.Nemesis))),
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			warlock.Pet.DemonicEmpowermentAura.Activate(sim)
		},
	})
}
