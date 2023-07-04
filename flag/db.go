package flag

import (
	"go_blog/global"
	"go_blog/models"
)

func MigratorTables() {
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollect{})
	if global.DB.Error != nil {
		global.Log.Fatal(global.DB.Error)
		return
	}
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.ArticleModel{},
		&models.TagModel{},
		&models.CommentModel{},
		&models.ImageModel{},
		&models.LoginRecordModel{},
	)
	if err != nil {
		global.Log.Fatal(err.Error())
	}
}
