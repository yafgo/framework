package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlural(t *testing.T) {
	for k, v := range map[string]string{
		"user":  "users",
		"users": "users",
		"class": "classes",
		"man":   "men",
		"woman": "women",
	} {
		assert.Equal(t, v, Plural(k))
	}
}

func TestSingular(t *testing.T) {
	for k, v := range map[string]string{
		"users":   "user",
		"classes": "class",
		"men":     "man",
		"women":   "woman",
	} {
		assert.Equal(t, v, Singular(k))
	}
}

func TestSnake(t *testing.T) {
	for k, v := range map[string]string{
		"helloWorld": "hello_world",
		"users":      "users",
		"MyName":     "my_name",
		"MySQL":      "my_sql",
		"OK":         "ok",
	} {
		assert.Equal(t, v, Snake(k))
	}
}

func TestCamel(t *testing.T) {
	for k, v := range map[string]string{
		"hello_world": "HelloWorld",
		"users":       "Users",
		"my_name":     "MyName",
		"my_sql":      "MySql",
		"ok":          "Ok",
	} {
		assert.Equal(t, v, Camel(k))
	}
}

func TestLowerCamel(t *testing.T) {
	for k, v := range map[string]string{
		"hello_world": "helloWorld",
		"HelloWorld":  "helloWorld",
		"users":       "users",
		"my_name":     "myName",
		"my_sql":      "mySql",
		"ok":          "ok",
	} {
		assert.Equal(t, v, LowerCamel(k))
	}
}
