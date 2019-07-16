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

type Ec2CapacityReservation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2CapacityReservationSpec   `json:"spec,omitempty"`
	Status            Ec2CapacityReservationStatus `json:"status,omitempty"`
}

type Ec2CapacityReservationSpec struct {
	AvailabilityZone string `json:"availability_zone"`
	// +optional
	EbsOptimized bool `json:"ebs_optimized,omitempty"`
	// +optional
	EndDate string `json:"end_date,omitempty"`
	// +optional
	EndDateType string `json:"end_date_type,omitempty"`
	// +optional
	EphemeralStorage bool `json:"ephemeral_storage,omitempty"`
	InstanceCount    int  `json:"instance_count"`
	// +optional
	InstanceMatchCriteria string `json:"instance_match_criteria,omitempty"`
	InstancePlatform      string `json:"instance_platform"`
	InstanceType          string `json:"instance_type"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Tenancy string `json:"tenancy,omitempty"`
}

type Ec2CapacityReservationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2CapacityReservationList is a list of Ec2CapacityReservations
type Ec2CapacityReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2CapacityReservation CRD objects
	Items []Ec2CapacityReservation `json:"items,omitempty"`
}
