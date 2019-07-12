package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleMonitoringUptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleMonitoringUptimeCheckConfigSpec   `json:"spec,omitempty"`
	Status            GoogleMonitoringUptimeCheckConfigStatus `json:"status,omitempty"`
}

type GoogleMonitoringUptimeCheckConfigSpecInternalCheckers struct {
	DisplayName   string `json:"display_name"`
	GcpZone       string `json:"gcp_zone"`
	Name          string `json:"name"`
	Network       string `json:"network"`
	PeerProjectId string `json:"peer_project_id"`
}

type GoogleMonitoringUptimeCheckConfigSpecMonitoredResource struct {
	Labels map[string]string `json:"labels"`
	Type   string            `json:"type"`
}

type GoogleMonitoringUptimeCheckConfigSpecResourceGroup struct {
	GroupId      string `json:"group_id"`
	ResourceType string `json:"resource_type"`
}

type GoogleMonitoringUptimeCheckConfigSpecTcpCheck struct {
	Port int `json:"port"`
}

type GoogleMonitoringUptimeCheckConfigSpecContentMatchers struct {
	Content string `json:"content"`
}

type GoogleMonitoringUptimeCheckConfigSpecHttpCheckAuthInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GoogleMonitoringUptimeCheckConfigSpecHttpCheck struct {
	Port        int                                              `json:"port"`
	UseSsl      bool                                             `json:"use_ssl"`
	AuthInfo    []GoogleMonitoringUptimeCheckConfigSpecHttpCheck `json:"auth_info"`
	Headers     map[string]string                                `json:"headers"`
	MaskHeaders bool                                             `json:"mask_headers"`
	Path        string                                           `json:"path"`
}

type GoogleMonitoringUptimeCheckConfigSpec struct {
	Name              string                                  `json:"name"`
	Project           string                                  `json:"project"`
	DisplayName       string                                  `json:"display_name"`
	InternalCheckers  []GoogleMonitoringUptimeCheckConfigSpec `json:"internal_checkers"`
	MonitoredResource []GoogleMonitoringUptimeCheckConfigSpec `json:"monitored_resource"`
	ResourceGroup     []GoogleMonitoringUptimeCheckConfigSpec `json:"resource_group"`
	Period            string                                  `json:"period"`
	SelectedRegions   []string                                `json:"selected_regions"`
	TcpCheck          []GoogleMonitoringUptimeCheckConfigSpec `json:"tcp_check"`
	Timeout           string                                  `json:"timeout"`
	ContentMatchers   []GoogleMonitoringUptimeCheckConfigSpec `json:"content_matchers"`
	HttpCheck         []GoogleMonitoringUptimeCheckConfigSpec `json:"http_check"`
	IsInternal        bool                                    `json:"is_internal"`
}



type GoogleMonitoringUptimeCheckConfigStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleMonitoringUptimeCheckConfigList is a list of GoogleMonitoringUptimeCheckConfigs
type GoogleMonitoringUptimeCheckConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleMonitoringUptimeCheckConfig CRD objects
	Items []GoogleMonitoringUptimeCheckConfig `json:"items,omitempty"`
}