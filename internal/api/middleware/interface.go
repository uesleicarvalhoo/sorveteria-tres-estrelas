package middleware

import "net/http"

type Middleware = func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
