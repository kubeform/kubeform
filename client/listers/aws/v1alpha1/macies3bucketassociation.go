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

// MacieS3BucketAssociationLister helps list MacieS3BucketAssociations.
type MacieS3BucketAssociationLister interface {
	// List lists all MacieS3BucketAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MacieS3BucketAssociation, err error)
	// MacieS3BucketAssociations returns an object that can list and get MacieS3BucketAssociations.
	MacieS3BucketAssociations(namespace string) MacieS3BucketAssociationNamespaceLister
	MacieS3BucketAssociationListerExpansion
}

// macieS3BucketAssociationLister implements the MacieS3BucketAssociationLister interface.
type macieS3BucketAssociationLister struct {
	indexer cache.Indexer
}

// NewMacieS3BucketAssociationLister returns a new MacieS3BucketAssociationLister.
func NewMacieS3BucketAssociationLister(indexer cache.Indexer) MacieS3BucketAssociationLister {
	return &macieS3BucketAssociationLister{indexer: indexer}
}

// List lists all MacieS3BucketAssociations in the indexer.
func (s *macieS3BucketAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.MacieS3BucketAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MacieS3BucketAssociation))
	})
	return ret, err
}

// MacieS3BucketAssociations returns an object that can list and get MacieS3BucketAssociations.
func (s *macieS3BucketAssociationLister) MacieS3BucketAssociations(namespace string) MacieS3BucketAssociationNamespaceLister {
	return macieS3BucketAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MacieS3BucketAssociationNamespaceLister helps list and get MacieS3BucketAssociations.
type MacieS3BucketAssociationNamespaceLister interface {
	// List lists all MacieS3BucketAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MacieS3BucketAssociation, err error)
	// Get retrieves the MacieS3BucketAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MacieS3BucketAssociation, error)
	MacieS3BucketAssociationNamespaceListerExpansion
}

// macieS3BucketAssociationNamespaceLister implements the MacieS3BucketAssociationNamespaceLister
// interface.
type macieS3BucketAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MacieS3BucketAssociations in the indexer for a given namespace.
func (s macieS3BucketAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MacieS3BucketAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MacieS3BucketAssociation))
	})
	return ret, err
}

// Get retrieves the MacieS3BucketAssociation from the indexer for a given namespace and name.
func (s macieS3BucketAssociationNamespaceLister) Get(name string) (*v1alpha1.MacieS3BucketAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("macies3bucketassociation"), name)
	}
	return obj.(*v1alpha1.MacieS3BucketAssociation), nil
}
