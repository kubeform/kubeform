package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DatasyncLocationNfs struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncLocationNfsSpec   `json:"spec,omitempty"`
	Status            DatasyncLocationNfsStatus `json:"status,omitempty"`
}

type DatasyncLocationNfsSpecOnPremConfig struct {
	AgentArns []string `json:"agentArns" tf:"agent_arns"`
}

type DatasyncLocationNfsSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	OnPremConfig   []DatasyncLocationNfsSpecOnPremConfig `json:"onPremConfig" tf:"on_prem_config"`
	ServerHostname string                                `json:"serverHostname" tf:"server_hostname"`
	Subdirectory   string                                `json:"subdirectory" tf:"subdirectory"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Uri string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type DatasyncLocationNfsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DatasyncLocationNfsSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatasyncLocationNfsList is a list of DatasyncLocationNfss
type DatasyncLocationNfsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatasyncLocationNfs CRD objects
	Items []DatasyncLocationNfs `json:"items,omitempty"`
}
