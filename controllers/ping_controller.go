package controllers

import "net/http"

const (
	pong = "pong"
)

var (
	PingControllerInterface pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct {

}


func (p *pingController) Ping(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}