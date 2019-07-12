package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	NameServers     []string             `json:"name_servers"`
	Tags            map[string]string    `json:"tags"`
	Name            string               `json:"name"`
	Comment         string               `json:"comment"`
	VpcId           string               `json:"vpc_id"`
	VpcRegion       string               `json:"vpc_region"`
	Vpc             []AwsRoute53ZoneSpec `json:"vpc"`
	ZoneId          string               `json:"zone_id"`
	DelegationSetId string               `json:"delegation_set_id"`
	ForceDestroy    bool                 `json:"force_destroy"`
}

type AwsRoute53ZoneStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53ZoneList is a list of AwsRoute53Zones
type AwsRoute53ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53Zone CRD objects
	Items []AwsRoute53Zone `json:"items,omitempty"`
}
