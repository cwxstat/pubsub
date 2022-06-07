package pub

import (
	"fmt"
	"testing"
)

func TestPub_Add(t *testing.T) {
	p := NewPub()
	for i := 1; i < 5; i++ {
		msg := MSG{Message: fmt.Sprintf("test_%d", i)}
		p.Add("1", msg)
	}

}
