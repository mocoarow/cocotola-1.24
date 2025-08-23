package controller

import (
	"context"

	"go.opentelemetry.io/otel/baggage"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

func AddBaggageMembers(ctx context.Context, values map[string]string) (context.Context, error) {
	bag := baggage.FromContext(ctx)
	for key, value := range values {
		member, err := baggage.NewMember(key, value)
		if err != nil {
			return nil, mbliberrors.Errorf("baggage.NewMember: %w", err)
		}
		if newBag, err := bag.SetMember(member); err == nil {
			bag = newBag
		} else {
			return nil, mbliberrors.Errorf("baggage.SetMember: %w", err)
		}
	}
	ctx = baggage.ContextWithBaggage(ctx, bag)
	return ctx, nil
}
