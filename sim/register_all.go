package sim

import (
	_ "github.com/Tereneckla/wotlk/sim/common"
	"github.com/Tereneckla/wotlk/sim/druid/balance"
	"github.com/Tereneckla/wotlk/sim/druid/feral"
	restoDruid "github.com/Tereneckla/wotlk/sim/druid/restoration"
	feralTank "github.com/Tereneckla/wotlk/sim/druid/tank"
	_ "github.com/Tereneckla/wotlk/sim/encounters"
	"github.com/Tereneckla/wotlk/sim/hunter"
	"github.com/Tereneckla/wotlk/sim/mage"
	holyPaladin "github.com/Tereneckla/wotlk/sim/paladin/holy"
	protectionPaladin "github.com/Tereneckla/wotlk/sim/paladin/protection"
	"github.com/Tereneckla/wotlk/sim/paladin/retribution"
	healingPriest "github.com/Tereneckla/wotlk/sim/priest/healing"
	"github.com/Tereneckla/wotlk/sim/priest/shadow"
	"github.com/Tereneckla/wotlk/sim/priest/smite"
	"github.com/Tereneckla/wotlk/sim/rogue"
	"github.com/Tereneckla/wotlk/sim/shaman/elemental"
	"github.com/Tereneckla/wotlk/sim/shaman/enhancement"
	restoShaman "github.com/Tereneckla/wotlk/sim/shaman/restoration"
	"github.com/Tereneckla/wotlk/sim/warlock"
	dpsWarrior "github.com/Tereneckla/wotlk/sim/warrior/dps"
	protectionWarrior "github.com/Tereneckla/wotlk/sim/warrior/protection"
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
}
