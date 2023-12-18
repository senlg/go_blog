package flag

import (
	"flag"
	"fmt"
	"os"
)

func Parse() {
	db := flag.Bool("db", true, "是否初始化数据库表")
	flag.Parse()
	if *db {
		Migrator(*db)
	} else {
		os.Exit(0)
	}
	// 初始化迁移表

}

func Migrator(isMigrator bool) {
	fmt.Printf("isMigrator: %v", isMigrator)
	// 无论如何因为自定义中间表所以要程序开始运行时进行关联
	MigratorTables()

}
