package main

// 1.提供一个API接口，用户传入一个长url
// 2.在go中生成一个唯一短码，如"aB3xY7"
// 3.将“短码 --> 长url”映射关系存储到redis中，设置过期时间
// 4.用户访问例如：http://my-domain/aB3xY7 时从redis中查找出原始URL，并将用户重定向过去

import (
	"gowebg/shorten"
	// "github.com/redis/go-redis/v9"
	"gowebg/redis"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	//任务1： 完成API接口，接收长URL
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/shorten", func(c *gin.Context) {
		var req struct {
			LongURL string `json:"long_url" binding:"required"`
		}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		// 生成短码
		shortCode := shorten.EncodeBase62(req.LongURL)
		shorturl := FindHost(req.LongURL) + "/" + shortCode

		c.JSON(200, gin.H{
			"original_url": req.LongURL,
			"short_code":   shortCode,
			"short_url":    shorturl,
		})

		// 存储到redis
		rdb := redis.NewClient()
		err := redis.ShortCodeimageUrl(rdb, shortCode, req.LongURL)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to store URL"})
			return
		}

	})

	//任务2： 访问短URL时，重定向到原始URL
	r.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")
		rdb := redis.NewClient()
		longURL, err := rdb.Get(c, shortCode).Result()
		if err != nil {
			c.JSON(404, gin.H{"error": "URL not found"})
			return
		}
		c.Redirect(302, longURL)
	})

	r.Run("127.0.0.1:8082")

}

func FindHost(url string) string {
	re := regexp.MustCompile(`^https?://([^/]+)`)
	m := re.FindStringSubmatch(url)
	if len(m) < 2 {
		return ""
	}
	return m[0]
}
