package server

import "github.com/gin-gonic/gin"

type API struct {
	listenAddr string
}

func NewAPI(listenAddr string) *API {
	return &API{
		listenAddr: listenAddr,
	}
}

func (a *API) Start() error {

	r := gin.Default()

	return r.Run(a.listenAddr)
}
