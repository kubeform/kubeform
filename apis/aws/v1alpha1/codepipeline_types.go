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

type Codepipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodepipelineSpec   `json:"spec,omitempty"`
	Status            CodepipelineStatus `json:"status,omitempty"`
}

type CodepipelineSpecArtifactStoreEncryptionKey struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CodepipelineSpecArtifactStore struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionKey *[]CodepipelineSpecArtifactStore `json:"encryption_key,omitempty"`
	Location      string                           `json:"location"`
	Type          string                           `json:"type"`
}

type CodepipelineSpecStageAction struct {
	Category string `json:"category"`
	// +optional
	Configuration map[string]string `json:"configuration,omitempty"`
	// +optional
	InputArtifacts []string `json:"input_artifacts,omitempty"`
	Name           string   `json:"name"`
	// +optional
	OutputArtifacts []string `json:"output_artifacts,omitempty"`
	Owner           string   `json:"owner"`
	Provider        string   `json:"provider"`
	// +optional
	RoleArn string `json:"role_arn,omitempty"`
	Version string `json:"version"`
}

type CodepipelineSpecStage struct {
	Action []CodepipelineSpecStage `json:"action"`
	Name   string                  `json:"name"`
}

type CodepipelineSpec struct {
	// +kubebuilder:validation:MaxItems=1
	ArtifactStore []CodepipelineSpec `json:"artifact_store"`
	Name          string             `json:"name"`
	RoleArn       string             `json:"role_arn"`
	// +kubebuilder:validation:MinItems=2
	Stage []CodepipelineSpec `json:"stage"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type CodepipelineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodepipelineList is a list of Codepipelines
type CodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Codepipeline CRD objects
	Items []Codepipeline `json:"items,omitempty"`
}
