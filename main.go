package main

import (
	"net/http"

	"github.com/Johnman67112/fx-api/ui"
	"go.uber.org/fx"
)

//version v1.0.0

func main() {
	fx.New(
		fx.Provide(ui.NewHTTPServer, ui.NewServeMux, ui.NewEchoHandler),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
