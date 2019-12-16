package param

import "fmt"

type (
	Namespace map[string]interface{}
	Params    map[string]*Namespace
)

func New() *Params {
	return &Params{}
}
func (p *Params) AddNs(name string, variables map[string]interface{}) (err error) {
	if _, ok := p.Ns(name); ok {
		return fmt.Errorf("namespace `%s` already exists", name)
	}
	ns := Namespace(variables)
	(*p)[name] = &ns
	return
}
func (p *Params) Ns(name string) (ns *Namespace, ok bool) {
	ns, ok = (*p)[name]
	return
}
func (ns *Namespace) Set(k string, v interface{}) (err error) {
	if old, ok := ns.Get(k); ok {
		return fmt.Errorf("params variable `%s` already has a value `%#v`, so that you can't change it to `%#v`.", k, old, v)
	}
	(*ns)[k] = v
	return
}
func (ns *Namespace) Get(k string) (v interface{}, ok bool) {
	v, ok = (*ns)[k]
	return
}
