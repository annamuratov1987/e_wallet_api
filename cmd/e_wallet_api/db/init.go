package main

import (
	"fmt"
	"github.com/annamuratov1987/e_wallet_api/pkg/ewallets"
	"github.com/annamuratov1987/e_wallet_api/pkg/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main()  {
	fmt.Println("Start database initialization...")

	dsn := "host=db user=root password=abc123 dbname=e-wallet-db port=5432 sslmode=disable TimeZone=Asia/Tashkent"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}

	err = db.AutoMigrate(users.User{}, ewallets.Ewallet{}, ewallets.EwalletOperation{})
	if err!=nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}

	fmt.Println("Database initialized successfully!")
}

