package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDatasyncLocationNfs struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDatasyncLocationNfsSpec   `json:"spec,omitempty"`
	Status            AwsDatasyncLocationNfsStatus `json:"status,omitempty"`
}

type AwsDatasyncLocationNfsSpecOnPremConfig struct {
	AgentArns []string `json:"agent_arns"`
}

type AwsDatasyncLocationNfsSpec struct {
	ServerHostname string                       `json:"server_hostname"`
	Subdirectory   string                       `json:"subdirectory"`
	Tags           map[string]string            `json:"tags"`
	Uri            string                       `json:"uri"`
	Arn            string                       `json:"arn"`
	OnPremConfig   []AwsDatasyncLocationNfsSpec `json:"on_prem_config"`
}

type AwsDatasyncLocationNfsStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDatasyncLocationNfsList is a list of AwsDatasyncLocationNfss
type AwsDatasyncLocationNfsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDatasyncLocationNfs CRD objects
	Items []AwsDatasyncLocationNfs `json:"items,omitempty"`
}
