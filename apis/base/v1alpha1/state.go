/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// xref: https://github.com/hashicorp/terraform/blob/d7bce857cf7e2ab8cd6378196ee95b9a935acb5c/states/statefile/version4.go#L497

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces
type State struct {
	Version          int64  `json:"version"`
	TerraformVersion string `json:"terraform_version"`
	Serial           uint64 `json:"serial"`
	Lineage          string `json:"lineage"`
}
