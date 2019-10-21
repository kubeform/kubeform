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

// ElasticsearchDomainPolicyLister helps list ElasticsearchDomainPolicies.
type ElasticsearchDomainPolicyLister interface {
	// List lists all ElasticsearchDomainPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchDomainPolicy, err error)
	// ElasticsearchDomainPolicies returns an object that can list and get ElasticsearchDomainPolicies.
	ElasticsearchDomainPolicies(namespace string) ElasticsearchDomainPolicyNamespaceLister
	ElasticsearchDomainPolicyListerExpansion
}

// elasticsearchDomainPolicyLister implements the ElasticsearchDomainPolicyLister interface.
type elasticsearchDomainPolicyLister struct {
	indexer cache.Indexer
}

// NewElasticsearchDomainPolicyLister returns a new ElasticsearchDomainPolicyLister.
func NewElasticsearchDomainPolicyLister(indexer cache.Indexer) ElasticsearchDomainPolicyLister {
	return &elasticsearchDomainPolicyLister{indexer: indexer}
}

// List lists all ElasticsearchDomainPolicies in the indexer.
func (s *elasticsearchDomainPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchDomainPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticsearchDomainPolicy))
	})
	return ret, err
}

// ElasticsearchDomainPolicies returns an object that can list and get ElasticsearchDomainPolicies.
func (s *elasticsearchDomainPolicyLister) ElasticsearchDomainPolicies(namespace string) ElasticsearchDomainPolicyNamespaceLister {
	return elasticsearchDomainPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElasticsearchDomainPolicyNamespaceLister helps list and get ElasticsearchDomainPolicies.
type ElasticsearchDomainPolicyNamespaceLister interface {
	// List lists all ElasticsearchDomainPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchDomainPolicy, err error)
	// Get retrieves the ElasticsearchDomainPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ElasticsearchDomainPolicy, error)
	ElasticsearchDomainPolicyNamespaceListerExpansion
}

// elasticsearchDomainPolicyNamespaceLister implements the ElasticsearchDomainPolicyNamespaceLister
// interface.
type elasticsearchDomainPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElasticsearchDomainPolicies in the indexer for a given namespace.
func (s elasticsearchDomainPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchDomainPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticsearchDomainPolicy))
	})
	return ret, err
}

// Get retrieves the ElasticsearchDomainPolicy from the indexer for a given namespace and name.
func (s elasticsearchDomainPolicyNamespaceLister) Get(name string) (*v1alpha1.ElasticsearchDomainPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elasticsearchdomainpolicy"), name)
	}
	return obj.(*v1alpha1.ElasticsearchDomainPolicy), nil
}