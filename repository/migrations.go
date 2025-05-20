package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/jSierraB3991/golang-multitenant/models"
	"gorm.io/gorm"
)

func (repo *Repository) Migrations() error {
	return MigrateSchemas(repo.db, strings.Split(repo.schemas, ","))
}

func MigrateSchemas(db *gorm.DB, schemas []string) error {
	for _, schema := range schemas {
		// Asegúrate de que el schema existe antes de migrar
		if err := ensureSchemaExists(db, schema); err != nil {
			return fmt.Errorf("schema '%s' creation failed: %w", schema, err)
		}

		dbTenant, err := db.Session(&gorm.Session{NewDB: true}).DB()
		if err != nil {
			log.Fatalf("could not create session for %s: %v", schema, err)
		}

		// establecer el search_path de forma segura
		_, err = dbTenant.Exec(`SET search_path TO ` + QuoteIdentifier(schema)) // o con una query preparada
		if err != nil {
			log.Fatalf("could not set search_path for %s: %v", schema, err)
		}

		if err := db.AutoMigrate(
			&models.User{}, // tus modelos aquí
		); err != nil {
			return fmt.Errorf("migration failed for schema '%s': %w", schema, err)
		}

		log.Printf("✅ Migrated schema: %s", schema)
	}
	return nil
}
func ensureSchemaExists(db *gorm.DB, schema string) error {
	return db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", QuoteIdentifier(schema))).Error
}

func QuoteIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}
