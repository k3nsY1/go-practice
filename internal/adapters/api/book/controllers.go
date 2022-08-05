package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/k3nsY1/PracticeGo/internal/adapters"
	"github.com/k3nsY1/PracticeGo/internal/book"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService book.Service
}

func NewHandler(service book.Service) adapters.Handler {
	return &handler{bookService: service}
}
func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}
func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// books:= 	h.bookService.GetAllBooks(context.Background(), 0, 0)
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}
