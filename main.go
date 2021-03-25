package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KiraiChang/fingo/infra/persistences/ent"
	"github.com/KiraiChang/fingo/infra/services/graph"
	"github.com/kataras/iris/v12"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//"log"
)

const defaultPort = "8080"

func main() {
	client, err := ent.Open("sqlite3", "file:foo.db?loc=auto&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := iris.New()
	// Set Logger level to "debug",
	// see your terminal and the created file.
	app.Logger().SetLevel("debug")

	//h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	h := handler.NewDefaultServer(graph.NewSchema(client))
	// GET: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	})

	app.Post("/query", func(ctx iris.Context) {
		h.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	})

	app.Listen("127.0.0.1:" + defaultPort)
}
