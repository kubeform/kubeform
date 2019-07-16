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

type PinpointEmailChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointEmailChannelSpec   `json:"spec,omitempty"`
	Status            PinpointEmailChannelStatus `json:"status,omitempty"`
}

type PinpointEmailChannelSpec struct {
	ApplicationId string `json:"application_id"`
	// +optional
	Enabled     bool   `json:"enabled,omitempty"`
	FromAddress string `json:"from_address"`
	Identity    string `json:"identity"`
	RoleArn     string `json:"role_arn"`
}

type PinpointEmailChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointEmailChannelList is a list of PinpointEmailChannels
type PinpointEmailChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointEmailChannel CRD objects
	Items []PinpointEmailChannel `json:"items,omitempty"`
}
