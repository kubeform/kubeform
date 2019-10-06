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

// DmsEndpointLister helps list DmsEndpoints.
type DmsEndpointLister interface {
	// List lists all DmsEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DmsEndpoint, err error)
	// DmsEndpoints returns an object that can list and get DmsEndpoints.
	DmsEndpoints(namespace string) DmsEndpointNamespaceLister
	DmsEndpointListerExpansion
}

// dmsEndpointLister implements the DmsEndpointLister interface.
type dmsEndpointLister struct {
	indexer cache.Indexer
}

// NewDmsEndpointLister returns a new DmsEndpointLister.
func NewDmsEndpointLister(indexer cache.Indexer) DmsEndpointLister {
	return &dmsEndpointLister{indexer: indexer}
}

// List lists all DmsEndpoints in the indexer.
func (s *dmsEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.DmsEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsEndpoint))
	})
	return ret, err
}

// DmsEndpoints returns an object that can list and get DmsEndpoints.
func (s *dmsEndpointLister) DmsEndpoints(namespace string) DmsEndpointNamespaceLister {
	return dmsEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DmsEndpointNamespaceLister helps list and get DmsEndpoints.
type DmsEndpointNamespaceLister interface {
	// List lists all DmsEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DmsEndpoint, err error)
	// Get retrieves the DmsEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DmsEndpoint, error)
	DmsEndpointNamespaceListerExpansion
}

// dmsEndpointNamespaceLister implements the DmsEndpointNamespaceLister
// interface.
type dmsEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DmsEndpoints in the indexer for a given namespace.
func (s dmsEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DmsEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsEndpoint))
	})
	return ret, err
}

// Get retrieves the DmsEndpoint from the indexer for a given namespace and name.
func (s dmsEndpointNamespaceLister) Get(name string) (*v1alpha1.DmsEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dmsendpoint"), name)
	}
	return obj.(*v1alpha1.DmsEndpoint), nil
}
