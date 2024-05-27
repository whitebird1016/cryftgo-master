// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/cryft-labs/cryftgo/ids"
	"github.com/cryft-labs/cryftgo/utils/json"
	"github.com/cryft-labs/cryftgo/utils/set"
)

type Info struct {
	IP                    string                 `json:"ip"`
	PublicIP              string                 `json:"publicIP,omitempty"`
	ID                    ids.NodeID             `json:"nodeID"`
	Version               string                 `json:"version"`
	LastSent              time.Time              `json:"lastSent"`
	LastReceived          time.Time              `json:"lastReceived"`
	ObservedUptime        json.Uint32            `json:"observedUptime"`
	ObservedSubnetUptimes map[ids.ID]json.Uint32 `json:"observedSubnetUptimes"`
	TrackedSubnets        set.Set[ids.ID]        `json:"trackedSubnets"`
	SupportedACPs         set.Set[uint32]        `json:"supportedACPs"`
	ObjectedACPs          set.Set[uint32]        `json:"objectedACPs"`
}
