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

type NodeStatusSpec struct {
	StatusDown int `json:"status_down"`
	StatusUp   int `json:"status_up"`
}

type LinodeNodebalancerConfigSpec struct {
	CheckAttempts  int               `json:"check_attempts"`
	Algorithm      string            `json:"algorithm"`
	SslFingerprint string            `json:"ssl_fingerprint"`
	SslCert        string            `json:"ssl_cert"`
	NodeStatus     map[string]string `json:"node_status"`
	NodebalancerId int               `json:"nodebalancer_id"`
	Port           int               `json:"port"`
	CheckPath      string            `json:"check_path"`
	CheckBody      string            `json:"check_body"`
	SslKey         string            `json:"ssl_key"`
	CheckTimeout   int               `json:"check_timeout"`
	Stickiness     string            `json:"stickiness"`
	CheckPassive   bool              `json:"check_passive"`
	SslCommonname  string            `json:"ssl_commonname"`
	Protocol       string            `json:"protocol"`
	CheckInterval  int               `json:"check_interval"`
	Check          string            `json:"check"`
	CipherSuite    string            `json:"cipher_suite"`
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