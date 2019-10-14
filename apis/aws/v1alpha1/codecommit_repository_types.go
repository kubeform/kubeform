package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type CodecommitRepository struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodecommitRepositorySpec   `json:"spec,omitempty"`
	Status            CodecommitRepositoryStatus `json:"status,omitempty"`
}

type CodecommitRepositorySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CloneURLHTTP string `json:"cloneURLHTTP,omitempty" tf:"clone_url_http,omitempty"`
	// +optional
	CloneURLSSH string `json:"cloneURLSSH,omitempty" tf:"clone_url_ssh,omitempty"`
	// +optional
	DefaultBranch string `json:"defaultBranch,omitempty" tf:"default_branch,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	RepositoryID   string `json:"repositoryID,omitempty" tf:"repository_id,omitempty"`
	RepositoryName string `json:"repositoryName" tf:"repository_name"`
}

type CodecommitRepositoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodecommitRepositorySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodecommitRepositoryList is a list of CodecommitRepositorys
type CodecommitRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodecommitRepository CRD objects
	Items []CodecommitRepository `json:"items,omitempty"`
}
