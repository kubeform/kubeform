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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	azurermv1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/azurerm/v1alpha1"
)

// AzurermApiManagementOpenidConnectProviderInformer provides access to a shared informer and lister for
// AzurermApiManagementOpenidConnectProviders.
type AzurermApiManagementOpenidConnectProviderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzurermApiManagementOpenidConnectProviderLister
}

type azurermApiManagementOpenidConnectProviderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAzurermApiManagementOpenidConnectProviderInformer constructs a new informer for AzurermApiManagementOpenidConnectProvider type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzurermApiManagementOpenidConnectProviderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzurermApiManagementOpenidConnectProviderInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAzurermApiManagementOpenidConnectProviderInformer constructs a new informer for AzurermApiManagementOpenidConnectProvider type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzurermApiManagementOpenidConnectProviderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzurermV1alpha1().AzurermApiManagementOpenidConnectProviders().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzurermV1alpha1().AzurermApiManagementOpenidConnectProviders().Watch(options)
			},
		},
		&azurermv1alpha1.AzurermApiManagementOpenidConnectProvider{},
		resyncPeriod,
		indexers,
	)
}

func (f *azurermApiManagementOpenidConnectProviderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzurermApiManagementOpenidConnectProviderInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azurermApiManagementOpenidConnectProviderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&azurermv1alpha1.AzurermApiManagementOpenidConnectProvider{}, f.defaultInformer)
}

func (f *azurermApiManagementOpenidConnectProviderInformer) Lister() v1alpha1.AzurermApiManagementOpenidConnectProviderLister {
	return v1alpha1.NewAzurermApiManagementOpenidConnectProviderLister(f.Informer().GetIndexer())
}
