package cache

import (
	"context"
	"testing"
	"time"

	"github.com/yafgo/framework/contracts/cache"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func instance() cache.Store {
	return &Redis{redis: redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}),
		prefix: "yafgo_cache:",
		ctx:    context.Background(),
	}
}

func TestGet(t *testing.T) {
	r := instance()

	assert.Equal(t, "default", r.Get("test-get", "default").(string))
	assert.Equal(t, "default", r.Get("test-get", func() interface{} {
		return "default"
	}).(string))
}

func TestGetBool(t *testing.T) {
	r := instance()

	assert.Equal(t, true, r.GetBool("test-get-bool", true))
	assert.Nil(t, r.Put("test-get-bool", true, 2*time.Second))
	assert.Equal(t, true, r.GetBool("test-get-bool", false))
}

func TestGetInt(t *testing.T) {
	r := instance()

	assert.Equal(t, 2, r.GetInt("test-get-int", 2))
	assert.Nil(t, r.Put("test-get-int", 3, 2*time.Second))
	assert.Equal(t, 3, r.GetInt("test-get-int", 2))
}

func TestGetString(t *testing.T) {
	r := instance()

	assert.Equal(t, "2", r.GetString("test-get-string", "2"))
	assert.Nil(t, r.Put("test-get-string", "3", 2*time.Second))
	assert.Equal(t, "3", r.GetString("test-get-string", "2"))
}

func TestHas(t *testing.T) {
	r := instance()

	assert.False(t, r.Has("test-has"))
	err := r.Put("test-has", "yafgo", 5*time.Second)
	assert.Nil(t, err)
	assert.True(t, r.Has("test-has"))
}

func TestPut(t *testing.T) {
	r := instance()

	assert.Nil(t, r.Put("test-put", "yafgo", 5*time.Second))
	assert.True(t, r.Has("test-put"))
	assert.Equal(t, "yafgo", r.Get("test-put", "default"))
}

func TestPull(t *testing.T) {
	r := instance()

	assert.Nil(t, r.Put("test-put", "yafgo", 5*time.Second))
	assert.True(t, r.Has("test-put"))
	assert.Equal(t, "yafgo", r.Get("test-put", "default"))
}

func TestAdd(t *testing.T) {
	r := instance()

	assert.True(t, r.Add("test-add", "yafgo", 5*time.Second))
	assert.True(t, r.Has("test-put"))
	assert.False(t, r.Add("test-add", "yafgo", 5*time.Second))
}

func TestRemember(t *testing.T) {
	r := instance()

	val, err := r.Remember("test-remember", 5*time.Second, func() interface{} {
		return "yafgo"
	})

	assert.Nil(t, err)
	assert.Equal(t, "yafgo", val.(string))
}

func TestRememberForever(t *testing.T) {
	r := instance()

	val, err := r.RememberForever("test-remember-forever", func() interface{} {
		return "yafgo"
	})

	assert.Nil(t, err)
	assert.Equal(t, "yafgo", val.(string))
}

func TestForever(t *testing.T) {
	r := instance()

	val := r.Forever("test-forever", "yafgo")

	assert.True(t, val)
	assert.Equal(t, "yafgo", r.Get("test-forever", nil))
}

func TestForget(t *testing.T) {
	r := instance()

	val := r.Forget("test-forget")
	assert.True(t, val)

	err := r.Put("test-forget", "yafgo", 5*time.Second)
	assert.Nil(t, err)
	assert.True(t, r.Forget("test-forget"))
}

func TestFlush(t *testing.T) {
	r := instance()

	err := r.Put("test-flush", "yafgo", 5*time.Second)
	assert.Nil(t, err)
	assert.Equal(t, "yafgo", r.Get("test-flush", nil).(string))

	r.Flush()
	assert.False(t, r.Has("test-flush"))
}
