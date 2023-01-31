package mock

import (
	cachemocks "github.com/yafgo/framework/contracts/cache/mocks"
	configmocks "github.com/yafgo/framework/contracts/config/mocks"
	consolemocks "github.com/yafgo/framework/contracts/console/mocks"
	filesystemmocks "github.com/yafgo/framework/contracts/filesystem/mocks"
	validatemocks "github.com/yafgo/framework/contracts/validation/mocks"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/log"
)

func Cache() *cachemocks.Store {
	mockCache := &cachemocks.Store{}
	facades.Cache = mockCache

	return mockCache
}

func Config() *configmocks.Config {
	mockConfig := &configmocks.Config{}
	facades.Config = mockConfig

	return mockConfig
}

func Artisan() *consolemocks.Artisan {
	mockArtisan := &consolemocks.Artisan{}
	facades.Artisan = mockArtisan

	return mockArtisan
}

func Log() {
	facades.Log = log.NewLogrus(nil, log.NewTestWriter())
}

func Storage() (*filesystemmocks.Storage, *filesystemmocks.Driver, *filesystemmocks.File) {
	mockStorage := &filesystemmocks.Storage{}
	mockDriver := &filesystemmocks.Driver{}
	mockFile := &filesystemmocks.File{}
	facades.Storage = mockStorage

	return mockStorage, mockDriver, mockFile
}

func Validation() (*validatemocks.Validation, *validatemocks.Validator, *validatemocks.Errors) {
	mockValidation := &validatemocks.Validation{}
	mockValidator := &validatemocks.Validator{}
	mockErrors := &validatemocks.Errors{}
	facades.Validation = mockValidation

	return mockValidation, mockValidator, mockErrors
}
