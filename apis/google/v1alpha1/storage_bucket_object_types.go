/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type StorageBucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketObjectSpec   `json:"spec,omitempty"`
	Status            StorageBucketObjectStatus `json:"status,omitempty"`
}

type StorageBucketObjectSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	CacheControl string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`
	// +optional
	Content string `json:"content,omitempty" tf:"content,omitempty"`
	// +optional
	ContentDisposition string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`
	// +optional
	ContentEncoding string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`
	// +optional
	ContentLanguage string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`
	// +optional
	ContentType string `json:"contentType,omitempty" tf:"content_type,omitempty"`
	// +optional
	Crc32c string `json:"crc32c,omitempty" tf:"crc32c,omitempty"`
	// +optional
	DetectMd5hash string `json:"detectMd5hash,omitempty" tf:"detect_md5hash,omitempty"`
	// +optional
	Md5hash string `json:"md5hash,omitempty" tf:"md5hash,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type StorageBucketObjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageBucketObjectSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketObjectList is a list of StorageBucketObjects
type StorageBucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucketObject CRD objects
	Items []StorageBucketObject `json:"items,omitempty"`
}
