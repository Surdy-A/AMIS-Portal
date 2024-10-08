package database

import (
	"fmt"
	"log"

	"github.com/Surdy-A/amis_portal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	cfg := config.Get()
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	cfg.DB.Username,
	// 	cfg.DB.Password,
	// 	cfg.DB.Host,
	// 	cfg.DB.Port,
	// 	cfg.DB.Name,
	// )

	//psqlInfo := cfg.DB.Pusername + "://" + cfg.DB.Pusername + ":" + cfg.DB.Password + "@" + cfg.DB.Host + "/" + cfg.DB.Name + "?sslmode=required"
	pdsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		cfg.DB.Host,
		cfg.DB.Pusername,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Postgresport,
	)
	fmt.Println(pdsn)

	//fmt.Println(psqlInfo)

	db, err := gorm.Open(postgres.Open(pdsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	DB = db
}
