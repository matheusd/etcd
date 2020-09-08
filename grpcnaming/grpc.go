package grpcnaming

type Operation uint8

const (
	Add Operation = iota
	Delete
)

type Update struct {
	Op       Operation
	Addr     string
	Metadata interface{}
}

type Watcher interface {
	Next() ([]*Update, error)
	Close()
}
