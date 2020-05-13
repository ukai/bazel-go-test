package sample

import "testing"

func TestLoad(t *testing.T) {
	_, err := Load("testdata/foo")
	if err != nil {
		t.Errorf("Load %v; want nil err", err)
	}
}
