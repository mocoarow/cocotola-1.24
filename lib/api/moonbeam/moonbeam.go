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
