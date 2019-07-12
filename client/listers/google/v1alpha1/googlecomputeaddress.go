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

// GoogleComputeAddressLister helps list GoogleComputeAddresses.
type GoogleComputeAddressLister interface {
	// List lists all GoogleComputeAddresses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeAddress, err error)
	// Get retrieves the GoogleComputeAddress from the index for a given name.
	Get(name string) (*v1alpha1.GoogleComputeAddress, error)
	GoogleComputeAddressListerExpansion
}

// googleComputeAddressLister implements the GoogleComputeAddressLister interface.
type googleComputeAddressLister struct {
	indexer cache.Indexer
}

// NewGoogleComputeAddressLister returns a new GoogleComputeAddressLister.
func NewGoogleComputeAddressLister(indexer cache.Indexer) GoogleComputeAddressLister {
	return &googleComputeAddressLister{indexer: indexer}
}

// List lists all GoogleComputeAddresses in the indexer.
func (s *googleComputeAddressLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeAddress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleComputeAddress))
	})
	return ret, err
}

// Get retrieves the GoogleComputeAddress from the index for a given name.
func (s *googleComputeAddressLister) Get(name string) (*v1alpha1.GoogleComputeAddress, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googlecomputeaddress"), name)
	}
	return obj.(*v1alpha1.GoogleComputeAddress), nil
}
