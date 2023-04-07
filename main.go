package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sxc/oishifood/controllers"
	"github.com/sxc/oishifood/models"
	"github.com/sxc/oishifood/templates"
	"github.com/sxc/oishifood/views"
	// _ "github.com/jackc/pgx/v4/stdlib"
)

// type PostgresConfig struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	Database string
// 	SSLMode  string
// }

// func (cfg PostgresConfig) String() string {
// 	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
// }

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	// Create a table...
	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT
	// 	);

	// 	CREATE TABLE IF NOT EXISTS orders (
	// 		id SERIAL PRIMARY KEY,
	// 		user_id INT NOT NULL,
	// 		amount INT,
	// 		description TEXT
	// 		);
	// 		`)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Created tables")

	// Insert a row...

	// _, err = db.Exec(`INSERT INTO users
	// (email, password_hash)
	// VALUES
	//  ('Jordan333@example.com', 'abc123');  Insert into orders (user_id, amount, description)
	//  values (1, 100, 'test');`)

	// if err != nil {
	// 	panic(err)
	// }

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

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
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS,
		"signup.gohtml", "tailwind.gohtml"))

	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS,
		"signin.gohtml", "tailwind.gohtml"))

	// usersC.Templates.Create = views.Must(views.ParseFS(templates.FS,
	// 	"signup.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.New)
	r.Get("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)

	defer db.Close()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Page Not Found not found", http.StatusNotFound)
	})
	fmt.Println("Server is running on port 3000")

	http.ListenAndServe(":3000", r)
}
