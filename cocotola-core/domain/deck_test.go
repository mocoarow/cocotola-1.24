//go:build small

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

func TestNewDeckModel_Valid(t *testing.T) {
	t.Parallel()

	baseModel, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)

	deckID, err := domain.NewDeckID(1)
	require.NoError(t, err)

	organizationID, err := userdomain.NewOrganizationID(1)
	require.NoError(t, err)

	spaceID, err := domain.NewSpaceID(1)
	require.NoError(t, err)

	ownerID, err := userdomain.NewAppUserID(1)
	require.NoError(t, err)

	folderID, err := domain.NewFolderID(1)
	require.NoError(t, err)

	deckModel, err := domain.NewDeckModel(
		baseModel,
		deckID,
		organizationID,
		spaceID,
		ownerID,
		folderID,
		"Test Deck",
		1,
		"en",
		"Test Description",
	)

	require.NoError(t, err)
	require.NotNil(t, deckModel)
	require.Equal(t, "Test Deck", deckModel.Name)
	require.Equal(t, 1, deckModel.TemplateID)
	require.Equal(t, "en", deckModel.Lang2)
	require.Equal(t, "Test Description", deckModel.Description)
}

func TestNewDeckModel_Invalid(t *testing.T) {
	t.Parallel()

	baseModel, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)

	deckID, err := domain.NewDeckID(1)
	require.NoError(t, err)

	organizationID, err := userdomain.NewOrganizationID(1)
	require.NoError(t, err)

	spaceID, err := domain.NewSpaceID(1)
	require.NoError(t, err)

	ownerID, err := userdomain.NewAppUserID(1)
	require.NoError(t, err)

	folderID, err := domain.NewFolderID(1)
	require.NoError(t, err)

	tests := []struct {
		name           string
		baseModel      *libdomain.BaseModel
		deckID         *domain.DeckID
		organizationID *userdomain.OrganizationID
		spaceID        *domain.SpaceID
		ownerID        *userdomain.AppUserID
		folderID       *domain.FolderID
		deckName       string
		templateID     int
		lang2          string
	}{
		{
			name:           "empty name",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "",
			templateID:     1,
			lang2:          "en",
		},
		{
			name:           "invalid template ID (zero)",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     0,
			lang2:          "en",
		},
		{
			name:           "invalid template ID (negative)",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     -1,
			lang2:          "en",
		},
		{
			name:           "empty lang2",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "",
		},
		{
			name:           "lang2 too short (1 char)",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "e",
		},
		{
			name:           "lang2 too long (3 chars)",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "eng",
		},
		{
			name:           "nil baseModel",
			baseModel:      nil,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "en",
		},
		{
			name:           "nil deckID",
			baseModel:      baseModel,
			deckID:         nil,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "en",
		},
		{
			name:           "nil organizationID",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: nil,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "en",
		},
		{
			name:           "nil spaceID",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        nil,
			ownerID:        ownerID,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "en",
		},
		{
			name:           "nil ownerID",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        nil,
			folderID:       folderID,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "en",
		},
		{
			name:           "nil folderID",
			baseModel:      baseModel,
			deckID:         deckID,
			organizationID: organizationID,
			spaceID:        spaceID,
			ownerID:        ownerID,
			folderID:       nil,
			deckName:       "Test Deck",
			templateID:     1,
			lang2:          "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := domain.NewDeckModel(
				tt.baseModel,
				tt.deckID,
				tt.organizationID,
				tt.spaceID,
				tt.ownerID,
				tt.folderID,
				tt.deckName,
				tt.templateID,
				tt.lang2,
				"Test Description",
			)

			require.Error(t, err)
			require.Contains(t, err.Error(), "validate deck model")
		})
	}
}
