package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/eddie023/accounting/pkg/server"
	"github.com/eddie023/accounting/pkg/xeroreports"
)

func main() {
	// Here you will define your configuration settings
	// and inject them to NewHandler function.
	cfg := struct {
		XeroBaseURL string
		Address     string
	}{
		XeroBaseURL: "http://localhost:4000",
		Address:     ":8080",
	}

	xeroBaseURL := os.Getenv("XERO_BASE_URL")
	if xeroBaseURL != "" {
		cfg.XeroBaseURL = xeroBaseURL
	}

	slog.Info("starting service", "xeroclient baseURL", cfg.XeroBaseURL)

	auth := server.NewAuth()
	xeroClient := xeroreports.NewXeroClient(cfg.XeroBaseURL)
	xeroReportsProvider := xeroreports.NewXeroReportsProvider(xeroClient)

	h := server.NewHandler(auth, xeroReportsProvider)
	api := http.Server{
		Addr:    cfg.Address,
		Handler: h,
	}

	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("startup", "service", "accounting", "status", "api router started", "host", "localhost", "port", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// blocking main and waiting for errors
	select {
	case err := <-serverErrors:
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
