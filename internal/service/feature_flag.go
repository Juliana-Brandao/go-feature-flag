package service

import (
	"github.com/Waelson/go-feature-flag/internal/repository"
	"github.com/Waelson/go-feature-flag/internal/util"
	"sync"
)

type FeatureFlagService struct {
	repo          *repository.FeatureFlagRepository
	MetricsRecord util.MetricsRecord
	featureFlags  map[string]bool
	mu            sync.RWMutex
}

func NewFeatureFlagService(repo *repository.FeatureFlagRepository, metricsRecord util.MetricsRecord) *FeatureFlagService {
	return &FeatureFlagService{
		repo:          repo,
		MetricsRecord: metricsRecord,
		featureFlags:  make(map[string]bool),
	}
}

func (ffs *FeatureFlagService) UpdateFeatureFlags() error {
	ffs.mu.Lock()
	defer ffs.mu.Unlock()

	flags, err := ffs.repo.GetAllFeatureFlags()
	if err != nil {
		return err
	}

	ffs.MetricsRecord.ResetGaugeFeatureFlag()

	for flagName, enabled := range flags {
		ffs.featureFlags[flagName] = enabled

		status := "enabled"
		if !enabled {
			status = "disabled"
		}
		ffs.MetricsRecord.WithLabelValues(flagName, status)
	}

	return nil
}

func (ffs *FeatureFlagService) IsFeatureEnabled(flag string) bool {
	ffs.mu.RLock()
	defer ffs.mu.RUnlock()

	enabled, exists := ffs.featureFlags[flag]
	return exists && enabled
}
