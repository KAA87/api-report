package store

import "github.com/KAA87/api-report.git/internal/app/model"

// CalculationDiagnosticTestsRepository ...
type CalculationDiagnosticTestsRepository struct {
	store *Store
}

// FindByDiagnosticTestId ...
func (r *CalculationDiagnosticTestsRepository) FindByDiagnosticTestId(diagnostic_test_id int) (*model.CalculationDiagnosticTests, error) {
	return nil, nil
}