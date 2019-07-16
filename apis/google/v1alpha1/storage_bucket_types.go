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

type StorageBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketSpec   `json:"spec,omitempty"`
	Status            StorageBucketStatus `json:"status,omitempty"`
}

type StorageBucketSpecCors struct {
	// +optional
	MaxAgeSeconds int `json:"max_age_seconds,omitempty"`
	// +optional
	Method []string `json:"method,omitempty"`
	// +optional
	Origin []string `json:"origin,omitempty"`
	// +optional
	ResponseHeader []string `json:"response_header,omitempty"`
}

type StorageBucketSpecEncryption struct {
	DefaultKmsKeyName string `json:"default_kms_key_name"`
}

type StorageBucketSpecLifecycleRuleAction struct {
	// +optional
	StorageClass string `json:"storage_class,omitempty"`
	Type         string `json:"type"`
}

type StorageBucketSpecLifecycleRuleCondition struct {
	// +optional
	Age int `json:"age,omitempty"`
	// +optional
	CreatedBefore string `json:"created_before,omitempty"`
	// +optional
	IsLive bool `json:"is_live,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchesStorageClass []string `json:"matches_storage_class,omitempty"`
	// +optional
	NumNewerVersions int `json:"num_newer_versions,omitempty"`
}

type StorageBucketSpecLifecycleRule struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Action []StorageBucketSpecLifecycleRule `json:"action"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Condition []StorageBucketSpecLifecycleRule `json:"condition"`
}

type StorageBucketSpecLogging struct {
	LogBucket string `json:"log_bucket"`
}

type StorageBucketSpecVersioning struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
}

type StorageBucketSpecWebsite struct {
	// +optional
	MainPageSuffix string `json:"main_page_suffix,omitempty"`
	// +optional
	NotFoundPage string `json:"not_found_page,omitempty"`
}

type StorageBucketSpec struct {
	// +optional
	Cors *[]StorageBucketSpec `json:"cors,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Encryption *[]StorageBucketSpec `json:"encryption,omitempty"`
	// +optional
	ForceDestroy bool `json:"force_destroy,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	LifecycleRule *[]StorageBucketSpec `json:"lifecycle_rule,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging *[]StorageBucketSpec `json:"logging,omitempty"`
	Name    string               `json:"name"`
	// +optional
	StorageClass string `json:"storage_class,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Versioning *[]StorageBucketSpec `json:"versioning,omitempty"`
	// +optional
	Website *[]StorageBucketSpec `json:"website,omitempty"`
}

type StorageBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketList is a list of StorageBuckets
type StorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucket CRD objects
	Items []StorageBucket `json:"items,omitempty"`
}
