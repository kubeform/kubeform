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

type LambdaLayerVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaLayerVersionSpec   `json:"spec,omitempty"`
	Status            LambdaLayerVersionStatus `json:"status,omitempty"`
}

type LambdaLayerVersionSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:UniqueItems=true
	CompatibleRuntimes []string `json:"compatible_runtimes,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Filename  string `json:"filename,omitempty"`
	LayerName string `json:"layer_name"`
	// +optional
	LicenseInfo string `json:"license_info,omitempty"`
	// +optional
	S3Bucket string `json:"s3_bucket,omitempty"`
	// +optional
	S3Key string `json:"s3_key,omitempty"`
	// +optional
	S3ObjectVersion string `json:"s3_object_version,omitempty"`
}

type LambdaLayerVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaLayerVersionList is a list of LambdaLayerVersions
type LambdaLayerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaLayerVersion CRD objects
	Items []LambdaLayerVersion `json:"items,omitempty"`
}
