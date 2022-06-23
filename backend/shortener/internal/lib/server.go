package lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"url_shortener/backend/lib"
	"url_shortener/backend/lib/errs"
	"url_shortener/backend/lib/request"
	"url_shortener/backend/shortener/api"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	server    *http.Server
	router    *http.ServeMux
	mtx       *sync.Mutex
	log       *lib.Logger
	db        *gorm.DB
	config    *lib.Config
	isStarted bool
}

func NewServer(config *lib.Config) *Server {
	return &Server{
		log:    lib.NewLogger().SetMethod("Server Setup"),
		mtx:    &sync.Mutex{},
		config: config,
	}
}

func (s *Server) ConfigureRouter() {
	router := http.NewServeMux()
	router.HandleFunc("/create", s.RequestCreateShortUrl)
	router.HandleFunc("/get", s.RequestGetLongtUrl)
	s.router = router
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
	s.log.Infof("Starting http server: %s", s)
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				s.log.Infof("Server closed under request: %v", err)
			} else {
				s.log.Fatalf("Server closed unexpect: %v", err)
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
		s.log.Errorf("Timeout: %v", ctx.Err())
		break
	case <-stop:
		s.log.Info("Finished")
	}
	return nil
}

// Utils

func (s *Server) InitMeta(req request.Request, r *http.Request) {
	req.SetHeader(request.SignatureHeader, r.Header.Get(request.SignatureHeader))
	req.SetHeader(request.RequestHeader, r.Header.Get(request.RequestHeader))
	req.SetMeta(s.db, *s.log)
}

// Controllers

func (s *Server) RequestCreateShortUrl(w http.ResponseWriter, r *http.Request) {
	req := &api.ReqCreateShortUrl{}

	s.InitMeta(req, r)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		lib.Respond(w, r, err)
		return
	}

	if err := req.Authorize(); err != nil {
		lib.Respond(w, r, err)
		return
	}

	if err := req.Validate(); err != nil {
		lib.Respond(w, r, err)
		return
	}

	rpl, err := req.Execute()
	if err != nil {
		lib.Respond(w, r, rpl)
		return
	}

	lib.Respond(w, r, rpl)
}

func (s *Server) RequestGetLongtUrl(w http.ResponseWriter, r *http.Request) {
	req := &api.ReqGetShortUrl{}

	s.InitMeta(req, r)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		lib.Respond(w, r, err)
		return
	}

	if err := req.Authorize(); err != nil {
		lib.Respond(w, r, err)
		return
	}

	if err := req.Validate(); err != nil {
		lib.Respond(w, r, err)
		return
	}

	rpl, err := req.Execute()
	if err != nil {
		lib.Respond(w, r, rpl)
		return
	}

	lib.Respond(w, r, rpl)
}

func (s *Server) execute(req request.Request) (request.Reply, *errs.Error) {
	return nil, nil 
}
