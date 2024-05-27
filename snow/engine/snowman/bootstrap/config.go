// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"github.com/cryft-labs/cryftgo/database"
	"github.com/cryft-labs/cryftgo/network/p2p"
	"github.com/cryft-labs/cryftgo/snow"
	"github.com/cryft-labs/cryftgo/snow/engine/common"
	"github.com/cryft-labs/cryftgo/snow/engine/common/tracker"
	"github.com/cryft-labs/cryftgo/snow/engine/snowman/block"
	"github.com/cryft-labs/cryftgo/snow/validators"
)

type Config struct {
	common.AllGetsServer

	Ctx     *snow.ConsensusContext
	Beacons validators.Manager

	SampleK          int
	StartupTracker   tracker.Startup
	Sender           common.Sender
	BootstrapTracker common.BootstrapTracker
	Timer            common.Timer

	// PeerTracker manages the set of nodes that we fetch the next block from.
	PeerTracker *p2p.PeerTracker

	// This node will only consider the first [AncestorsMaxContainersReceived]
	// containers in an ancestors message it receives.
	AncestorsMaxContainersReceived int

	// Database used to track the fetched, but not yet executed, blocks during
	// bootstrapping.
	DB database.Database

	VM block.ChainVM

	Bootstrapped func()
}
