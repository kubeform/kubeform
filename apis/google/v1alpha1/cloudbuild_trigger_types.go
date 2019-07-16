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

type CloudbuildTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudbuildTriggerSpec   `json:"spec,omitempty"`
	Status            CloudbuildTriggerStatus `json:"status,omitempty"`
}

type CloudbuildTriggerSpecBuildStep struct {
	// +optional
	Args string `json:"args,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
}

type CloudbuildTriggerSpecBuild struct {
	// +optional
	Images []string `json:"images,omitempty"`
	// +optional
	Step *[]CloudbuildTriggerSpecBuild `json:"step,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty"`
}

type CloudbuildTriggerSpecTriggerTemplate struct {
	// +optional
	BranchName string `json:"branch_name,omitempty"`
	// +optional
	CommitSha string `json:"commit_sha,omitempty"`
	// +optional
	Dir string `json:"dir,omitempty"`
	// +optional
	Project string `json:"project,omitempty"`
	// +optional
	RepoName string `json:"repo_name,omitempty"`
	// +optional
	TagName string `json:"tag_name,omitempty"`
}

type CloudbuildTriggerSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Build *[]CloudbuildTriggerSpec `json:"build,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Filename string `json:"filename,omitempty"`
	// +optional
	Substitutions map[string]string `json:"substitutions,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TriggerTemplate *[]CloudbuildTriggerSpec `json:"trigger_template,omitempty"`
}

type CloudbuildTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudbuildTriggerList is a list of CloudbuildTriggers
type CloudbuildTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudbuildTrigger CRD objects
	Items []CloudbuildTrigger `json:"items,omitempty"`
}
