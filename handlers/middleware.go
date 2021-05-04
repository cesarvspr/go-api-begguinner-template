package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cesarvspr/golang-modules/product-api/data"
)


func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(
				rw,
				fmt.Sprintf("error validatindg product", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
