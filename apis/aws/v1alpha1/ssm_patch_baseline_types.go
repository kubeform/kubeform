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

type SsmPatchBaseline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmPatchBaselineSpec   `json:"spec,omitempty"`
	Status            SsmPatchBaselineStatus `json:"status,omitempty"`
}

type SsmPatchBaselineSpecApprovalRulePatchFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SsmPatchBaselineSpecApprovalRule struct {
	ApproveAfterDays int `json:"approve_after_days"`
	// +optional
	ComplianceLevel string `json:"compliance_level,omitempty"`
	// +optional
	EnableNonSecurity bool `json:"enable_non_security,omitempty"`
	// +kubebuilder:validation:MaxItems=10
	PatchFilter []SsmPatchBaselineSpecApprovalRule `json:"patch_filter"`
}

type SsmPatchBaselineSpecGlobalFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SsmPatchBaselineSpec struct {
	// +optional
	ApprovalRule *[]SsmPatchBaselineSpec `json:"approval_rule,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ApprovedPatches []string `json:"approved_patches,omitempty"`
	// +optional
	ApprovedPatchesComplianceLevel string `json:"approved_patches_compliance_level,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	GlobalFilter *[]SsmPatchBaselineSpec `json:"global_filter,omitempty"`
	Name         string                  `json:"name"`
	// +optional
	OperatingSystem string `json:"operating_system,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RejectedPatches []string `json:"rejected_patches,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SsmPatchBaselineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmPatchBaselineList is a list of SsmPatchBaselines
type SsmPatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmPatchBaseline CRD objects
	Items []SsmPatchBaseline `json:"items,omitempty"`
}
