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

type AwsSagemakerModel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSagemakerModelSpec   `json:"spec,omitempty"`
	Status            AwsSagemakerModelStatus `json:"status,omitempty"`
}

type AwsSagemakerModelSpecPrimaryContainer struct {
	Image             string            `json:"image"`
	ModelDataUrl      string            `json:"model_data_url"`
	Environment       map[string]string `json:"environment"`
	ContainerHostname string            `json:"container_hostname"`
}

type AwsSagemakerModelSpecVpcConfig struct {
	Subnets          []string `json:"subnets"`
	SecurityGroupIds []string `json:"security_group_ids"`
}

type AwsSagemakerModelSpecContainer struct {
	ContainerHostname string            `json:"container_hostname"`
	Image             string            `json:"image"`
	ModelDataUrl      string            `json:"model_data_url"`
	Environment       map[string]string `json:"environment"`
}

type AwsSagemakerModelSpec struct {
	PrimaryContainer       []AwsSagemakerModelSpec `json:"primary_container"`
	VpcConfig              []AwsSagemakerModelSpec `json:"vpc_config"`
	ExecutionRoleArn       string                  `json:"execution_role_arn"`
	EnableNetworkIsolation bool                    `json:"enable_network_isolation"`
	Container              []AwsSagemakerModelSpec `json:"container"`
	Tags                   map[string]string       `json:"tags"`
	Arn                    string                  `json:"arn"`
	Name                   string                  `json:"name"`
}



type AwsSagemakerModelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSagemakerModelList is a list of AwsSagemakerModels
type AwsSagemakerModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSagemakerModel CRD objects
	Items []AwsSagemakerModel `json:"items,omitempty"`
}