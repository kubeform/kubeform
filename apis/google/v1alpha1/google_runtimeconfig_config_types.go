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

type GoogleRuntimeconfigConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleRuntimeconfigConfigSpec   `json:"spec,omitempty"`
	Status            GoogleRuntimeconfigConfigStatus `json:"status,omitempty"`
}

type GoogleRuntimeconfigConfigSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Project     string `json:"project"`
}



type GoogleRuntimeconfigConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleRuntimeconfigConfigList is a list of GoogleRuntimeconfigConfigs
type GoogleRuntimeconfigConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleRuntimeconfigConfig CRD objects
	Items []GoogleRuntimeconfigConfig `json:"items,omitempty"`
}