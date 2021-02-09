package main

import (
	"net/http"

	"github.com/moosecanswim/learn-go/code/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
