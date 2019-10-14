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

// Route53ResolverEndpointLister helps list Route53ResolverEndpoints.
type Route53ResolverEndpointLister interface {
	// List lists all Route53ResolverEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Route53ResolverEndpoint, err error)
	// Route53ResolverEndpoints returns an object that can list and get Route53ResolverEndpoints.
	Route53ResolverEndpoints(namespace string) Route53ResolverEndpointNamespaceLister
	Route53ResolverEndpointListerExpansion
}

// route53ResolverEndpointLister implements the Route53ResolverEndpointLister interface.
type route53ResolverEndpointLister struct {
	indexer cache.Indexer
}

// NewRoute53ResolverEndpointLister returns a new Route53ResolverEndpointLister.
func NewRoute53ResolverEndpointLister(indexer cache.Indexer) Route53ResolverEndpointLister {
	return &route53ResolverEndpointLister{indexer: indexer}
}

// List lists all Route53ResolverEndpoints in the indexer.
func (s *route53ResolverEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.Route53ResolverEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route53ResolverEndpoint))
	})
	return ret, err
}

// Route53ResolverEndpoints returns an object that can list and get Route53ResolverEndpoints.
func (s *route53ResolverEndpointLister) Route53ResolverEndpoints(namespace string) Route53ResolverEndpointNamespaceLister {
	return route53ResolverEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Route53ResolverEndpointNamespaceLister helps list and get Route53ResolverEndpoints.
type Route53ResolverEndpointNamespaceLister interface {
	// List lists all Route53ResolverEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Route53ResolverEndpoint, err error)
	// Get retrieves the Route53ResolverEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Route53ResolverEndpoint, error)
	Route53ResolverEndpointNamespaceListerExpansion
}

// route53ResolverEndpointNamespaceLister implements the Route53ResolverEndpointNamespaceLister
// interface.
type route53ResolverEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Route53ResolverEndpoints in the indexer for a given namespace.
func (s route53ResolverEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Route53ResolverEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route53ResolverEndpoint))
	})
	return ret, err
}

// Get retrieves the Route53ResolverEndpoint from the indexer for a given namespace and name.
func (s route53ResolverEndpointNamespaceLister) Get(name string) (*v1alpha1.Route53ResolverEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("route53resolverendpoint"), name)
	}
	return obj.(*v1alpha1.Route53ResolverEndpoint), nil
}
