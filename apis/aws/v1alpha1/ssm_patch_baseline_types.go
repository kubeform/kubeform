package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	Key    string   `json:"key" tf:"key"`
	Values []string `json:"values" tf:"values"`
}

type SsmPatchBaselineSpecApprovalRule struct {
	ApproveAfterDays int `json:"approveAfterDays" tf:"approve_after_days"`
	// +optional
	ComplianceLevel string `json:"complianceLevel,omitempty" tf:"compliance_level,omitempty"`
	// +optional
	EnableNonSecurity bool `json:"enableNonSecurity,omitempty" tf:"enable_non_security,omitempty"`
	// +kubebuilder:validation:MaxItems=10
	PatchFilter []SsmPatchBaselineSpecApprovalRulePatchFilter `json:"patchFilter" tf:"patch_filter"`
}

type SsmPatchBaselineSpecGlobalFilter struct {
	Key    string   `json:"key" tf:"key"`
	Values []string `json:"values" tf:"values"`
}

type SsmPatchBaselineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApprovalRule []SsmPatchBaselineSpecApprovalRule `json:"approvalRule,omitempty" tf:"approval_rule,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ApprovedPatches []string `json:"approvedPatches,omitempty" tf:"approved_patches,omitempty"`
	// +optional
	ApprovedPatchesComplianceLevel string `json:"approvedPatchesComplianceLevel,omitempty" tf:"approved_patches_compliance_level,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	GlobalFilter []SsmPatchBaselineSpecGlobalFilter `json:"globalFilter,omitempty" tf:"global_filter,omitempty"`
	Name         string                             `json:"name" tf:"name"`
	// +optional
	OperatingSystem string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RejectedPatches []string `json:"rejectedPatches,omitempty" tf:"rejected_patches,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type SsmPatchBaselineStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *SsmPatchBaselineSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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