package main

type Model struct {
	Interfaces map[string]iface
}

type iface struct {
	Name       string
	Namespaces []string
	Methods    map[string]method
}

type method struct {
	Name       string
	ReturnType string
	Arguments  []argument
}

type argument struct {
	Name string
	Type string
}

func NewModel() Model {
	return Model{
		Interfaces: make(map[string]iface),
	}
}

func NewInterface() iface {
	return iface{
		Name:       "",
		Namespaces: []string{},
		Methods:    make(map[string]method),
	}
}
