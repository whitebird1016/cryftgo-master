// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txstest

import (
	"time"

	"github.com/cryft-labs/cryftgo/snow"
	"github.com/cryft-labs/cryftgo/vms/platformvm/config"
	"github.com/cryft-labs/cryftgo/vms/platformvm/txs"
	"github.com/cryft-labs/cryftgo/vms/platformvm/txs/fee"
	"github.com/cryft-labs/cryftgo/wallet/chain/p/builder"
)

func newContext(
	ctx *snow.Context,
	cfg *config.Config,
	timestamp time.Time,
) *builder.Context {
	var (
		feeCalc         = fee.NewStaticCalculator(cfg.StaticFeeConfig, cfg.UpgradeConfig)
		createSubnetFee = feeCalc.CalculateFee(&txs.CreateSubnetTx{}, timestamp)
		createChainFee  = feeCalc.CalculateFee(&txs.CreateChainTx{}, timestamp)
	)

	return &builder.Context{
		NetworkID:                     ctx.NetworkID,
		AVAXAssetID:                   ctx.AVAXAssetID,
		BaseTxFee:                     cfg.StaticFeeConfig.TxFee,
		CreateSubnetTxFee:             createSubnetFee,
		TransformSubnetTxFee:          cfg.StaticFeeConfig.TransformSubnetTxFee,
		CreateBlockchainTxFee:         createChainFee,
		AddPrimaryNetworkValidatorFee: cfg.StaticFeeConfig.AddPrimaryNetworkValidatorFee,
		AddPrimaryNetworkDelegatorFee: cfg.StaticFeeConfig.AddPrimaryNetworkDelegatorFee,
		AddSubnetValidatorFee:         cfg.StaticFeeConfig.AddSubnetValidatorFee,
		AddSubnetDelegatorFee:         cfg.StaticFeeConfig.AddSubnetDelegatorFee,
	}
}
