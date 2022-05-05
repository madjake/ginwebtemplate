package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {

	config := loadJSONConfig("configs/local.json")

	if config.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	if config.TLS {
		r.Use(useTLS(config))
	}

	r.SetTrustedProxies(nil)

	RegisterRoutes(config, r)

	if config.TLS {
		r.RunTLS(config.Host+":"+config.Port, config.TLSCertFile, config.TLSKeyFile)
	} else {
		r.Run(config.Host + ":" + config.Port)
	}
}

func useTLS(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     config.Host + ":" + config.Port,
		})

		err := middleware.Process(c.Writer, c.Request)

		if err != nil {
			panic(err)
		}

		c.Next()
	}
}
