package simulator

type Getter interface {
	Get(key Key) (*Object, error)
}

type Finder interface {
	Find(filter Object) ([]Object, error)
}

type Setter interface {
	Set(object Object) error
}

type Deleter interface {
	Delete(key Key) error
}

type Watcher interface {
	Watch(filter Object, ch chan Reconcile) error
}

// FIXME: WTF is this name?
type Gimulator interface {
	Getter
	Finder
	Setter
	Deleter
	Watcher
}
