package service

import (
	"github.com/Waelson/go-feature-flag/internal/repository"
	"sync"
)

type FeatureFlagService struct {
	repo         *repository.FeatureFlagRepository
	featureFlags map[string]bool
	mu           sync.RWMutex
}

func NewFeatureFlagService(repo *repository.FeatureFlagRepository) *FeatureFlagService {
	return &FeatureFlagService{
		repo:         repo,
		featureFlags: make(map[string]bool),
	}
}

func (ffs *FeatureFlagService) UpdateFeatureFlags() error {
	ffs.mu.Lock()
	defer ffs.mu.Unlock()

	flags, err := ffs.repo.GetAllFeatureFlags()
	if err != nil {
		return err
	}
	ffs.featureFlags = flags
	return nil
}

func (ffs *FeatureFlagService) IsFeatureEnabled(flag string) bool {
	ffs.mu.RLock()
	defer ffs.mu.RUnlock()

	enabled, exists := ffs.featureFlags[flag]
	return exists && enabled
}
