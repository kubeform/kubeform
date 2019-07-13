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

type GoogleCloudbuildTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleCloudbuildTriggerSpec   `json:"spec,omitempty"`
	Status            GoogleCloudbuildTriggerStatus `json:"status,omitempty"`
}

type GoogleCloudbuildTriggerSpecBuildStep struct {
	Args string `json:"args"`
	Name string `json:"name"`
}

type GoogleCloudbuildTriggerSpecBuild struct {
	Images []string                           `json:"images"`
	Step   []GoogleCloudbuildTriggerSpecBuild `json:"step"`
	Tags   []string                           `json:"tags"`
}

type GoogleCloudbuildTriggerSpecTriggerTemplate struct {
	Project    string `json:"project"`
	RepoName   string `json:"repo_name"`
	TagName    string `json:"tag_name"`
	BranchName string `json:"branch_name"`
	CommitSha  string `json:"commit_sha"`
	Dir        string `json:"dir"`
}

type GoogleCloudbuildTriggerSpec struct {
	Filename        string                        `json:"filename"`
	Build           []GoogleCloudbuildTriggerSpec `json:"build"`
	Description     string                        `json:"description"`
	Substitutions   map[string]string             `json:"substitutions"`
	TriggerTemplate []GoogleCloudbuildTriggerSpec `json:"trigger_template"`
	Project         string                        `json:"project"`
}



type GoogleCloudbuildTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleCloudbuildTriggerList is a list of GoogleCloudbuildTriggers
type GoogleCloudbuildTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleCloudbuildTrigger CRD objects
	Items []GoogleCloudbuildTrigger `json:"items,omitempty"`
}