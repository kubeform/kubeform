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

type AwsGlueConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueConnectionSpec   `json:"spec,omitempty"`
	Status            AwsGlueConnectionStatus `json:"status,omitempty"`
}

type AwsGlueConnectionSpecPhysicalConnectionRequirements struct {
	AvailabilityZone    string   `json:"availability_zone"`
	SecurityGroupIdList []string `json:"security_group_id_list"`
	SubnetId            string   `json:"subnet_id"`
}

type AwsGlueConnectionSpec struct {
	CatalogId                      string                  `json:"catalog_id"`
	ConnectionProperties           map[string]string       `json:"connection_properties"`
	ConnectionType                 string                  `json:"connection_type"`
	Description                    string                  `json:"description"`
	MatchCriteria                  []string                `json:"match_criteria"`
	Name                           string                  `json:"name"`
	PhysicalConnectionRequirements []AwsGlueConnectionSpec `json:"physical_connection_requirements"`
}



type AwsGlueConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlueConnectionList is a list of AwsGlueConnections
type AwsGlueConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueConnection CRD objects
	Items []AwsGlueConnection `json:"items,omitempty"`
}