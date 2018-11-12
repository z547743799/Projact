package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	// This handler will match /user/john but will not match neither /user/ or /user.
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	// This handler will match /users/42
	// but will not match
	// neither /users or /users/.
	app.Get("/users/{id:long}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User with ID: %d", id)
	})

	// This handler will match /user/john/send
	// but will not match /user/john/
	app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	app.Run(iris.Addr(":8080"))
}
