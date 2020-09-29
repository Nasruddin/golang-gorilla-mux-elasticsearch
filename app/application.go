package app

import (
	"github.com/gorilla/mux"
	"github.com/nasruddin/golang/elastic/clients/elasticsearch"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication()  {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Addr: ":9205",
		/*WriteTimeout: 500 * time.Second,
		ReadTimeout: 2 * time.Second,*/
		//IdleTimeout: 60 * time.Second,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
