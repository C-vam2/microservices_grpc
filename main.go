package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"log"
	"net/http"

	"github.com/microservices_grpc/data"
	"github.com/microservices_grpc/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	v := &data.Validation{}
	ph := handlers.NewProducts(l, v)

	router := gin.Default()

	router.GET("/", ph.ListAll)
	router.PUT("/:id", ph.MiddlewareValicateProduct(), ph.Update)
	router.POST("/", ph.MiddlewareValicateProduct(), ph.Create)
	router.StaticFile("/swagger.yaml", "./swagger.yaml")

	s := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
