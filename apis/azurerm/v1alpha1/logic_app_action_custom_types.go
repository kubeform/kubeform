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

type LogicAppActionCustom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppActionCustomSpec   `json:"spec,omitempty"`
	Status            LogicAppActionCustomStatus `json:"status,omitempty"`
}

type LogicAppActionCustomSpec struct {
	Body       string `json:"body"`
	LogicAppId string `json:"logic_app_id"`
	Name       string `json:"name"`
}

type LogicAppActionCustomStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppActionCustomList is a list of LogicAppActionCustoms
type LogicAppActionCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppActionCustom CRD objects
	Items []LogicAppActionCustom `json:"items,omitempty"`
}
