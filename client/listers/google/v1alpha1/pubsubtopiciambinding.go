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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PubsubTopicIamBindingLister helps list PubsubTopicIamBindings.
type PubsubTopicIamBindingLister interface {
	// List lists all PubsubTopicIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamBinding, err error)
	// PubsubTopicIamBindings returns an object that can list and get PubsubTopicIamBindings.
	PubsubTopicIamBindings(namespace string) PubsubTopicIamBindingNamespaceLister
	PubsubTopicIamBindingListerExpansion
}

// pubsubTopicIamBindingLister implements the PubsubTopicIamBindingLister interface.
type pubsubTopicIamBindingLister struct {
	indexer cache.Indexer
}

// NewPubsubTopicIamBindingLister returns a new PubsubTopicIamBindingLister.
func NewPubsubTopicIamBindingLister(indexer cache.Indexer) PubsubTopicIamBindingLister {
	return &pubsubTopicIamBindingLister{indexer: indexer}
}

// List lists all PubsubTopicIamBindings in the indexer.
func (s *pubsubTopicIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubTopicIamBinding))
	})
	return ret, err
}

// PubsubTopicIamBindings returns an object that can list and get PubsubTopicIamBindings.
func (s *pubsubTopicIamBindingLister) PubsubTopicIamBindings(namespace string) PubsubTopicIamBindingNamespaceLister {
	return pubsubTopicIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PubsubTopicIamBindingNamespaceLister helps list and get PubsubTopicIamBindings.
type PubsubTopicIamBindingNamespaceLister interface {
	// List lists all PubsubTopicIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamBinding, err error)
	// Get retrieves the PubsubTopicIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PubsubTopicIamBinding, error)
	PubsubTopicIamBindingNamespaceListerExpansion
}

// pubsubTopicIamBindingNamespaceLister implements the PubsubTopicIamBindingNamespaceLister
// interface.
type pubsubTopicIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PubsubTopicIamBindings in the indexer for a given namespace.
func (s pubsubTopicIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubTopicIamBinding))
	})
	return ret, err
}

// Get retrieves the PubsubTopicIamBinding from the indexer for a given namespace and name.
func (s pubsubTopicIamBindingNamespaceLister) Get(name string) (*v1alpha1.PubsubTopicIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pubsubtopiciambinding"), name)
	}
	return obj.(*v1alpha1.PubsubTopicIamBinding), nil
}
