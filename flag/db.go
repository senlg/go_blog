package flag

import (
	"go_blog/global"
	"go_blog/models"
)

func MigratorTables() {
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollect{}) // 自定义中间结构体必须要手动指定中间表来告诉gorm 谁和谁链接
	// global.DB.SetupJoinTable(&models.ArticleModel{}, "UserModels", &models.UserCollect{})  // 自定义中间结构体必须要手动指定中间表来告诉gorm 谁和谁链接
	if global.DB.Error != nil {
		global.Log.Fatal(global.DB.Error)
		return
	}
	//根据定义的模型结构自动创建或更新数据库表。所以这段代码的作用是将定义的模型结构映射到数据库表，并使用 InnoDB 引擎进行创建或更新。
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.ArticleModel{},
		&models.TagModel{},
		&models.CommentModel{},
		&models.ImageModel{},
		&models.LoginRecordModel{},
		&models.AgreeModel{},
		&models.MenuModel{},
		&models.UserCollect{},
	)
	if err != nil {
		global.Log.Fatal(err.Error())
	}
}
