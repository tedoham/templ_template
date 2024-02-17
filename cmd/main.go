package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/tedoham/myride/internal/templates"
)

// func init() {
// 	// templates.Load()
// 	// serve static css files
// 	http.Handle("../static/css/styles.css", http.StripPrefix("../static/css/styles.css", http.FileServer(http.Dir("../static/css/styles.css"))))
// }

// Initialize the file server for static content in the init function
func init() {
	// Serve CSS statically
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))
}

func main() {

	// Use a template that doesn't take parameters.
	http.Handle("/", templ.Handler(templates.Hello("yooooooooooo.....")))

	// Use a template that accesses data or handles form posts.
	// http.Handle("/posts", NewPostsHandler())

	// Start the server.
	fmt.Println("listening on http://localhost:3000")
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}

// func main() {
// 	http.HandleFunc("/index", IndexHandler)

// 	http.ListenAndServe(":3000", nil)
// }

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello from route /index")
// }

// func main() {
// 	app := echo.New()
// 	app.GET("/", HomeHandler)

// 	app.Static("/web/css", "web/css")
// 	app.Static("/web/static", "web/static")
// 	app.Static("/web/fonts", "web/fonts")
// 	app.Logger.Fatal(app.Start(":3000"))
// }

// // This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
// func Render(ctx echo.Context, statusCode int, t templ.Component) error {
// 	ctx.Response().Writer.WriteHeader(statusCode)
// 	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
// 	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
// }

// func HomeHandler(c echo.Context) error {
// 	return Render(c, http.StatusOK, templates.Hello("yoooooopppppppppp"))
// }
