/*
Copyright 2019 The Kubeform Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// IotThingPrincipalAttachmentLister helps list IotThingPrincipalAttachments.
type IotThingPrincipalAttachmentLister interface {
	// List lists all IotThingPrincipalAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IotThingPrincipalAttachment, err error)
	// IotThingPrincipalAttachments returns an object that can list and get IotThingPrincipalAttachments.
	IotThingPrincipalAttachments(namespace string) IotThingPrincipalAttachmentNamespaceLister
	IotThingPrincipalAttachmentListerExpansion
}

// iotThingPrincipalAttachmentLister implements the IotThingPrincipalAttachmentLister interface.
type iotThingPrincipalAttachmentLister struct {
	indexer cache.Indexer
}

// NewIotThingPrincipalAttachmentLister returns a new IotThingPrincipalAttachmentLister.
func NewIotThingPrincipalAttachmentLister(indexer cache.Indexer) IotThingPrincipalAttachmentLister {
	return &iotThingPrincipalAttachmentLister{indexer: indexer}
}

// List lists all IotThingPrincipalAttachments in the indexer.
func (s *iotThingPrincipalAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.IotThingPrincipalAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotThingPrincipalAttachment))
	})
	return ret, err
}

// IotThingPrincipalAttachments returns an object that can list and get IotThingPrincipalAttachments.
func (s *iotThingPrincipalAttachmentLister) IotThingPrincipalAttachments(namespace string) IotThingPrincipalAttachmentNamespaceLister {
	return iotThingPrincipalAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IotThingPrincipalAttachmentNamespaceLister helps list and get IotThingPrincipalAttachments.
type IotThingPrincipalAttachmentNamespaceLister interface {
	// List lists all IotThingPrincipalAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IotThingPrincipalAttachment, err error)
	// Get retrieves the IotThingPrincipalAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IotThingPrincipalAttachment, error)
	IotThingPrincipalAttachmentNamespaceListerExpansion
}

// iotThingPrincipalAttachmentNamespaceLister implements the IotThingPrincipalAttachmentNamespaceLister
// interface.
type iotThingPrincipalAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IotThingPrincipalAttachments in the indexer for a given namespace.
func (s iotThingPrincipalAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IotThingPrincipalAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotThingPrincipalAttachment))
	})
	return ret, err
}

// Get retrieves the IotThingPrincipalAttachment from the indexer for a given namespace and name.
func (s iotThingPrincipalAttachmentNamespaceLister) Get(name string) (*v1alpha1.IotThingPrincipalAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iotthingprincipalattachment"), name)
	}
	return obj.(*v1alpha1.IotThingPrincipalAttachment), nil
}
