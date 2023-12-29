package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

type EthConfig struct {
	Mnemonic string `fig:"mnemonic"`
	//RequestsNumber     int64          `fig:"requests_number"`
	AddressesNumber    int            `fig:"addresses_number"`
	StartNumber        int            `fig:"start_number"`
	RPC                string         `fig:"rpc,required"`
	Vault              common.Address `fig:"vault"`
	ConstitutionVoting common.Address `fig:"constitution_voting"`
	ProposalID         *big.Int       `fig:"proposal_id"`
}

func GetConfig() (EthConfig, error) {
	var result EthConfig

	err := figure.
		Out(&result).
		With(figure.BaseHooks, figure.EthereumHooks).
		From(kv.MustGetStringMap(kv.MustFromEnv(), "ethereum")).
		Please()

	return result, err
}
