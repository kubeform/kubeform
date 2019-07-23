package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApiID      string `json:"apiID" tf:"api_id"`
	DataSource string `json:"dataSource" tf:"data_source"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	FunctionVersion         string `json:"functionVersion,omitempty" tf:"function_version,omitempty"`
	Name                    string `json:"name" tf:"name"`
	RequestMappingTemplate  string `json:"requestMappingTemplate" tf:"request_mapping_template"`
	ResponseMappingTemplate string `json:"responseMappingTemplate" tf:"response_mapping_template"`
}

type AppsyncFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
