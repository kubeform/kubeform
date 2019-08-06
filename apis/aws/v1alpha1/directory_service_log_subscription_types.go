package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DirectoryServiceLogSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryServiceLogSubscriptionSpec   `json:"spec,omitempty"`
	Status            DirectoryServiceLogSubscriptionStatus `json:"status,omitempty"`
}

type DirectoryServiceLogSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DirectoryID  string `json:"directoryID" tf:"directory_id"`
	LogGroupName string `json:"logGroupName" tf:"log_group_name"`
}

type DirectoryServiceLogSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DirectoryServiceLogSubscriptionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DirectoryServiceLogSubscriptionList is a list of DirectoryServiceLogSubscriptions
type DirectoryServiceLogSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DirectoryServiceLogSubscription CRD objects
	Items []DirectoryServiceLogSubscription `json:"items,omitempty"`
}
