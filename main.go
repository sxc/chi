package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sxc/oishifood/controllers"
	"github.com/sxc/oishifood/migrations"
	"github.com/sxc/oishifood/models"
	"github.com/sxc/oishifood/templates"
	"github.com/sxc/oishifood/views"
)

func main() {

	// Setup the database...
	cfg := models.DefaultPostgresConfig()
	fmt.Println(cfg.String())
	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	sessionService := models.SessionService{
		DB: db,
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Setup the routes...

	r := chi.NewRouter()

	// These middleware are used for every request

	r.Use(csrfMw)
	r.Use(umw.SetUser)
	// now we setup routes

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	userService := models.UserService{
		DB: db,
	}
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS,
		"signup.gohtml", "tailwind.gohtml"))

	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS,
		"signin.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Get("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})

	// defer db.Close()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Page Not Found not found", http.StatusNotFound)
	})

	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	fmt.Println("Server is running on port 3000")

	csrfKey := []byte("very-secret")
	// TODO: Fix this before deploying to production
	// csrfMw := csrf.Protect(csrfKey, csrf.Secure(false))
	// r.Use(csrfMiddleware)
	// http.ListenAndServe(":3000", csrfMiddleware(r))
	http.ListenAndServe(":3000", r)
}
