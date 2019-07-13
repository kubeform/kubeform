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

type GoogleMonitoringUptimeCheckConfigSpecContentMatchers struct {
	Content string `json:"content"`
}

type GoogleMonitoringUptimeCheckConfigSpecInternalCheckers struct {
	DisplayName   string `json:"display_name"`
	GcpZone       string `json:"gcp_zone"`
	Name          string `json:"name"`
	Network       string `json:"network"`
	PeerProjectId string `json:"peer_project_id"`
}

type GoogleMonitoringUptimeCheckConfigSpecResourceGroup struct {
	GroupId      string `json:"group_id"`
	ResourceType string `json:"resource_type"`
}

type GoogleMonitoringUptimeCheckConfigSpecHttpCheckAuthInfo struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type GoogleMonitoringUptimeCheckConfigSpecHttpCheck struct {
	MaskHeaders bool                                             `json:"mask_headers"`
	Path        string                                           `json:"path"`
	Port        int                                              `json:"port"`
	UseSsl      bool                                             `json:"use_ssl"`
	AuthInfo    []GoogleMonitoringUptimeCheckConfigSpecHttpCheck `json:"auth_info"`
	Headers     map[string]string                                `json:"headers"`
}

type GoogleMonitoringUptimeCheckConfigSpecMonitoredResource struct {
	Labels map[string]string `json:"labels"`
	Type   string            `json:"type"`
}

type GoogleMonitoringUptimeCheckConfigSpecTcpCheck struct {
	Port int `json:"port"`
}

type GoogleMonitoringUptimeCheckConfigSpec struct {
	Timeout           string                                  `json:"timeout"`
	ContentMatchers   []GoogleMonitoringUptimeCheckConfigSpec `json:"content_matchers"`
	InternalCheckers  []GoogleMonitoringUptimeCheckConfigSpec `json:"internal_checkers"`
	ResourceGroup     []GoogleMonitoringUptimeCheckConfigSpec `json:"resource_group"`
	SelectedRegions   []string                                `json:"selected_regions"`
	Name              string                                  `json:"name"`
	Project           string                                  `json:"project"`
	DisplayName       string                                  `json:"display_name"`
	HttpCheck         []GoogleMonitoringUptimeCheckConfigSpec `json:"http_check"`
	IsInternal        bool                                    `json:"is_internal"`
	MonitoredResource []GoogleMonitoringUptimeCheckConfigSpec `json:"monitored_resource"`
	Period            string                                  `json:"period"`
	TcpCheck          []GoogleMonitoringUptimeCheckConfigSpec `json:"tcp_check"`
}



type GoogleMonitoringUptimeCheckConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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