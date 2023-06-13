package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use: "app",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("RUNNING APP")
	},
}
