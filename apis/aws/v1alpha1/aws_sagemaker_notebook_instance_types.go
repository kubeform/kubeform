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

type AwsSagemakerNotebookInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSagemakerNotebookInstanceSpec   `json:"spec,omitempty"`
	Status            AwsSagemakerNotebookInstanceStatus `json:"status,omitempty"`
}

type AwsSagemakerNotebookInstanceSpec struct {
	RoleArn             string            `json:"role_arn"`
	InstanceType        string            `json:"instance_type"`
	SecurityGroups      []string          `json:"security_groups"`
	LifecycleConfigName string            `json:"lifecycle_config_name"`
	Tags                map[string]string `json:"tags"`
	Arn                 string            `json:"arn"`
	Name                string            `json:"name"`
	SubnetId            string            `json:"subnet_id"`
	KmsKeyId            string            `json:"kms_key_id"`
}



type AwsSagemakerNotebookInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSagemakerNotebookInstanceList is a list of AwsSagemakerNotebookInstances
type AwsSagemakerNotebookInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSagemakerNotebookInstance CRD objects
	Items []AwsSagemakerNotebookInstance `json:"items,omitempty"`
}