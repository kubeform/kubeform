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

type GoogleComputeFirewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeFirewallSpec   `json:"spec,omitempty"`
	Status            GoogleComputeFirewallStatus `json:"status,omitempty"`
}

type GoogleComputeFirewallSpecAllow struct {
	Protocol string   `json:"protocol"`
	Ports    []string `json:"ports"`
}

type GoogleComputeFirewallSpecDeny struct {
	Protocol string   `json:"protocol"`
	Ports    []string `json:"ports"`
}

type GoogleComputeFirewallSpec struct {
	Name                  string                      `json:"name"`
	Disabled              bool                        `json:"disabled"`
	SourceServiceAccounts []string                    `json:"source_service_accounts"`
	EnableLogging         bool                        `json:"enable_logging"`
	Priority              int                         `json:"priority"`
	CreationTimestamp     string                      `json:"creation_timestamp"`
	SelfLink              string                      `json:"self_link"`
	Network               string                      `json:"network"`
	Allow                 []GoogleComputeFirewallSpec `json:"allow"`
	Deny                  []GoogleComputeFirewallSpec `json:"deny"`
	Description           string                      `json:"description"`
	SourceTags            []string                    `json:"source_tags"`
	Project               string                      `json:"project"`
	TargetTags            []string                    `json:"target_tags"`
	DestinationRanges     []string                    `json:"destination_ranges"`
	Direction             string                      `json:"direction"`
	SourceRanges          []string                    `json:"source_ranges"`
	TargetServiceAccounts []string                    `json:"target_service_accounts"`
}

type GoogleComputeFirewallStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeFirewallList is a list of GoogleComputeFirewalls
type GoogleComputeFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeFirewall CRD objects
	Items []GoogleComputeFirewall `json:"items,omitempty"`
}
