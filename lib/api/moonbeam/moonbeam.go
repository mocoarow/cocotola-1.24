package moonbeam

import (
	"encoding/json"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type SpaceID struct {
	Value *mbuserdomain.SpaceID `validate:"required,gte=1"`
}

func (m *SpaceID) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	value, err := mbuserdomain.NewSpaceID(v)
	if err != nil {
		return err
	}
	m.Value = value
	return nil
}

func (m SpaceID) MarshalJSON() ([]byte, error) {
	if m.Value == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(m.Value.Value)
}

type Lang2 struct {
	Value *mblibdomain.Lang2 `validate:"required"`
}

func (m *Lang2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	value, err := mblibdomain.NewLang2(v)
	if err != nil {
		return err
	}
	m.Value = value
	return nil
}

func (m Lang2) MarshalJSON() ([]byte, error) {
	if m.Value == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(m.Value.String())
}

type SpaceIDs []int

func (s *SpaceIDs) UnmarshalJSON(data []byte) error {
	var ids []int
	if err := json.Unmarshal(data, &ids); err != nil {
		return err
	}
	*s = SpaceIDs(ids)
	return nil
}

func (s SpaceIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal([]int(s))
}

func (s SpaceIDs) ToSpaceIDs() ([]*mbuserdomain.SpaceID, error) {
	spaceIDs := make([]*mbuserdomain.SpaceID, 0, len(s))
	for _, id := range s {
		spaceID, err := mbuserdomain.NewSpaceID(id)
		if err != nil {
			return nil, err
		}
		spaceIDs = append(spaceIDs, spaceID)
	}
	return spaceIDs, nil
}
