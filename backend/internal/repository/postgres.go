package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sort"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(databaseURL string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&LearnerWorkspaceModel{},
		&LearnerPreferenceModel{},
		&VisualThemePaletteModel{},
		&TagModel{},
		&SourceMaterialModel{},
		&MaterialVersionModel{},
		&KnowledgePointModel{},
		&KnowledgeRelationshipModel{},
		&PromptPresetModel{},
		&PromptSnapshotModel{},
		&AIJobModel{},
		&AIDraftModel{},
	)
}

func RunMigrations(ctx context.Context, db *sql.DB, dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)
	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (version text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil {
		return err
	}
	for _, file := range files {
		var exists bool
		if err := db.QueryRowContext(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version=$1)`, file).Scan(&exists); err != nil {
			return err
		}
		if exists {
			continue
		}
		content, err := os.ReadFile(dir + "/" + file)
		if err != nil {
			return err
		}
		if _, err := db.ExecContext(ctx, string(content)); err != nil {
			return fmt.Errorf("migration %s: %w", file, err)
		}
		if _, err := db.ExecContext(ctx, `INSERT INTO schema_migrations(version) VALUES($1)`, file); err != nil {
			return err
		}
	}
	return nil
}
