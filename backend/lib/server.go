package lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	server    *http.Server
	router    *httprouter.Router
	mtx       *sync.Mutex
	logger    *Logger
	db        *gorm.DB
	config    *Config
	isStarted bool
}

func NewServer(config *Config) *Server {
	return &Server{
		router: httprouter.New(),
		logger: NewLogger(),
		mtx:    &sync.Mutex{},
		config: config,
	}
}

func (s *Server) CreateEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type request struct {
		UrlLong string `json:"url_long"`
	}
	req := &request{}
	w.Header().Set("Content-Type", "application/json")
	l := s.logger.SetMethod("CreateEndpoint")
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		l.Error(err)
		s.error(w, r, http.StatusBadRequest, err)
	}
	l.Infof("%+v", req)
	s.respond(w, r, http.StatusOK, nil)
}

func (s *Server) RedirectEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	l := s.logger.SetMethod("RedirectEndpoint")
	l.Info("id: ", ps.ByName("id"))
	// получили айди и загрузили из бд далее редирект по лонгу
	//http.Redirect(w, r, "https://yandex.ru", http.StatusMovedPermanently)
	s.respond(w, r, http.StatusAccepted, map[string]interface{}{"redirect": "https://yandex.ru"})
}

func (s *Server) ExpandEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//l := s.logger.SetMethod("ExpandEndpoint")
	// читаем из бд по айди
	// возвращаем лонг
	s.error(w, r, http.StatusInternalServerError, errors.New("ERROR_INTERNAL"))
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) ConnectToDB() error {
	// connect to postgres
	var err error
	s.db, err = gorm.Open(postgres.Open(s.config.GetDBConnection()), &gorm.Config{})
	if err != nil {
		return errors.New("unable connect to pg database")
	}
	return nil
}

func (s *Server) String() string {
	return fmt.Sprintf("%s:%s", s.config.Api.Host, s.config.Api.Port)
}

func (s *Server) Run() error {
	s.logger.SetMethod("Main")
	// run the server
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.isStarted {
		return errors.New("Server is already started")
	}
	s.isStarted = true
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.config.Api.Host, s.config.Api.Port),
		Handler: s.router,
	}
	s.logger.Infof("Starting http server: %s", s)
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				s.logger.Infof("Server closed under request: %v", err)
			} else {
				s.logger.Fatalf("Server closed unexpect: %v", err)
			}
			s.isStarted = false
		}
	}()
	time.Sleep(10 * time.Millisecond)
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.SetMethod("Main")

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if !s.isStarted || s.server == nil {
		return errors.New("Server is not started")
	}

	stop := make(chan bool)
	go func() {
		_ = s.server.Shutdown(ctx)
		stop <- true
	}()

	select {
	case <-ctx.Done():
		s.logger.Errorf("Timeout: %v", ctx.Err())
		break
	case <-stop:
		s.logger.Info("Finished")
	}
	return nil
}
