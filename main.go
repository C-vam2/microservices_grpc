package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"log"
	"net/http"

	"github.com/microservices_grpc/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	router := gin.Default()

	router.GET("/", ph.GetProducts)
	router.PUT("/:id", ph.MiddlewareProductValidation, ph.UpdateProduct)
	router.POST("/", ph.MiddlewareProductValidation, ph.AddProduct)

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
