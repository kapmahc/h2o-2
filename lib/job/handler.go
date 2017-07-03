package job

// Handler handler
type Handler func([]byte) error

var handlers = make(map[string]Handler)

// Register register handler by name
func Register(n string, h Handler) {
	handlers[n] = h
}
