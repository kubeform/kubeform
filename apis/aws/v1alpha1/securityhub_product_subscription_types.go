package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SecurityhubProductSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityhubProductSubscriptionSpec   `json:"spec,omitempty"`
	Status            SecurityhubProductSubscriptionStatus `json:"status,omitempty"`
}

type SecurityhubProductSubscriptionSpec struct {
	ProductArn  string                    `json:"productArn" tf:"product_arn"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SecurityhubProductSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecurityhubProductSubscriptionList is a list of SecurityhubProductSubscriptions
type SecurityhubProductSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityhubProductSubscription CRD objects
	Items []SecurityhubProductSubscription `json:"items,omitempty"`
}
