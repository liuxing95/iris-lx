package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	bookApi := app.Party("/books")

	{
		bookApi.Use(iris.Compression)

		// GET: http://localhost:8080/books
		bookApi.Get("/", list)
		// POST: http://localhost:8080/books
		bookApi.Post("/", create)
	}

	app.Listen(":8080")
}

type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}
	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
