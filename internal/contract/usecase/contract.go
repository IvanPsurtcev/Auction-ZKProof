package usecase

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
	"metamask-auth/config"
	"metamask-auth/internal/contract"
)

type contractUseCase struct {
	logger *zap.Logger
	cfg    *config.Config
}

func NewUseCase(logger *zap.Logger, cfg *config.Config) contract.UseCase {
	return &contractUseCase{
		logger: logger,
		cfg:    cfg,
	}
}

func (u *contractUseCase) StartAuction(ctx context.Context) error {
	address := common.HexToAddress(u.cfg.Contract.SellerAddress)

	privateKey, err := crypto.HexToECDSA(u.cfg.Contract.SellerPK)
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	ethereumClient, err := ethclient.Dial(u.cfg.Contract.Rpc)
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	nonce, err := ethereumClient.PendingNonceAt(ctx, address)
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	gasPrice, err := ethereumClient.SuggestGasPrice(ctx)
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	//gasTip, err := ethereumClient.SuggestGasTipCap(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//estimateGas, err := ethereumClient.EstimateGas(ctx, ethereum.CallMsg{
	//	To:       &address,
	//	GasPrice: gasPrice,
	//	Value:    big.NewInt(0),
	//	Data:     []byte{},
	//},
	//)

	seller, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(u.cfg.Contract.ChainID))
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	myContract, err := contract.NewContract(common.HexToAddress(u.cfg.Contract.ContractAddress), ethereumClient)
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	tx, err := myContract.StartAuction(&bind.TransactOpts{
		From:     address,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   seller.Signer,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: 100000,
		Context:  ctx,
		NoSend:   true,
	})
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(u.cfg.Contract.ChainID)), privateKey)
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	err = ethereumClient.SendTransaction(ctx, signedTx)
	if err != nil {
		u.logger.Error("StartAuction err", zap.Error(err))
		return err
	}

	return nil
}

func (u *contractUseCase) EndAuction(ctx context.Context) error {
	address := common.HexToAddress(u.cfg.Contract.SellerAddress)

	privateKey, err := crypto.HexToECDSA(u.cfg.Contract.SellerPK)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	ethereumClient, err := ethclient.Dial(u.cfg.Contract.Rpc)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	nonce, err := ethereumClient.PendingNonceAt(ctx, address)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	gasPrice, err := ethereumClient.SuggestGasPrice(ctx)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	//gasTip, err := ethereumClient.SuggestGasTipCap(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//estimateGas, err := ethereumClient.EstimateGas(ctx, ethereum.CallMsg{
	//	To:       &address,
	//	GasPrice: gasPrice,
	//	Value:    big.NewInt(0),
	//	Data:     []byte{},
	//},
	//)

	seller, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(u.cfg.Contract.ChainID))
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	myContract, err := contract.NewContract(common.HexToAddress(u.cfg.Contract.ContractAddress), ethereumClient)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	tx, err := myContract.StartAuction(&bind.TransactOpts{
		From:     address,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   seller.Signer,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: 100000,
		Context:  ctx,
		NoSend:   true,
	})
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(u.cfg.Contract.ChainID)), privateKey)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	err = ethereumClient.SendTransaction(ctx, signedTx)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	return nil
}

func (u *contractUseCase) MakeBid(ctx context.Context, bid *contract.Bid) error {
	address := common.HexToAddress(u.cfg.Contract.User2Address)

	privateKey, err := crypto.HexToECDSA(u.cfg.Contract.User2PK)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	ethereumClient, err := ethclient.Dial(u.cfg.Contract.Rpc)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	nonce, err := ethereumClient.PendingNonceAt(ctx, address)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	gasPrice, err := ethereumClient.SuggestGasPrice(ctx)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	//gasTip, err := ethereumClient.SuggestGasTipCap(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//estimateGas, err := ethereumClient.EstimateGas(ctx, ethereum.CallMsg{
	//	To:       &address,
	//	GasPrice: gasPrice,
	//	Value:    big.NewInt(0),
	//	Data:     []byte{},
	//},
	//)

	user2, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(u.cfg.Contract.ChainID))
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	myContract, err := contract.NewContract(common.HexToAddress(u.cfg.Contract.ContractAddress), ethereumClient)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	tx, err := myContract.MakeBid(&bind.TransactOpts{
		From:     address,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   user2.Signer,
		Value:    big.NewInt(10000000000000000),
		GasPrice: gasPrice,
		GasLimit: 100000,
		Context:  ctx,
		NoSend:   true,
	}, &bid.OneBid)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(u.cfg.Contract.ChainID)), privateKey)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	err = ethereumClient.SendTransaction(ctx, signedTx)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	return nil
}

func (u *contractUseCase) BidWin(ctx context.Context) error {
	address := common.HexToAddress(u.cfg.Contract.User3Address)

	privateKey, err := crypto.HexToECDSA(u.cfg.Contract.User3PK)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	ethereumClient, err := ethclient.Dial(u.cfg.Contract.Rpc)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	nonce, err := ethereumClient.PendingNonceAt(ctx, address)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	gasPrice, err := ethereumClient.SuggestGasPrice(ctx)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	//gasTip, err := ethereumClient.SuggestGasTipCap(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//estimateGas, err := ethereumClient.EstimateGas(ctx, ethereum.CallMsg{
	//	To:       &address,
	//	GasPrice: gasPrice,
	//	Value:    big.NewInt(0),
	//	Data:     []byte{},
	//},
	//)

	user3, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(u.cfg.Contract.ChainID))
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	myContract, err := contract.NewContract(common.HexToAddress(u.cfg.Contract.ContractAddress), ethereumClient)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	tx, err := myContract.StartAuction(&bind.TransactOpts{
		From:     address,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   user3.Signer,
		Value:    big.NewInt(10000000000000000),
		GasPrice: gasPrice,
		GasLimit: 100000,
		Context:  ctx,
		NoSend:   true,
	})
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(u.cfg.Contract.ChainID)), privateKey)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	err = ethereumClient.SendTransaction(ctx, signedTx)
	if err != nil {
		u.logger.Error("EndAuction err", zap.Error(err))
		return err
	}

	return nil
}
