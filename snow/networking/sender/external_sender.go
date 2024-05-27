// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sender

import (
	"github.com/cryft-labs/cryftgo/ids"
	"github.com/cryft-labs/cryftgo/message"
	"github.com/cryft-labs/cryftgo/snow/engine/common"
	"github.com/cryft-labs/cryftgo/subnets"
	"github.com/cryft-labs/cryftgo/utils/set"
)

// ExternalSender sends consensus messages to other validators
// Right now this is implemented in the networking package
type ExternalSender interface {
	Send(
		msg message.OutboundMessage,
		config common.SendConfig,
		subnetID ids.ID,
		allower subnets.Allower,
	) set.Set[ids.NodeID]
}
