package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
	"voter/contracts/constitutionvoting"
	"voter/contracts/qvault"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbsobreira/gotron-sdk/pkg/keys/hd"
	"github.com/tyler-smith/go-bip39"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Account struct {
	Sk      *ecdsa.PrivateKey
	Pk      *ecdsa.PublicKey
	Address common.Address
	Balance *big.Int
	Nonce   *uint64
}

func FromMnemonicSeed(mnemonic string, index int) (*btcec.PrivateKey, *btcec.PublicKey) {
	seed := bip39.NewSeed(mnemonic, "")
	master, ch := hd.ComputeMastersFromSeed(seed, []byte("Bitcoin seed"))
	private, _ := hd.DerivePrivateKeyForPath(
		btcec.S256(),
		master,
		ch,
		fmt.Sprintf("44'/60'/0'/0/%d", index),
	)

	return btcec.PrivKeyFromBytes(private[:])
}

func main() {
	config, err := GetConfig()
	if err != nil {
		panic(errors.Wrap(err, "wrong config"))
	}

	client, err := ethclient.Dial(config.RPC)
	if err != nil {
		panic(errors.Wrap(err, "failed to create connection"))
	}
	defer client.Close()

	accounts := make(map[int]Account)

	fmt.Println("Start generating addresses and getting balances")
	now := time.Now()

	for i := 0; i < config.AddressesNumber; i++ {
		sk, pk := FromMnemonicSeed(config.Mnemonic, i+config.StartNumber)
		address := crypto.PubkeyToAddress(*pk.ToECDSA())

		var err error
		balance := big.NewInt(0)

		for j := 0; j < 5; j++ {
			balance, err = client.BalanceAt(context.Background(), address, nil)
			if err == nil {
				break
			}
			fmt.Println(err)
			time.Sleep(time.Second)
		}

		accounts[i] = Account{
			Sk:      sk.ToECDSA(),
			Pk:      pk.ToECDSA(),
			Address: address,
			Balance: balance,
		}
	}

	fmt.Println(time.Since(now))

	fmt.Println("Finish generating addresses and getting balances")

	sent := 0
	sentPrev := -1

	chainID, _ := client.ChainID(context.Background())

	fmt.Println("Start sending txs")
	fmt.Println("Start time ", time.Now())
	now = time.Now()

	notSentAddr := make(map[int]Account)

	vault, err := qvault.NewContracts(config.Vault, client)
	if err != nil {
		panic(errors.Wrap(err, "failed to create vault instance"))
	}

	voting, err := constitutionvoting.NewConstitutionvoting(config.ConstitutionVoting, client)
	if err != nil {
		panic(errors.Wrap(err, "failed to create voting instance"))
	}

	for sent < config.AddressesNumber {
		if sentPrev != sent {
			fmt.Printf("start processing address: %d %s\n", sent, accounts[sent].Address.String())
			sentPrev = sent
		}

		var nonce uint64
		nonce, err = client.PendingNonceAt(context.Background(), accounts[sent].Address)
		if err != nil {
			fmt.Println(err)
			continue
		}

		authOpts, err := bind.NewKeyedTransactorWithChainID(accounts[sent].Sk, chainID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		authOpts.NoSend = true

		depositOpts := *authOpts
		depositOpts.Value = big.NewInt(1)

		depositTx, err := vault.Deposit(&depositOpts)
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to estimate DepositTX"))
			continue
		}
		depositCost := big.NewInt(int64(depositTx.Gas()))
		depositCost.Mul(depositCost, depositTx.GasPrice())

		//lockOpts := *authOpts
		//
		//lockTx, err := vault.Lock(&lockOpts, big.NewInt(1))
		//if err != nil {
		//	fmt.Println(errors.Wrap(err, "failed to estimate LockTX"))
		//	continue
		//}
		//lockCost := big.NewInt(int64(lockTx.Gas()))
		//lockCost.Mul(lockCost, lockTx.GasPrice())

		// cannot estimate tx without executing previous one
		lockCost := big.NewInt(13000000000000000)

		//voteOpts := *authOpts
		//voteOpts.Nonce = big.NewInt(int64(nonce) + 2)
		//
		//voteTx, err := voting.VoteFor(&voteOpts, config.ProposalID)
		//if err != nil {
		//	fmt.Println(errors.Wrap(err, "failed to estimate VoteTX"))
		//	continue
		//}
		//voteCost := big.NewInt(int64(voteTx.Gas()))
		//voteCost.Mul(voteCost, voteTx.GasPrice())

		// cannot estimate tx without executing previous one
		voteCost := big.NewInt(12000000000000000)

		feeSum := new(big.Int).Add(depositCost, lockCost)
		feeSum = feeSum.Add(feeSum, voteCost)
		// adjust fee sum
		feeSum = feeSum.Mul(feeSum, big.NewInt(15))
		feeSum = feeSum.Div(feeSum, big.NewInt(10))

		toSend := new(big.Int).Sub(accounts[sent].Balance, feeSum)

		// send deposit tx
		newDepositTx, err := vault.Deposit(&bind.TransactOpts{
			From:     authOpts.From,
			Nonce:    authOpts.Nonce,
			Signer:   authOpts.Signer,
			Value:    toSend,
			GasPrice: depositTx.GasPrice(),
			GasLimit: depositTx.Gas(),
			Context:  authOpts.Context,
			NoSend:   false,
		})
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to create DepositTx"))
			continue
		}

		receipt, err := bind.WaitMined(context.Background(), client, newDepositTx)
		if err != nil {
			fmt.Printf("%d failed DepositTx wait mined %s\n", sent, newDepositTx.Hash())
			fmt.Println(err)
			notSentAddr[sent] = accounts[sent]
			sent++
			continue
		}
		fmt.Printf("%d DepositTx wait mined %s block: %d\n", sent, newDepositTx.Hash(), receipt.BlockNumber.Int64())
		sleepRand(4000, 2000, 2000)

		//send lock tx
		newLockTx, err := vault.Lock(&bind.TransactOpts{
			From:     authOpts.From,
			Nonce:    big.NewInt(int64(nonce) + 1),
			Signer:   authOpts.Signer,
			Value:    nil,
			GasPrice: depositTx.GasPrice(),
			GasLimit: 300000, // todo: add randomisation
			Context:  authOpts.Context,
			NoSend:   false,
		}, toSend)
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to create LockTx"))
			notSentAddr[sent] = accounts[sent]
			sent++
			continue
		}

		receipt, err = bind.WaitMined(context.Background(), client, newLockTx)
		if err != nil {
			fmt.Printf("%d failed LockTx wait mined %s\n", sent, newLockTx.Hash())
			fmt.Println(err)
			notSentAddr[sent] = accounts[sent]
			sent++
			continue
		}
		fmt.Printf("%d LockTx wait mined %s block: %d\n", sent, newLockTx.Hash(), receipt.BlockNumber.Int64())
		sleepRand(4000, 2000, 2000)

		// send vote tx
		voteOpts := *authOpts
		voteOpts.Nonce = big.NewInt(int64(nonce) + 2)
		voteTx, err := voting.VoteFor(&voteOpts, config.ProposalID)
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to create VoteTX"))
			notSentAddr[sent] = accounts[sent]
			sent++
			continue
		}

		voteTx, err = types.SignTx(voteTx, types.NewLondonSigner(chainID), accounts[sent].Sk)
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to sign VoteTx"))
			notSentAddr[sent] = accounts[sent]
			sent++
			continue
		}

		err = client.SendTransaction(context.Background(), voteTx)
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to send VoteTx"))
			notSentAddr[sent] = accounts[sent]
			sent++
			continue
		}

		receipt, err = bind.WaitMined(context.Background(), client, voteTx)
		if err != nil {
			fmt.Printf("%d failed VoteTx wait mined %s\n", sent, voteTx.Hash())
			fmt.Println(err)
			notSentAddr[sent] = accounts[sent]
			sent++
			continue
		}
		fmt.Printf("%d VoteTx wait mined %s block: %d\n", sent, voteTx.Hash(), receipt.BlockNumber.Int64())

		sent++

		sleepRand(7500, 3000, 2000)
	}

	fmt.Println(now)
	fmt.Println(time.Since(now))

	fmt.Println("Finish sending txs")
	fmt.Println("Not sent:")
	for _, account := range notSentAddr {
		fmt.Println(account.Address.String())
	}
}

func sleepRand(maxDuration, minDuration, delta int64) {
	randDelay := new(big.Int)
	randDelay, err := rand.Int(rand.Reader, big.NewInt(maxDuration))
	if err != nil {
		randDelay = big.NewInt(minDuration)
	}
	randDelay.Add(randDelay, big.NewInt(delta))
	time.Sleep(time.Duration(randDelay.Int64() * int64(time.Millisecond)))
}
