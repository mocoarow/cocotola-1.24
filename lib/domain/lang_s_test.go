//go:build small

package domain_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/mocoarow/cocotola-1.24/lib/domain"
	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
)

func TestNewLang2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		args          string
		want          *domain.Lang2
		wantErr       bool
		wantErrDetail error
	}{
		{
			name:    "en",
			args:    "en",
			want:    domain.Lang2EN,
			wantErr: false,
		},
		{
			name:    "ja",
			args:    "ja",
			want:    domain.Lang2JA,
			wantErr: false,
		},
		{
			name:          "empty string",
			args:          "",
			wantErr:       true,
			wantErrDetail: mblibdomain.ErrInvalidArgument,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := domain.NewLang2(tt.args)
			if !tt.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				if tt.wantErrDetail != nil && !errors.Is(err, tt.wantErrDetail) {
					t.Errorf("NewLang2() err = %v, wantErrDetail %v", err, tt.wantErrDetail)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLang2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLang5_ToLang2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lang5 *domain.Lang5
		lang2 *domain.Lang2
	}{
		{
			name:  domain.Lang5ENUS.String(),
			lang5: domain.Lang5ENUS,
			lang2: domain.Lang2EN,
		},
		{
			name:  domain.Lang5JAJP.String(),
			lang5: domain.Lang5JAJP,
			lang2: domain.Lang2JA,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.lang5.ToLang2(), tt.lang2)
		})
	}
}
