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

type GoogleStorageDefaultObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageDefaultObjectAccessControlSpec   `json:"spec,omitempty"`
	Status            GoogleStorageDefaultObjectAccessControlStatus `json:"status,omitempty"`
}

type GoogleStorageDefaultObjectAccessControlSpecProjectTeam struct {
	ProjectNumber string `json:"project_number"`
	Team          string `json:"team"`
}

type GoogleStorageDefaultObjectAccessControlSpec struct {
	ProjectTeam []GoogleStorageDefaultObjectAccessControlSpec `json:"project_team"`
	Bucket      string                                        `json:"bucket"`
	Entity      string                                        `json:"entity"`
	Role        string                                        `json:"role"`
	Domain      string                                        `json:"domain"`
	EntityId    string                                        `json:"entity_id"`
	Object      string                                        `json:"object"`
	Email       string                                        `json:"email"`
	Generation  int                                           `json:"generation"`
}

type GoogleStorageDefaultObjectAccessControlStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageDefaultObjectAccessControlList is a list of GoogleStorageDefaultObjectAccessControls
type GoogleStorageDefaultObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageDefaultObjectAccessControl CRD objects
	Items []GoogleStorageDefaultObjectAccessControl `json:"items,omitempty"`
}
