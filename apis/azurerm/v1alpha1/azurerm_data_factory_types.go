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

type AzurermDataFactory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactorySpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryStatus `json:"status,omitempty"`
}

type AzurermDataFactorySpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermDataFactorySpecGithubConfiguration struct {
	BranchName     string `json:"branch_name"`
	GitUrl         string `json:"git_url"`
	RepositoryName string `json:"repository_name"`
	RootFolder     string `json:"root_folder"`
	AccountName    string `json:"account_name"`
}

type AzurermDataFactorySpecVstsConfiguration struct {
	AccountName    string `json:"account_name"`
	BranchName     string `json:"branch_name"`
	ProjectName    string `json:"project_name"`
	RepositoryName string `json:"repository_name"`
	RootFolder     string `json:"root_folder"`
	TenantId       string `json:"tenant_id"`
}

type AzurermDataFactorySpec struct {
	ResourceGroupName   string                   `json:"resource_group_name"`
	Identity            []AzurermDataFactorySpec `json:"identity"`
	GithubConfiguration []AzurermDataFactorySpec `json:"github_configuration"`
	VstsConfiguration   []AzurermDataFactorySpec `json:"vsts_configuration"`
	Tags                map[string]string        `json:"tags"`
	Name                string                   `json:"name"`
	Location            string                   `json:"location"`
}

type AzurermDataFactoryStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataFactoryList is a list of AzurermDataFactorys
type AzurermDataFactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactory CRD objects
	Items []AzurermDataFactory `json:"items,omitempty"`
}
