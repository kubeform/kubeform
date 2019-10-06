package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeTargetPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetPoolSpec   `json:"spec,omitempty"`
	Status            ComputeTargetPoolStatus `json:"status,omitempty"`
}

type ComputeTargetPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BackupPool string `json:"backupPool,omitempty" tf:"backup_pool,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	FailoverRatio json.Number `json:"failoverRatio,omitempty" tf:"failover_ratio,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthChecks []string `json:"healthChecks,omitempty" tf:"health_checks,omitempty"`
	// +optional
	Instances []string `json:"instances,omitempty" tf:"instances,omitempty"`
	Name      string   `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	SessionAffinity string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`
}

type ComputeTargetPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeTargetPoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetPoolList is a list of ComputeTargetPools
type ComputeTargetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetPool CRD objects
	Items []ComputeTargetPool `json:"items,omitempty"`
}
