package controller

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	"go.opentelemetry.io/otel/baggage"
)

func newBaggage(ctx context.Context, values map[string]string) (*baggage.Baggage, error) {
	members := make([]baggage.Member, 0, len(values))
	for key, value := range values {
		member, err := baggage.NewMember(key, value)
		if err != nil {
			return nil, mbliberrors.Errorf("baggage.NewMember: %w", err)
		}
		members = append(members, member)
	}

	bag, err := baggage.New(members...)
	if err != nil {
		return nil, mbliberrors.Errorf("baggage.New: %w", err)
	}

	return &bag, nil
}
