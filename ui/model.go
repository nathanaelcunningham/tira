package ui

import "github.com/nathanaelcunningham/tira/pane"

type Model struct {
    ServersPane pane.Model
    Cursor int
    Ready bool
}

func NewModel() Model {
    return Model {
        ServersPane: pane.Model{},
        Cursor: 0,
        Ready: false,
    }
}
