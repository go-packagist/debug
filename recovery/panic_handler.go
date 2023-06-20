package recovery

type PanicHandler struct{}

var _ Handler = (*PanicHandler)(nil)

func (p *PanicHandler) Report(v interface{}) {
	panic(v)
}
