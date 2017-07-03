package i18n

import "net/http"

const (
	// LOCALE locale key
	LOCALE = "locale"
)

// Detect locale from http request
func Detect(r *http.Request) string {
	// 1. Check URL arguments.
	if lang := r.URL.Query().Get(LOCALE); lang != "" {
		return lang
	}

	// 2. Get language information from cookies.
	if ck, er := r.Cookie(LOCALE); er == nil {
		return ck.Value
	}

	// 3. Get language information from 'Accept-Language'.
	if al := r.Header.Get("Accept-Language"); len(al) > 4 {
		return al[:5] // Only compare first 5 letters.
	}

	return ""
}
