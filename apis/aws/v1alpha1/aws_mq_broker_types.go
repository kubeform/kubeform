package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqBroker struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMqBrokerSpec   `json:"spec,omitempty"`
	Status            AwsMqBrokerStatus `json:"status,omitempty"`
}

type AwsMqBrokerSpecLogs struct {
	General bool `json:"general"`
	Audit   bool `json:"audit"`
}

type AwsMqBrokerSpecConfiguration struct {
	Revision int    `json:"revision"`
	Id       string `json:"id"`
}

type AwsMqBrokerSpecInstances struct {
	ConsoleUrl string   `json:"console_url"`
	IpAddress  string   `json:"ip_address"`
	Endpoints  []string `json:"endpoints"`
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

type AwsMqBrokerSpec struct {
	DeploymentMode             string            `json:"deployment_mode"`
	HostInstanceType           string            `json:"host_instance_type"`
	Logs                       []AwsMqBrokerSpec `json:"logs"`
	SecurityGroups             []string          `json:"security_groups"`
	SubnetIds                  []string          `json:"subnet_ids"`
	Configuration              []AwsMqBrokerSpec `json:"configuration"`
	EngineType                 string            `json:"engine_type"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	Tags                       map[string]string `json:"tags"`
	Instances                  []AwsMqBrokerSpec `json:"instances"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	BrokerName                 string            `json:"broker_name"`
	EngineVersion              string            `json:"engine_version"`
	MaintenanceWindowStartTime []AwsMqBrokerSpec `json:"maintenance_window_start_time"`
	User                       []AwsMqBrokerSpec `json:"user"`
	Arn                        string            `json:"arn"`
}



type AwsMqBrokerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqBrokerList is a list of AwsMqBrokers
type AwsMqBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMqBroker CRD objects
	Items []AwsMqBroker `json:"items,omitempty"`
}