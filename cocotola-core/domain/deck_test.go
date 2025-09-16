//go:build small

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

func TestNewDeckModel_Valid(t *testing.T) {
	t.Parallel()

	baseModel, err := mblibdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)

	deckID, err := domain.NewDeckID(1)
	require.NoError(t, err)

	organizationID, err := mbuserdomain.NewOrganizationID(1)
	require.NoError(t, err)

	spaceID, err := mbuserdomain.NewSpaceID(1)
	require.NoError(t, err)

	ownerID, err := mbuserdomain.NewUserID(1)
	require.NoError(t, err)

	folderID, err := domain.NewFolderID(1)
	require.NoError(t, err)

	templateID, err := domain.NewTemplateID(1)
	require.NoError(t, err)

	lang2, err := libdomain.NewLang2("en")
	require.NoError(t, err)

	deckModel, err := domain.NewDeckModel(
		baseModel,
		deckID,
		organizationID,
		spaceID,
		folderID,
		"Test Deck",
		templateID,
		lang2,
		"Test Description",
		ownerID,
	)

	require.NoError(t, err)
	require.NotNil(t, deckModel)
	require.Equal(t, "Test Deck", deckModel.Name)
	require.Equal(t, 1, deckModel.TemplateID.Int())
	require.Equal(t, "en", deckModel.Lang2.String())
	require.Equal(t, "Test Description", deckModel.Description)
}

func TestNewDeckModel_Invalid(t *testing.T) {
	t.Parallel()

	baseModel, err := mblibdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)

	deckID, err := domain.NewDeckID(1)
	require.NoError(t, err)

	organizationID, err := mbuserdomain.NewOrganizationID(1)
	require.NoError(t, err)

	spaceID, err := mbuserdomain.NewSpaceID(1)
	require.NoError(t, err)

	folderID, err := domain.NewFolderID(1)
	require.NoError(t, err)

	templateID, err := domain.NewTemplateID(1)
	require.NoError(t, err)

	lang2, err := libdomain.NewLang2("en")
	require.NoError(t, err)

	ownerID, err := mbuserdomain.NewUserID(1)
	require.NoError(t, err)

	tests := []struct {
		name           string
		baseModel      *mblibdomain.BaseModel
		deckID         *domain.DeckID
		organizationID *mbuserdomain.OrganizationID
		spaceID        *mbuserdomain.SpaceID
		folderID       *domain.FolderID
		deckName       string
		templateID     *domain.TemplateID
		lang2          *libdomain.Lang2
		ownerID        *mbuserdomain.UserID
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
			templateID:     templateID,
			lang2:          lang2,
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
			templateID:     templateID,
			lang2:          nil,
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
			templateID:     templateID,
			lang2:          lang2,
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
			templateID:     templateID,
			lang2:          lang2,
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
			templateID:     templateID,
			lang2:          lang2,
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
			templateID:     templateID,
			lang2:          lang2,
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
			templateID:     templateID,
			lang2:          lang2,
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
			templateID:     templateID,
			lang2:          lang2,
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
				tt.folderID,
				tt.deckName,
				tt.templateID,
				tt.lang2,
				"Test Description",
				tt.ownerID,
			)

			require.Error(t, err)
		})
	}
}
