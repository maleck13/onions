package web

// here we depend on external things such as http, this is an outer layer
import (
	"net/http"

	"github.com/maleck13/onions/pkg/shop/catalog"
)

type CatalogHandler struct {
	Catalog *catalog.Service
}

func (ch *CatalogHandler) Catalog(rw http.ResponseWriter, req *http.Request) {}
