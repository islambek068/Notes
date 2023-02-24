package routes

import (
	"SuperDuper/controllers"
	"SuperDuper/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()
	/*chi is object that maps URLs to functions*/
	router.Use(middleware.Nosurf)
	router.Get("/", controllers.HomeHandler)
	router.Mount("/todo", todoHandler())
	router.Delete("/todo/delete-completed", controllers.DeleteCompleted)
	//serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}

func todoHandler() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", controllers.FetchTodos)
		r.Post("/", controllers.CreateTodo)
		r.Put("/{id}", controllers.UpdateTodo)
		r.Delete("/{id}", controllers.DeleteOneTodo)
	})

	return rg
}
