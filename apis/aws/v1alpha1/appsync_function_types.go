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

type AppsyncFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncFunctionSpec   `json:"spec,omitempty"`
	Status            AppsyncFunctionStatus `json:"status,omitempty"`
}

type AppsyncFunctionSpec struct {
	ApiId      string `json:"api_id"`
	DataSource string `json:"data_source"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	FunctionVersion         string `json:"function_version,omitempty"`
	Name                    string `json:"name"`
	RequestMappingTemplate  string `json:"request_mapping_template"`
	ResponseMappingTemplate string `json:"response_mapping_template"`
}

type AppsyncFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppsyncFunctionList is a list of AppsyncFunctions
type AppsyncFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppsyncFunction CRD objects
	Items []AppsyncFunction `json:"items,omitempty"`
}
