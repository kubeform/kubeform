package util

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/flect"
)

func TestName(t *testing.T) {
	fmt.Println(flect.Camelize("abc_tls30"))
}
