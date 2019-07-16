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

type ServiceAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountSpec   `json:"spec,omitempty"`
	Status            ServiceAccountStatus `json:"status,omitempty"`
}

type ServiceAccountSpec struct {
	AccountId string `json:"account_id"`
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	// +optional
	// Deprecated
	PolicyData string `json:"policy_data,omitempty"`
}

type ServiceAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceAccountList is a list of ServiceAccounts
type ServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceAccount CRD objects
	Items []ServiceAccount `json:"items,omitempty"`
}
