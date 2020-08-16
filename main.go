package main

import (
	_ "github.com/Wenchy/tableau-demo/testpb"
	"github.com/Wenchy/tableau/pkg/tableau"
)

func main() {
	// tableau.Convert("test", "./testdata/", "./output/")
	tbx := tableau.NewTableaux(&tableau.Options{
		ProtoPackageName:          "test",
		InputPath:                 "./testdata/",
		OutputPath:                "./output/",
		OutputFilenameAsSnakeCase: false,
		OutputFormat:              tableau.JSON,
		OutputPretty:              true,
	})
	tbx.Convert()
}
