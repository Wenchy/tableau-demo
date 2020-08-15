package main

import (
	_ "github.com/Wenchy/tableau-demo/testpb"
	"github.com/Wenchy/tableau/pkg/tableau"
)

func main() {
	tableau.Convert("test", "./testdata/")
}
