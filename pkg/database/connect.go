package database

import (
	"fmt"

	"github.com/Surdy-A/amis_portal/pkg/config"
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

	psqlInfo := cfg.DB.Pusername + "://" + cfg.DB.Pusername + ":" + cfg.DB.Password + "@" + cfg.DB.Host + "/" + cfg.DB.Name + "?sslmode=required"

	// pdsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	cfg.DB.Host,
	// 	cfg.DB.Pusername,
	// 	cfg.DB.Password,
	// 	cfg.DB.Name,
	// 	cfg.DB.Postgresport,
	// )
	// fmt.Println(dsn)
	fmt.Println(psqlInfo)

	//db, err := gorm.Open(postgres.Open("postgresql://amis:qaibOry3utzA0aZnyh7pYU36apWGqpIp@amis-portal.onrender.com/amis_db"), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal("Cannot connect to database")
	// 	return
	// }

	// DB = db
}
