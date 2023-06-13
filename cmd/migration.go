package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use: "migration",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("RUNNING MIGRATION")
	},
}
