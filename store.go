package eventflow

type Store interface {
	Save(event Event) error
}
