package main

import (
	"encoding/binary"
	"fmt"
	"sync"
)

var mu sync.Mutex

// Initialization
type hist struct {
	req_val      int
	delta        int
	trust        float64
	total_req    int
	accepted_req int
	decision     string
}

type Counter struct {
	id      int
	value   int
	history map[int][]hist
}

func NewCounter(id int) *Counter {
	return &Counter{id: id, value: 0, history: make(map[int][]hist)}
}

// Get methods
func (c *Counter) Id() int {
	return c.id
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) History() map[int][]hist {
	return c.history
}

// Counter oerations
func (c *Counter) Inc() {
	mu.Lock()
	c.value++
	mu.Unlock()
}

func (c *Counter) Dec() {
	mu.Lock()
	c.value--
	mu.Unlock()
}

// Communication methods
func (c *Counter) ToByteArray() []byte {

	a1 := make([]byte, 64)
	a2 := make([]byte, 64)

	binary.LittleEndian.PutUint64(a1, uint64(c.Id()))
	binary.LittleEndian.PutUint64(a2, uint64(c.Value()))
	return append(a1, a2...)
}

func FromByteArray(bytes []byte) *Counter {

	var r1 = binary.LittleEndian.Uint64(bytes[0:(len(bytes) / 2)])
	var r2 = binary.LittleEndian.Uint64(bytes[len(bytes)/2:])
	id := int64(r1)
	c := NewCounter(int(id))
	c.value = int(int64(r2))
	return c

}

// Additional methods
func (c *Counter) Print() string {
	res := fmt.Sprintf("%s%d:%d", "Counter_", c.id, c.value)
	return res
}

// Merge methods
func Merge(c *Counter, o *Counter) {
	mu.Lock()
	if Decision(c, o) {
		c.value = o.value
	}
	mu.Unlock()
}

func main() {
	fmt.Print("hello")
}
