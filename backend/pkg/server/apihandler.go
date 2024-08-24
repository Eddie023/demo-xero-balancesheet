package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type TokenAuthorizer interface {
	Authenticate(token string) error
}

type BalanceSheetReportProvider interface {
	BuildReport(ctx context.Context, query map[string]string) (Report, error)
}

// decouple API response with internal platform struct
// for the sake of simplicity this has been skipped.
type JSONSerializer interface {
	BalanceSheetReport(report Report) ([]byte, error)
}

type HandlerV1 struct {
	authenticator              TokenAuthorizer
	balanceSheetReportProvider BalanceSheetReportProvider
}

func NewHandler(authenticator TokenAuthorizer, bsrProvider BalanceSheetReportProvider) http.Handler {
	r := chi.NewRouter()

	h := HandlerV1{
		authenticator:              authenticator,
		balanceSheetReportProvider: bsrProvider,
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(cors.AllowAll().Handler)

	r.Get("/api/reports/balance-sheet", h.getBalanceSheet)

	return r
}

func (h HandlerV1) getBalanceSheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// simple demonstration on how we can utilize dependency injection to decouple
	// diffferent responsibility and make code more testable.
	err := h.authenticator.Authenticate("dummy-user-scopes")
	if err != nil {
		ErrorResponse(ctx, w, fmt.Errorf("authenticating: %w", err))
		return
	}

	report, err := h.balanceSheetReportProvider.BuildReport(ctx, nil)
	if err != nil {
		ErrorResponse(ctx, w, fmt.Errorf("buiding balance sheet reports: %w", err))
		return
	}

	JSONResponse(ctx, w, report, http.StatusOK)
}

// JSONResponse is a simple utility function to write JSON-encoded response.
func JSONResponse(ctx context.Context, w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("marshalling json", "err", err)
	}

	if _, err := w.Write(jsonData); err != nil {
		slog.Error("writing response", "err", err)
	}
}

// ErrorResponse is a simple utility function to return error from our API calls.
// For the purpose of this demo, the functionality is kept simple.
func ErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	slog.ErrorContext(ctx, "api_request_failed", "err", err)

	type errResponse struct {
		Error  string            `json:"error"`
		Fields map[string]string `json:"fields,omitempty"`
	}

	w.WriteHeader(http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(errResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	})
	if err != nil {
		slog.Error("failed to encode", "err", err.Error())
	}
}
