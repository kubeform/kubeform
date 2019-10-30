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

// ApplicationInsightsLister helps list ApplicationInsightses.
type ApplicationInsightsLister interface {
	// List lists all ApplicationInsightses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsights, err error)
	// ApplicationInsightses returns an object that can list and get ApplicationInsightses.
	ApplicationInsightses(namespace string) ApplicationInsightsNamespaceLister
	ApplicationInsightsListerExpansion
}

// applicationInsightsLister implements the ApplicationInsightsLister interface.
type applicationInsightsLister struct {
	indexer cache.Indexer
}

// NewApplicationInsightsLister returns a new ApplicationInsightsLister.
func NewApplicationInsightsLister(indexer cache.Indexer) ApplicationInsightsLister {
	return &applicationInsightsLister{indexer: indexer}
}

// List lists all ApplicationInsightses in the indexer.
func (s *applicationInsightsLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsights, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationInsights))
	})
	return ret, err
}

// ApplicationInsightses returns an object that can list and get ApplicationInsightses.
func (s *applicationInsightsLister) ApplicationInsightses(namespace string) ApplicationInsightsNamespaceLister {
	return applicationInsightsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationInsightsNamespaceLister helps list and get ApplicationInsightses.
type ApplicationInsightsNamespaceLister interface {
	// List lists all ApplicationInsightses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsights, err error)
	// Get retrieves the ApplicationInsights from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApplicationInsights, error)
	ApplicationInsightsNamespaceListerExpansion
}

// applicationInsightsNamespaceLister implements the ApplicationInsightsNamespaceLister
// interface.
type applicationInsightsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationInsightses in the indexer for a given namespace.
func (s applicationInsightsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsights, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationInsights))
	})
	return ret, err
}

// Get retrieves the ApplicationInsights from the indexer for a given namespace and name.
func (s applicationInsightsNamespaceLister) Get(name string) (*v1alpha1.ApplicationInsights, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationinsights"), name)
	}
	return obj.(*v1alpha1.ApplicationInsights), nil
}
