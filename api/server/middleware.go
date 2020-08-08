package server

import (
	"github.com/vietnamz/prime-generator/api/server/httputils"
)

func (s *Server) handlerWithGlobalMiddlewares( handler httputils.APIFunc) httputils.APIFunc {
	next := handler
	for _, m := range s.middlewares {
		next = m.WrapHandler(next)
	}
	return next
}
