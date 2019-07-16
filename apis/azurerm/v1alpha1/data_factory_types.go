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

type DataFactory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactorySpec   `json:"spec,omitempty"`
	Status            DataFactoryStatus `json:"status,omitempty"`
}

type DataFactorySpecGithubConfiguration struct {
	AccountName    string `json:"account_name"`
	BranchName     string `json:"branch_name"`
	GitUrl         string `json:"git_url"`
	RepositoryName string `json:"repository_name"`
	RootFolder     string `json:"root_folder"`
}

type DataFactorySpecVstsConfiguration struct {
	AccountName    string `json:"account_name"`
	BranchName     string `json:"branch_name"`
	ProjectName    string `json:"project_name"`
	RepositoryName string `json:"repository_name"`
	RootFolder     string `json:"root_folder"`
	TenantId       string `json:"tenant_id"`
}

type DataFactorySpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	GithubConfiguration *[]DataFactorySpec `json:"github_configuration,omitempty"`
	Location            string             `json:"location"`
	Name                string             `json:"name"`
	ResourceGroupName   string             `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VstsConfiguration *[]DataFactorySpec `json:"vsts_configuration,omitempty"`
}

type DataFactoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryList is a list of DataFactorys
type DataFactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactory CRD objects
	Items []DataFactory `json:"items,omitempty"`
}
