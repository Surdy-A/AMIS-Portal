package bootstrap

import (
	"github.com/Surdy-A/amis_portal/pkg/config"
	"github.com/Surdy-A/amis_portal/pkg/database"
	"github.com/Surdy-A/amis_portal/pkg/html"
	"github.com/Surdy-A/amis_portal/pkg/routing"
	"github.com/Surdy-A/amis_portal/pkg/sessions"
	"github.com/Surdy-A/amis_portal/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
