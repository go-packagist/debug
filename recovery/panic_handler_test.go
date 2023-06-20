package recovery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicHandler_Report(t *testing.T) {
	p := &PanicHandler{}
	defer func() {
		if r := recover(); r != nil {
			assert.Panics(t, func() {
				p.Report(r)
			})
		}
	}()

	panic("test panic")
}
