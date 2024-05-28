package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	usecaseprovider "github.com/pandhu/hehemock/app/providers/usecase"
	"github.com/pandhu/hehemock/config"
	routes "github.com/pandhu/hehemock/routes/api"
)

// HTTPServer represents http server
type HTTPServer struct {
	router http.Handler
}

// Serve serves the http requests to http server
func (hs *HTTPServer) Serve() {
	conf := config.All()
	srv := &http.Server{
		ReadTimeout:       10 * time.Minute,
		ReadHeaderTimeout: 60 * time.Second,
		Addr:              fmt.Sprintf(":%d", conf.Server.Port),
		Handler:           hs.router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// The server will listen to the SIGINT and SIGTERM
	// SIGINT will listen to CTRL-C.
	// SIGTERM will be caught if kill command executed.
	//
	// See:
	// - https://en.wikipedia.org/wiki/Unix_signal
	// - https://www.quora.com/What-is-the-difference-between-the-SIGINT-and-SIGTERM-signals-in-Linux
	// - http://programmergamer.blogspot.co.id/2013/05/clarification-on-sigint-sigterm-sigkill.html
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// InitServer initialize server & setup routes
func InitServer(uc *usecaseprovider.Usecase) *HTTPServer {
	srv := &HTTPServer{}

	srv.router = routes.InitRouter(uc)
	return srv
}
