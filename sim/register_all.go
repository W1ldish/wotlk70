package sim

import (
	_ "github.com/Tereneckla/wotlk70/sim/common"
	dpsDeathKnight "github.com/Tereneckla/wotlk70/sim/deathknight/dps"
	tankDeathKnight "github.com/Tereneckla/wotlk70/sim/deathknight/tank"
	"github.com/Tereneckla/wotlk70/sim/druid/balance"
	"github.com/Tereneckla/wotlk70/sim/druid/feral"
	restoDruid "github.com/Tereneckla/wotlk70/sim/druid/restoration"
	feralTank "github.com/Tereneckla/wotlk70/sim/druid/tank"
	_ "github.com/Tereneckla/wotlk70/sim/encounters"
	"github.com/Tereneckla/wotlk70/sim/hunter"
	"github.com/Tereneckla/wotlk70/sim/mage"
	holyPaladin "github.com/Tereneckla/wotlk70/sim/paladin/holy"
	protectionPaladin "github.com/Tereneckla/wotlk70/sim/paladin/protection"
	"github.com/Tereneckla/wotlk70/sim/paladin/retribution"
	healingPriest "github.com/Tereneckla/wotlk70/sim/priest/healing"
	"github.com/Tereneckla/wotlk70/sim/priest/shadow"
	"github.com/Tereneckla/wotlk70/sim/priest/smite"
	"github.com/Tereneckla/wotlk70/sim/rogue"
	"github.com/Tereneckla/wotlk70/sim/shaman/elemental"
	"github.com/Tereneckla/wotlk70/sim/shaman/enhancement"
	restoShaman "github.com/Tereneckla/wotlk70/sim/shaman/restoration"
	"github.com/Tereneckla/wotlk70/sim/warlock"
	dpsWarrior "github.com/Tereneckla/wotlk70/sim/warrior/dps"
	protectionWarrior "github.com/Tereneckla/wotlk70/sim/warrior/protection"
)

var registered = false

func RegisterAll() {
	if registered {
		return
	}
	registered = true

	balance.RegisterBalanceDruid()
	feral.RegisterFeralDruid()
	feralTank.RegisterFeralTankDruid()
	restoDruid.RegisterRestorationDruid()
	elemental.RegisterElementalShaman()
	enhancement.RegisterEnhancementShaman()
	restoShaman.RegisterRestorationShaman()
	hunter.RegisterHunter()
	mage.RegisterMage()
	healingPriest.RegisterHealingPriest()
	shadow.RegisterShadowPriest()
	smite.RegisterSmitePriest()
	rogue.RegisterRogue()
	dpsWarrior.RegisterDpsWarrior()
	protectionWarrior.RegisterProtectionWarrior()
	holyPaladin.RegisterHolyPaladin()
	protectionPaladin.RegisterProtectionPaladin()
	retribution.RegisterRetributionPaladin()
	warlock.RegisterWarlock()
	dpsDeathKnight.RegisterDpsDeathknight()
	tankDeathKnight.RegisterTankDeathknight()
}
