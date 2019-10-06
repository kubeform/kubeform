package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Cloud9EnvironmentEc2 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Cloud9EnvironmentEc2Spec   `json:"spec,omitempty"`
	Status            Cloud9EnvironmentEc2Status `json:"status,omitempty"`
}

type Cloud9EnvironmentEc2Spec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AutomaticStopTimeMinutes int `json:"automaticStopTimeMinutes,omitempty" tf:"automatic_stop_time_minutes,omitempty"`
	// +optional
	Description  string `json:"description,omitempty" tf:"description,omitempty"`
	InstanceType string `json:"instanceType" tf:"instance_type"`
	Name         string `json:"name" tf:"name"`
	// +optional
	OwnerArn string `json:"ownerArn,omitempty" tf:"owner_arn,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type Cloud9EnvironmentEc2Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *Cloud9EnvironmentEc2Spec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Cloud9EnvironmentEc2List is a list of Cloud9EnvironmentEc2s
type Cloud9EnvironmentEc2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cloud9EnvironmentEc2 CRD objects
	Items []Cloud9EnvironmentEc2 `json:"items,omitempty"`
}
