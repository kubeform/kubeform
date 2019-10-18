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
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MonitoringUptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringUptimeCheckConfigSpec   `json:"spec,omitempty"`
	Status            MonitoringUptimeCheckConfigStatus `json:"status,omitempty"`
}

type MonitoringUptimeCheckConfigSpecContentMatchers struct {
	// +optional
	Content string `json:"content,omitempty" tf:"content,omitempty"`
}

type MonitoringUptimeCheckConfigSpecHttpCheckAuthInfo struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type MonitoringUptimeCheckConfigSpecHttpCheck struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthInfo []MonitoringUptimeCheckConfigSpecHttpCheckAuthInfo `json:"authInfo,omitempty" tf:"auth_info,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty" tf:"headers,omitempty"`
	// +optional
	MaskHeaders bool `json:"maskHeaders,omitempty" tf:"mask_headers,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	UseSSL bool `json:"useSSL,omitempty" tf:"use_ssl,omitempty"`
}

type MonitoringUptimeCheckConfigSpecInternalCheckers struct {
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	GcpZone string `json:"gcpZone,omitempty" tf:"gcp_zone,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	PeerProjectID string `json:"peerProjectID,omitempty" tf:"peer_project_id,omitempty"`
}

type MonitoringUptimeCheckConfigSpecMonitoredResource struct {
	Labels map[string]string `json:"labels" tf:"labels"`
	Type   string            `json:"type" tf:"type"`
}

type MonitoringUptimeCheckConfigSpecResourceGroup struct {
	// +optional
	GroupID string `json:"groupID,omitempty" tf:"group_id,omitempty"`
	// +optional
	ResourceType string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type MonitoringUptimeCheckConfigSpecTcpCheck struct {
	Port int `json:"port" tf:"port"`
}

type MonitoringUptimeCheckConfigSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	ContentMatchers []MonitoringUptimeCheckConfigSpecContentMatchers `json:"contentMatchers,omitempty" tf:"content_matchers,omitempty"`
	DisplayName     string                                           `json:"displayName" tf:"display_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpCheck []MonitoringUptimeCheckConfigSpecHttpCheck `json:"httpCheck,omitempty" tf:"http_check,omitempty"`
	// +optional
	InternalCheckers []MonitoringUptimeCheckConfigSpecInternalCheckers `json:"internalCheckers,omitempty" tf:"internal_checkers,omitempty"`
	// +optional
	IsInternal bool `json:"isInternal,omitempty" tf:"is_internal,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MonitoredResource []MonitoringUptimeCheckConfigSpecMonitoredResource `json:"monitoredResource,omitempty" tf:"monitored_resource,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Period string `json:"period,omitempty" tf:"period,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ResourceGroup []MonitoringUptimeCheckConfigSpecResourceGroup `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`
	// +optional
	SelectedRegions []string `json:"selectedRegions,omitempty" tf:"selected_regions,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TcpCheck []MonitoringUptimeCheckConfigSpecTcpCheck `json:"tcpCheck,omitempty" tf:"tcp_check,omitempty"`
	Timeout  string                                    `json:"timeout" tf:"timeout"`
}

type MonitoringUptimeCheckConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitoringUptimeCheckConfigSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitoringUptimeCheckConfigList is a list of MonitoringUptimeCheckConfigs
type MonitoringUptimeCheckConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitoringUptimeCheckConfig CRD objects
	Items []MonitoringUptimeCheckConfig `json:"items,omitempty"`
}
