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

func TestCryptoPriceService_GetCryptoPrice(t *testing.T) {

	ctx := context.Background()

	t.Run("price is successfully fetched from cache", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceForTest(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: true,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Return(models.NewCrypto(
			constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(0)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, gomock.Any()).Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("price is successfully fetched from downstream", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceForTest(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Times(1).Return(models.Crypto{},
			fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(1).Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, gomock.Any()).Times(1).Return(nil)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("both storage and downstream layers throws errors", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceForTest(mockCryptoClientInterface, mockCryptoStorageInterface)

		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Times(1).Return(models.Crypto{},
			fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(1).Return(models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, gomock.Any()).Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, models.CryptoPriceServiceResponse{}, got)
		assert.Equal(t, err, fmt.Errorf("intentional error"))
	})

	t.Run("price is not set in cache, downstream call succeeds but setting price in cache fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceForTest(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER).Times(1).Return(models.Crypto{},
			fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(1).Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(ctx, gomock.Any()).Times(1).Return(
			fmt.Errorf("intentional error"))

		got, err := mockCryptoPriceService.GetCryptoPrice(ctx, constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})
}
