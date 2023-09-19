package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

var visitNumber = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "visit_number_total",
    Help: "Total visit number",
})

func init() {
  prometheus.MustRegister(visitNumber)
}

func getHelloWorld(c *gin.Context) {
  visitNumber.Inc()
  c.IndentedJSON(http.StatusOK, "hello world")
}

func prometheusHandler() gin.HandlerFunc {
  h := promhttp.Handler()

  return func(c *gin.Context) {
      h.ServeHTTP(c.Writer, c.Request)
  }
}

func main() {
  r := gin.Default()
  r.GET("/", getHelloWorld)
  r.GET("/metrics", prometheusHandler())
  r.Run("0.0.0.0:8080")
}