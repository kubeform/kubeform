package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDatasyncLocationS3 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDatasyncLocationS3Spec   `json:"spec,omitempty"`
	Status            AwsDatasyncLocationS3Status `json:"status,omitempty"`
}

type AwsDatasyncLocationS3SpecS3Config struct {
	BucketAccessRoleArn string `json:"bucket_access_role_arn"`
}

type AwsDatasyncLocationS3Spec struct {
	S3Config     []AwsDatasyncLocationS3Spec `json:"s3_config"`
	Subdirectory string                      `json:"subdirectory"`
	Tags         map[string]string           `json:"tags"`
	Uri          string                      `json:"uri"`
	Arn          string                      `json:"arn"`
	S3BucketArn  string                      `json:"s3_bucket_arn"`
}

type AwsDatasyncLocationS3Status struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDatasyncLocationS3List is a list of AwsDatasyncLocationS3s
type AwsDatasyncLocationS3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDatasyncLocationS3 CRD objects
	Items []AwsDatasyncLocationS3 `json:"items,omitempty"`
}
