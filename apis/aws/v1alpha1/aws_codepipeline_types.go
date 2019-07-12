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

type AwsCodepipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodepipelineSpec   `json:"spec,omitempty"`
	Status            AwsCodepipelineStatus `json:"status,omitempty"`
}

type AwsCodepipelineSpecArtifactStoreEncryptionKey struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type AwsCodepipelineSpecArtifactStore struct {
	Location      string                             `json:"location"`
	Type          string                             `json:"type"`
	EncryptionKey []AwsCodepipelineSpecArtifactStore `json:"encryption_key"`
}

type AwsCodepipelineSpecStageAction struct {
	OutputArtifacts []string          `json:"output_artifacts"`
	Name            string            `json:"name"`
	Category        string            `json:"category"`
	Version         string            `json:"version"`
	Provider        string            `json:"provider"`
	InputArtifacts  []string          `json:"input_artifacts"`
	RoleArn         string            `json:"role_arn"`
	RunOrder        int               `json:"run_order"`
	Configuration   map[string]string `json:"configuration"`
	Owner           string            `json:"owner"`
}

type AwsCodepipelineSpecStage struct {
	Name   string                     `json:"name"`
	Action []AwsCodepipelineSpecStage `json:"action"`
}

type AwsCodepipelineSpec struct {
	Arn           string                `json:"arn"`
	Name          string                `json:"name"`
	RoleArn       string                `json:"role_arn"`
	ArtifactStore []AwsCodepipelineSpec `json:"artifact_store"`
	Stage         []AwsCodepipelineSpec `json:"stage"`
	Tags          map[string]string     `json:"tags"`
}

type AwsCodepipelineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCodepipelineList is a list of AwsCodepipelines
type AwsCodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodepipeline CRD objects
	Items []AwsCodepipeline `json:"items,omitempty"`
}
