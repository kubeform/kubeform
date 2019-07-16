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

type KmsCryptoKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsCryptoKeySpec   `json:"spec,omitempty"`
	Status            KmsCryptoKeyStatus `json:"status,omitempty"`
}

type KmsCryptoKeySpec struct {
	KeyRing string `json:"key_ring"`
	Name    string `json:"name"`
	// +optional
	RotationPeriod string `json:"rotation_period,omitempty"`
}

type KmsCryptoKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsCryptoKeyList is a list of KmsCryptoKeys
type KmsCryptoKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsCryptoKey CRD objects
	Items []KmsCryptoKey `json:"items,omitempty"`
}
