package holy

import (
	"time"

	"github.com/Tereneckla/wotlk/sim/core"
)

func (holy *HolyPaladin) OnGCDReady(sim *core.Simulation) {
	holy.WaitUntil(sim, sim.CurrentTime+time.Second*5)
}
