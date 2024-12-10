package server

import "flag"

var config = flag.String("conf", "./config/app.yaml", "app config file")

func main() {
	flag.Parse()

	println("Hello world!")
}
