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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// ContainerAnalysisNoteLister helps list ContainerAnalysisNotes.
type ContainerAnalysisNoteLister interface {
	// List lists all ContainerAnalysisNotes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerAnalysisNote, err error)
	// ContainerAnalysisNotes returns an object that can list and get ContainerAnalysisNotes.
	ContainerAnalysisNotes(namespace string) ContainerAnalysisNoteNamespaceLister
	ContainerAnalysisNoteListerExpansion
}

// containerAnalysisNoteLister implements the ContainerAnalysisNoteLister interface.
type containerAnalysisNoteLister struct {
	indexer cache.Indexer
}

// NewContainerAnalysisNoteLister returns a new ContainerAnalysisNoteLister.
func NewContainerAnalysisNoteLister(indexer cache.Indexer) ContainerAnalysisNoteLister {
	return &containerAnalysisNoteLister{indexer: indexer}
}

// List lists all ContainerAnalysisNotes in the indexer.
func (s *containerAnalysisNoteLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerAnalysisNote, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerAnalysisNote))
	})
	return ret, err
}

// ContainerAnalysisNotes returns an object that can list and get ContainerAnalysisNotes.
func (s *containerAnalysisNoteLister) ContainerAnalysisNotes(namespace string) ContainerAnalysisNoteNamespaceLister {
	return containerAnalysisNoteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ContainerAnalysisNoteNamespaceLister helps list and get ContainerAnalysisNotes.
type ContainerAnalysisNoteNamespaceLister interface {
	// List lists all ContainerAnalysisNotes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerAnalysisNote, err error)
	// Get retrieves the ContainerAnalysisNote from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ContainerAnalysisNote, error)
	ContainerAnalysisNoteNamespaceListerExpansion
}

// containerAnalysisNoteNamespaceLister implements the ContainerAnalysisNoteNamespaceLister
// interface.
type containerAnalysisNoteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ContainerAnalysisNotes in the indexer for a given namespace.
func (s containerAnalysisNoteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerAnalysisNote, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerAnalysisNote))
	})
	return ret, err
}

// Get retrieves the ContainerAnalysisNote from the indexer for a given namespace and name.
func (s containerAnalysisNoteNamespaceLister) Get(name string) (*v1alpha1.ContainerAnalysisNote, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("containeranalysisnote"), name)
	}
	return obj.(*v1alpha1.ContainerAnalysisNote), nil
}