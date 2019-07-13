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

type AwsRamResourceShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRamResourceShareSpec   `json:"spec,omitempty"`
	Status            AwsRamResourceShareStatus `json:"status,omitempty"`
}

type AwsRamResourceShareSpec struct {
	Arn                     string            `json:"arn"`
	Name                    string            `json:"name"`
	AllowExternalPrincipals bool              `json:"allow_external_principals"`
	Tags                    map[string]string `json:"tags"`
}



type AwsRamResourceShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRamResourceShareList is a list of AwsRamResourceShares
type AwsRamResourceShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRamResourceShare CRD objects
	Items []AwsRamResourceShare `json:"items,omitempty"`
}