package http

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yafgo/framework/contracts/http"
)

func Background() http.Context {
	return NewGinContext(&gin.Context{})
}

type GinContext struct {
	instance *gin.Context
}

func NewGinContext(ctx *gin.Context) http.Context {
	return &GinContext{ctx}
}

func (c *GinContext) Request() http.Request {
	return NewGinRequest(c)
}

func (c *GinContext) Response() http.Response {
	responseOrigin := c.Value("responseOrigin")
	if responseOrigin != nil {
		return NewGinResponse(c.instance, responseOrigin.(http.ResponseOrigin))
	}

	return NewGinResponse(c.instance, &BodyWriter{ResponseWriter: c.instance.Writer})
}

func (c *GinContext) WithValue(key string, value any) {
	c.instance.Set(key, value)
}

func (c *GinContext) Context() context.Context {
	ctx := context.Background()
	for key, value := range c.instance.Keys {
		ctx = context.WithValue(ctx, key, value)
	}

	return ctx
}

func (c *GinContext) Deadline() (deadline time.Time, ok bool) {
	return c.instance.Deadline()
}

func (c *GinContext) Done() <-chan struct{} {
	return c.instance.Done()
}

func (c *GinContext) Err() error {
	return c.instance.Err()
}

func (c *GinContext) Value(key any) any {
	return c.instance.Value(key)
}

func (c *GinContext) Instance() *gin.Context {
	return c.instance
}
