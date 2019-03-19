package eventflow

type Store interface {
	//Events(rowId string) ([]Event, error)
	Save(event Event) error
}
