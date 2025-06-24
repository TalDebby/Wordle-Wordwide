package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	languagesCrud "server/crud"
	"server/middleware"
	"server/routes"
	languagesService "server/service"
	"sync"
	"time"
)

func initHttpServer() *http.Server {
	languagesRepository := languagesCrud.NewJsonWordsRepository()
	languagesService := languagesService.NewLanguagesService(languagesRepository)

	mux := http.NewServeMux()

	routes.AddRoutes(mux, languagesService)

	defaultMiddlewars := middleware.CreateStack(
		middleware.Logging,
		middleware.Recovery,
		middleware.CORS,
	)

	httpServer := &http.Server{
		Addr:    "localhost:8080",
		Handler: defaultMiddlewars(mux),
	}

	return httpServer
}

func startHttpServer(httpServer *http.Server) {
	fmt.Println("Starting Server")
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}

func shutdownHttpServer(httpServer *http.Server, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// Wait for the context to be done (e.g., on interrupt signal)
	<-ctx.Done()
	fmt.Println("Shutting down server...")
	shutdownCtx := context.Background()
	shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	httpServer := initHttpServer()

	var wg sync.WaitGroup
	wg.Add(1)

	go startHttpServer(httpServer)
	go shutdownHttpServer(httpServer, ctx, &wg)

	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error running server: %v\n", err)
		os.Exit(1)
	}
}
