package test

import (
	"fmt"
	"github.com/kaito2/Introduction-to-search-engine/src/models"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestExtractNextToken(t *testing.T) {
	g, err := ioutil.ReadFile(filepath.Join("testdata", "wagahai-ha-neko-dearu.golden"))
	if err != nil {
		t.Fatalf("failed reading .golden: %s", err)
	}
	str := string(g)
	iStr := []rune(str)
	var token []rune
	N := 4
	fmt.Println(iStr)
	for len(iStr) > 0 {
		token, iStr = modules.ExtractNextToken(N, iStr)
		fmt.Println(string(token))
	}
}
