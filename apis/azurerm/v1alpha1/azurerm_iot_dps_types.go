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

type AzurermIotDps struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermIotDpsSpec   `json:"spec,omitempty"`
	Status            AzurermIotDpsStatus `json:"status,omitempty"`
}

type AzurermIotDpsSpecSku struct {
	Tier     string `json:"tier"`
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}

type AzurermIotDpsSpec struct {
	ResourceGroupName string              `json:"resource_group_name"`
	Location          string              `json:"location"`
	Sku               []AzurermIotDpsSpec `json:"sku"`
	Tags              map[string]string   `json:"tags"`
	Name              string              `json:"name"`
}

type AzurermIotDpsStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermIotDpsList is a list of AzurermIotDpss
type AzurermIotDpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermIotDps CRD objects
	Items []AzurermIotDps `json:"items,omitempty"`
}
