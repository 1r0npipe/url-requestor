package server

import (
	"net/http"
	"os"
	"time"

	"github.com/1r0npipe/url-requestor/pkg/config"
	server_errors "github.com/1r0npipe/url-requestor/pkg/errors"
	"github.com/1r0npipe/url-requestor/pkg/handler"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type RequestServer struct {
	logger     *zerolog.Logger
	httpServer *http.Server
	config     *config.Config
}

func Init(config *config.Config, logger *zerolog.Logger) *RequestServer {
	return &RequestServer{
		logger: logger,
		httpServer: &http.Server{
			Addr:         config.Server.Address + ":" + config.Server.Port,
			Handler:      nil,
			ReadTimeout:  time.Duration(config.Server.Timeout) * time.Second,
			WriteTimeout: time.Duration(config.Server.Timeout) * time.Second,
		},
		config: config,
	}
}
func (s *RequestServer) Handler() http.Handler {

	// section for router description
	router := mux.NewRouter()
	router.HandleFunc("/healthz", s.HealthCheck).Methods("GET")
	router.HandleFunc("/request", s.LoggerWrapper(s.HandleRequests)).Methods("POST")
	return router
}

func (s *RequestServer) HandleRequests(w http.ResponseWriter, r *http.Request) {
	handler.URLRequestsHandler(w, r)
}

func (s *RequestServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	handler.HealthCheck(w, r)
}

func (s *RequestServer) Run() error {

	s.httpServer.Handler = s.Handler()
	if err := s.httpServer.ListenAndServe(); err != nil {
		return server_errors.ErrStartHTTPServer
	}
	return nil
}
func (s *RequestServer) LoggerWrapper(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Info().Msgf("request from %s with body: %s", r.RemoteAddr, r.Body)
		next(rw, r)
	}
}

func (s *RequestServer) Close() {

}

func NewLogger(logLevel string) *zerolog.Logger {
	var level zerolog.Level
	switch {
	case logLevel == "debug":
		level = zerolog.DebugLevel
	case logLevel == "error":
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &logger
}
