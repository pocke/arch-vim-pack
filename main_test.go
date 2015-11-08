package main

import (
	"testing"

	"github.com/pocke/httpmock"
)

func TestGetLatestVimVersion(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := Asset("testdata/vim_releases.html")
	if err != nil {
		t.Fatal(err)
	}
	httpmock.RegisterResponder("GET", "https://github.com/vim/vim/releases",
		httpmock.NewBytesResponder(200, b))

	v, err := GetLatestVimVersion()
	if err != nil {
		t.Fatal(err)
	}
	if v != "7.4.909" {
		t.Fatalf("Expected 7.4.909, but got %s", v)
	}
}
