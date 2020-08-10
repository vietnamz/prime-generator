package prime

import (
	"bytes"
	"context"
	"net/http"
)

func (s *primeRoute) primeHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Pragma", "no-cache")

	if r.Method == http.MethodHead {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Content-Length", "0")
	}
	key, ok := r.URL.Query()["number"]
	if ok == false {
		w.WriteHeader(400)
		_, err := w.Write(bytes.NewBufferString("number is required").Bytes())
		return err
	}
	result := s.D.PrimeSrv.TakeLargestPrimes(key[0])
	_, err := w.Write(bytes.NewBufferString(result).Bytes())
	return err
}