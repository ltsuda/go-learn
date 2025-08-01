package contextz

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log some error
		}
		_, err = fmt.Fprint(w, data)
		if err != nil {
			return // todo: log some error
		}
	}
}
