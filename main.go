package cache

import (
	"os"
	"strconv"

	"github.com/devisettymahidhar315/cache/api"
	"github.com/gin-gonic/gin"
)

func Hello() *gin.Engine {
	var value string
	if len(os.Args) < 2 {
		value = "2"
	} else {
		// Read the argument
		value = os.Args[1]

	}
	length, _ := strconv.Atoi(value)

	// Initialize the Gin router and routes
	r := gin.Default()
	// multi cache
	r.GET("/:key", api.GetCacheValue)
	r.DELETE("/:key", api.DeleteCacheValue)
	r.POST("/:key/:value/:time", func(ctx *gin.Context) {
		api.SetCacheValue(ctx, length) // Pass the length to the handler
	})
	r.GET("/print", api.PrintCache)
	r.DELETE("/all", api.DeleteAll)

	//redis
	r.GET("/redis/:key", api.RedisGet)
	r.DELETE("/redis/:key", api.RedisDeleteValue)
	r.POST("/redis/:key/:value/:time", func(ctx *gin.Context) {
		api.RedisSetValue(ctx, length) // Pass the length to the handler
	})
	r.GET("/redis/print", api.RedisPrint)
	r.DELETE("/redis/all", api.RedisDeleteAll)

	//inmemory
	r.GET("/inmemory/:key", api.InmemoryGet)
	r.DELETE("/inmemory/:key", api.InmemoryDeleteValue)
	r.POST("/inmemory/:key/:value/:time", func(ctx *gin.Context) {
		api.InmemorySetValue(ctx, length) // Pass the length to the handler
	})
	r.GET("/inmemory/print", api.InmemoryPrint)
	r.DELETE("/inmemory/all", api.InmemoryDeleteAll)

	return r
}
