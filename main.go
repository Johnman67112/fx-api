package main

import (
	"net/http"

	"github.com/Johnman67112/fx-api/ui"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(ui.NewHTTPServer, ui.NewServeMux, ui.NewEchoHandler),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
