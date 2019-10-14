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

// CloudformationStackLister helps list CloudformationStacks.
type CloudformationStackLister interface {
	// List lists all CloudformationStacks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudformationStack, err error)
	// CloudformationStacks returns an object that can list and get CloudformationStacks.
	CloudformationStacks(namespace string) CloudformationStackNamespaceLister
	CloudformationStackListerExpansion
}

// cloudformationStackLister implements the CloudformationStackLister interface.
type cloudformationStackLister struct {
	indexer cache.Indexer
}

// NewCloudformationStackLister returns a new CloudformationStackLister.
func NewCloudformationStackLister(indexer cache.Indexer) CloudformationStackLister {
	return &cloudformationStackLister{indexer: indexer}
}

// List lists all CloudformationStacks in the indexer.
func (s *cloudformationStackLister) List(selector labels.Selector) (ret []*v1alpha1.CloudformationStack, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudformationStack))
	})
	return ret, err
}

// CloudformationStacks returns an object that can list and get CloudformationStacks.
func (s *cloudformationStackLister) CloudformationStacks(namespace string) CloudformationStackNamespaceLister {
	return cloudformationStackNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudformationStackNamespaceLister helps list and get CloudformationStacks.
type CloudformationStackNamespaceLister interface {
	// List lists all CloudformationStacks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudformationStack, err error)
	// Get retrieves the CloudformationStack from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudformationStack, error)
	CloudformationStackNamespaceListerExpansion
}

// cloudformationStackNamespaceLister implements the CloudformationStackNamespaceLister
// interface.
type cloudformationStackNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudformationStacks in the indexer for a given namespace.
func (s cloudformationStackNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudformationStack, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudformationStack))
	})
	return ret, err
}

// Get retrieves the CloudformationStack from the indexer for a given namespace and name.
func (s cloudformationStackNamespaceLister) Get(name string) (*v1alpha1.CloudformationStack, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudformationstack"), name)
	}
	return obj.(*v1alpha1.CloudformationStack), nil
}
