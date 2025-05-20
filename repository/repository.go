package repository

import (
	"context"
	"fmt"
	"regexp"

	"github.com/jSierraB3991/golang-multitenant/libs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(pg_url string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(pg_url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Repository struct {
	db      *gorm.DB
	schemas string
}

func NewRepository(db *gorm.DB, schemas string) *Repository {
	return &Repository{
		db:      db,
		schemas: schemas,
	}
}

func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(libs.ContextTenantKey).(string)
	return v, ok
}

func (p *Repository) WithTenant(ctx context.Context) (*gorm.DB, error) {
	tenant, isOk := FromContext(ctx)
	if !isOk {
		return nil, fmt.Errorf("tenant ID not found in context")
	}
	// Validar de nuevo (doble capa)
	if ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, tenant); !ok {
		return nil, fmt.Errorf("invalid tenant ID")
	}

	// Creamos una nueva sesión segura
	tx := p.db.Session(&gorm.Session{
		NewDB: true,
	})

	// Validación del nombre del schema
	if !regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`).MatchString(tenant) {
		return nil, fmt.Errorf("invalid tenant ID")
	}

	// Quoteo seguro
	quotedSchema := fmt.Sprintf(`"%s"`, tenant)

	// Ejecutar con interpolación controlada (porque no se puede parametrizar)
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", quotedSchema)).Error; err != nil {
		return nil, err
	}

	return tx, nil
}
