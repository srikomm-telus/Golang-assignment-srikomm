package main

import (
	mockclient "Golang-assignment-srikomm/client/mocks"
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/models"
	mockstore "Golang-assignment-srikomm/store/mocks"
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCryptoPriceServiceV2_GetCryptoPrice(t *testing.T) {

	ctx := context.Background()

	t.Run("price is successfully fetched from cache", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCoinDeskClient := mockclient.NewMockCoinDeskClientInterface(mockCtrl)
		mockCryptonatorClient := mockclient.NewMockCryptonatorClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCoinDeskClient, mockCryptonatorClient, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}
		response := &models.CryptoPriceServiceResponse{
			IsFromCache: true,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Times(1).Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any(), gomock.Any()).Times(0)
		mockCoinDeskClient.EXPECT().GetBTCCurrentPrice().Times(0)
		mockCryptonatorClient.EXPECT().GetETHCurrentPrice().Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("BTC price is successfully fetched from downstream coinDeskClient", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCoinDeskClient := mockclient.NewMockCoinDeskClientInterface(mockCtrl)
		mockCryptonatorClient := mockclient.NewMockCryptonatorClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCoinDeskClient, mockCryptonatorClient, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}
		mockCrypto := models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice)

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCoinDeskClient.EXPECT().GetBTCCurrentPrice().Times(1).Return(mockCrypto, nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, mockCrypto).Times(1).Return(nil)
		mockCryptonatorClient.EXPECT().GetETHCurrentPrice().Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("failed to fetch BTC price as no cache & downstream coinDeskClient throws error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCoinDeskClient := mockclient.NewMockCoinDeskClientInterface(mockCtrl)
		mockCryptonatorClient := mockclient.NewMockCryptonatorClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCoinDeskClient, mockCryptonatorClient, mockCryptoStorageInterface)

		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCoinDeskClient.EXPECT().GetBTCCurrentPrice().Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any(), gomock.Any()).Times(0)
		mockCryptonatorClient.EXPECT().GetETHCurrentPrice().Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, models.CryptoPriceServiceResponse{}, got)
		assert.Equal(t, fmt.Errorf("intentional error"), err)
	})

	t.Run("BTC price is successfully fetched from downstream coinDeskClient, but setting price in cache fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCoinDeskClient := mockclient.NewMockCoinDeskClientInterface(mockCtrl)
		mockCryptonatorClient := mockclient.NewMockCryptonatorClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCoinDeskClient, mockCryptonatorClient, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}
		mockCrypto := models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice)

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCoinDeskClient.EXPECT().GetBTCCurrentPrice().Times(1).Return(mockCrypto, nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, mockCrypto).Times(1).Return(
			fmt.Errorf("intentional error"))
		mockCryptonatorClient.EXPECT().GetETHCurrentPrice().Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("ETH price is successfully fetched from downstream cryptonatorClient", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCoinDeskClient := mockclient.NewMockCoinDeskClientInterface(mockCtrl)
		mockCryptonatorClient := mockclient.NewMockCryptonatorClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCoinDeskClient, mockCryptonatorClient, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}
		mockCrypto := models.NewCrypto(constants.ETHEREUM_IDENTIFIER, mockCryptoPrice)

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.ETHEREUM_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.ETHEREUM_IDENTIFIER).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptonatorClient.EXPECT().GetETHCurrentPrice().Times(1).Return(mockCrypto, nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, mockCrypto).Times(1).Return(nil)
		mockCoinDeskClient.EXPECT().GetBTCCurrentPrice().Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.ETHEREUM_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("failed to fetch ETH price as no cache & downstream cryptonatorClient throws error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCoinDeskClient := mockclient.NewMockCoinDeskClientInterface(mockCtrl)
		mockCryptonatorClient := mockclient.NewMockCryptonatorClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCoinDeskClient, mockCryptonatorClient, mockCryptoStorageInterface)

		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.ETHEREUM_IDENTIFIER).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptonatorClient.EXPECT().GetETHCurrentPrice().Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any(), gomock.Any()).Times(0)
		mockCoinDeskClient.EXPECT().GetBTCCurrentPrice().Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.ETHEREUM_IDENTIFIER)
		assert.Equal(t, models.CryptoPriceServiceResponse{}, got)
		assert.Equal(t, fmt.Errorf("intentional error"), err)
	})

	t.Run("ETH price is successfully fetched from downstream cryptonatorClient, but setting price in cache fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCoinDeskClient := mockclient.NewMockCoinDeskClientInterface(mockCtrl)
		mockCryptonatorClient := mockclient.NewMockCryptonatorClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCoinDeskClient, mockCryptonatorClient, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}
		mockCrypto := models.NewCrypto(constants.ETHEREUM_IDENTIFIER, mockCryptoPrice)

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.ETHEREUM_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.ETHEREUM_IDENTIFIER).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptonatorClient.EXPECT().GetETHCurrentPrice().Times(1).Return(mockCrypto, nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, mockCrypto).Times(1).Return(
			fmt.Errorf("intentional error"))
		mockCoinDeskClient.EXPECT().GetBTCCurrentPrice().Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.ETHEREUM_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

}
