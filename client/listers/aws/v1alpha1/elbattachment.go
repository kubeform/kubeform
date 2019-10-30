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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ElbAttachmentLister helps list ElbAttachments.
type ElbAttachmentLister interface {
	// List lists all ElbAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ElbAttachment, err error)
	// ElbAttachments returns an object that can list and get ElbAttachments.
	ElbAttachments(namespace string) ElbAttachmentNamespaceLister
	ElbAttachmentListerExpansion
}

// elbAttachmentLister implements the ElbAttachmentLister interface.
type elbAttachmentLister struct {
	indexer cache.Indexer
}

// NewElbAttachmentLister returns a new ElbAttachmentLister.
func NewElbAttachmentLister(indexer cache.Indexer) ElbAttachmentLister {
	return &elbAttachmentLister{indexer: indexer}
}

// List lists all ElbAttachments in the indexer.
func (s *elbAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.ElbAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElbAttachment))
	})
	return ret, err
}

// ElbAttachments returns an object that can list and get ElbAttachments.
func (s *elbAttachmentLister) ElbAttachments(namespace string) ElbAttachmentNamespaceLister {
	return elbAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElbAttachmentNamespaceLister helps list and get ElbAttachments.
type ElbAttachmentNamespaceLister interface {
	// List lists all ElbAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ElbAttachment, err error)
	// Get retrieves the ElbAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ElbAttachment, error)
	ElbAttachmentNamespaceListerExpansion
}

// elbAttachmentNamespaceLister implements the ElbAttachmentNamespaceLister
// interface.
type elbAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElbAttachments in the indexer for a given namespace.
func (s elbAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ElbAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElbAttachment))
	})
	return ret, err
}

// Get retrieves the ElbAttachment from the indexer for a given namespace and name.
func (s elbAttachmentNamespaceLister) Get(name string) (*v1alpha1.ElbAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elbattachment"), name)
	}
	return obj.(*v1alpha1.ElbAttachment), nil
}
