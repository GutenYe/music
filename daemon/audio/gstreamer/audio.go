package gstreamer

import (
	//"fmt"
	"../../audio"
	"github.com/ziutek/gst"
	"reflect"
)

var STATE_MAP = map[gst.State]int{
	gst.STATE_VOID_PENDING: audio.STATE_STOP,
	gst.STATE_NULL: audio.STATE_STOP,
	gst.STATE_READY: audio.STATE_STOP,
	gst.STATE_PAUSED: audio.STATE_PAUSED,
	gst.STATE_PLAYING: audio.STATE_PLAYING,
}

func init() {
	audio.Register("gstreamer", &Backend{})
}

type Backend struct {
}

func (b *Backend) New(backendName string) audio.Audio {
	a := &Audio{name: backendName}
	pipe := gst.ElementFactoryMake("playbin", "audio")

	bus := pipe.GetBus()
	bus.AddSignalWatch()
	bus.Connect("message", (*Audio).onGstMessage, a)

	a.pipe = pipe
	a.bus = bus

	return a
}

type Audio struct {
	name   string
	pipe  *gst.Element
	bus   *gst.Bus
	cb, p0 reflect.Value
}

func (a *Audio) OnMessage(cb_func, param0 interface{}) {
	a.cb = reflect.ValueOf(cb_func)
	a.p0 = reflect.ValueOf(param0)
}

func (a *Audio) onMessage(msg int) {
	rps := []reflect.Value{}
	if a.p0.Kind() != reflect.Invalid {
		rps = append(rps, a.p0)
	}
	rps = append(rps, reflect.ValueOf(msg))
	a.cb.Call(rps)
}

func (a *Audio) onGstMessage(bus *gst.Bus, msg *gst.Message) {
	switch msg.GetType() {
	case gst.MESSAGE_EOS:
		a.onMessage(audio.MESSAGE_EOS)
		//p.pipe.SetState(gst.STATE_NULL)
	case gst.MESSAGE_ERROR:
		//fmt.Println("ERROR")
		// FIXME Error: free(): invalid pointer: 0x0000000003a7c280 ***
		//err, debug := msg.ParseError()
		//fmt.Printf("Error: %s (debug: %s)\n", err, debug)
		//p.pipe.SetState(gst.STATE_NULL)
		a.onMessage(audio.MESSAGE_ERROR)
	}
}

func (a *Audio) Play(uri string) {
	a.pipe.SetState(gst.STATE_NULL)
	a.pipe.SetProperty("uri", uri)
	a.pipe.SetState(gst.STATE_PLAYING)
}

func (a *Audio) Pause() {
	a.pipe.SetState(gst.STATE_PAUSED)
}

func (a *Audio) Resume() {
	a.pipe.SetState(gst.STATE_PLAYING)
}

func (a *Audio) Stop() {
	a.pipe.SetState(gst.STATE_NULL)
}

func (a *Audio) GetState() int {
	state, _, _ := a.pipe.GetState(gst.CLOCK_TIME_NONE)
	return STATE_MAP[state]
}

func (a *Audio) GetName() string {
	return a.name
}
