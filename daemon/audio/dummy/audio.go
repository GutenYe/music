package dummy

import (
	//"fmt"
	"../../audio"
	"reflect"
)

func init() {
	audio.Register("dummy", &Backend{})
}

type Backend struct {
}

func (b *Backend) New(backendName string) audio.Audio {
	return &Audio{name: backendName}
}

type Audio struct {
	name   string
	cb, p0 reflect.Value
}

func (a *Audio) OnMessage(cb_func, param0 interface{}) {
	a.cb = reflect.ValueOf(cb_func)
	a.p0 = reflect.ValueOf(param0)
}

func (a *Audio) Play(uri string) {
}

func (a *Audio) Pause() {
}

func (a *Audio) Resume() {
}

func (a *Audio) Stop() {
}

func (a *Audio) GetState() int {
	return 0
}

func (a *Audio) GetName() string {
	return a.name
}
