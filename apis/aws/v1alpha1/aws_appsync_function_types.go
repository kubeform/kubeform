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

type AwsAppsyncFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppsyncFunctionSpec   `json:"spec,omitempty"`
	Status            AwsAppsyncFunctionStatus `json:"status,omitempty"`
}

type AwsAppsyncFunctionSpec struct {
	ApiId                   string `json:"api_id"`
	DataSource              string `json:"data_source"`
	RequestMappingTemplate  string `json:"request_mapping_template"`
	Description             string `json:"description"`
	FunctionVersion         string `json:"function_version"`
	Name                    string `json:"name"`
	ResponseMappingTemplate string `json:"response_mapping_template"`
	Arn                     string `json:"arn"`
	FunctionId              string `json:"function_id"`
}



type AwsAppsyncFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppsyncFunctionList is a list of AwsAppsyncFunctions
type AwsAppsyncFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppsyncFunction CRD objects
	Items []AwsAppsyncFunction `json:"items,omitempty"`
}