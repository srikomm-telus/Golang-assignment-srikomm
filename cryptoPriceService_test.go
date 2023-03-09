package main

import (
	mockclient "Golang-assignment-srikomm/client/mocks"
	"Golang-assignment-srikomm/constants"
	"Golang-assignment-srikomm/models"
	mockstore "Golang-assignment-srikomm/store/mocks"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCryptoPriceService_GetCryptoPrice(t *testing.T) {
	t.Run("price is successfully fetched from cache", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceService(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: true,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER).Return(models.NewCrypto(
			constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("price is successfully fetched from downstream", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceService(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER).Return(models.Crypto{},
			fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any()).Times(1).Return(true, nil)

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})

	t.Run("both storage and downstream layers throws errors", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceService(mockCryptoClientInterface, mockCryptoStorageInterface)

		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER).Return(models.Crypto{},
			fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Return(models.Crypto{}, fmt.Errorf("intentional error"))
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any()).Times(0)

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, nil, got)
		assert.Equal(t, err, fmt.Errorf("intentional error"))
	})

	t.Run("cache doesn't work, downstream call succeeds but setting price in cache fails", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCryptoStorageInterface := mockstore.NewMockCryptoStorageInterface(mockCtrl)
		mockCryptoClientInterface := mockclient.NewMockCryptoClientInterface(mockCtrl)
		mockCryptoPriceService := NewCryptoPriceService(mockCryptoClientInterface, mockCryptoStorageInterface)
		mockCryptoPrice := map[string]string{
			"USD": "30",
			"EUR": "40",
		}

		response := &models.CryptoPriceServiceResponse{
			IsFromCache: false,
			CryptoName:  constants.BITCOIN_IDENTIFIER,
			Data:        mockCryptoPrice,
		}
		mockCryptoStorageInterface.EXPECT().GetCryptoPrice(constants.BITCOIN_IDENTIFIER).Return(models.Crypto{},
			fmt.Errorf("intentional error"))
		mockCryptoClientInterface.EXPECT().GetCurrentPrice().Return(
			models.NewCrypto(constants.BITCOIN_IDENTIFIER, mockCryptoPrice), nil)
		mockCryptoStorageInterface.EXPECT().SetCryptoPrice(gomock.Any()).Times(1).Return(false,
			fmt.Errorf("intentional error"))

		got, err := mockCryptoPriceService.GetCryptoPrice(constants.BITCOIN_IDENTIFIER)
		assert.Equal(t, response, got)
		assert.Equal(t, nil, err)
	})
}
