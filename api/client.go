/*
A Client for the API.
It initializes the application from Config and contains structures that will be used later.
Method RunServer() is using for the initialization and running.
*/
package api

import (
	"fmt"
	"github.com/cegorah/auth_service/internal/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	Router      *mux.Router
	Middlewares map[string]mux.MiddlewareFunc
	appContext  *LocalContext
}

func LogNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	http.Error(w, "", 404)
	return
}

func (c *Client) RunServer() {
	cfg, err := config.FromEnv("PROD_API")
	errHandler(err)
	fCfg, err := config.FromFile(os.Getenv("CONFIG_PATH") + "/config.json")
	errHandler(err)
	cfg.ExtendFromFile(&fCfg)
	c.appContext = &LocalContext{}
	c.appContext.initContext(&cfg)
	rt := mux.NewRouter()
	srv := &http.Server{
		Handler:      rt,
		Addr:         cfg.ApiCfg.ConnectionString,
		WriteTimeout: time.Duration(cfg.ApiCfg.WriteTimeoutSecond) * time.Second,
		ReadTimeout:  time.Duration(cfg.ApiCfg.ReadTimeoutSecond) * time.Second,
	}
	rt.HandleFunc("/code", c.PostValidationCode()).Methods("POST")
	rt.HandleFunc("/token", c.PostToken()).Methods("POST")
	rt.NotFoundHandler = http.HandlerFunc(LogNotFound)
	log.Fatal(srv.ListenAndServe())
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
