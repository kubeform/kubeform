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

type AwsMqBroker struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMqBrokerSpec   `json:"spec,omitempty"`
	Status            AwsMqBrokerStatus `json:"status,omitempty"`
}

type AwsMqBrokerSpecConfiguration struct {
	Id       string `json:"id"`
	Revision int    `json:"revision"`
}

type AwsMqBrokerSpecMaintenanceWindowStartTime struct {
	DayOfWeek string `json:"day_of_week"`
	TimeOfDay string `json:"time_of_day"`
	TimeZone  string `json:"time_zone"`
}

type AwsMqBrokerSpecUser struct {
	ConsoleAccess bool     `json:"console_access"`
	Groups        []string `json:"groups"`
	Password      string   `json:"password"`
	Username      string   `json:"username"`
}

type AwsMqBrokerSpecInstances struct {
	ConsoleUrl string   `json:"console_url"`
	IpAddress  string   `json:"ip_address"`
	Endpoints  []string `json:"endpoints"`
}

type AwsMqBrokerSpecLogs struct {
	General bool `json:"general"`
	Audit   bool `json:"audit"`
}

type AwsMqBrokerSpec struct {
	EngineVersion              string            `json:"engine_version"`
	BrokerName                 string            `json:"broker_name"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	Configuration              []AwsMqBrokerSpec `json:"configuration"`
	EngineType                 string            `json:"engine_type"`
	HostInstanceType           string            `json:"host_instance_type"`
	MaintenanceWindowStartTime []AwsMqBrokerSpec `json:"maintenance_window_start_time"`
	Arn                        string            `json:"arn"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	SubnetIds                  []string          `json:"subnet_ids"`
	User                       []AwsMqBrokerSpec `json:"user"`
	Instances                  []AwsMqBrokerSpec `json:"instances"`
	SecurityGroups             []string          `json:"security_groups"`
	Logs                       []AwsMqBrokerSpec `json:"logs"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	Tags                       map[string]string `json:"tags"`
	DeploymentMode             string            `json:"deployment_mode"`
}

type AwsMqBrokerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsMqBrokerList is a list of AwsMqBrokers
type AwsMqBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMqBroker CRD objects
	Items []AwsMqBroker `json:"items,omitempty"`
}
