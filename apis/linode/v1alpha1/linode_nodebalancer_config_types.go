package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LinodeNodebalancerConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeNodebalancerConfigSpec   `json:"spec,omitempty"`
	Status            LinodeNodebalancerConfigStatus `json:"status,omitempty"`
}

type LinodeNodebalancerConfigSpecNodeStatus struct {
	StatusUp   int `json:"status_up"`
	StatusDown int `json:"status_down"`
}

type LinodeNodebalancerConfigSpec struct {
	Algorithm      string            `json:"algorithm"`
	Check          string            `json:"check"`
	CheckBody      string            `json:"check_body"`
	CipherSuite    string            `json:"cipher_suite"`
	SslFingerprint string            `json:"ssl_fingerprint"`
	SslCert        string            `json:"ssl_cert"`
	SslKey         string            `json:"ssl_key"`
	Port           int               `json:"port"`
	CheckInterval  int               `json:"check_interval"`
	CheckTimeout   int               `json:"check_timeout"`
	Stickiness     string            `json:"stickiness"`
	CheckPath      string            `json:"check_path"`
	Protocol       string            `json:"protocol"`
	CheckAttempts  int               `json:"check_attempts"`
	CheckPassive   bool              `json:"check_passive"`
	SslCommonname  string            `json:"ssl_commonname"`
	NodeStatus     map[string]string `json:"node_status"`
	NodebalancerId int               `json:"nodebalancer_id"`
}

type LinodeNodebalancerConfigStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinodeNodebalancerConfigList is a list of LinodeNodebalancerConfigs
type LinodeNodebalancerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeNodebalancerConfig CRD objects
	Items []LinodeNodebalancerConfig `json:"items,omitempty"`
}
