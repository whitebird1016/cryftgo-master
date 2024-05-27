// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package subnet

import (
	"math"
	"time"

	"github.com/cryft-labs/cryftgo/tests/fixture/tmpnet"
	"github.com/cryft-labs/cryftgo/utils/crypto/secp256k1"
	"github.com/cryft-labs/cryftgo/vms/example/xsvm"
	"github.com/cryft-labs/cryftgo/vms/example/xsvm/genesis"
)

func NewXSVMOrPanic(name string, key *secp256k1.PrivateKey, nodes ...*tmpnet.Node) *tmpnet.Subnet {
	if len(nodes) == 0 {
		panic("a subnet must be validated by at least one node")
	}

	genesisBytes, err := genesis.Codec.Marshal(genesis.CodecVersion, &genesis.Genesis{
		Timestamp: time.Now().Unix(),
		Allocations: []genesis.Allocation{
			{
				Address: key.Address(),
				Balance: math.MaxUint64,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	return &tmpnet.Subnet{
		Name: name,
		Chains: []*tmpnet.Chain{
			{
				VMID:         xsvm.ID,
				Genesis:      genesisBytes,
				PreFundedKey: key,
			},
		},
		ValidatorIDs: tmpnet.NodesToIDs(nodes...),
	}
}
