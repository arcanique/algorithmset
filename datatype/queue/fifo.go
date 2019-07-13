package queue

import "fmt"

const (
	MAX_CAP = 65535
)

type Fifo struct {
	cap uint16
	len uint16
	// element
	element []interface{}
}

func New() *Fifo {
	return &Fifo{
		cap:     0,
		len:     0,
		element: make([]interface{},0),
	}
}

func NewFromSlice(s []interface{}) (*Fifo,error) {
	if len(s) > MAX_CAP {
		return nil, fmt.Errorf("")
	}
	return &Fifo{
		cap: uint16(len(s)),
		len: uint16(len(s)),
		element: s,
	}, nil
}

// Insert to tail
func (f *Fifo)Enqueue(ele interface{})error {
	f.element = append(f.element, ele)
	return nil
}

// move from head
func (f *Fifo)Dequeue()(interface{}, error) {
	if f.len > 0 {
		f.element = f.element[1:]
		return f.element[0],nil
	}
	return nil,fmt.Errorf("FIFO is empty")
}

func (f *Fifo)Length() uint16 {
	return f.len
}

func (f *Fifo)Clear() {
	f = New()
}

func (f *Fifo)IsEmpty() bool {
	return f.len == 0
}
