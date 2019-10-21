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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// PinpointEmailChannelLister helps list PinpointEmailChannels.
type PinpointEmailChannelLister interface {
	// List lists all PinpointEmailChannels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointEmailChannel, err error)
	// PinpointEmailChannels returns an object that can list and get PinpointEmailChannels.
	PinpointEmailChannels(namespace string) PinpointEmailChannelNamespaceLister
	PinpointEmailChannelListerExpansion
}

// pinpointEmailChannelLister implements the PinpointEmailChannelLister interface.
type pinpointEmailChannelLister struct {
	indexer cache.Indexer
}

// NewPinpointEmailChannelLister returns a new PinpointEmailChannelLister.
func NewPinpointEmailChannelLister(indexer cache.Indexer) PinpointEmailChannelLister {
	return &pinpointEmailChannelLister{indexer: indexer}
}

// List lists all PinpointEmailChannels in the indexer.
func (s *pinpointEmailChannelLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointEmailChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointEmailChannel))
	})
	return ret, err
}

// PinpointEmailChannels returns an object that can list and get PinpointEmailChannels.
func (s *pinpointEmailChannelLister) PinpointEmailChannels(namespace string) PinpointEmailChannelNamespaceLister {
	return pinpointEmailChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PinpointEmailChannelNamespaceLister helps list and get PinpointEmailChannels.
type PinpointEmailChannelNamespaceLister interface {
	// List lists all PinpointEmailChannels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointEmailChannel, err error)
	// Get retrieves the PinpointEmailChannel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PinpointEmailChannel, error)
	PinpointEmailChannelNamespaceListerExpansion
}

// pinpointEmailChannelNamespaceLister implements the PinpointEmailChannelNamespaceLister
// interface.
type pinpointEmailChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PinpointEmailChannels in the indexer for a given namespace.
func (s pinpointEmailChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointEmailChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointEmailChannel))
	})
	return ret, err
}

// Get retrieves the PinpointEmailChannel from the indexer for a given namespace and name.
func (s pinpointEmailChannelNamespaceLister) Get(name string) (*v1alpha1.PinpointEmailChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pinpointemailchannel"), name)
	}
	return obj.(*v1alpha1.PinpointEmailChannel), nil
}