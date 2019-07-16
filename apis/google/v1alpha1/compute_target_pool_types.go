package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	// +optional
	BackupPool string `json:"backup_pool,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	FailoverRatio json.Number `json:"failover_ratio,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthChecks []string `json:"health_checks,omitempty"`
	Name         string   `json:"name"`
	// +optional
	SessionAffinity string `json:"session_affinity,omitempty"`
}

type ComputeTargetPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
