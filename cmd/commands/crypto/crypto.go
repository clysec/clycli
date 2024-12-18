package crypto

import (
	"github.com/v3-nu/vv/cmd/utils"
	"github.com/spf13/cobra"
)

var ExportCommands = utils.CommandGroup{
	Command: &cobra.Command{
		Use:     "crypto",
		Aliases: []string{"ssl", "crypt", "tls"},
		Short:   "Crypto related commands",
	},
	Children: []func(*cobra.Command){
		PfxToPem,
	},
}
