package main

import (
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// HttpHandleFunc is http handle function
type HttpHandleFunc func(w http.ResponseWriter, r *http.Request)

// NewRouters create Routers from reader
func NewRouters(r io.Reader) (*Routers, error) {
	if r == nil {
		return nil, errors.New("reader is nil")
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var rs Routers
	if err := yaml.Unmarshal(data, &rs); err != nil {
		return nil, err
	}

	return &rs, nil
}

func main() {
	rf := flag.String("router", "", "router file (support yaml)")
	flag.Parse()

	if len(*rf) == 0 {
		logrus.Errorf("router file couldn't be empty. please set -router")
		os.Exit(0)
	}

	f, err := os.OpenFile(*rf, os.O_RDONLY, os.ModePerm)
	if err != nil {
		logrus.Errorf("open file error: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	rs, err := NewRouters(f)
	if err != nil {
		logrus.Errorf("read routers error: %v", err)
		os.Exit(1)
	}

	logrus.Infof("get routers: %+v", rs)

	// handle routers
	for _, r := range rs.Rs {
		http.HandleFunc("/"+r.URI, handleRouter(r))
	}

	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRouter(router Router) HttpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("router uri: %s, response file: %s", router.URI, router.Response)

		w.Write([]byte("success"))
	}
}
