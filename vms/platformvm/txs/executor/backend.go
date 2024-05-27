// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/cryft-labs/cryftgo/snow"
	"github.com/cryft-labs/cryftgo/snow/uptime"
	"github.com/cryft-labs/cryftgo/utils"
	"github.com/cryft-labs/cryftgo/utils/timer/mockable"
	"github.com/cryft-labs/cryftgo/vms/platformvm/config"
	"github.com/cryft-labs/cryftgo/vms/platformvm/fx"
	"github.com/cryft-labs/cryftgo/vms/platformvm/reward"
	"github.com/cryft-labs/cryftgo/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Calculator
	Rewards      reward.Calculator
	Bootstrapped *utils.Atomic[bool]
}
