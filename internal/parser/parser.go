package parser

import "encoding/json"

type Plan struct{
	ResourceChanges []ResourceChange `json:"resource_changes"`
	PlannedValues PlannedValues `json:"planned_values"`
}

type ResourceChange struct{
	Address string `json:"address"`
	ModuleAddress string `json:"module_address"`
	Type string `json:"type"`
	Change Change `json:"change"`
}

type Change struct {
	Actions []string `json:"actions"`
}

type PlannedValues struct{}

const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionReplace = "replace"
	ActionNoOp = "no-op"
)

func Parse(data []byte) (*Plan, error) {
	var plan Plan 
	if err := json.Unmarshal(data,&plan); err != nil {
		return nil, err
	}
	return &plan, nil
}

func FilterChanges(plan *Plan, filter string) []ResourceChange {
	if filter == "" {
		return plan.ResourceChanges
	}
	var result []ResourceChange
      for _, rc := range plan.ResourceChanges {
          if rc.Type == filter {
              result = append(result, rc)
          }
      }
      return result
}