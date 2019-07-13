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

type AwsRoute53Zone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53ZoneSpec   `json:"spec,omitempty"`
	Status            AwsRoute53ZoneStatus `json:"status,omitempty"`
}

type AwsRoute53ZoneSpecVpc struct {
	VpcId     string `json:"vpc_id"`
	VpcRegion string `json:"vpc_region"`
}

type AwsRoute53ZoneSpec struct {
	VpcId           string               `json:"vpc_id"`
	ZoneId          string               `json:"zone_id"`
	DelegationSetId string               `json:"delegation_set_id"`
	NameServers     []string             `json:"name_servers"`
	Tags            map[string]string    `json:"tags"`
	ForceDestroy    bool                 `json:"force_destroy"`
	Vpc             []AwsRoute53ZoneSpec `json:"vpc"`
	Comment         string               `json:"comment"`
	VpcRegion       string               `json:"vpc_region"`
	Name            string               `json:"name"`
}



type AwsRoute53ZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRoute53ZoneList is a list of AwsRoute53Zones
type AwsRoute53ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53Zone CRD objects
	Items []AwsRoute53Zone `json:"items,omitempty"`
}