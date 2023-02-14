package server

import (
	"golang.org/x/net/context"
	"http-type-one/controller"
	db2 "http-type-one/db"
	"http-type-one/repository"
	"http-type-one/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Server() {
	db := db2.NewDB()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := controller.NewHandler(services)

	// init router
	r := handlers.InitRouter()

	//server configuration
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
