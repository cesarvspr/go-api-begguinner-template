package handlers

import (
	"net/http"

	"github.com/cesarvspr/golang-modules/product-api/data"
)

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProduct()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall json", http.StatusInternalServerError)
	}
}
