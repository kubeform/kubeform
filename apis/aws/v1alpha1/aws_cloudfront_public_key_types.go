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

type AwsCloudfrontPublicKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudfrontPublicKeySpec   `json:"spec,omitempty"`
	Status            AwsCloudfrontPublicKeyStatus `json:"status,omitempty"`
}

type AwsCloudfrontPublicKeySpec struct {
	CallerReference string `json:"caller_reference"`
	Comment         string `json:"comment"`
	EncodedKey      string `json:"encoded_key"`
	Etag            string `json:"etag"`
	Name            string `json:"name"`
	NamePrefix      string `json:"name_prefix"`
}



type AwsCloudfrontPublicKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudfrontPublicKeyList is a list of AwsCloudfrontPublicKeys
type AwsCloudfrontPublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudfrontPublicKey CRD objects
	Items []AwsCloudfrontPublicKey `json:"items,omitempty"`
}