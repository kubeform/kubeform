package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanProjectSpec   `json:"spec,omitempty"`
	Status            DigitaloceanProjectStatus `json:"status,omitempty"`
}

type DigitaloceanProjectSpec struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Environment string   `json:"environment"`
	OwnerUuid   string   `json:"owner_uuid"`
	Purpose     string   `json:"purpose"`
	OwnerId     int      `json:"owner_id"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Resources   []string `json:"resources"`
}

type DigitaloceanProjectStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanProjectList is a list of DigitaloceanProjects
type DigitaloceanProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanProject CRD objects
	Items []DigitaloceanProject `json:"items,omitempty"`
}
