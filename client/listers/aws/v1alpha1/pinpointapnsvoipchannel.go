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

// PinpointApnsVoipChannelLister helps list PinpointApnsVoipChannels.
type PinpointApnsVoipChannelLister interface {
	// List lists all PinpointApnsVoipChannels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipChannel, err error)
	// PinpointApnsVoipChannels returns an object that can list and get PinpointApnsVoipChannels.
	PinpointApnsVoipChannels(namespace string) PinpointApnsVoipChannelNamespaceLister
	PinpointApnsVoipChannelListerExpansion
}

// pinpointApnsVoipChannelLister implements the PinpointApnsVoipChannelLister interface.
type pinpointApnsVoipChannelLister struct {
	indexer cache.Indexer
}

// NewPinpointApnsVoipChannelLister returns a new PinpointApnsVoipChannelLister.
func NewPinpointApnsVoipChannelLister(indexer cache.Indexer) PinpointApnsVoipChannelLister {
	return &pinpointApnsVoipChannelLister{indexer: indexer}
}

// List lists all PinpointApnsVoipChannels in the indexer.
func (s *pinpointApnsVoipChannelLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointApnsVoipChannel))
	})
	return ret, err
}

// PinpointApnsVoipChannels returns an object that can list and get PinpointApnsVoipChannels.
func (s *pinpointApnsVoipChannelLister) PinpointApnsVoipChannels(namespace string) PinpointApnsVoipChannelNamespaceLister {
	return pinpointApnsVoipChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PinpointApnsVoipChannelNamespaceLister helps list and get PinpointApnsVoipChannels.
type PinpointApnsVoipChannelNamespaceLister interface {
	// List lists all PinpointApnsVoipChannels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipChannel, err error)
	// Get retrieves the PinpointApnsVoipChannel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PinpointApnsVoipChannel, error)
	PinpointApnsVoipChannelNamespaceListerExpansion
}

// pinpointApnsVoipChannelNamespaceLister implements the PinpointApnsVoipChannelNamespaceLister
// interface.
type pinpointApnsVoipChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PinpointApnsVoipChannels in the indexer for a given namespace.
func (s pinpointApnsVoipChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointApnsVoipChannel))
	})
	return ret, err
}

// Get retrieves the PinpointApnsVoipChannel from the indexer for a given namespace and name.
func (s pinpointApnsVoipChannelNamespaceLister) Get(name string) (*v1alpha1.PinpointApnsVoipChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pinpointapnsvoipchannel"), name)
	}
	return obj.(*v1alpha1.PinpointApnsVoipChannel), nil
}