package repository

import (
	"database/sql"
	"fmt"
)

type FeatureFlagRepository struct {
	db *sql.DB
}

func NewFeatureFlagRepository(db *sql.DB) *FeatureFlagRepository {
	return &FeatureFlagRepository{db: db}
}

func (repo *FeatureFlagRepository) GetAllFeatureFlags() (map[string]bool, error) {
	flags := make(map[string]bool)

	rows, err := repo.db.Query("SELECT flag_name, enabled FROM feature_flags")
	if err != nil {
		return nil, fmt.Errorf("erro ao obter feature flags: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var flagName string
		var enabled bool
		if err := rows.Scan(&flagName, &enabled); err != nil {
			return nil, fmt.Errorf("erro ao escanear resultado: %v", err)
		}
		flags[flagName] = enabled
	}

	return flags, rows.Err()
}

func (repo *FeatureFlagRepository) UpdateFeatureFlagStatus(flagName string, enabled bool) error {
	query := "UPDATE feature_flags SET enabled = $1 WHERE flag_name = $2"
	result, err := repo.db.Exec(query, enabled, flagName)
	if err != nil {
		return fmt.Errorf("erro ao atualizar o status da feature flag: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar as linhas afetadas: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("feature flag '%s' não encontrada", flagName)
	}

	return nil
}
