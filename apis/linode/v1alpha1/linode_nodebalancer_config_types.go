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
	CheckPath      string            `json:"check_path"`
	SslFingerprint string            `json:"ssl_fingerprint"`
	NodebalancerId int               `json:"nodebalancer_id"`
	Check          string            `json:"check"`
	CheckBody      string            `json:"check_body"`
	SslCommonname  string            `json:"ssl_commonname"`
	SslKey         string            `json:"ssl_key"`
	Protocol       string            `json:"protocol"`
	CheckInterval  int               `json:"check_interval"`
	CheckTimeout   int               `json:"check_timeout"`
	CheckAttempts  int               `json:"check_attempts"`
	Algorithm      string            `json:"algorithm"`
	CheckPassive   bool              `json:"check_passive"`
	Port           int               `json:"port"`
	Stickiness     string            `json:"stickiness"`
	CipherSuite    string            `json:"cipher_suite"`
	SslCert        string            `json:"ssl_cert"`
	NodeStatus     map[string]string `json:"node_status"`
}

type LinodeNodebalancerConfigStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeNodebalancerConfigList is a list of LinodeNodebalancerConfigs
type LinodeNodebalancerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeNodebalancerConfig CRD objects
	Items []LinodeNodebalancerConfig `json:"items,omitempty"`
}
