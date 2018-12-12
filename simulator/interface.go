package simulator

type Getter interface {
	Get(key string) (interface{}, error)
}

type Finder interface {
	Find(filter interface{}) ([]interface{}, error)
}

type Setter interface {
	Set(key string, object interface{}) error
}

type Deleter interface {
	Delete(key string) error
}

type Watcher interface {
	Watch(key string, ch chan Reconcile) error
}

// FIXME: WTF is this name?
type Gimulator interface {
	Getter
	Finder
	Setter
	Deleter
	Watcher
}
