// package main

// import (
// 	"fiber-postgres/repo"
// 	"fiber-postgres/router"

// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	// Connect Database
// 	db := repo.ConnectDatabase()
// 	defer db.Close()

// 	// Init Fiber App
// 	app := fiber.New()

// 	// Register router
// 	router.InitRouter(app)

// 	// Fiber App listen port
// 	app.Listen(":3000")
// }

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		m := map[string]interface{}{
			"time": time.Now().UnixMilli(),
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(m)
	})
	
	fmt.Println("Run on port :8000")
	http.ListenAndServe(":8000", r)
}
