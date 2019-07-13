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

type GoogleStorageBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketSpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketStatus `json:"status,omitempty"`
}

type GoogleStorageBucketSpecLifecycleRuleAction struct {
	Type         string `json:"type"`
	StorageClass string `json:"storage_class"`
}

type GoogleStorageBucketSpecLifecycleRuleCondition struct {
	Age                 int      `json:"age"`
	CreatedBefore       string   `json:"created_before"`
	IsLive              bool     `json:"is_live"`
	MatchesStorageClass []string `json:"matches_storage_class"`
	NumNewerVersions    int      `json:"num_newer_versions"`
}

type GoogleStorageBucketSpecLifecycleRule struct {
	Action    []GoogleStorageBucketSpecLifecycleRule `json:"action"`
	Condition []GoogleStorageBucketSpecLifecycleRule `json:"condition"`
}

type GoogleStorageBucketSpecWebsite struct {
	MainPageSuffix string `json:"main_page_suffix"`
	NotFoundPage   string `json:"not_found_page"`
}

type GoogleStorageBucketSpecVersioning struct {
	Enabled bool `json:"enabled"`
}

type GoogleStorageBucketSpecLogging struct {
	LogBucket       string `json:"log_bucket"`
	LogObjectPrefix string `json:"log_object_prefix"`
}

type GoogleStorageBucketSpecEncryption struct {
	DefaultKmsKeyName string `json:"default_kms_key_name"`
}

type GoogleStorageBucketSpecCors struct {
	ResponseHeader []string `json:"response_header"`
	MaxAgeSeconds  int      `json:"max_age_seconds"`
	Origin         []string `json:"origin"`
	Method         []string `json:"method"`
}

type GoogleStorageBucketSpec struct {
	Labels        map[string]string         `json:"labels"`
	Url           string                    `json:"url"`
	LifecycleRule []GoogleStorageBucketSpec `json:"lifecycle_rule"`
	Website       []GoogleStorageBucketSpec `json:"website"`
	ForceDestroy  bool                      `json:"force_destroy"`
	Project       string                    `json:"project"`
	SelfLink      string                    `json:"self_link"`
	StorageClass  string                    `json:"storage_class"`
	Name          string                    `json:"name"`
	Versioning    []GoogleStorageBucketSpec `json:"versioning"`
	Logging       []GoogleStorageBucketSpec `json:"logging"`
	Encryption    []GoogleStorageBucketSpec `json:"encryption"`
	Location      string                    `json:"location"`
	PredefinedAcl string                    `json:"predefined_acl"`
	Cors          []GoogleStorageBucketSpec `json:"cors"`
}



type GoogleStorageBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageBucketList is a list of GoogleStorageBuckets
type GoogleStorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucket CRD objects
	Items []GoogleStorageBucket `json:"items,omitempty"`
}