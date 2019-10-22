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

// PubsubTopicLister helps list PubsubTopics.
type PubsubTopicLister interface {
	// List lists all PubsubTopics in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubTopic, err error)
	// PubsubTopics returns an object that can list and get PubsubTopics.
	PubsubTopics(namespace string) PubsubTopicNamespaceLister
	PubsubTopicListerExpansion
}

// pubsubTopicLister implements the PubsubTopicLister interface.
type pubsubTopicLister struct {
	indexer cache.Indexer
}

// NewPubsubTopicLister returns a new PubsubTopicLister.
func NewPubsubTopicLister(indexer cache.Indexer) PubsubTopicLister {
	return &pubsubTopicLister{indexer: indexer}
}

// List lists all PubsubTopics in the indexer.
func (s *pubsubTopicLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubTopic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubTopic))
	})
	return ret, err
}

// PubsubTopics returns an object that can list and get PubsubTopics.
func (s *pubsubTopicLister) PubsubTopics(namespace string) PubsubTopicNamespaceLister {
	return pubsubTopicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PubsubTopicNamespaceLister helps list and get PubsubTopics.
type PubsubTopicNamespaceLister interface {
	// List lists all PubsubTopics in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubTopic, err error)
	// Get retrieves the PubsubTopic from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PubsubTopic, error)
	PubsubTopicNamespaceListerExpansion
}

// pubsubTopicNamespaceLister implements the PubsubTopicNamespaceLister
// interface.
type pubsubTopicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PubsubTopics in the indexer for a given namespace.
func (s pubsubTopicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubTopic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubTopic))
	})
	return ret, err
}

// Get retrieves the PubsubTopic from the indexer for a given namespace and name.
func (s pubsubTopicNamespaceLister) Get(name string) (*v1alpha1.PubsubTopic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pubsubtopic"), name)
	}
	return obj.(*v1alpha1.PubsubTopic), nil
}
