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

// GoogleComputeFirewallLister helps list GoogleComputeFirewalls.
type GoogleComputeFirewallLister interface {
	// List lists all GoogleComputeFirewalls in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeFirewall, err error)
	// Get retrieves the GoogleComputeFirewall from the index for a given name.
	Get(name string) (*v1alpha1.GoogleComputeFirewall, error)
	GoogleComputeFirewallListerExpansion
}

// googleComputeFirewallLister implements the GoogleComputeFirewallLister interface.
type googleComputeFirewallLister struct {
	indexer cache.Indexer
}

// NewGoogleComputeFirewallLister returns a new GoogleComputeFirewallLister.
func NewGoogleComputeFirewallLister(indexer cache.Indexer) GoogleComputeFirewallLister {
	return &googleComputeFirewallLister{indexer: indexer}
}

// List lists all GoogleComputeFirewalls in the indexer.
func (s *googleComputeFirewallLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeFirewall, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleComputeFirewall))
	})
	return ret, err
}

// Get retrieves the GoogleComputeFirewall from the index for a given name.
func (s *googleComputeFirewallLister) Get(name string) (*v1alpha1.GoogleComputeFirewall, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googlecomputefirewall"), name)
	}
	return obj.(*v1alpha1.GoogleComputeFirewall), nil
}
