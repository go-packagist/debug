package recovery

type Handler interface {
	// The Report the given panic.
	Report(interface{})
}
