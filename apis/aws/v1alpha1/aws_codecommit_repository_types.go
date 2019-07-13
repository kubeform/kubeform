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

type AwsCodecommitRepository struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodecommitRepositorySpec   `json:"spec,omitempty"`
	Status            AwsCodecommitRepositoryStatus `json:"status,omitempty"`
}

type AwsCodecommitRepositorySpec struct {
	DefaultBranch  string            `json:"default_branch"`
	Tags           map[string]string `json:"tags"`
	RepositoryName string            `json:"repository_name"`
	Description    string            `json:"description"`
	Arn            string            `json:"arn"`
	RepositoryId   string            `json:"repository_id"`
	CloneUrlHttp   string            `json:"clone_url_http"`
	CloneUrlSsh    string            `json:"clone_url_ssh"`
}



type AwsCodecommitRepositoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCodecommitRepositoryList is a list of AwsCodecommitRepositorys
type AwsCodecommitRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodecommitRepository CRD objects
	Items []AwsCodecommitRepository `json:"items,omitempty"`
}