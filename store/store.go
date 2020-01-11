package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Store ...
type Store struct {
	config *Config
	db *sql.DB
	calculationDiagnosticTestsRepository *CalculationDiagnosticTestsRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// CalculationDiagnosticTests ...
func (s *Store) CalculationDiagnosticTests() *CalculationDiagnosticTestsRepository {
	if s.calculationDiagnosticTestsRepository != nil {
		return s.calculationDiagnosticTestsRepository
	}

	s.calculationDiagnosticTestsRepository = &CalculationDiagnosticTestsRepository{
		store: s,
	}

	return s.calculationDiagnosticTestsRepository
}