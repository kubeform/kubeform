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

type DigitaloceanSpacesBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanSpacesBucketSpec   `json:"spec,omitempty"`
	Status            DigitaloceanSpacesBucketStatus `json:"status,omitempty"`
}

type DigitaloceanSpacesBucketSpec struct {
	Name             string `json:"name"`
	Urn              string `json:"urn"`
	Region           string `json:"region"`
	Acl              string `json:"acl"`
	BucketDomainName string `json:"bucket_domain_name"`
	ForceDestroy     bool   `json:"force_destroy"`
}



type DigitaloceanSpacesBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanSpacesBucketList is a list of DigitaloceanSpacesBuckets
type DigitaloceanSpacesBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanSpacesBucket CRD objects
	Items []DigitaloceanSpacesBucket `json:"items,omitempty"`
}