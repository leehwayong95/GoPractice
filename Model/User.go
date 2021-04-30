package PushModel

type Push struct {
	ProgressId   interface{}
	ProgressName interface{}
	PushTitle    interface{}
	PushContents interface{}
}

//Newpush is create push struct
func (p Push) Newpush(
	title string,
	data string,
	pname string,
	pid string,
) *Push {
	p = Push{pid, pname, title, data}
	return &p
}
