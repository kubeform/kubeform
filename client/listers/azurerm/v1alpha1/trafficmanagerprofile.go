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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TrafficManagerProfileLister helps list TrafficManagerProfiles.
type TrafficManagerProfileLister interface {
	// List lists all TrafficManagerProfiles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TrafficManagerProfile, err error)
	// TrafficManagerProfiles returns an object that can list and get TrafficManagerProfiles.
	TrafficManagerProfiles(namespace string) TrafficManagerProfileNamespaceLister
	TrafficManagerProfileListerExpansion
}

// trafficManagerProfileLister implements the TrafficManagerProfileLister interface.
type trafficManagerProfileLister struct {
	indexer cache.Indexer
}

// NewTrafficManagerProfileLister returns a new TrafficManagerProfileLister.
func NewTrafficManagerProfileLister(indexer cache.Indexer) TrafficManagerProfileLister {
	return &trafficManagerProfileLister{indexer: indexer}
}

// List lists all TrafficManagerProfiles in the indexer.
func (s *trafficManagerProfileLister) List(selector labels.Selector) (ret []*v1alpha1.TrafficManagerProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TrafficManagerProfile))
	})
	return ret, err
}

// TrafficManagerProfiles returns an object that can list and get TrafficManagerProfiles.
func (s *trafficManagerProfileLister) TrafficManagerProfiles(namespace string) TrafficManagerProfileNamespaceLister {
	return trafficManagerProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TrafficManagerProfileNamespaceLister helps list and get TrafficManagerProfiles.
type TrafficManagerProfileNamespaceLister interface {
	// List lists all TrafficManagerProfiles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TrafficManagerProfile, err error)
	// Get retrieves the TrafficManagerProfile from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TrafficManagerProfile, error)
	TrafficManagerProfileNamespaceListerExpansion
}

// trafficManagerProfileNamespaceLister implements the TrafficManagerProfileNamespaceLister
// interface.
type trafficManagerProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TrafficManagerProfiles in the indexer for a given namespace.
func (s trafficManagerProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TrafficManagerProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TrafficManagerProfile))
	})
	return ret, err
}

// Get retrieves the TrafficManagerProfile from the indexer for a given namespace and name.
func (s trafficManagerProfileNamespaceLister) Get(name string) (*v1alpha1.TrafficManagerProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("trafficmanagerprofile"), name)
	}
	return obj.(*v1alpha1.TrafficManagerProfile), nil
}
