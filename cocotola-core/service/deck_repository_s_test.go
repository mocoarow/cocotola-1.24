//go:build small

package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

func TestNewAddDeckParameter_shouldReturnParameter_whenValidInput(t *testing.T) {
	t.Parallel()

	spaceID, err := mbuserdomain.NewSpaceID(1)
	require.NoError(t, err)

	folderID, err := domain.NewFolderID(1)
	require.NoError(t, err)

	templateID, err := domain.NewTemplateID(1)
	require.NoError(t, err)

	lang2, err := mblibdomain.NewLang2("en")
	require.NoError(t, err)

	param, err := service.NewAddDeckParameter(
		spaceID,
		folderID,
		templateID,
		"Test Deck",
		lang2,
		"Test Description",
	)

	require.NoError(t, err)
	require.NotNil(t, param)
	require.Equal(t, spaceID, param.SpaceID)
	require.Equal(t, folderID, param.FolderID)
	require.Equal(t, templateID, param.TemplateID)
	require.Equal(t, "Test Deck", param.Name)
	require.Equal(t, lang2, param.Lang2)
	require.Equal(t, "Test Description", param.Description)
}

func TestNewAddDeckParameter_shouldReturnError_whenInvalidInput(t *testing.T) {
	t.Parallel()

	spaceID, err := mbuserdomain.NewSpaceID(1)
	require.NoError(t, err)

	folderID, err := domain.NewFolderID(1)
	require.NoError(t, err)

	templateID, err := domain.NewTemplateID(1)
	require.NoError(t, err)

	lang2, err := mblibdomain.NewLang2("en")
	require.NoError(t, err)

	tests := []struct {
		name        string
		spaceID     *mbuserdomain.SpaceID
		folderID    *domain.FolderID
		templateID  *domain.TemplateID
		deckName    string
		lang2       *mblibdomain.Lang2
		description string
	}{
		{
			name:        "nil spaceID",
			spaceID:     nil,
			folderID:    folderID,
			templateID:  templateID,
			deckName:    "Test Deck",
			lang2:       lang2,
			description: "Test Description",
		},
		{
			name:        "nil folderID",
			spaceID:     spaceID,
			folderID:    nil,
			templateID:  templateID,
			deckName:    "Test Deck",
			lang2:       lang2,
			description: "Test Description",
		},
		{
			name:        "nil templateID",
			spaceID:     spaceID,
			folderID:    folderID,
			templateID:  nil,
			deckName:    "Test Deck",
			lang2:       lang2,
			description: "Test Description",
		},
		{
			name:        "empty name",
			spaceID:     spaceID,
			folderID:    folderID,
			templateID:  templateID,
			deckName:    "",
			lang2:       lang2,
			description: "Test Description",
		},
		{
			name:        "nil lang2",
			spaceID:     spaceID,
			folderID:    folderID,
			templateID:  templateID,
			deckName:    "Test Deck",
			lang2:       nil,
			description: "Test Description",
		},
		{
			name:        "description too long",
			spaceID:     spaceID,
			folderID:    folderID,
			templateID:  templateID,
			deckName:    "Test Deck",
			lang2:       lang2,
			description: "This description is way too long and exceeds the maximum allowed length of 100 characters for deck descriptions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := service.NewAddDeckParameter(
				tt.spaceID,
				tt.folderID,
				tt.templateID,
				tt.deckName,
				tt.lang2,
				tt.description,
			)

			require.Error(t, err)
		})
	}
}
