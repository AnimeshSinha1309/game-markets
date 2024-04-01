package main

import (
    "context"
	"net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
        val, err1 := rdb.Get(ctx, "counter").Result()
        if err1 != nil {
            val = "1"
        }
        previous, err2 := strconv.Atoi(val)
        if err2 != nil {
            panic(err2)
        }

        err3 := rdb.Set(ctx, "counter", previous + 1, 0).Err()
        if err3 != nil {
            panic(err3)
        }

        c.JSON(http.StatusOK, gin.H{
            "message": previous + 1,
        })
    })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
