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

type GoogleComputeFirewallSpecDeny struct {
	Protocol string   `json:"protocol"`
	Ports    []string `json:"ports"`
}

type GoogleComputeFirewallSpecAllow struct {
	Protocol string   `json:"protocol"`
	Ports    []string `json:"ports"`
}

type GoogleComputeFirewallSpec struct {
	Description           string                      `json:"description"`
	Disabled              bool                        `json:"disabled"`
	EnableLogging         bool                        `json:"enable_logging"`
	Deny                  []GoogleComputeFirewallSpec `json:"deny"`
	Priority              int                         `json:"priority"`
	SourceTags            []string                    `json:"source_tags"`
	SelfLink              string                      `json:"self_link"`
	Name                  string                      `json:"name"`
	Network               string                      `json:"network"`
	DestinationRanges     []string                    `json:"destination_ranges"`
	Direction             string                      `json:"direction"`
	SourceServiceAccounts []string                    `json:"source_service_accounts"`
	TargetTags            []string                    `json:"target_tags"`
	Project               string                      `json:"project"`
	Allow                 []GoogleComputeFirewallSpec `json:"allow"`
	SourceRanges          []string                    `json:"source_ranges"`
	TargetServiceAccounts []string                    `json:"target_service_accounts"`
	CreationTimestamp     string                      `json:"creation_timestamp"`
}



type GoogleComputeFirewallStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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