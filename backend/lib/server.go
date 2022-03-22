package lib

import (
	"context"
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
		logger: NewLogger().SetMethod("Server Setup"),
		mtx:    &sync.Mutex{},
		config: config,
	}
}

func (s *Server) ConfigureRouter(router *httprouter.Router) *Server {
	s.router = router
	return s
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
