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

type AwsDatasyncAgent struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDatasyncAgentSpec   `json:"spec,omitempty"`
	Status            AwsDatasyncAgentStatus `json:"status,omitempty"`
}

type AwsDatasyncAgentSpec struct {
	Arn           string            `json:"arn"`
	ActivationKey string            `json:"activation_key"`
	IpAddress     string            `json:"ip_address"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
}

type AwsDatasyncAgentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDatasyncAgentList is a list of AwsDatasyncAgents
type AwsDatasyncAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDatasyncAgent CRD objects
	Items []AwsDatasyncAgent `json:"items,omitempty"`
}
