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

type DirectoryServiceLogSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryServiceLogSubscriptionSpec   `json:"spec,omitempty"`
	Status            DirectoryServiceLogSubscriptionStatus `json:"status,omitempty"`
}

type DirectoryServiceLogSubscriptionSpec struct {
	DirectoryId  string `json:"directory_id"`
	LogGroupName string `json:"log_group_name"`
}

type DirectoryServiceLogSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
