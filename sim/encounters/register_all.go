package encounters

import (
	"github.com/Tereneckla/wotlk70/sim/core"
	"github.com/Tereneckla/wotlk70/sim/encounters/naxxramas"
	"github.com/Tereneckla/wotlk70/sim/encounters/ulduar"
)

func init() {
	naxxramas.Register()
	ulduar.Register()
}

func AddSingleTargetBossEncounter(presetTarget *core.PresetTarget) {
	core.AddPresetTarget(presetTarget)
	core.AddPresetEncounter(presetTarget.Config.Name, []string{
		presetTarget.Path(),
	})
}
