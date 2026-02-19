package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
      Use:   "tfdrift",
      Short: "Visualise Terraform plan changes at a glance",
  }

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: An error while executing '%s'\n", err)
		os.Exit(1)
	}
}
