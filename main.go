package main

import (
	"MyTEnsApp/jobs"
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("MyTEnsApp").
		SetVersion("1.0.0").
		SetDescription(
			"Tool untuk mengambil log file pada OS Linux di dalam suatu directory",
		)

	commando.
		Register(nil).
		AddArgument("fileName", "File Location", "").                                                                                 // get file name
		AddFlag("type, t", "Convert log file into plaintext or json. [valid: no-type, \"text\", \"json\"]", commando.String, "text"). // get opsi output format
		AddFlag("output, o", "Set new output location [valid: no-option, <valid directory>]", commando.String, "same").               // get opsi output location
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			filePathTxt := args["fileName"].Value
			filePath := filePathTxt

			fileTypeFinal := fmt.Sprintf("%s", flags["type"].Value)

			fileLocFinal := fmt.Sprintf("%s", flags["output"].Value)

			// menjalankan import log file
			err := jobs.Convert(filePath, fileTypeFinal, fileLocFinal)
			if err != nil {
				fmt.Println(err)
			}
		})

	commando.Parse(nil)
}
