package main

import (
	"crypto/tls"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	r := gin.Default()
	client := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetHeader("Host", "api.openai.com").
		SetBaseURL("https://52.152.96.252").
		SetHeader("Content-Type", "application/json")
	r.Use(func(c *gin.Context) {
		req := client.R().SetContext(c.Request.Context())
		for k, l := range c.Request.Header {
			for _, v := range l {
				req.SetHeader(k, v)
			}
		}
		resp, err := req.SetBody(c.Request.Body).Post(c.Request.URL.Path)
		if err != nil {
			c.Error(err)
			return
		}
		c.Data(resp.StatusCode(), "application/json", resp.Body())
	})
	r.Run(":8087")
}
