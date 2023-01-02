// Recipes API
//
// これはサンプルAPIです。
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: Self Note<selfnote@example.com> https://selfnote.work/
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
// main.go
package main

import "recipe-api/routes"

func main() {
	router := routes.NewRoutes()
	router.Run()
}
