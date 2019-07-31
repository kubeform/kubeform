package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:UniqueItems=true
	CompatibleRuntimes []string `json:"compatibleRuntimes,omitempty" tf:"compatible_runtimes,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Filename  string `json:"filename,omitempty" tf:"filename,omitempty"`
	LayerName string `json:"layerName" tf:"layer_name"`
	// +optional
	LicenseInfo string `json:"licenseInfo,omitempty" tf:"license_info,omitempty"`
	// +optional
	S3Bucket string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`
	// +optional
	S3Key string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`
	// +optional
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`
	// +optional
	SourceCodeHash string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`
}

type LambdaLayerVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
