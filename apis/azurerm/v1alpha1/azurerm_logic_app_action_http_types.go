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

type AzurermLogicAppActionHttp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogicAppActionHttpSpec   `json:"spec,omitempty"`
	Status            AzurermLogicAppActionHttpStatus `json:"status,omitempty"`
}

type AzurermLogicAppActionHttpSpec struct {
	Name       string            `json:"name"`
	LogicAppId string            `json:"logic_app_id"`
	Method     string            `json:"method"`
	Uri        string            `json:"uri"`
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
}

type AzurermLogicAppActionHttpStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLogicAppActionHttpList is a list of AzurermLogicAppActionHttps
type AzurermLogicAppActionHttpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogicAppActionHttp CRD objects
	Items []AzurermLogicAppActionHttp `json:"items,omitempty"`
}
