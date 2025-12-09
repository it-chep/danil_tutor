package get_states

type State struct {
	State int    `json:"state"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
}

type Response struct {
	States []State `json:"states"`
}
