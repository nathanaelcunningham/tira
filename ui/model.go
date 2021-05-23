package ui

type Model struct {
    Cursor int
}

func NewModel() Model {
    return Model {
        Cursor: 0,
    }
}
