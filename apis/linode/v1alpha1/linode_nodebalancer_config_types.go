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
	SslCert        string            `json:"ssl_cert"`
	CheckBody      string            `json:"check_body"`
	CheckPassive   bool              `json:"check_passive"`
	CipherSuite    string            `json:"cipher_suite"`
	SslFingerprint string            `json:"ssl_fingerprint"`
	Port           int               `json:"port"`
	CheckTimeout   int               `json:"check_timeout"`
	Check          string            `json:"check"`
	Stickiness     string            `json:"stickiness"`
	SslCommonname  string            `json:"ssl_commonname"`
	NodebalancerId int               `json:"nodebalancer_id"`
	CheckInterval  int               `json:"check_interval"`
	CheckAttempts  int               `json:"check_attempts"`
	NodeStatus     map[string]string `json:"node_status"`
	Protocol       string            `json:"protocol"`
	Algorithm      string            `json:"algorithm"`
	SslKey         string            `json:"ssl_key"`
}



type LinodeNodebalancerConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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