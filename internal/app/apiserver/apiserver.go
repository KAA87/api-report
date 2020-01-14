package apiserver

import (
	"encoding/json"
	"github.com/KAA87/api-report.git/store"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr,s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter()  {
	s.router.HandleFunc("/hello",s.handleHello())

	s.router.HandleFunc("/get-test",s.handleGetTest())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,"Hello")
	}
}

func (s *APIServer) handleGetTest() http.HandlerFunc  {

	type request struct {
		TestId int `json:"testId,string"`
		GeneralState int8 `json:"generalState,string"`
		Dynamics int8 `json:"dynamics,string"`
		Projection int8 `json:"projection,string"`
		Spine int8 `json:"spine,string"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w,r,http.StatusBadRequest,err)
			return
		}

		t, err := s.store.CalculationDiagnosticTests().FindByDiagnosticTestId(req.TestId)
		if err != nil {
			s.error(w,r,http.StatusNotFound,err)
			return
		}

		s.respond(w,r,http.StatusOK,t)
	}
}

func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w,r,code,map[string]string{"error":err.Error()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}){
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}