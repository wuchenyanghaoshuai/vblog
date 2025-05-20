package init

import (
	"fmt"
//	"vblog/apps/cmd"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Init vblog service",
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	  fmt.Println("init ........")
	},
  }

func init() {
	
}