package main

import (
    "github.com/folkersj/celeritas"
    "myapp/handlers"
)

type application struct {
    App      *celeritas.Celeritas
    Handlers *handlers.Handlers
}

func main() {
    c := initApplication()
    c.App.ListenAndServe()
}
