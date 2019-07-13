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

type AwsCloud9EnvironmentEc2 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloud9EnvironmentEc2Spec   `json:"spec,omitempty"`
	Status            AwsCloud9EnvironmentEc2Status `json:"status,omitempty"`
}

type AwsCloud9EnvironmentEc2Spec struct {
	AutomaticStopTimeMinutes int    `json:"automatic_stop_time_minutes"`
	Description              string `json:"description"`
	OwnerArn                 string `json:"owner_arn"`
	SubnetId                 string `json:"subnet_id"`
	Arn                      string `json:"arn"`
	Type                     string `json:"type"`
	Name                     string `json:"name"`
	InstanceType             string `json:"instance_type"`
}



type AwsCloud9EnvironmentEc2Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloud9EnvironmentEc2List is a list of AwsCloud9EnvironmentEc2s
type AwsCloud9EnvironmentEc2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloud9EnvironmentEc2 CRD objects
	Items []AwsCloud9EnvironmentEc2 `json:"items,omitempty"`
}