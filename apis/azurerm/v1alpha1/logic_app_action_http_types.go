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

type LogicAppActionHttp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppActionHttpSpec   `json:"spec,omitempty"`
	Status            LogicAppActionHttpStatus `json:"status,omitempty"`
}

type LogicAppActionHttpSpec struct {
	// +optional
	Body string `json:"body,omitempty"`
	// +optional
	Headers    map[string]string `json:"headers,omitempty"`
	LogicAppId string            `json:"logic_app_id"`
	Method     string            `json:"method"`
	Name       string            `json:"name"`
	Uri        string            `json:"uri"`
}

type LogicAppActionHttpStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppActionHttpList is a list of LogicAppActionHttps
type LogicAppActionHttpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppActionHttp CRD objects
	Items []LogicAppActionHttp `json:"items,omitempty"`
}
