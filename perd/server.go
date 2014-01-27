package perd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Server is a simple http server who listens for incoming requests.
// When request comes, he evals its body using Runner
type Server interface {
	Run()
}

type config struct {
	listen string
}

// NewServer returns new server
func NewServer(listen string, workers map[string]int64, timeout int64) Server {
	runners := map[string]Runner{
		"ruby":   NewRunner(Ruby, workers["ruby"], timeout),
		"nodejs": NewRunner(Nodejs, workers["nodejs"], timeout),
		"golang": NewRunner(Golang, workers["golang"], timeout),
		"python": NewRunner(Python, workers["python"], timeout),
		"c":      NewRunner(C, workers["c"], timeout),
	}
	return &server{&config{listen}, runners}
}

type server struct {
	config  *config
	runners map[string]Runner
}

var (
	ErrUndefinedLang = errors.New("Undefined Language.")
)

func (s *server) Run() {
	// Root path

	http.HandleFunc("/api/evaluate", s.evaluateHandler)

	http.HandleFunc("/api/evaluate/ruby", s.rubyHandler)
	http.HandleFunc("/api/evaluate/nodejs", s.nodejsHandler)
	http.HandleFunc("/api/evaluate/golang", s.golangHandler)
	http.HandleFunc("/api/evaluate/python", s.pythonHandler)
	http.HandleFunc("/api/evaluate/c", s.cHandler)

	log.Println("Listen http on", s.config.listen)
	http.ListenAndServe(s.config.listen, nil)
}

func (s *server) langHandler(w http.ResponseWriter, r *http.Request, lang string) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	res, err := s.eval(lang, string(body))

	if err != nil {
		log.Println(err)
		return
	}

	w.Write(res.Bytes())
}

func (s *server) nodejsHandler(w http.ResponseWriter, r *http.Request) {
	s.langHandler(w, r, "nodejs")
}

func (s *server) rubyHandler(w http.ResponseWriter, r *http.Request) {
	s.langHandler(w, r, "ruby")
}

func (s *server) golangHandler(w http.ResponseWriter, r *http.Request) {
	s.langHandler(w, r, "golang")
}

func (s *server) pythonHandler(w http.ResponseWriter, r *http.Request) {
	s.langHandler(w, r, "python")
}

func (s *server) cHandler(w http.ResponseWriter, r *http.Request) {
	s.langHandler(w, r, "c")
}

type RequestJson struct {
	Lang string `json:"language"`
	Code string `json:"code"`
}

func (s *server) evaluateHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte
	var res Result

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		return
	}

	js := &RequestJson{}
	err = json.Unmarshal(body, js)

	if err != nil {
		log.Println(err)
		return
	}

	res, err = s.eval(js.Lang, js.Code)

	if err != nil {
		log.Println(err)
		return
	}

	w.Write(res.Bytes())
}

func (s *server) eval(lang, code string) (Result, error) {
	runner, ok := s.runners[lang]
	if !ok {
		return nil, ErrUndefinedLang
	}
	return runner.Eval(code), nil
}
