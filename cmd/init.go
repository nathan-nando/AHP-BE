package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Mode struct {
	AppMode string
	DbMode  string
}

var rootCmd = &cobra.Command{Run: func(cmd *cobra.Command, args []string) {
	logrus.Info("use - h to show available commands")
},
}

func Init() *Mode {
	appCmd.Flags().String("appMode", "local", "SELECT APP MODE")
	appCmd.Flags().String("dbMode", "mysql", "SELECT DB MODE")
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(migrationCmd)

	err := rootCmd.Execute()
	if err != nil {
		logrus.Error(err)
	}

	return &Mode{
		AppMode: appCmd.Flag("appMode").Value.String(),
		DbMode:  appCmd.Flag("dbMode").Value.String(),
	}
}
