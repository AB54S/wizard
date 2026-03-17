/*
Copyright © 2026 Abbaas Mohamud ab54s@github.io
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func Execute() error {
	rootCmd := &cobra.Command{
		Version: 	"v0.0.1",
		Use:   		"wizard",
		Short: 		"Wizard is you terminal agent.",
		Long:  		`Wizard displays system stats (CPU, GPU, Memory, Kernel, DE) in the terminal`,
		Example: 	"wizard",

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(initialize())
	return rootCmd.ExecuteContext(context.Background())
}
