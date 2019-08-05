package apis

import "encoding/json"

type State struct {
	Version          json.Number `json:"version"`
	TerraformVersion string      `json:"terraform_version"`
	Serial           uint64      `json:"serial"`
	Lineage          string      `json:"lineage"`
}
