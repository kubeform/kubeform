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

// GoogleProjectIamBindingLister helps list GoogleProjectIamBindings.
type GoogleProjectIamBindingLister interface {
	// List lists all GoogleProjectIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleProjectIamBinding, err error)
	// Get retrieves the GoogleProjectIamBinding from the index for a given name.
	Get(name string) (*v1alpha1.GoogleProjectIamBinding, error)
	GoogleProjectIamBindingListerExpansion
}

// googleProjectIamBindingLister implements the GoogleProjectIamBindingLister interface.
type googleProjectIamBindingLister struct {
	indexer cache.Indexer
}

// NewGoogleProjectIamBindingLister returns a new GoogleProjectIamBindingLister.
func NewGoogleProjectIamBindingLister(indexer cache.Indexer) GoogleProjectIamBindingLister {
	return &googleProjectIamBindingLister{indexer: indexer}
}

// List lists all GoogleProjectIamBindings in the indexer.
func (s *googleProjectIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleProjectIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleProjectIamBinding))
	})
	return ret, err
}

// Get retrieves the GoogleProjectIamBinding from the index for a given name.
func (s *googleProjectIamBindingLister) Get(name string) (*v1alpha1.GoogleProjectIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googleprojectiambinding"), name)
	}
	return obj.(*v1alpha1.GoogleProjectIamBinding), nil
}
