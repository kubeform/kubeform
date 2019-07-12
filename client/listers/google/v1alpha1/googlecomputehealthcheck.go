/*
Copyright 2019 The KubeDB Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// GoogleComputeHealthCheckLister helps list GoogleComputeHealthChecks.
type GoogleComputeHealthCheckLister interface {
	// List lists all GoogleComputeHealthChecks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeHealthCheck, err error)
	// Get retrieves the GoogleComputeHealthCheck from the index for a given name.
	Get(name string) (*v1alpha1.GoogleComputeHealthCheck, error)
	GoogleComputeHealthCheckListerExpansion
}

// googleComputeHealthCheckLister implements the GoogleComputeHealthCheckLister interface.
type googleComputeHealthCheckLister struct {
	indexer cache.Indexer
}

// NewGoogleComputeHealthCheckLister returns a new GoogleComputeHealthCheckLister.
func NewGoogleComputeHealthCheckLister(indexer cache.Indexer) GoogleComputeHealthCheckLister {
	return &googleComputeHealthCheckLister{indexer: indexer}
}

// List lists all GoogleComputeHealthChecks in the indexer.
func (s *googleComputeHealthCheckLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeHealthCheck, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleComputeHealthCheck))
	})
	return ret, err
}

// Get retrieves the GoogleComputeHealthCheck from the index for a given name.
func (s *googleComputeHealthCheckLister) Get(name string) (*v1alpha1.GoogleComputeHealthCheck, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googlecomputehealthcheck"), name)
	}
	return obj.(*v1alpha1.GoogleComputeHealthCheck), nil
}
