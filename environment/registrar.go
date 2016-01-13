package environment

import (
	"sync"
)

type Registrar struct {
	values map[string]interface{}
	sync.Mutex
}

func (r *Registrar) Register(k string, v interface{}) {
	if r == nil {
		return
	}

	r.Lock()
	defer r.Unlock()
	r.values[k] = v
}

func (r *Registrar) Lookup(k string) interface{} {
	if r == nil {
		return nil
	}

	r.Lock()
	defer r.Unlock()
	return r.values[k]
}
