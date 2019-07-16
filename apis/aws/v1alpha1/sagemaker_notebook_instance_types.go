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

type SagemakerNotebookInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerNotebookInstanceSpec   `json:"spec,omitempty"`
	Status            SagemakerNotebookInstanceStatus `json:"status,omitempty"`
}

type SagemakerNotebookInstanceSpec struct {
	InstanceType string `json:"instance_type"`
	// +optional
	KmsKeyId string `json:"kms_key_id,omitempty"`
	// +optional
	LifecycleConfigName string `json:"lifecycle_config_name,omitempty"`
	Name                string `json:"name"`
	RoleArn             string `json:"role_arn"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SagemakerNotebookInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SagemakerNotebookInstanceList is a list of SagemakerNotebookInstances
type SagemakerNotebookInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SagemakerNotebookInstance CRD objects
	Items []SagemakerNotebookInstance `json:"items,omitempty"`
}
