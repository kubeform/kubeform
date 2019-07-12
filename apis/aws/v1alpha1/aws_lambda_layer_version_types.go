package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaLayerVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaLayerVersionSpec   `json:"spec,omitempty"`
	Status            AwsLambdaLayerVersionStatus `json:"status,omitempty"`
}

type AwsLambdaLayerVersionSpec struct {
	S3ObjectVersion    string   `json:"s3_object_version"`
	SourceCodeSize     int      `json:"source_code_size"`
	Version            string   `json:"version"`
	LayerName          string   `json:"layer_name"`
	S3Bucket           string   `json:"s3_bucket"`
	CompatibleRuntimes []string `json:"compatible_runtimes"`
	Arn                string   `json:"arn"`
	LayerArn           string   `json:"layer_arn"`
	CreatedDate        string   `json:"created_date"`
	SourceCodeHash     string   `json:"source_code_hash"`
	S3Key              string   `json:"s3_key"`
	Description        string   `json:"description"`
	LicenseInfo        string   `json:"license_info"`
	Filename           string   `json:"filename"`
}

type AwsLambdaLayerVersionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaLayerVersionList is a list of AwsLambdaLayerVersions
type AwsLambdaLayerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLambdaLayerVersion CRD objects
	Items []AwsLambdaLayerVersion `json:"items,omitempty"`
}
