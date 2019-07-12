package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Configuration   map[string]string `json:"configuration"`
	Category        string            `json:"category"`
	Owner           string            `json:"owner"`
	Provider        string            `json:"provider"`
	Version         string            `json:"version"`
	OutputArtifacts []string          `json:"output_artifacts"`
	Name            string            `json:"name"`
	RoleArn         string            `json:"role_arn"`
	InputArtifacts  []string          `json:"input_artifacts"`
	RunOrder        int               `json:"run_order"`
}

type AwsCodepipelineSpecStage struct {
	Action []AwsCodepipelineSpecStage `json:"action"`
	Name   string                     `json:"name"`
}

type AwsCodepipelineSpec struct {
	Arn           string                `json:"arn"`
	Name          string                `json:"name"`
	RoleArn       string                `json:"role_arn"`
	ArtifactStore []AwsCodepipelineSpec `json:"artifact_store"`
	Stage         []AwsCodepipelineSpec `json:"stage"`
}

type AwsCodepipelineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodepipelineList is a list of AwsCodepipelines
type AwsCodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodepipeline CRD objects
	Items []AwsCodepipeline `json:"items,omitempty"`
}
