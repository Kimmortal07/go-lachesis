// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package ethapi implements the general Ethereum API functions.
package ethapi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	notify "github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Fantom-foundation/go-lachesis/evm_core"
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// General Ethereum API
	ProtocolVersion() int
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() ethdb.Database
	NotifyMux() *notify.TypeMux
	AccountManager() *accounts.Manager
	ExtRPCEnabled() bool
	RPCGasCap() *big.Int // global gas cap for eth_call over rpc: DoS protection

	// Blockchain API
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*evm_core.EvmHeader, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*evm_core.EvmHeader, error)
	BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*evm_core.EvmBlock, error)
	StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *evm_core.EvmHeader, error)
	GetHeader(ctx context.Context, hash common.Hash) *evm_core.EvmHeader
	GetBlock(ctx context.Context, hash common.Hash) (*evm_core.EvmBlock, error)
	GetReceipts(ctx context.Context, number rpc.BlockNumber) (types.Receipts, error)
	GetTd(hash common.Hash) *big.Int
	GetEVM(ctx context.Context, msg evm_core.Message, state *state.StateDB, header *evm_core.EvmHeader) (*vm.EVM, func() error, error)

	// Transaction pool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, uint64, uint64, error)
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeNewTxsNotify(chan<- evm_core.NewTxsNotify) notify.Subscription

	ChainConfig() *params.ChainConfig
	CurrentBlock() *evm_core.EvmBlock

	// Lachesis debug API
	GetEvent(ctx context.Context, shortEventID string) (*inter.Event, error)
	GetEventHeader(ctx context.Context, shortEventID string) (*inter.EventHeaderData, error)
	GetConsensusTime(ctx context.Context, shortEventID string) (inter.Timestamp, error)
	GetHeads(ctx context.Context) hash.Events
}

func GetAPIs(apiBackend Backend) []rpc.API {
	nonceLock := new(AddrLocker)
	return []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(apiBackend),
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(apiBackend, nonceLock),
			Public:    false,
		},
	}
}
