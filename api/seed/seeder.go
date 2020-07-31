package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/srigalamilitan/fullstack_learn/models"
)

var users = []models.User{
	models.User{
		NickName: "Gita",
		Email:    "gita@putracode.com",
		Password: "password",
	},
	models.User{
		NickName: "Gandhi",
		Email:    "gandhi@putracode.com",
		Password: "password123",
	},
}
var posts = []models.Post{
	models.Post{
		Title:   "title 1",
		Content: "content 1",
	},
	models.Post{
		Title:   "title 2",
		Content: "content 2",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop tables: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table : %v", err)
	}
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching Foreign key error: %v", err)
	}
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table : %v", err)
		}
		posts[i].AuthorId = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table : %v", err)
		}
	}

}
