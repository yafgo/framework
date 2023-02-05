package mock

import (
	accessmocks "github.com/yafgo/framework/contracts/auth/access/mocks"
	authmocks "github.com/yafgo/framework/contracts/auth/mocks"
	cachemocks "github.com/yafgo/framework/contracts/cache/mocks"
	configmocks "github.com/yafgo/framework/contracts/config/mocks"
	consolemocks "github.com/yafgo/framework/contracts/console/mocks"
	ormmocks "github.com/yafgo/framework/contracts/database/orm/mocks"
	eventmocks "github.com/yafgo/framework/contracts/event/mocks"
	filesystemmocks "github.com/yafgo/framework/contracts/filesystem/mocks"
	grpcmocks "github.com/yafgo/framework/contracts/grpc/mocks"
	mailmocks "github.com/yafgo/framework/contracts/mail/mocks"
	queuemocks "github.com/yafgo/framework/contracts/queue/mocks"
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

func Orm() (*ormmocks.Orm, *ormmocks.DB, *ormmocks.Transaction, *ormmocks.Association) {
	mockOrm := &ormmocks.Orm{}
	facades.Orm = mockOrm

	return mockOrm, &ormmocks.DB{}, &ormmocks.Transaction{}, &ormmocks.Association{}
}

func Event() (*eventmocks.Instance, *eventmocks.Task) {
	mockEvent := &eventmocks.Instance{}
	facades.Event = mockEvent

	return mockEvent, &eventmocks.Task{}
}

func Log() {
	facades.Log = log.NewApplication(log.NewTestWriter())
}

func Mail() *mailmocks.Mail {
	mockMail := &mailmocks.Mail{}
	facades.Mail = mockMail

	return mockMail
}

func Queue() (*queuemocks.Queue, *queuemocks.Task) {
	mockQueue := &queuemocks.Queue{}
	facades.Queue = mockQueue

	return mockQueue, &queuemocks.Task{}
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

func Auth() *authmocks.Auth {
	mockAuth := &authmocks.Auth{}
	facades.Auth = mockAuth

	return mockAuth
}

func Gate() *accessmocks.Gate {
	mockGate := &accessmocks.Gate{}
	facades.Gate = mockGate

	return mockGate
}

func Grpc() *grpcmocks.Grpc {
	mockGrpc := &grpcmocks.Grpc{}
	facades.Grpc = mockGrpc

	return mockGrpc
}
