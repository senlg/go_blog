package flag

import (
	"flag"
	"go_blog/global"
)

func Parse() {
	db := flag.Bool("db", false, "初始化数据库表")
	flag.Parse()
	// 初始化迁移表
	Migrator(*db)
}

func Migrator(isMigrator bool) {
	if isMigrator {
		MigratorTables()
		global.Log.Exit(1)
	} else {

	}
}
