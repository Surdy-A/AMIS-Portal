package bootstrap

import (
	"github.com/Surdy-A/amis_portal/internal/database/seeder"
	"github.com/Surdy-A/amis_portal/pkg/config"
	"github.com/Surdy-A/amis_portal/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
