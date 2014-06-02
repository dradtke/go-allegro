package console

import (
	"container/list"
	"fmt"
	al "github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/example/ongoing/subsystems/graphics"
	"io"
	"os"
)

const (
	SUMMARY_LENGTH = 30
	BLINK_SPEED    = 0.6
)

var (
	log      list.List
	visible  bool
	is_blunk bool
	blinker  *al.Timer
	cmd      string

	color_map map[string]al.Color
)

type message struct {
	level string
	text  string
}

func (m message) String() string {
	return fmt.Sprintf("[%s] %s", m.level, m.text)
}

func (m message) Line() graphics.Line {
	return graphics.Line{Text: m.String(), Color: color_map[m.level]}
}

func getSummary() (sum []message) {
	sum = make([]message, 0, SUMMARY_LENGTH)
	for e, i := log.Back(), 0; e != nil && i < SUMMARY_LENGTH; e, i = e.Prev(), i+1 {
		sum = append(sum, e.Value.(message))
	}
	return
}

func Debug(msg string) {
	log.PushBack(message{level: "DEBUG", text: msg})
}

func Debugf(msg string, v ...interface{}) {
	log.PushBack(message{level: "DEBUG", text: fmt.Sprintf(msg, v...)})
}

func Info(msg string) {
	log.PushBack(message{level: "INFO", text: msg})
}

func Infof(msg string, v ...interface{}) {
	log.PushBack(message{level: "INFO", text: fmt.Sprintf(msg, v...)})
}

func Error(msg string) {
	log.PushBack(message{level: "ERROR", text: msg})
}

func Errorf(msg string, v ...interface{}) {
	log.PushBack(message{level: "ERROR", text: fmt.Sprintf(msg, v...)})
}

func Toggle() {
	visible = !visible
}

func Visible() bool {
	return visible
}

func SetVisible(v bool) {
	visible = v
}

func Blink() {
	is_blunk = !is_blunk
}

func Render() {
	if !visible {
		return
	}
	sum := getSummary()
	lines := make([]graphics.Line, len(sum))
	for i, msg := range sum {
		lines[i] = msg.Line()
	}
	graphics.RenderConsole(lines, cmd, is_blunk)
}

func Blinker() *al.Timer {
	return blinker
}

func Init(eventQueue *al.EventQueue) {
	var err error
	if blinker, err = al.CreateTimer(BLINK_SPEED); err != nil {
		panic(err)
	}
	eventQueue.Register(blinker)
	blinker.Start()

	color_map = map[string]al.Color{
		"DEBUG": al.MapRGB(0, 0, 255),
		"INFO":  al.MapRGB(0, 255, 0),
		"ERROR": al.MapRGB(255, 0, 0),
	}
}

func Save(filename string) {
	if f, err := os.Create(filename); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	} else {
		for e := log.Front(); e != nil; e = e.Next() {
			io.WriteString(f, e.Value.(message).String()+"\n")
		}
		f.Close()
	}
}

func WriteCmd(c string) {
	cmd += c
}

func BackspaceCmd() {
	if cmd == "" {
		return
	}
	cmd = cmd[:len(cmd)-1]
}

func SubmitCmd() {
	if cmd == "" {
		return
	}
	Debugf("Submitted command: %s", cmd)
	cmd = ""
}
