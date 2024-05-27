// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/cryft-labs/cryftgo/api/health"
	"github.com/cryft-labs/cryftgo/ids"
	"github.com/cryft-labs/cryftgo/message"
	"github.com/cryft-labs/cryftgo/proto/pb/p2p"
	"github.com/cryft-labs/cryftgo/snow/networking/benchlist"
	"github.com/cryft-labs/cryftgo/snow/networking/handler"
	"github.com/cryft-labs/cryftgo/snow/networking/timeout"
	"github.com/cryft-labs/cryftgo/utils/logging"
	"github.com/cryft-labs/cryftgo/utils/set"
)

// Router routes consensus messages to the Handler of the consensus
// engine that the messages are intended for
type Router interface {
	ExternalHandler
	InternalHandler

	Initialize(
		nodeID ids.NodeID,
		log logging.Logger,
		timeouts timeout.Manager,
		shutdownTimeout time.Duration,
		criticalChains set.Set[ids.ID],
		sybilProtectionEnabled bool,
		trackedSubnets set.Set[ids.ID],
		onFatal func(exitCode int),
		healthConfig HealthConfig,
		metricsNamespace string,
		metricsRegisterer prometheus.Registerer,
	) error
	Shutdown(context.Context)
	AddChain(ctx context.Context, chain handler.Handler)
	health.Checker
}

// InternalHandler deals with messages internal to this node
type InternalHandler interface {
	benchlist.Benchable

	RegisterRequest(
		ctx context.Context,
		nodeID ids.NodeID,
		sourceChainID ids.ID,
		destinationChainID ids.ID,
		requestID uint32,
		op message.Op,
		failedMsg message.InboundMessage,
		engineType p2p.EngineType,
	)
}
