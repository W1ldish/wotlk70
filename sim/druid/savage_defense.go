package druid

import (
	"time"

	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
)

func (druid *Druid) registerSavageDefensePassive() {
	if !druid.InForm(Bear) {
		return
	}

	druid.SavageDefenseAura = druid.RegisterAura(core.Aura{
		Label:    "Savage Defense",
		ActionID: core.ActionID{SpellID: 62606},
		Duration: 10 * time.Second,
	})

	druid.AddDynamicDamageTakenModifier(func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if druid.SavageDefenseAura.IsActive() && (result.Damage > 0) {
			result.Damage = core.MaxFloat(0, result.Damage-0.25*druid.GetStat(stats.AttackPower))
			druid.SavageDefenseAura.Deactivate(sim)
		}
	})

	core.MakeProcTriggerAura(&druid.Unit, core.ProcTrigger{
		Name:     "Savage Defense Trigger",
		Callback: core.CallbackOnSpellHitDealt | core.CallbackOnPeriodicDamageDealt,
		ProcMask: core.ProcMaskMelee | core.ProcMaskSpellDamage,
		Harmful:  true,
		Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
			if result.Outcome.Matches(core.OutcomeCrit) {
				druid.SavageDefenseAura.Activate(sim)
			}
		},
	})
}
