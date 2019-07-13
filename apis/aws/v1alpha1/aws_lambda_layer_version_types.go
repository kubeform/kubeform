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

type AwsLambdaLayerVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaLayerVersionSpec   `json:"spec,omitempty"`
	Status            AwsLambdaLayerVersionStatus `json:"status,omitempty"`
}

type AwsLambdaLayerVersionSpec struct {
	LayerName          string   `json:"layer_name"`
	S3Bucket           string   `json:"s3_bucket"`
	LicenseInfo        string   `json:"license_info"`
	LayerArn           string   `json:"layer_arn"`
	SourceCodeSize     int      `json:"source_code_size"`
	S3ObjectVersion    string   `json:"s3_object_version"`
	CompatibleRuntimes []string `json:"compatible_runtimes"`
	Description        string   `json:"description"`
	CreatedDate        string   `json:"created_date"`
	S3Key              string   `json:"s3_key"`
	SourceCodeHash     string   `json:"source_code_hash"`
	Version            string   `json:"version"`
	Filename           string   `json:"filename"`
	Arn                string   `json:"arn"`
}



type AwsLambdaLayerVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLambdaLayerVersionList is a list of AwsLambdaLayerVersions
type AwsLambdaLayerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLambdaLayerVersion CRD objects
	Items []AwsLambdaLayerVersion `json:"items,omitempty"`
}