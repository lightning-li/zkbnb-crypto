/*
 * Copyright © 2022 ZkBNB Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package types

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
	eddsaConstraints "github.com/consensys/gnark/std/signature/eddsa"
)

type (
	Variable             = frontend.Variable
	API                  = frontend.API
	MiMC                 = mimc.MiMC
	PublicKeyConstraints = eddsaConstraints.PublicKey
	PublicKey            = eddsa.PublicKey
)

const (
	ZeroInt    = uint64(0)
	DefaultInt = int64(-1)

	NbAccountAssetsPerAccount = 2
	NbAccountsPerTx           = 4
	NbGasAssetsPerTx          = 2 // at most two assets transferred to gas account

	NbRoots = 2 // account root, nft root

	PubDataSizePerTx = 6

	PubDataBitsSizePerTx = 968 // registerZNS

	NftContentHashBytesSize = 16

	OfferSizePerAsset = 128

	ChainId = 1
)

const (
	TxTypeEmptyTx = iota
	TxTypeRegisterZns
	TxTypeDeposit
	TxTypeDepositNft
	TxTypeTransfer
	TxTypeWithdraw
	TxTypeCreateCollection
	TxTypeMintNft
	TxTypeTransferNft
	TxTypeAtomicMatch
	TxTypeCancelOffer
	TxTypeWithdrawNft
	TxTypeFullExit
	TxTypeFullExitNft
)

const (
	RateBase = 10000
)

var (
	EmptyAssetRoot, _ = new(big.Int).SetString("0f683e832fd8381ddac14f6fdfe32d8c3e5a1e679ead0b06755ccb61eb6b4c5d", 16)
)
