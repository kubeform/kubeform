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

// Ec2ClientVPNEndpointLister helps list Ec2ClientVPNEndpoints.
type Ec2ClientVPNEndpointLister interface {
	// List lists all Ec2ClientVPNEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2ClientVPNEndpoint, err error)
	// Ec2ClientVPNEndpoints returns an object that can list and get Ec2ClientVPNEndpoints.
	Ec2ClientVPNEndpoints(namespace string) Ec2ClientVPNEndpointNamespaceLister
	Ec2ClientVPNEndpointListerExpansion
}

// ec2ClientVPNEndpointLister implements the Ec2ClientVPNEndpointLister interface.
type ec2ClientVPNEndpointLister struct {
	indexer cache.Indexer
}

// NewEc2ClientVPNEndpointLister returns a new Ec2ClientVPNEndpointLister.
func NewEc2ClientVPNEndpointLister(indexer cache.Indexer) Ec2ClientVPNEndpointLister {
	return &ec2ClientVPNEndpointLister{indexer: indexer}
}

// List lists all Ec2ClientVPNEndpoints in the indexer.
func (s *ec2ClientVPNEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2ClientVPNEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2ClientVPNEndpoint))
	})
	return ret, err
}

// Ec2ClientVPNEndpoints returns an object that can list and get Ec2ClientVPNEndpoints.
func (s *ec2ClientVPNEndpointLister) Ec2ClientVPNEndpoints(namespace string) Ec2ClientVPNEndpointNamespaceLister {
	return ec2ClientVPNEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Ec2ClientVPNEndpointNamespaceLister helps list and get Ec2ClientVPNEndpoints.
type Ec2ClientVPNEndpointNamespaceLister interface {
	// List lists all Ec2ClientVPNEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2ClientVPNEndpoint, err error)
	// Get retrieves the Ec2ClientVPNEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Ec2ClientVPNEndpoint, error)
	Ec2ClientVPNEndpointNamespaceListerExpansion
}

// ec2ClientVPNEndpointNamespaceLister implements the Ec2ClientVPNEndpointNamespaceLister
// interface.
type ec2ClientVPNEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ec2ClientVPNEndpoints in the indexer for a given namespace.
func (s ec2ClientVPNEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2ClientVPNEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2ClientVPNEndpoint))
	})
	return ret, err
}

// Get retrieves the Ec2ClientVPNEndpoint from the indexer for a given namespace and name.
func (s ec2ClientVPNEndpointNamespaceLister) Get(name string) (*v1alpha1.Ec2ClientVPNEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ec2clientvpnendpoint"), name)
	}
	return obj.(*v1alpha1.Ec2ClientVPNEndpoint), nil
}