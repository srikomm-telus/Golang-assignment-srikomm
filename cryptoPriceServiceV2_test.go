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
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: true,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx).Times(1).Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(0)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any(), gomock.Any()).Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("price is successfully fetched from downstream client", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(1).Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any(), ctx).Times(1).Return(true, nil)

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("both storage and downstream layers throws errors", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any(), ctx).Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx)
		assert.Equal(t, nil, got)
		assert.Equal(t, fmt.Errorf("intentional error"), err)
	})

	t.Run("price is not set in cache, downstream call succeeds but setting price in cache fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceServiceV2ForTest(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx).Times(1).Return(
			models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Times(1).Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any(), gomock.Any()).Times(1).Return(false,
			fmt.Errorf("intentional error"))

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER, ctx)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})
}
