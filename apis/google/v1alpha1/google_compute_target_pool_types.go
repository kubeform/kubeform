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

type GoogleComputeTargetPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeTargetPoolSpec   `json:"spec,omitempty"`
	Status            GoogleComputeTargetPoolStatus `json:"status,omitempty"`
}

type GoogleComputeTargetPoolSpec struct {
	FailoverRatio   float64  `json:"failover_ratio"`
	SelfLink        string   `json:"self_link"`
	BackupPool      string   `json:"backup_pool"`
	Description     string   `json:"description"`
	HealthChecks    []string `json:"health_checks"`
	Instances       []string `json:"instances"`
	Project         string   `json:"project"`
	Region          string   `json:"region"`
	SessionAffinity string   `json:"session_affinity"`
	Name            string   `json:"name"`
}

type GoogleComputeTargetPoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeTargetPoolList is a list of GoogleComputeTargetPools
type GoogleComputeTargetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeTargetPool CRD objects
	Items []GoogleComputeTargetPool `json:"items,omitempty"`
}
