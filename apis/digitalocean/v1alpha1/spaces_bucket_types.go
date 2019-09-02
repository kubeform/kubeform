package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SpacesBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpacesBucketSpec   `json:"spec,omitempty"`
	Status            SpacesBucketStatus `json:"status,omitempty"`
}

type SpacesBucketSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Canned ACL applied on bucket creation
	// +optional
	Acl string `json:"acl,omitempty" tf:"acl,omitempty"`
	// The FQDN of the bucket
	// +optional
	BucketDomainName string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`
	// Unless true, the bucket will only be destroyed if empty
	// +optional
	ForceDestroy bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	// Bucket name
	Name string `json:"name" tf:"name"`
	// Bucket region
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// the uniform resource name for the bucket
	// +optional
	Urn string `json:"urn,omitempty" tf:"urn,omitempty"`
}



type SpacesBucketStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *SpacesBucketSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpacesBucketList is a list of SpacesBuckets
type SpacesBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpacesBucket CRD objects
	Items []SpacesBucket `json:"items,omitempty"`
}