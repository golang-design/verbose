package verbose_test

import (
	"testing"

	"golang.design/x/verbose"
)

func TestHello(t *testing.T) {
	if verbose.Hello() != "world" {
		t.Fatal("Hello does not say world!")
	}
}
