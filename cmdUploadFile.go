// +build !rm_basic_commands allcommands uploadcmd

package main

import (
	"fmt"
)

func init() {
	command := Command{
		Cmd:         []string{"upload", "u"},
		Description: "$filePath $fileName - Upload file with optional name",
		Help:        "",
		Exec:        cmdUploadFile,
	}

	RegisterCommand(command)
}

func cmdUploadFile(cmd []string) {
	filePath := cmd[1]
	var fileName string
	if len(cmd) == 3 {
		fileName = cmd[2]
	} else {
		fileName = ""
	}
	chat := k.NewChat(channel)
	_, err := chat.Upload(fileName, filePath)
	if err != nil {
		printToView("Feed", fmt.Sprintf("There was an error uploading %s to %s", filePath, channel.Name))
	} else {
		printToView("Feed", fmt.Sprintf("Uploaded %s to %s", filePath, channel.Name))
	}
}
