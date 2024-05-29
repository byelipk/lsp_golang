package analysis

type State struct {
    Documents map[string]string
}

func NewState() State {
    return State{
        Documents: make(map[string]string),
    }
}
