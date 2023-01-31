package log

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"

	configmocks "github.com/yafgo/framework/contracts/config/mocks"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/support/file"
	"github.com/yafgo/framework/support/time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var singleLog = "storage/logs/yafgo.log"
var dailyLog = fmt.Sprintf("storage/logs/yafgo-%s.log", time.Now().Format("2006-01-02"))

func initMockConfig() *configmocks.Config {
	mockConfig := &configmocks.Config{}
	facades.Config = mockConfig

	mockConfig.On("GetString", "logging.default").Return("stack").Once()
	mockConfig.On("GetString", "logging.channels.stack.driver").Return("stack").Once()
	mockConfig.On("Get", "logging.channels.stack.channels").Return([]string{"single", "daily"}).Once()
	mockConfig.On("GetString", "logging.channels.daily.driver").Return("daily").Once()
	mockConfig.On("GetString", "logging.channels.daily.path").Return(singleLog).Once()
	mockConfig.On("GetInt", "logging.channels.daily.days").Return(7).Once()
	mockConfig.On("GetString", "logging.channels.single.driver").Return("single").Once()
	mockConfig.On("GetString", "logging.channels.single.path").Return(singleLog).Once()

	return mockConfig
}

func mockDriverConfig(mockConfig *configmocks.Config) {
	mockConfig.On("GetString", "logging.channels.daily.level").Return("debug").Once()
	mockConfig.On("GetString", "logging.channels.single.level").Return("debug").Once()
	mockConfig.On("GetString", "app.timezone").Return("UTC")
	mockConfig.On("GetString", "app.env").Return("test")
}

func initFacadesLog() {
	logrusInstance := logrusInstance()
	facades.Log = NewLogrus(logrusInstance, NewWriter(logrusInstance.WithContext(context.Background())))
}

type LogrusTestSuite struct {
	suite.Suite
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(LogrusTestSuite))
}

func (s *LogrusTestSuite) SetupTest() {

}

func (s *LogrusTestSuite) TestLogrus() {
	var mockConfig *configmocks.Config

	beforeEach := func() {
		mockConfig = initMockConfig()
	}

	tests := []struct {
		name   string
		setup  func()
		assert func(name string)
	}{
		{
			name: "WithContext",
			setup: func() {
				mockConfig.On("GetString", "logging.channels.daily.level").Return("debug").Once()
				mockConfig.On("GetString", "logging.channels.single.level").Return("debug").Once()

				initFacadesLog()
			},
			assert: func(name string) {
				writer := facades.Log.WithContext(context.Background())
				assert.Equal(s.T(), reflect.TypeOf(writer).String(), reflect.TypeOf(&Writer{}).String(), name)
			},
		},
		{
			name: "Debug",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Debug("Yafgo")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.debug: Yafgo"))
				assert.True(s.T(), file.Contain(dailyLog, "test.debug: Yafgo"))
			},
		},
		{
			name: "No Debug",
			setup: func() {
				mockConfig.On("GetString", "logging.channels.daily.level").Return("info").Once()
				mockConfig.On("GetString", "logging.channels.single.level").Return("info").Once()

				initFacadesLog()
				facades.Log.Debug("Yafgo")
			},
			assert: func(name string) {
				assert.False(s.T(), file.Exists(dailyLog))
				assert.False(s.T(), file.Exists(singleLog))
			},
		},
		{
			name: "Debugf",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Debugf("Yafgo: %s", "World")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.debug: Yafgo: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.debug: Yafgo: World"))
			},
		},
		{
			name: "Info",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Info("Yafgo")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.info: Yafgo"))
				assert.True(s.T(), file.Contain(dailyLog, "test.info: Yafgo"))
			},
		},
		{
			name: "Infof",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Infof("Yafgo: %s", "World")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.info: Yafgo: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.info: Yafgo: World"))
			},
		},
		{
			name: "Warning",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Warning("Yafgo")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.warning: Yafgo"))
				assert.True(s.T(), file.Contain(dailyLog, "test.warning: Yafgo"))
			},
		},
		{
			name: "Warningf",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Warningf("Yafgo: %s", "World")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.warning: Yafgo: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.warning: Yafgo: World"))
			},
		},
		{
			name: "Error",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Error("Yafgo")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.error: Yafgo"))
				assert.True(s.T(), file.Contain(dailyLog, "test.error: Yafgo"))
			},
		},
		{
			name: "Errorf",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
				facades.Log.Errorf("Yafgo: %s", "World")
			},
			assert: func(name string) {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.error: Yafgo: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.error: Yafgo: World"))
			},
		},
		{
			name: "Panic",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
			},
			assert: func(name string) {
				assert.Panics(s.T(), func() {
					facades.Log.Panic("Yafgo")
				})
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.panic: Yafgo"))
				assert.True(s.T(), file.Contain(dailyLog, "test.panic: Yafgo"))
			},
		},
		{
			name: "Panicf",
			setup: func() {
				mockDriverConfig(mockConfig)

				initFacadesLog()
			},
			assert: func(name string) {
				assert.Panics(s.T(), func() {
					facades.Log.Panicf("Yafgo: %s", "World")
				})
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.panic: Yafgo: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.panic: Yafgo: World"))
			},
		},
	}

	for _, test := range tests {
		beforeEach()
		test.setup()
		test.assert(test.name)
		mockConfig.AssertExpectations(s.T())
		file.Remove("storage")
	}
}

func (s *LogrusTestSuite) TestTestWriter() {
	facades.Log = NewLogrus(nil, NewTestWriter())
	assert.Equal(s.T(), facades.Log.WithContext(context.Background()), &TestWriter{})
	assert.NotPanics(s.T(), func() {
		facades.Log.Debug("Yafgo")
		facades.Log.Debugf("Yafgo")
		facades.Log.Info("Yafgo")
		facades.Log.Infof("Yafgo")
		facades.Log.Warning("Yafgo")
		facades.Log.Warningf("Yafgo")
		facades.Log.Error("Yafgo")
		facades.Log.Errorf("Yafgo")
		facades.Log.Fatal("Yafgo")
		facades.Log.Fatalf("Yafgo")
		facades.Log.Panic("Yafgo")
		facades.Log.Panicf("Yafgo")
	})
}

func TestLogrus_Fatal(t *testing.T) {
	mockConfig := initMockConfig()
	mockDriverConfig(mockConfig)
	initFacadesLog()

	if os.Getenv("FATAL") == "1" {
		facades.Log.Fatal("Yafgo")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogrus_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()

	assert.EqualError(t, err, "exit status 1")
	assert.True(t, file.Exists(dailyLog))
	assert.True(t, file.Exists(singleLog))
	assert.True(t, file.Contain(singleLog, "test.fatal: Yafgo"))
	assert.True(t, file.Contain(dailyLog, "test.fatal: Yafgo"))
	file.Remove("storage")
}

func TestLogrus_Fatalf(t *testing.T) {
	mockConfig := initMockConfig()
	mockDriverConfig(mockConfig)
	initFacadesLog()

	if os.Getenv("FATAL") == "1" {
		facades.Log.Fatalf("Yafgo")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogrus_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()

	assert.EqualError(t, err, "exit status 1")
	assert.True(t, file.Exists(dailyLog))
	assert.True(t, file.Exists(singleLog))
	assert.True(t, file.Contain(singleLog, "test.fatal: Yafgo"))
	assert.True(t, file.Contain(dailyLog, "test.fatal: Yafgo"))
	file.Remove("storage")
}
