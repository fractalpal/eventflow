package eventflow

type Commander interface {
	Command(interface{}) (Event, error)
}
