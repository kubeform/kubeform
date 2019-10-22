package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type SagemakerNotebookInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerNotebookInstanceSpec   `json:"spec,omitempty"`
	Status            SagemakerNotebookInstanceStatus `json:"status,omitempty"`
}

type SagemakerNotebookInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn          string `json:"arn,omitempty" tf:"arn,omitempty"`
	InstanceType string `json:"instanceType" tf:"instance_type"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	LifecycleConfigName string `json:"lifecycleConfigName,omitempty" tf:"lifecycle_config_name,omitempty"`
	Name                string `json:"name" tf:"name"`
	RoleArn             string `json:"roleArn" tf:"role_arn"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SagemakerNotebookInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SagemakerNotebookInstanceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
