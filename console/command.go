package console

import (
	"github.com/globalxtreme/gobaseconf/console/command"
	"github.com/spf13/cobra"
)

func Commands(cobraCmd *cobra.Command, newCommands []BaseInterface) {
	addCommand(cobraCmd, &command.DeleteLogFileCommand{})

	for _, newCommand := range newCommands {
		addCommand(cobraCmd, newCommand)
	}
}

func addCommand(cmd *cobra.Command, newCmd BaseInterface) {
	newCmd.Command(cmd)
}
