package app

import (
	"sync"

	"github.com/sdaros/passgo/environment"
)

type (
	App struct {
		*environment.Env
		*Registrar
	}
	Registrar struct {
		values map[string]interface{}
		sync.Mutex
	}
)

// Passgo returns the App instance that is responsible for
// parsing the environment and registering things during
// its runtime
func Passgo(env *environment.Env, registrar *Registrar) *App {
	// nil environment initialises an empty environment
	if env == nil {
		env = environment.Null()
	}
	if registrar == nil {
		registrar = new(Registrar)
		registrar.values = make(map[string]interface{})
	}
	return &App{Registrar: registrar, Env: env}

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

func Null() *App {
	return Passgo(nil, nil)
}
