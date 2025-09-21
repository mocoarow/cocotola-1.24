package controller

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"

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

// AddBaggageToCurrentSpan adds current baggage members as attributes to the current span
func AddBaggageToCurrentSpan(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	bag := baggage.FromContext(ctx)
	for _, member := range bag.Members() {
		span.SetAttributes(attribute.String("baggage."+member.Key(), member.Value()))
	}
}
