package migrations

import (
	"blog/global"
	_ "blog/inits"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	//migrate create -ext sql -dir migrations create_users_table

	var migrationDir = flag.String("migration.files", "./migrations", "Directory where the migration files are located ?")
	if err := global.SqlDB.Ping();err!=nil{
		panic("failed to connect database")
	}
	// 开始数据迁移
	// 开始数据迁移
	driver, err := mysql.WithInstance(global.SqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir), // file://path/to/directory
		"mysql", driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated")
}
