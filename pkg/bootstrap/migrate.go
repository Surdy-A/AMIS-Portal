package bootstrap

import (
	"github.com/Surdy-A/amis_portal/internal/database/migration"
	"github.com/Surdy-A/amis_portal/pkg/config"
	"github.com/Surdy-A/amis_portal/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
