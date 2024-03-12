package seeder

import (
	"log"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
	"gorm.io/gorm"
)

func userTableSeeder(conn *gorm.DB) {

	var users = []model.User{
		{Name: "Joko", Email: "joko@joko.com", Password: "joko123"},
		{Name: "Joni", Email: "joni@joni.com", Password: "joni123"},
		{Name: "koko", Email: "koko@koko.com", Password: "koko123"},
	}

	if err := conn.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data users")
}
