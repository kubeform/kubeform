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

type AwsLambdaAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaAliasSpec   `json:"spec,omitempty"`
	Status            AwsLambdaAliasStatus `json:"status,omitempty"`
}

type AwsLambdaAliasSpecRoutingConfig struct {
	AdditionalVersionWeights map[string]float64 `json:"additional_version_weights"`
}

type AwsLambdaAliasSpec struct {
	InvokeArn       string               `json:"invoke_arn"`
	RoutingConfig   []AwsLambdaAliasSpec `json:"routing_config"`
	Description     string               `json:"description"`
	FunctionName    string               `json:"function_name"`
	FunctionVersion string               `json:"function_version"`
	Name            string               `json:"name"`
	Arn             string               `json:"arn"`
}

type AwsLambdaAliasStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLambdaAliasList is a list of AwsLambdaAliass
type AwsLambdaAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLambdaAlias CRD objects
	Items []AwsLambdaAlias `json:"items,omitempty"`
}
