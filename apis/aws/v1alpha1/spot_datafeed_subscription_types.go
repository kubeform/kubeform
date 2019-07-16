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

type SpotDatafeedSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotDatafeedSubscriptionSpec   `json:"spec,omitempty"`
	Status            SpotDatafeedSubscriptionStatus `json:"status,omitempty"`
}

type SpotDatafeedSubscriptionSpec struct {
	Bucket string `json:"bucket"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
}

type SpotDatafeedSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpotDatafeedSubscriptionList is a list of SpotDatafeedSubscriptions
type SpotDatafeedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpotDatafeedSubscription CRD objects
	Items []SpotDatafeedSubscription `json:"items,omitempty"`
}
