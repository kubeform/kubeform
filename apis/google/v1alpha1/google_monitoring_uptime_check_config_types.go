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

type GoogleMonitoringUptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleMonitoringUptimeCheckConfigSpec   `json:"spec,omitempty"`
	Status            GoogleMonitoringUptimeCheckConfigStatus `json:"status,omitempty"`
}

type GoogleMonitoringUptimeCheckConfigSpecResourceGroup struct {
	GroupId      string `json:"group_id"`
	ResourceType string `json:"resource_type"`
}

type GoogleMonitoringUptimeCheckConfigSpecTcpCheck struct {
	Port int `json:"port"`
}

type GoogleMonitoringUptimeCheckConfigSpecHttpCheckAuthInfo struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type GoogleMonitoringUptimeCheckConfigSpecHttpCheck struct {
	AuthInfo    []GoogleMonitoringUptimeCheckConfigSpecHttpCheck `json:"auth_info"`
	Headers     map[string]string                                `json:"headers"`
	MaskHeaders bool                                             `json:"mask_headers"`
	Path        string                                           `json:"path"`
	Port        int                                              `json:"port"`
	UseSsl      bool                                             `json:"use_ssl"`
}

type GoogleMonitoringUptimeCheckConfigSpecInternalCheckers struct {
	Network       string `json:"network"`
	PeerProjectId string `json:"peer_project_id"`
	DisplayName   string `json:"display_name"`
	GcpZone       string `json:"gcp_zone"`
	Name          string `json:"name"`
}

type GoogleMonitoringUptimeCheckConfigSpecMonitoredResource struct {
	Labels map[string]string `json:"labels"`
	Type   string            `json:"type"`
}

type GoogleMonitoringUptimeCheckConfigSpecContentMatchers struct {
	Content string `json:"content"`
}

type GoogleMonitoringUptimeCheckConfigSpec struct {
	ResourceGroup     []GoogleMonitoringUptimeCheckConfigSpec `json:"resource_group"`
	SelectedRegions   []string                                `json:"selected_regions"`
	TcpCheck          []GoogleMonitoringUptimeCheckConfigSpec `json:"tcp_check"`
	HttpCheck         []GoogleMonitoringUptimeCheckConfigSpec `json:"http_check"`
	InternalCheckers  []GoogleMonitoringUptimeCheckConfigSpec `json:"internal_checkers"`
	IsInternal        bool                                    `json:"is_internal"`
	MonitoredResource []GoogleMonitoringUptimeCheckConfigSpec `json:"monitored_resource"`
	Period            string                                  `json:"period"`
	Name              string                                  `json:"name"`
	DisplayName       string                                  `json:"display_name"`
	Timeout           string                                  `json:"timeout"`
	ContentMatchers   []GoogleMonitoringUptimeCheckConfigSpec `json:"content_matchers"`
	Project           string                                  `json:"project"`
}

type GoogleMonitoringUptimeCheckConfigStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleMonitoringUptimeCheckConfigList is a list of GoogleMonitoringUptimeCheckConfigs
type GoogleMonitoringUptimeCheckConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleMonitoringUptimeCheckConfig CRD objects
	Items []GoogleMonitoringUptimeCheckConfig `json:"items,omitempty"`
}
