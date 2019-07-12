package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmPatchBaseline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmPatchBaselineSpec   `json:"spec,omitempty"`
	Status            AwsSsmPatchBaselineStatus `json:"status,omitempty"`
}

type AwsSsmPatchBaselineSpecApprovalRulePatchFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmPatchBaselineSpecApprovalRule struct {
	ApproveAfterDays  int                                   `json:"approve_after_days"`
	ComplianceLevel   string                                `json:"compliance_level"`
	EnableNonSecurity bool                                  `json:"enable_non_security"`
	PatchFilter       []AwsSsmPatchBaselineSpecApprovalRule `json:"patch_filter"`
}

type AwsSsmPatchBaselineSpecGlobalFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmPatchBaselineSpec struct {
	ApprovalRule                   []AwsSsmPatchBaselineSpec `json:"approval_rule"`
	Name                           string                    `json:"name"`
	Description                    string                    `json:"description"`
	GlobalFilter                   []AwsSsmPatchBaselineSpec `json:"global_filter"`
	ApprovedPatches                []string                  `json:"approved_patches"`
	RejectedPatches                []string                  `json:"rejected_patches"`
	OperatingSystem                string                    `json:"operating_system"`
	ApprovedPatchesComplianceLevel string                    `json:"approved_patches_compliance_level"`
	Tags                           map[string]string         `json:"tags"`
}



type AwsSsmPatchBaselineStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmPatchBaselineList is a list of AwsSsmPatchBaselines
type AwsSsmPatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmPatchBaseline CRD objects
	Items []AwsSsmPatchBaseline `json:"items,omitempty"`
}