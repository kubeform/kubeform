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

type ContainerRegistry struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerRegistrySpec   `json:"spec,omitempty"`
	Status            ContainerRegistryStatus `json:"status,omitempty"`
}

type ContainerRegistrySpecNetworkRuleSetIpRule struct {
	Action  string `json:"action" tf:"action"`
	IpRange string `json:"ipRange" tf:"ip_range"`
}

type ContainerRegistrySpecNetworkRuleSet struct {
	// +optional
	DefaultAction string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpRule []ContainerRegistrySpecNetworkRuleSetIpRule `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`
}

type ContainerRegistrySpecStorageAccount struct {
	AccessKey string `json:"-" sensitive:"true" tf:"access_key"`
	Name      string `json:"name" tf:"name"`
}

type ContainerRegistrySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AdminEnabled bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`
	// +optional
	AdminPassword string `json:"-" sensitive:"true" tf:"admin_password,omitempty"`
	// +optional
	AdminUsername string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	GeoreplicationLocations []string `json:"georeplicationLocations,omitempty" tf:"georeplication_locations,omitempty"`
	Location                string   `json:"location" tf:"location"`
	// +optional
	LoginServer string `json:"loginServer,omitempty" tf:"login_server,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkRuleSet    []ContainerRegistrySpecNetworkRuleSet `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`
	ResourceGroupName string                                `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	StorageAccount []ContainerRegistrySpecStorageAccount `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ContainerRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ContainerRegistrySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerRegistryList is a list of ContainerRegistrys
type ContainerRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerRegistry CRD objects
	Items []ContainerRegistry `json:"items,omitempty"`
}
