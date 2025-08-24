//go:build small

package service_test

import (
	"testing"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

func TestA(t *testing.T) {
	t.Parallel()
	service.A()
}

func TestB(t *testing.T) {
	t.Parallel()
	service.B()
}

func TestC(t *testing.T) {
	t.Parallel()
	service.C()
}
