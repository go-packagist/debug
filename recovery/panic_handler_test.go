package recovery

import "testing"

func TestPanicHandler_Report(t *testing.T) {
	p := &PanicHandler{}
	defer func() {
		if r := recover(); r != nil {
			p.Report(r)
		}
	}()
}
