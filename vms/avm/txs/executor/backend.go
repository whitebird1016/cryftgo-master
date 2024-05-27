// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"reflect"

	"github.com/cryft-labs/cryftgo/codec"
	"github.com/cryft-labs/cryftgo/ids"
	"github.com/cryft-labs/cryftgo/snow"
	"github.com/cryft-labs/cryftgo/vms/avm/config"
	"github.com/cryft-labs/cryftgo/vms/avm/fxs"
)

type Backend struct {
	Ctx           *snow.Context
	Config        *config.Config
	Fxs           []*fxs.ParsedFx
	TypeToFxIndex map[reflect.Type]int
	Codec         codec.Manager
	// Note: FeeAssetID may be different than ctx.AVAXAssetID if this AVM is
	// running in a subnet.
	FeeAssetID   ids.ID
	Bootstrapped bool
}
