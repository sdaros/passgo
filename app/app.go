package app

import (
	"github.com/sdaros/passgo/environment"
	"sync"
)

const (
	// buildMetadata is replaced when package is built using -ldflags -X
	// ex: go build -ldflags "-X main.buildMetadata=`git rev-parse HEAD`"
	buildMetadata = "<placeholder>"
	version       = "0.1.0"
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

func Null() *App {
	return Passgo(nil, nil)
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

func Version() string { return version + "+" + buildMetadata }
