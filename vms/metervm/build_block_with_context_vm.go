// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"context"

	"github.com/cryft-labs/cryftgo/snow/consensus/snowman"
	"github.com/cryft-labs/cryftgo/snow/engine/snowman/block"
)

func (vm *blockVM) BuildBlockWithContext(ctx context.Context, blockCtx *block.Context) (snowman.Block, error) {
	if vm.buildBlockVM == nil {
		return vm.BuildBlock(ctx)
	}

	start := vm.clock.Time()
	blk, err := vm.buildBlockVM.BuildBlockWithContext(ctx, blockCtx)
	end := vm.clock.Time()
	duration := float64(end.Sub(start))
	if err != nil {
		vm.blockMetrics.buildBlockWithContextErr.Observe(duration)
		return nil, err
	}
	vm.blockMetrics.buildBlockWithContext.Observe(duration)
	return &meterBlock{
		Block: blk,
		vm:    vm,
	}, nil
}
