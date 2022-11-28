package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrat-muslim/note-app/api/v1"
	"github.com/ibrat-muslim/note-app/config"
	"github.com/ibrat-muslim/note-app/storage"
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	_ = v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	return router
}
