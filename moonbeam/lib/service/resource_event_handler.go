package service

import "context"

type ResourceEventHandler interface {
	OnAdd(ctx context.Context, obj interface{})
	OnUpdate(ctx context.Context, oldObj, newObj interface{})
	OnDelete(ctx context.Context, sobj interface{})
}

type ResourceEventHandlerFuncs struct {
	AddFunc    func(ctx context.Context, obj interface{})
	UpdateFunc func(ctx context.Context, oldObj, newObj interface{})
	DeleteFunc func(ctx context.Context, obj interface{})
}

// OnAdd calls AddFunc if it's not nil.
func (r ResourceEventHandlerFuncs) OnAdd(ctx context.Context, obj interface{}) {
	if r.AddFunc != nil {
		r.AddFunc(ctx, obj)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r ResourceEventHandlerFuncs) OnUpdate(ctx context.Context, oldObj, newObj interface{}) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(ctx, oldObj, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r ResourceEventHandlerFuncs) OnDelete(ctx context.Context, obj interface{}) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(ctx, obj)
	}
}
