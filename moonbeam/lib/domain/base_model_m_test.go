//go:build medium

package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBaseModel_shouldReturnBaseModel_whenGivenValidParameters(t *testing.T) {
	createdAt := time.Now()
	updatedAt := createdAt.Add(time.Hour)

	result, err := NewBaseModel(1, createdAt, updatedAt, 100, 200)

	require.NoError(t, err)
	assert.Equal(t, 1, result.Version)
	assert.Equal(t, createdAt, result.CreatedAt)
	assert.Equal(t, updatedAt, result.UpdatedAt)
	assert.Equal(t, 100, result.CreatedBy)
	assert.Equal(t, 200, result.UpdatedBy)
}

func TestNewBaseModel_shouldReturnBaseModel_whenGivenZeroCreatedByAndUpdatedBy(t *testing.T) {
	createdAt := time.Now()
	updatedAt := createdAt.Add(time.Hour)

	result, err := NewBaseModel(1, createdAt, updatedAt, 0, 0)

	require.NoError(t, err)
	assert.Equal(t, 0, result.CreatedBy)
	assert.Equal(t, 0, result.UpdatedBy)
}

func TestNewBaseModel_shouldReturnError_whenGivenInvalidVersion(t *testing.T) {
	createdAt := time.Now()
	updatedAt := createdAt.Add(time.Hour)

	result, err := NewBaseModel(0, createdAt, updatedAt, 100, 200)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validate base model")
}

func TestNewBaseModel_shouldReturnError_whenGivenNegativeVersion(t *testing.T) {
	createdAt := time.Now()
	updatedAt := createdAt.Add(time.Hour)

	result, err := NewBaseModel(-1, createdAt, updatedAt, 100, 200)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validate base model")
}

func TestNewBaseModel_shouldReturnError_whenGivenNegativeCreatedBy(t *testing.T) {
	createdAt := time.Now()
	updatedAt := createdAt.Add(time.Hour)

	result, err := NewBaseModel(1, createdAt, updatedAt, -1, 200)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validate base model")
}

func TestNewBaseModel_shouldReturnError_whenGivenNegativeUpdatedBy(t *testing.T) {
	createdAt := time.Now()
	updatedAt := createdAt.Add(time.Hour)

	result, err := NewBaseModel(1, createdAt, updatedAt, 100, -1)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validate base model")
}
