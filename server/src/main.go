package main

import (
	ctrl "server/src/controller"
)

const port = ":9000"

func main() {
	ctrl.StartServer(port)
}
