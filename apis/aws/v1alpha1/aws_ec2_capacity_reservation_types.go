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

type AwsEc2CapacityReservation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2CapacityReservationSpec   `json:"spec,omitempty"`
	Status            AwsEc2CapacityReservationStatus `json:"status,omitempty"`
}

type AwsEc2CapacityReservationSpec struct {
	EndDateType           string            `json:"end_date_type"`
	EphemeralStorage      bool              `json:"ephemeral_storage"`
	InstanceCount         int               `json:"instance_count"`
	InstanceMatchCriteria string            `json:"instance_match_criteria"`
	InstancePlatform      string            `json:"instance_platform"`
	InstanceType          string            `json:"instance_type"`
	AvailabilityZone      string            `json:"availability_zone"`
	EbsOptimized          bool              `json:"ebs_optimized"`
	Tenancy               string            `json:"tenancy"`
	EndDate               string            `json:"end_date"`
	Tags                  map[string]string `json:"tags"`
}



type AwsEc2CapacityReservationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2CapacityReservationList is a list of AwsEc2CapacityReservations
type AwsEc2CapacityReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2CapacityReservation CRD objects
	Items []AwsEc2CapacityReservation `json:"items,omitempty"`
}