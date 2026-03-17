/*
Copyright © 2026 Abbaas Mohamud ab54s@github.io
*/
package cmd

import (

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func initialize() *cobra.Command {
	init := &cobra.Command{
		Use:   		"init",
		Short: 		"Initilizaes your wizard",
		Long:  		`Sets up local agent and runs a scan of host system`,
		Example: 	"wizard init",

		RunE: func(cmd *cobra.Command, args []string) error { 
			cmd.Help()
			return nil
		},
	}
	
	return init
}
