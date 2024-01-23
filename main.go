package main

import (
	"log"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Packs    []Pack `json:"packs"`
}

type Pack struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Name        string    `json:"name"`
	TotalWeight float64   `json:"total_weight"`
	Sections    []Section `json:"sections"`
}

type Section struct {
	Id          int     `json:"id"`
	PackId      int     `json:"pack_id"`
	Name        string  `json:"name"`
	TotalWeight float64 `json:"total_weight"`
	Items       []Item  `json:"items"`
}

type Item struct {
	Id        int     `json:"id"`
	SectionId int     `json:"section_id"`
	Name      string  `json:"name"`
	Weight    float64 `json:"weight"`
	Quantity  int     `json:"quantity"`
}

func main() {
	// Create a database connection, close it before the program exits (defer)
	db, err := sql.Open("sqlite3", "file:luwa.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Create the packs table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS packs (id INTEGER PRIMARY KEY, user_id INTEGER, name TEXT, total_weight REAL)")
	if err != nil {
		log.Fatal(err)
	}

	// Create the sections table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sections (id INTEGER PRIMARY KEY, pack_id INTEGER, name TEXT, total_weight REAL)")
	if err != nil {
		log.Fatal(err)
	}

	// Create the items table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY, section_id INTEGER, name TEXT, weight REAL, quantity INTEGER)")
	if err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	// ! api health check
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.SendString("*** much healthy ***")
	})

	log.Fatal(app.Listen(":3000"))
}
