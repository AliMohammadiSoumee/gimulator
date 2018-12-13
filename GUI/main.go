package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
)

const (
	Height = 600
	Width  = 800
)

var (
	ui     lorca.UI
	drawer worldDrawer
)

func fuck(msg string)    { ui.Eval(fmt.Sprintf(`console.log("%s")`, msg)) }
func render(html string) { ui.Eval(fmt.Sprintf("render(`%s`);", html)) }
func width() int         { return ui.Eval(`width()`).Int() }
func height() int        { return ui.Eval(`height()`).Int() }

func main() {
	var err error

	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	ui, err = lorca.New("", "", Width, Height, args...)

	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	err = ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	if err != nil {
		panic(err)
	}

	fmt.Println(width())
	drawer = worldDrawer{
		World:  World{},
		width:  width(),
		height: height(),
	}

	html := drawer.DrawField()
	render(html)

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
