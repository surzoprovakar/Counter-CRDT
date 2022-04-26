package main

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {

	counter1, counter2 := NewCounter(23), NewCounter(39)

	counter1.Inc()
	counter1.Inc()

	counter2.Inc()
	counter2.Inc()
	counter2.Inc()

	counter1.Dec()
	counter2.Dec()

	b1 := counter1.ToByteArray()
	b2 := counter2.ToByteArray()

	counter3 := FromByteArray(b1)
	counter4 := FromByteArray(b2)

	fmt.Println(counter1.Print())
	fmt.Println(counter2.Print())

	fmt.Println(counter3.Print())
	fmt.Println(counter4.Print())

	if counter1.Value() != counter3.Value() ||
		counter1.Id() != counter3.Id() {
		t.Errorf("counter1 and counter3 are not the same")
	}

	if counter2.Value() != counter4.Value() ||
		counter2.Id() != counter4.Id() {
		t.Errorf("counter2 and counter4 are not the same")
	}

	// if counter1.Value() != 1 {
	// 	t.Errorf("expected total count to be %d, actual: %d",
	// 		1, counter1.Value())
	// }

	// if counter2.Value() != 2 {
	// 	t.Errorf("expected total count to be %d, actual: %d",
	// 		2, counter2.Value())
	// }

	// counter1.Merge(counter2)

	// if counter1.Value() != 2 {
	// 	t.Errorf("expected total count to be %d, actual: %d",
	// 		2, counter1.Value())
	// }

}
