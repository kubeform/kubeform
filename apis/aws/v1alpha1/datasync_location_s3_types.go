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

type DatasyncLocationS3 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncLocationS3Spec   `json:"spec,omitempty"`
	Status            DatasyncLocationS3Status `json:"status,omitempty"`
}

type DatasyncLocationS3SpecS3Config struct {
	BucketAccessRoleArn string `json:"bucket_access_role_arn"`
}

type DatasyncLocationS3Spec struct {
	S3BucketArn string `json:"s3_bucket_arn"`
	// +kubebuilder:validation:MaxItems=1
	S3Config     []DatasyncLocationS3Spec `json:"s3_config"`
	Subdirectory string                   `json:"subdirectory"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DatasyncLocationS3Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatasyncLocationS3List is a list of DatasyncLocationS3s
type DatasyncLocationS3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatasyncLocationS3 CRD objects
	Items []DatasyncLocationS3 `json:"items,omitempty"`
}
