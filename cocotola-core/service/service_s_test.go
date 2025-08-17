//go:build small

package service_test

import (
	"testing"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

func TestA(t *testing.T) {
	service.A()
}

func TestB(t *testing.T) {
	service.B()
}

func TestC(t *testing.T) {
	service.C()
}
