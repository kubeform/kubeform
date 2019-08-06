package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	MaxAgeSeconds int `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
	// +optional
	Method []string `json:"method,omitempty" tf:"method,omitempty"`
	// +optional
	Origin []string `json:"origin,omitempty" tf:"origin,omitempty"`
	// +optional
	ResponseHeader []string `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type StorageBucketSpecEncryption struct {
	DefaultKmsKeyName string `json:"defaultKmsKeyName" tf:"default_kms_key_name"`
}

type StorageBucketSpecLifecycleRuleAction struct {
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
	Type         string `json:"type" tf:"type"`
}

type StorageBucketSpecLifecycleRuleCondition struct {
	// +optional
	Age int `json:"age,omitempty" tf:"age,omitempty"`
	// +optional
	CreatedBefore string `json:"createdBefore,omitempty" tf:"created_before,omitempty"`
	// +optional
	IsLive bool `json:"isLive,omitempty" tf:"is_live,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty" tf:"matches_storage_class,omitempty"`
	// +optional
	NumNewerVersions int `json:"numNewerVersions,omitempty" tf:"num_newer_versions,omitempty"`
}

type StorageBucketSpecLifecycleRule struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Action []StorageBucketSpecLifecycleRuleAction `json:"action" tf:"action"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Condition []StorageBucketSpecLifecycleRuleCondition `json:"condition" tf:"condition"`
}

type StorageBucketSpecLogging struct {
	LogBucket string `json:"logBucket" tf:"log_bucket"`
	// +optional
	LogObjectPrefix string `json:"logObjectPrefix,omitempty" tf:"log_object_prefix,omitempty"`
}

type StorageBucketSpecVersioning struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type StorageBucketSpecWebsite struct {
	// +optional
	MainPageSuffix string `json:"mainPageSuffix,omitempty" tf:"main_page_suffix,omitempty"`
	// +optional
	NotFoundPage string `json:"notFoundPage,omitempty" tf:"not_found_page,omitempty"`
}

type StorageBucketSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Cors []StorageBucketSpecCors `json:"cors,omitempty" tf:"cors,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Encryption []StorageBucketSpecEncryption `json:"encryption,omitempty" tf:"encryption,omitempty"`
	// +optional
	ForceDestroy bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	LifecycleRule []StorageBucketSpecLifecycleRule `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging []StorageBucketSpecLogging `json:"logging,omitempty" tf:"logging,omitempty"`
	Name    string                     `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Versioning []StorageBucketSpecVersioning `json:"versioning,omitempty" tf:"versioning,omitempty"`
	// +optional
	Website []StorageBucketSpecWebsite `json:"website,omitempty" tf:"website,omitempty"`
}

type StorageBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageBucketSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
