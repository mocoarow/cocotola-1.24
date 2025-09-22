package moonbeam

import (
	"encoding/json"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type OrganizationID struct {
	Value *mbuserdomain.OrganizationID `validate:"required,gte=1"`
}

func (m *OrganizationID) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err //nolint:wrapcheck
	}
	value, err := mbuserdomain.NewOrganizationID(v)
	if err != nil {
		return mbliberrors.Errorf("mbuserdomain.NewOrganizationID: %w", err)
	}
	m.Value = value
	return nil
}

func (m OrganizationID) MarshalJSON() ([]byte, error) {
	if m.Value == nil {
		return json.Marshal(nil) //nolint:wrapcheck
	}
	return json.Marshal(m.Value.Value) //nolint:wrapcheck
}

type UserID struct {
	Value *mbuserdomain.UserID `validate:"required,gte=1"`
}

func (m *UserID) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err //nolint:wrapcheck
	}
	value, err := mbuserdomain.NewUserID(v)
	if err != nil {
		return mbliberrors.Errorf("mbuserdomain.NewOrganizationID: %w", err)
	}
	m.Value = value
	return nil
}

func (m UserID) MarshalJSON() ([]byte, error) {
	if m.Value == nil {
		return json.Marshal(nil) //nolint:wrapcheck
	}
	return json.Marshal(m.Value.Value) //nolint:wrapcheck
}

type SpaceID struct {
	Value *mbuserdomain.SpaceID `validate:"required,gte=1"`
}

func (m *SpaceID) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err //nolint:wrapcheck
	}
	value, err := mbuserdomain.NewSpaceID(v)
	if err != nil {
		return mbliberrors.Errorf("mbuserdomain.NewSpaceID: %w", err)
	}
	m.Value = value
	return nil
}

func (m SpaceID) MarshalJSON() ([]byte, error) {
	if m.Value == nil {
		return json.Marshal(nil) //nolint:wrapcheck
	}
	return json.Marshal(m.Value.Value) //nolint:wrapcheck
}

type Lang2 struct {
	Value *mblibdomain.Lang2 `validate:"required"`
}

func (m *Lang2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err //nolint:wrapcheck
	}
	value, err := mblibdomain.NewLang2(v)
	if err != nil {
		return mbliberrors.Errorf("mblibdomain.NewLang2: %w", err)
	}
	m.Value = value
	return nil
}

func (m Lang2) MarshalJSON() ([]byte, error) {
	if m.Value == nil {
		return json.Marshal(nil) //nolint:wrapcheck
	}
	return json.Marshal(m.Value.String()) //nolint:wrapcheck
}

type SpaceIDs []int

func (s *SpaceIDs) UnmarshalJSON(data []byte) error {
	var ids []int
	if err := json.Unmarshal(data, &ids); err != nil {
		return err //nolint:wrapcheck
	}
	*s = SpaceIDs(ids)
	return nil
}

func (s SpaceIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal([]int(s)) //nolint:wrapcheck
}

func (s SpaceIDs) ToSpaceIDs() ([]*mbuserdomain.SpaceID, error) {
	spaceIDs := make([]*mbuserdomain.SpaceID, 0, len(s))
	for _, id := range s {
		spaceID, err := mbuserdomain.NewSpaceID(id)
		if err != nil {
			return nil, mbliberrors.Errorf("mbuserdomain.NewSpaceID: %w", err)
		}
		spaceIDs = append(spaceIDs, spaceID)
	}
	return spaceIDs, nil
}
