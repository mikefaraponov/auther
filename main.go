package main

import (
	"github.com/labstack/echo"
	"github.com/mikefaraponov/auther/bootstrap"
	"github.com/mikefaraponov/auther/resources"
	"github.com/mikefaraponov/auther/server"
	"go.uber.org/fx"
)


func main() {
	app := fx.New(
		fx.Provide(resources.NewSentryHook),
		fx.Provide(resources.NewLogrusLogger),
		fx.Provide(resources.NewLogrusErrorLevel),
		fx.Provide(resources.NewRavenGo),
		fx.Provide(resources.NewPusherClient),
		fx.Provide(resources.NewEnvironment),
		fx.Provide(resources.NewHTTPClient),
		fx.Provide(echo.New),
		fx.Provide(server.New),
		fx.Invoke(bootstrap.Invoke),
	)
	app.Run()
}
