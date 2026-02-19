package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/faith/color"
	"github.com/jbrcoleman/terraform-drift-cli/internal/parser"
)

var (
	green = color.New(color.FgGreen).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
    red    = color.New(color.FgRed).SprintfFunc()
    cyan   = color.New(color.FgCyan).SprintfFunc()
)

type JSONResource struct {
	Address string `json:"address"`
	Module string `json:"module,omitempty"`
	Type string `json:"type"`
	Action string `json:"action"`
}

func Render(changes []parser.ResourceChange, summary bool) {
        groups := groupByAction(changes)

        if creates := groups[parser.ActionCreate]; len(creates) > 0 {
                fmt.Printf("\n  %s (%d)\n", green("+ create"), len(creates))
                if !summary {
                        for _, r := range creates {
                                fmt.Printf("    %s  %s\n", green("+"), r.Address)
                        }
                }
        }

        if updates := groups[parser.ActionUpdate]; len(updates) > 0 {
                fmt.Printf("\n  %s (%d)\n", yellow("~ update"), len(updates))
                if !summary {
                        for _, r := range updates {
                                fmt.Printf("    %s  %s\n", yellow("~"), r.Address)
                        }
                }
        }

        if replaces := groups[parser.ActionReplace]; len(replaces) > 0 {
                fmt.Printf("\n  %s (%d)\n", cyan("± replace"), len(replaces))
                if !summary {
                        for _, r := range replaces {
                                fmt.Printf("    %s  %s\n", cyan("±"), r.Address)
                                if r.ModuleAddress != "" {
                                        fmt.Printf("               module: %s\n", r.ModuleAddress)
                                }
                        }
                }
        }

        if deletes := groups[parser.ActionDelete]; len(deletes) > 0 {
                fmt.Printf("\n  %s (%d)\n", red("- destroy"), len(deletes))
                if !summary {
                        for _, r := range deletes {
                                fmt.Printf("    %s  %s\n", red("-"), r.Address)
                        }
                }
        }

        fmt.Printf("\nPlan: %d to add, %d to change, %d to replace, %d to destroy.\n\n",
                len(groups[parser.ActionCreate]),
                len(groups[parser.ActionUpdate]),
                len(groups[parser.ActionReplace]),
                len(groups[parser.ActionDelete]),
        )
  }

  func RenderJSON(changes []parser.ResourceChange) error {
        out := make([]JSONResource, 0, len(changes))
        for _, rc := range changes {
                out = append(out, JSONResource{
                        Address: rc.Address,
                        Module:  rc.ModuleAddress,
                        Type:    rc.Type,
                        Action:  resolveAction(rc.Change.Actions),
                })
        }
        return json.NewEncoder(os.Stdout).Encode(out)
  }

  func resolveAction(actions []string) string {
        if len(actions) == 2 {
                return parser.ActionReplace
        }
        if len(actions) == 1 {
                return actions[0]
        }
        return parser.ActionNoOp
  }