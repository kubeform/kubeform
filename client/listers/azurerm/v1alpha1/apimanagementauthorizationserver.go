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

// ApiManagementAuthorizationServerLister helps list ApiManagementAuthorizationServers.
type ApiManagementAuthorizationServerLister interface {
	// List lists all ApiManagementAuthorizationServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementAuthorizationServer, err error)
	// Get retrieves the ApiManagementAuthorizationServer from the index for a given name.
	Get(name string) (*v1alpha1.ApiManagementAuthorizationServer, error)
	ApiManagementAuthorizationServerListerExpansion
}

// apiManagementAuthorizationServerLister implements the ApiManagementAuthorizationServerLister interface.
type apiManagementAuthorizationServerLister struct {
	indexer cache.Indexer
}

// NewApiManagementAuthorizationServerLister returns a new ApiManagementAuthorizationServerLister.
func NewApiManagementAuthorizationServerLister(indexer cache.Indexer) ApiManagementAuthorizationServerLister {
	return &apiManagementAuthorizationServerLister{indexer: indexer}
}

// List lists all ApiManagementAuthorizationServers in the indexer.
func (s *apiManagementAuthorizationServerLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementAuthorizationServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementAuthorizationServer))
	})
	return ret, err
}

// Get retrieves the ApiManagementAuthorizationServer from the index for a given name.
func (s *apiManagementAuthorizationServerLister) Get(name string) (*v1alpha1.ApiManagementAuthorizationServer, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementauthorizationserver"), name)
	}
	return obj.(*v1alpha1.ApiManagementAuthorizationServer), nil
}
