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

type ServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountKeySpec   `json:"spec,omitempty"`
	Status            ServiceAccountKeyStatus `json:"status,omitempty"`
}

type ServiceAccountKeySpec struct {
	// +optional
	KeyAlgorithm string `json:"key_algorithm,omitempty"`
	// +optional
	PgpKey string `json:"pgp_key,omitempty"`
	// +optional
	PrivateKeyType string `json:"private_key_type,omitempty"`
	// +optional
	PublicKeyType    string `json:"public_key_type,omitempty"`
	ServiceAccountId string `json:"service_account_id"`
}

type ServiceAccountKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceAccountKeyList is a list of ServiceAccountKeys
type ServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceAccountKey CRD objects
	Items []ServiceAccountKey `json:"items,omitempty"`
}
