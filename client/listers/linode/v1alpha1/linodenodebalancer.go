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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// LinodeNodebalancerLister helps list LinodeNodebalancers.
type LinodeNodebalancerLister interface {
	// List lists all LinodeNodebalancers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LinodeNodebalancer, err error)
	// Get retrieves the LinodeNodebalancer from the index for a given name.
	Get(name string) (*v1alpha1.LinodeNodebalancer, error)
	LinodeNodebalancerListerExpansion
}

// linodeNodebalancerLister implements the LinodeNodebalancerLister interface.
type linodeNodebalancerLister struct {
	indexer cache.Indexer
}

// NewLinodeNodebalancerLister returns a new LinodeNodebalancerLister.
func NewLinodeNodebalancerLister(indexer cache.Indexer) LinodeNodebalancerLister {
	return &linodeNodebalancerLister{indexer: indexer}
}

// List lists all LinodeNodebalancers in the indexer.
func (s *linodeNodebalancerLister) List(selector labels.Selector) (ret []*v1alpha1.LinodeNodebalancer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LinodeNodebalancer))
	})
	return ret, err
}

// Get retrieves the LinodeNodebalancer from the index for a given name.
func (s *linodeNodebalancerLister) Get(name string) (*v1alpha1.LinodeNodebalancer, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("linodenodebalancer"), name)
	}
	return obj.(*v1alpha1.LinodeNodebalancer), nil
}
