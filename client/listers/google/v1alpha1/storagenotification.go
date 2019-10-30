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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageNotificationLister helps list StorageNotifications.
type StorageNotificationLister interface {
	// List lists all StorageNotifications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageNotification, err error)
	// StorageNotifications returns an object that can list and get StorageNotifications.
	StorageNotifications(namespace string) StorageNotificationNamespaceLister
	StorageNotificationListerExpansion
}

// storageNotificationLister implements the StorageNotificationLister interface.
type storageNotificationLister struct {
	indexer cache.Indexer
}

// NewStorageNotificationLister returns a new StorageNotificationLister.
func NewStorageNotificationLister(indexer cache.Indexer) StorageNotificationLister {
	return &storageNotificationLister{indexer: indexer}
}

// List lists all StorageNotifications in the indexer.
func (s *storageNotificationLister) List(selector labels.Selector) (ret []*v1alpha1.StorageNotification, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageNotification))
	})
	return ret, err
}

// StorageNotifications returns an object that can list and get StorageNotifications.
func (s *storageNotificationLister) StorageNotifications(namespace string) StorageNotificationNamespaceLister {
	return storageNotificationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageNotificationNamespaceLister helps list and get StorageNotifications.
type StorageNotificationNamespaceLister interface {
	// List lists all StorageNotifications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageNotification, err error)
	// Get retrieves the StorageNotification from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageNotification, error)
	StorageNotificationNamespaceListerExpansion
}

// storageNotificationNamespaceLister implements the StorageNotificationNamespaceLister
// interface.
type storageNotificationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageNotifications in the indexer for a given namespace.
func (s storageNotificationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageNotification, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageNotification))
	})
	return ret, err
}

// Get retrieves the StorageNotification from the indexer for a given namespace and name.
func (s storageNotificationNamespaceLister) Get(name string) (*v1alpha1.StorageNotification, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagenotification"), name)
	}
	return obj.(*v1alpha1.StorageNotification), nil
}
