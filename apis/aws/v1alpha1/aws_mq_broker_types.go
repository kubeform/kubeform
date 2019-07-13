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

type AwsMqBrokerSpecUser struct {
	ConsoleAccess bool     `json:"console_access"`
	Groups        []string `json:"groups"`
	Password      string   `json:"password"`
	Username      string   `json:"username"`
}

type AwsMqBrokerSpecLogs struct {
	General bool `json:"general"`
	Audit   bool `json:"audit"`
}

type AwsMqBrokerSpecMaintenanceWindowStartTime struct {
	TimeOfDay string `json:"time_of_day"`
	TimeZone  string `json:"time_zone"`
	DayOfWeek string `json:"day_of_week"`
}

type AwsMqBrokerSpecInstances struct {
	Endpoints  []string `json:"endpoints"`
	ConsoleUrl string   `json:"console_url"`
	IpAddress  string   `json:"ip_address"`
}

type AwsMqBrokerSpecConfiguration struct {
	Id       string `json:"id"`
	Revision int    `json:"revision"`
}

type AwsMqBrokerSpec struct {
	HostInstanceType           string            `json:"host_instance_type"`
	Tags                       map[string]string `json:"tags"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	SubnetIds                  []string          `json:"subnet_ids"`
	User                       []AwsMqBrokerSpec `json:"user"`
	Arn                        string            `json:"arn"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	DeploymentMode             string            `json:"deployment_mode"`
	EngineType                 string            `json:"engine_type"`
	Logs                       []AwsMqBrokerSpec `json:"logs"`
	MaintenanceWindowStartTime []AwsMqBrokerSpec `json:"maintenance_window_start_time"`
	SecurityGroups             []string          `json:"security_groups"`
	Instances                  []AwsMqBrokerSpec `json:"instances"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	BrokerName                 string            `json:"broker_name"`
	Configuration              []AwsMqBrokerSpec `json:"configuration"`
	EngineVersion              string            `json:"engine_version"`
}



type AwsMqBrokerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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