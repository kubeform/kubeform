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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// ApplicationInsightsApiKeyLister helps list ApplicationInsightsApiKeys.
type ApplicationInsightsApiKeyLister interface {
	// List lists all ApplicationInsightsApiKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsApiKey, err error)
	// Get retrieves the ApplicationInsightsApiKey from the index for a given name.
	Get(name string) (*v1alpha1.ApplicationInsightsApiKey, error)
	ApplicationInsightsApiKeyListerExpansion
}

// applicationInsightsApiKeyLister implements the ApplicationInsightsApiKeyLister interface.
type applicationInsightsApiKeyLister struct {
	indexer cache.Indexer
}

// NewApplicationInsightsApiKeyLister returns a new ApplicationInsightsApiKeyLister.
func NewApplicationInsightsApiKeyLister(indexer cache.Indexer) ApplicationInsightsApiKeyLister {
	return &applicationInsightsApiKeyLister{indexer: indexer}
}

// List lists all ApplicationInsightsApiKeys in the indexer.
func (s *applicationInsightsApiKeyLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsApiKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationInsightsApiKey))
	})
	return ret, err
}

// Get retrieves the ApplicationInsightsApiKey from the index for a given name.
func (s *applicationInsightsApiKeyLister) Get(name string) (*v1alpha1.ApplicationInsightsApiKey, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationinsightsapikey"), name)
	}
	return obj.(*v1alpha1.ApplicationInsightsApiKey), nil
}