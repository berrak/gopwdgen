package main

import (
	"os"
	"testing"

	gopwdgen "github.com/berrak/gopwdgen/pkg/lib"
)

func TestInit(t *testing.T) {

	os.Clearenv()
	cmdline := os.Args
	var ps gopwdgen.ParsedStruct
	want := ps
	got := gopwdgen.Init(cmdline, ps, Version)
	if got != want {
		t.Errorf("want %+v but got %+v", ps, got)
	}

}
