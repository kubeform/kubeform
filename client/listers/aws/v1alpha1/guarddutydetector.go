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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GuarddutyDetectorLister helps list GuarddutyDetectors.
type GuarddutyDetectorLister interface {
	// List lists all GuarddutyDetectors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GuarddutyDetector, err error)
	// GuarddutyDetectors returns an object that can list and get GuarddutyDetectors.
	GuarddutyDetectors(namespace string) GuarddutyDetectorNamespaceLister
	GuarddutyDetectorListerExpansion
}

// guarddutyDetectorLister implements the GuarddutyDetectorLister interface.
type guarddutyDetectorLister struct {
	indexer cache.Indexer
}

// NewGuarddutyDetectorLister returns a new GuarddutyDetectorLister.
func NewGuarddutyDetectorLister(indexer cache.Indexer) GuarddutyDetectorLister {
	return &guarddutyDetectorLister{indexer: indexer}
}

// List lists all GuarddutyDetectors in the indexer.
func (s *guarddutyDetectorLister) List(selector labels.Selector) (ret []*v1alpha1.GuarddutyDetector, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GuarddutyDetector))
	})
	return ret, err
}

// GuarddutyDetectors returns an object that can list and get GuarddutyDetectors.
func (s *guarddutyDetectorLister) GuarddutyDetectors(namespace string) GuarddutyDetectorNamespaceLister {
	return guarddutyDetectorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GuarddutyDetectorNamespaceLister helps list and get GuarddutyDetectors.
type GuarddutyDetectorNamespaceLister interface {
	// List lists all GuarddutyDetectors in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GuarddutyDetector, err error)
	// Get retrieves the GuarddutyDetector from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GuarddutyDetector, error)
	GuarddutyDetectorNamespaceListerExpansion
}

// guarddutyDetectorNamespaceLister implements the GuarddutyDetectorNamespaceLister
// interface.
type guarddutyDetectorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GuarddutyDetectors in the indexer for a given namespace.
func (s guarddutyDetectorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GuarddutyDetector, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GuarddutyDetector))
	})
	return ret, err
}

// Get retrieves the GuarddutyDetector from the indexer for a given namespace and name.
func (s guarddutyDetectorNamespaceLister) Get(name string) (*v1alpha1.GuarddutyDetector, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("guarddutydetector"), name)
	}
	return obj.(*v1alpha1.GuarddutyDetector), nil
}
