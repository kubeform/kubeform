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

// VpcPeeringConnectionAccepterLister helps list VpcPeeringConnectionAccepters.
type VpcPeeringConnectionAccepterLister interface {
	// List lists all VpcPeeringConnectionAccepters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnectionAccepter, err error)
	// VpcPeeringConnectionAccepters returns an object that can list and get VpcPeeringConnectionAccepters.
	VpcPeeringConnectionAccepters(namespace string) VpcPeeringConnectionAccepterNamespaceLister
	VpcPeeringConnectionAccepterListerExpansion
}

// vpcPeeringConnectionAccepterLister implements the VpcPeeringConnectionAccepterLister interface.
type vpcPeeringConnectionAccepterLister struct {
	indexer cache.Indexer
}

// NewVpcPeeringConnectionAccepterLister returns a new VpcPeeringConnectionAccepterLister.
func NewVpcPeeringConnectionAccepterLister(indexer cache.Indexer) VpcPeeringConnectionAccepterLister {
	return &vpcPeeringConnectionAccepterLister{indexer: indexer}
}

// List lists all VpcPeeringConnectionAccepters in the indexer.
func (s *vpcPeeringConnectionAccepterLister) List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnectionAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcPeeringConnectionAccepter))
	})
	return ret, err
}

// VpcPeeringConnectionAccepters returns an object that can list and get VpcPeeringConnectionAccepters.
func (s *vpcPeeringConnectionAccepterLister) VpcPeeringConnectionAccepters(namespace string) VpcPeeringConnectionAccepterNamespaceLister {
	return vpcPeeringConnectionAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcPeeringConnectionAccepterNamespaceLister helps list and get VpcPeeringConnectionAccepters.
type VpcPeeringConnectionAccepterNamespaceLister interface {
	// List lists all VpcPeeringConnectionAccepters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnectionAccepter, err error)
	// Get retrieves the VpcPeeringConnectionAccepter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VpcPeeringConnectionAccepter, error)
	VpcPeeringConnectionAccepterNamespaceListerExpansion
}

// vpcPeeringConnectionAccepterNamespaceLister implements the VpcPeeringConnectionAccepterNamespaceLister
// interface.
type vpcPeeringConnectionAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcPeeringConnectionAccepters in the indexer for a given namespace.
func (s vpcPeeringConnectionAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnectionAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcPeeringConnectionAccepter))
	})
	return ret, err
}

// Get retrieves the VpcPeeringConnectionAccepter from the indexer for a given namespace and name.
func (s vpcPeeringConnectionAccepterNamespaceLister) Get(name string) (*v1alpha1.VpcPeeringConnectionAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcpeeringconnectionaccepter"), name)
	}
	return obj.(*v1alpha1.VpcPeeringConnectionAccepter), nil
}
