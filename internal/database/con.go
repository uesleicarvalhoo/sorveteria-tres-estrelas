package database

import (
	"fmt"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresConnection(user, passwd, database, host string, port int) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=America/Sao_Paulo",
		host, user, passwd, database, port,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresConnectionWithMigration(user, passwd, database, host string, port int) (*gorm.DB, error) {
	db, err := NewPostgresConnection(user, passwd, database, host, port)
	if err != nil {
		return nil, err
	}

	if err := MigratePostgres(db, database); err != nil {
		return nil, err
	}

	return db, nil
}
