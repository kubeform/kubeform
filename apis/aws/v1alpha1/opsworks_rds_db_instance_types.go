package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type OpsworksRdsDbInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksRdsDbInstanceSpec   `json:"spec,omitempty"`
	Status            OpsworksRdsDbInstanceStatus `json:"status,omitempty"`
}

type OpsworksRdsDbInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	DbPassword       string `json:"-" sensitive:"true" tf:"db_password"`
	DbUser           string `json:"dbUser" tf:"db_user"`
	RdsDbInstanceArn string `json:"rdsDbInstanceArn" tf:"rds_db_instance_arn"`
	StackID          string `json:"stackID" tf:"stack_id"`
}

type OpsworksRdsDbInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksRdsDbInstanceList is a list of OpsworksRdsDbInstances
type OpsworksRdsDbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksRdsDbInstance CRD objects
	Items []OpsworksRdsDbInstance `json:"items,omitempty"`
}
