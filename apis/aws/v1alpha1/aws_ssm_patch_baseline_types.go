package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AwsSsmPatchBaseline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmPatchBaselineSpec   `json:"spec,omitempty"`
	Status            AwsSsmPatchBaselineStatus `json:"status,omitempty"`
}

type AwsSsmPatchBaselineSpecGlobalFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
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

type AwsSsmPatchBaselineSpec struct {
	GlobalFilter                   []AwsSsmPatchBaselineSpec `json:"global_filter"`
	RejectedPatches                []string                  `json:"rejected_patches"`
	Tags                           map[string]string         `json:"tags"`
	Name                           string                    `json:"name"`
	Description                    string                    `json:"description"`
	ApprovalRule                   []AwsSsmPatchBaselineSpec `json:"approval_rule"`
	ApprovedPatches                []string                  `json:"approved_patches"`
	OperatingSystem                string                    `json:"operating_system"`
	ApprovedPatchesComplianceLevel string                    `json:"approved_patches_compliance_level"`
}



type AwsSsmPatchBaselineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmPatchBaselineList is a list of AwsSsmPatchBaselines
type AwsSsmPatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmPatchBaseline CRD objects
	Items []AwsSsmPatchBaseline `json:"items,omitempty"`
}