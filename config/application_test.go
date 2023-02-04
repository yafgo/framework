package config

import (
	"testing"
	"time"

	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	file_helper "github.com/yafgo/framework/support/file"
	"github.com/yafgo/framework/testing/file"
)

type ApplicationTestSuite struct {
	suite.Suite
	config    *Application
	testCount uint32
}

func TestApplicationTestSuite(t *testing.T) {
	envPath := "./../.env"
	if !file_helper.Exists(envPath) {
		file.CreateEnv(envPath)
	}

	suite.Run(t, &ApplicationTestSuite{
		config: NewApplication(envPath),
	})

	assert.Equal(t, true, file_helper.Remove(envPath))

	color.Grayln("SetupTest Done")
}

func (s *ApplicationTestSuite) SetupTest() {
	color.Grayln("SetupTest")
}

func (s *ApplicationTestSuite) TearDownTest() {
	s.testCount++
	color.Grayf("TearDownTest test count:%d\n", s.testCount)
}

func (s *ApplicationTestSuite) TestEnv() {
	s.Equal("yafgo", s.config.GetString("APP_NAME"))
	s.Equal("127.0.0.1", s.config.GetString("DB_HOST", "127.0.0.1"))
}

func (s *ApplicationTestSuite) TestAdd() {
	s.config.Add("app", map[string]any{
		"env": "local",
	})

	s.Equal("local", s.config.GetString("app.env"))
}

func (s *ApplicationTestSuite) TestGet() {
	s.Equal("yafgo", s.config.Get("APP_NAME").(string))
}

func (s *ApplicationTestSuite) TestGetString() {
	s.config.Add("database", map[string]any{
		"default": s.config.Env("DB_CONNECTION", "mysql"),
		"connections": map[string]any{
			"mysql": map[string]any{
				"host": s.config.Env("DB_HOST", "127.0.0.1"),
			},
		},
	})

	s.Equal("yafgo", s.config.GetString("APP_NAME"))
	s.Equal("127.0.0.1", s.config.GetString("database.connections.mysql.host"))
	s.Equal("mysql", s.config.GetString("database.default"))
}

func (s *ApplicationTestSuite) TestGetInt() {
	s.Equal(3306, s.config.GetInt("DB_PORT"))
}

func (s *ApplicationTestSuite) TestGetFloat64() {
	s.Equal(3.1415926, s.config.GetFloat64("MY_PI"))
}

func (s *ApplicationTestSuite) TestGetDuration() {
	s.Equal(time.Second*30, s.config.GetDuration("MY_TIMEOUT"))
	s.Equal((time.Hour*2)+(time.Minute*30), s.config.GetDuration("CACHE_DURATION"))
}

func (s *ApplicationTestSuite) TestGetBool() {
	s.Equal(true, s.config.GetBool("APP_DEBUG"))
}
