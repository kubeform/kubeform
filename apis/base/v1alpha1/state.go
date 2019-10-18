package v1alpha1

import "encoding/json"

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces
type State struct {
	Version          json.Number `json:"version"`
	TerraformVersion string      `json:"terraform_version"`
	Serial           uint64      `json:"serial"`
	Lineage          string      `json:"lineage"`
}
