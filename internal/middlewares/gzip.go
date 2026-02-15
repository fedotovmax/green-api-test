package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/fedotovmax/green-api-test/internal/keys"
)

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !strings.Contains(r.Header.Get(keys.HeaderAcceptEncoding), keys.EncodingGzip) {
			next.ServeHTTP(w, r)
			return
		}

		if strings.Contains(r.Header.Get(keys.HeaderAccept), keys.ContentTypeTextEventStream) {
			next.ServeHTTP(w, r)
			return
		}

		gz := gzip.NewWriter(w)
		defer gz.Close()

		w.Header().Set(keys.HeaderContentEncoding, keys.EncodingGzip)
		wrw := &gzipResponseWriter{Writer: gz, ResponseWriter: w}

		next.ServeHTTP(wrw, r)
	})
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
