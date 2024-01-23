package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var signature = map[string]interface{}{
			"id":        "65ac024b964ac52cb250f0a0",
			"roles":     []string{"Admin", "User"},
			"firstName": "Tran",
			"lastName":  "Tinh",
		}

		ctx := context.WithValue(r.Context(), "signature", signature)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
