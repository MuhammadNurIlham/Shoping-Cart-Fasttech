package database

import (
	"fasttech/models"
	"fasttech/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Cart{})

	if err != nil {
		fmt.Println(err);
		panic("Migration Failed");
	}

	fmt.Println("Migration Database Success")
}