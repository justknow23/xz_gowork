package enumint

import "insurance/pkg/tools"

type Label struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

type Enum struct {
	maps map[int]string
}

func (e *Enum) Has(name int) bool {
	_, ok := e.maps[name]
	return ok
}

func (e *Enum) Get(name int) string {
	if e.Has(name) {
		return e.maps[name]
	}
	return ""
}

func (e *Enum) GetAll() map[int]string {
	return e.maps
}

func (e *Enum) GetLabel() []Label {
	var list []Label
	//for k, v := range e.maps {
	//	list = append(list, Label{Value: k, Label: v})
	//}
	m := tools.SortIntMap(e.maps)
	for _, v := range m {
		list = append(list, Label{Value: v.Key, Label: v.Val})
	}
	return list
}
