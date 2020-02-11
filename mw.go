package mw

// WithHeader returns an http.Handler that wraps the given http.Handler
// and adds a header with the given key and value
func WithHeader(key, value string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header.Add(key, value)
		next(w, r)
	})
}
