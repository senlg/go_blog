package flag

import (
	"flag"
)

func Parse() {
	db := flag.Bool("db", false, "初始化数据库表")
	flag.Parse()
	// 初始化迁移表
	Migrator(*db)
}

func Migrator(isMigrator bool) {
	// 无论如何因为自定义中间表所以要程序开始运行时进行关联
	MigratorTables()

}
