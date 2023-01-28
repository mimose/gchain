package help

import (
	"testing"
	"time"
)

func TestStopWatch_all(t *testing.T) {
	sw := NewStopWatch()
	err := sw.Start("test1")
	if err != nil {
		t.Log(err)
	}
	time.Sleep(time.Second * 2)
	err = sw.Stop()
	if err != nil {
		t.Log(err)
	}
	err = sw.Start("test2")
	if err != nil {
		t.Log(err)
	}
	time.Sleep(time.Second * 1)
	err = sw.Stop()
	if err != nil {
		t.Log(err)
	}
	err = sw.Print()
	if err != nil {
		t.Log(err)
	}
}
