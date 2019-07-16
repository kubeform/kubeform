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

type SagemakerModel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerModelSpec   `json:"spec,omitempty"`
	Status            SagemakerModelStatus `json:"status,omitempty"`
}

type SagemakerModelSpecContainer struct {
	// +optional
	ContainerHostname string `json:"container_hostname,omitempty"`
	// +optional
	Environment map[string]string `json:"environment,omitempty"`
	Image       string            `json:"image"`
	// +optional
	ModelDataUrl string `json:"model_data_url,omitempty"`
}

type SagemakerModelSpecPrimaryContainer struct {
	// +optional
	ContainerHostname string `json:"container_hostname,omitempty"`
	// +optional
	Environment map[string]string `json:"environment,omitempty"`
	Image       string            `json:"image"`
	// +optional
	ModelDataUrl string `json:"model_data_url,omitempty"`
}

type SagemakerModelSpecVpcConfig struct {
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids"`
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets"`
}

type SagemakerModelSpec struct {
	// +optional
	Container *[]SagemakerModelSpec `json:"container,omitempty"`
	// +optional
	EnableNetworkIsolation bool   `json:"enable_network_isolation,omitempty"`
	ExecutionRoleArn       string `json:"execution_role_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PrimaryContainer *[]SagemakerModelSpec `json:"primary_container,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcConfig *[]SagemakerModelSpec `json:"vpc_config,omitempty"`
}

type SagemakerModelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SagemakerModelList is a list of SagemakerModels
type SagemakerModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SagemakerModel CRD objects
	Items []SagemakerModel `json:"items,omitempty"`
}
