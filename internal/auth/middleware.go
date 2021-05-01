package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/krish8learn/Graphql_Mysql_GO/internal/pkg/jwtLogin"
	"github.com/krish8learn/Graphql_Mysql_GO/internal/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenstr := header

			username, err := jwtLogin.ParseToken(tokenstr)
			if err != nil {
				http.Error(w, "Invalid Token", http.StatusForbidden)
			}

			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
			}
			user.ID = strconv.Itoa(id)
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
