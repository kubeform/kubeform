package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStackSetInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudformationStackSetInstanceSpec   `json:"spec,omitempty"`
	Status            AwsCloudformationStackSetInstanceStatus `json:"status,omitempty"`
}

type AwsCloudformationStackSetInstanceSpec struct {
	StackSetName       string            `json:"stack_set_name"`
	AccountId          string            `json:"account_id"`
	ParameterOverrides map[string]string `json:"parameter_overrides"`
	RetainStack        bool              `json:"retain_stack"`
	Region             string            `json:"region"`
	StackId            string            `json:"stack_id"`
}

type AwsCloudformationStackSetInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudformationStackSetInstanceList is a list of AwsCloudformationStackSetInstances
type AwsCloudformationStackSetInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudformationStackSetInstance CRD objects
	Items []AwsCloudformationStackSetInstance `json:"items,omitempty"`
}
