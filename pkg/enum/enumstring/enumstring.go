package enumstring


type Label struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Enum struct {
	maps map[string]string
}

func (e *Enum) Has(name string) bool {
	_, ok := e.maps[name]
	return ok
}

func (e *Enum) Get(name string) string {
	if e.Has(name) {
		return e.maps[name]
	}
	return ""
}

func (e *Enum) GetAll() map[string]string {
	return e.maps
}

func (e *Enum) GetLabel() []Label {
	var list []Label
	for k, v := range e.maps {
		list = append(list, Label{Value: k, Label: v})
	}
	return list
}
