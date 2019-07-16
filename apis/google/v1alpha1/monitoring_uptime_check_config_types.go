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

type MonitoringUptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringUptimeCheckConfigSpec   `json:"spec,omitempty"`
	Status            MonitoringUptimeCheckConfigStatus `json:"status,omitempty"`
}

type MonitoringUptimeCheckConfigSpecContentMatchers struct {
	// +optional
	Content string `json:"content,omitempty"`
}

type MonitoringUptimeCheckConfigSpecHttpCheckAuthInfo struct {
	// +optional
	Password string `json:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty"`
}

type MonitoringUptimeCheckConfigSpecHttpCheck struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthInfo *[]MonitoringUptimeCheckConfigSpecHttpCheck `json:"auth_info,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty"`
	// +optional
	MaskHeaders bool `json:"mask_headers,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	UseSsl bool `json:"use_ssl,omitempty"`
}

type MonitoringUptimeCheckConfigSpecInternalCheckers struct {
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	// +optional
	GcpZone string `json:"gcp_zone,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Network string `json:"network,omitempty"`
	// +optional
	PeerProjectId string `json:"peer_project_id,omitempty"`
}

type MonitoringUptimeCheckConfigSpecMonitoredResource struct {
	Labels map[string]string `json:"labels"`
	Type   string            `json:"type"`
}

type MonitoringUptimeCheckConfigSpecResourceGroup struct {
	// +optional
	GroupId string `json:"group_id,omitempty"`
	// +optional
	ResourceType string `json:"resource_type,omitempty"`
}

type MonitoringUptimeCheckConfigSpecTcpCheck struct {
	Port int `json:"port"`
}

type MonitoringUptimeCheckConfigSpec struct {
	// +optional
	ContentMatchers *[]MonitoringUptimeCheckConfigSpec `json:"content_matchers,omitempty"`
	DisplayName     string                             `json:"display_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpCheck *[]MonitoringUptimeCheckConfigSpec `json:"http_check,omitempty"`
	// +optional
	InternalCheckers *[]MonitoringUptimeCheckConfigSpec `json:"internal_checkers,omitempty"`
	// +optional
	IsInternal bool `json:"is_internal,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MonitoredResource *[]MonitoringUptimeCheckConfigSpec `json:"monitored_resource,omitempty"`
	// +optional
	Period string `json:"period,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ResourceGroup *[]MonitoringUptimeCheckConfigSpec `json:"resource_group,omitempty"`
	// +optional
	SelectedRegions []string `json:"selected_regions,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TcpCheck *[]MonitoringUptimeCheckConfigSpec `json:"tcp_check,omitempty"`
	Timeout  string                             `json:"timeout"`
}

type MonitoringUptimeCheckConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
