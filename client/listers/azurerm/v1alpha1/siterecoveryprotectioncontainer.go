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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SiteRecoveryProtectionContainerLister helps list SiteRecoveryProtectionContainers.
type SiteRecoveryProtectionContainerLister interface {
	// List lists all SiteRecoveryProtectionContainers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryProtectionContainer, err error)
	// SiteRecoveryProtectionContainers returns an object that can list and get SiteRecoveryProtectionContainers.
	SiteRecoveryProtectionContainers(namespace string) SiteRecoveryProtectionContainerNamespaceLister
	SiteRecoveryProtectionContainerListerExpansion
}

// siteRecoveryProtectionContainerLister implements the SiteRecoveryProtectionContainerLister interface.
type siteRecoveryProtectionContainerLister struct {
	indexer cache.Indexer
}

// NewSiteRecoveryProtectionContainerLister returns a new SiteRecoveryProtectionContainerLister.
func NewSiteRecoveryProtectionContainerLister(indexer cache.Indexer) SiteRecoveryProtectionContainerLister {
	return &siteRecoveryProtectionContainerLister{indexer: indexer}
}

// List lists all SiteRecoveryProtectionContainers in the indexer.
func (s *siteRecoveryProtectionContainerLister) List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryProtectionContainer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SiteRecoveryProtectionContainer))
	})
	return ret, err
}

// SiteRecoveryProtectionContainers returns an object that can list and get SiteRecoveryProtectionContainers.
func (s *siteRecoveryProtectionContainerLister) SiteRecoveryProtectionContainers(namespace string) SiteRecoveryProtectionContainerNamespaceLister {
	return siteRecoveryProtectionContainerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SiteRecoveryProtectionContainerNamespaceLister helps list and get SiteRecoveryProtectionContainers.
type SiteRecoveryProtectionContainerNamespaceLister interface {
	// List lists all SiteRecoveryProtectionContainers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryProtectionContainer, err error)
	// Get retrieves the SiteRecoveryProtectionContainer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SiteRecoveryProtectionContainer, error)
	SiteRecoveryProtectionContainerNamespaceListerExpansion
}

// siteRecoveryProtectionContainerNamespaceLister implements the SiteRecoveryProtectionContainerNamespaceLister
// interface.
type siteRecoveryProtectionContainerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SiteRecoveryProtectionContainers in the indexer for a given namespace.
func (s siteRecoveryProtectionContainerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryProtectionContainer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SiteRecoveryProtectionContainer))
	})
	return ret, err
}

// Get retrieves the SiteRecoveryProtectionContainer from the indexer for a given namespace and name.
func (s siteRecoveryProtectionContainerNamespaceLister) Get(name string) (*v1alpha1.SiteRecoveryProtectionContainer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("siterecoveryprotectioncontainer"), name)
	}
	return obj.(*v1alpha1.SiteRecoveryProtectionContainer), nil
}
