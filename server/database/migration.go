package database

import (
	"fmt"

	"github.com/thxrhmn/chat-app/models"
	"github.com/thxrhmn/chat-app/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Chat{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Successful")
}
