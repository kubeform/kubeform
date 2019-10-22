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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DefaultVpcDHCPOptionsLister helps list DefaultVpcDHCPOptionses.
type DefaultVpcDHCPOptionsLister interface {
	// List lists all DefaultVpcDHCPOptionses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultVpcDHCPOptions, err error)
	// DefaultVpcDHCPOptionses returns an object that can list and get DefaultVpcDHCPOptionses.
	DefaultVpcDHCPOptionses(namespace string) DefaultVpcDHCPOptionsNamespaceLister
	DefaultVpcDHCPOptionsListerExpansion
}

// defaultVpcDHCPOptionsLister implements the DefaultVpcDHCPOptionsLister interface.
type defaultVpcDHCPOptionsLister struct {
	indexer cache.Indexer
}

// NewDefaultVpcDHCPOptionsLister returns a new DefaultVpcDHCPOptionsLister.
func NewDefaultVpcDHCPOptionsLister(indexer cache.Indexer) DefaultVpcDHCPOptionsLister {
	return &defaultVpcDHCPOptionsLister{indexer: indexer}
}

// List lists all DefaultVpcDHCPOptionses in the indexer.
func (s *defaultVpcDHCPOptionsLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultVpcDHCPOptions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultVpcDHCPOptions))
	})
	return ret, err
}

// DefaultVpcDHCPOptionses returns an object that can list and get DefaultVpcDHCPOptionses.
func (s *defaultVpcDHCPOptionsLister) DefaultVpcDHCPOptionses(namespace string) DefaultVpcDHCPOptionsNamespaceLister {
	return defaultVpcDHCPOptionsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DefaultVpcDHCPOptionsNamespaceLister helps list and get DefaultVpcDHCPOptionses.
type DefaultVpcDHCPOptionsNamespaceLister interface {
	// List lists all DefaultVpcDHCPOptionses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultVpcDHCPOptions, err error)
	// Get retrieves the DefaultVpcDHCPOptions from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DefaultVpcDHCPOptions, error)
	DefaultVpcDHCPOptionsNamespaceListerExpansion
}

// defaultVpcDHCPOptionsNamespaceLister implements the DefaultVpcDHCPOptionsNamespaceLister
// interface.
type defaultVpcDHCPOptionsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DefaultVpcDHCPOptionses in the indexer for a given namespace.
func (s defaultVpcDHCPOptionsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultVpcDHCPOptions, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultVpcDHCPOptions))
	})
	return ret, err
}

// Get retrieves the DefaultVpcDHCPOptions from the indexer for a given namespace and name.
func (s defaultVpcDHCPOptionsNamespaceLister) Get(name string) (*v1alpha1.DefaultVpcDHCPOptions, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("defaultvpcdhcpoptions"), name)
	}
	return obj.(*v1alpha1.DefaultVpcDHCPOptions), nil
}
