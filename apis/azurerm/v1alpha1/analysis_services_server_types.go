package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AnalysisServicesServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalysisServicesServerSpec   `json:"spec,omitempty"`
	Status            AnalysisServicesServerStatus `json:"status,omitempty"`
}

type AnalysisServicesServerSpecIpv4FirewallRule struct {
	Name       string `json:"name" tf:"name"`
	RangeEnd   string `json:"rangeEnd" tf:"range_end"`
	RangeStart string `json:"rangeStart" tf:"range_start"`
}

type AnalysisServicesServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdminUsers []string `json:"adminUsers,omitempty" tf:"admin_users,omitempty"`
	// +optional
	EnablePowerBiService bool `json:"enablePowerBiService,omitempty" tf:"enable_power_bi_service,omitempty"`
	// +optional
	Ipv4FirewallRule []AnalysisServicesServerSpecIpv4FirewallRule `json:"ipv4FirewallRule,omitempty" tf:"ipv4_firewall_rule,omitempty"`
	Location         string                                       `json:"location" tf:"location"`
	Name             string                                       `json:"name" tf:"name"`
	// +optional
	QuerypoolConnectionMode string `json:"querypoolConnectionMode,omitempty" tf:"querypool_connection_mode,omitempty"`
	ResourceGroupName       string `json:"resourceGroupName" tf:"resource_group_name"`
	Sku                     string `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AnalysisServicesServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AnalysisServicesServerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AnalysisServicesServerList is a list of AnalysisServicesServers
type AnalysisServicesServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AnalysisServicesServer CRD objects
	Items []AnalysisServicesServer `json:"items,omitempty"`
}
