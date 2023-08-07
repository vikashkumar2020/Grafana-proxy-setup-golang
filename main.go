package main

import (
	// "net/http"
	"net/http/httputil"
	"net/url"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	// Replace with the Grafana server URL
	grafanaURL, err := url.Parse("http://localhost:4000")
	if err != nil {
		panic(err)
	}

	// Create a reverse proxy to forward requests to the Grafana API
	proxy := httputil.NewSingleHostReverseProxy(grafanaURL)

	// Set up the Gin router
	router := gin.Default()

	// Use the CORS middleware
	router.Use(func(c *gin.Context) {
		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:3000"}, // Replace with your frontend domain
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Content-Type", "Authorization"},
		})
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	// Attach the reverse proxy to the router
	router.POST("/api/dashboards/db", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// Attach the reverse proxy to the router
	router.GET("/api/dashboards/uid/:uid", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// Start the server
	port := ":8000" // Replace with your desired port number
	println("Proxy server running on port", port)
	err = router.Run(port)
	if err != nil {
		panic(err)
	}
}

