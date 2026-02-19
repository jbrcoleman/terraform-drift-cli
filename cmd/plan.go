package cmd

import (
	"fmt"
      "os"

      "github.com/jbrcoleman/terraform-drift-cli/internal/output"
      "github.com/jbrcoleman/terraform-drift-cli/internal/parser"
      "github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use: "plan",
	Short: "Parse and display a Terraform plan",
	RunE: runPlan,
}

func init () {
	planCmd.Flags().StringP("file", "f", "", "Path to plan JSON file")
	planCmd.Flags().Bool("summary",false, "Show summary only")
	planCmd.Flags().String("filter", "", "Filter by resource type")
    planCmd.Flags().Bool("json", false, "Output as JSON")
    planCmd.Flags().Bool("run", false, "Run terraform plan automatically")


	rootCmd.AddCommand(planCmd)
}

func runPlan(cmd *cobra.Command, args []string) error {
	file, _    := cmd.Flags().GetString("file")
    summary, _ := cmd.Flags().GetBool("summary")
    filter, _  := cmd.Flags().GetString("filter")
    jsonOut, _ := cmd.Flags().GetBool("json")
    run, _     := cmd.Flags().GetBool("run")

	// wire up --file
      var data []byte
      var err error
      if file != "" {
          data, err = os.ReadFile(file)
          if err != nil {
              return fmt.Errorf("could not read file: %w", err)
          }
      }

      plan, err := parser.Parse(data)
      if err != nil {
          return fmt.Errorf("could not parse plan: %w", err)
      }

      changes := parser.FilterChanges(plan, filter)

      _ = summary
      _ = jsonOut
      _ = run

      output.Render(changes, false)
      return nil
}