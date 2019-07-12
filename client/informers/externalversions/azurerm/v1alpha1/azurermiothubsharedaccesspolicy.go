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

// AzurermIothubSharedAccessPolicyInformer provides access to a shared informer and lister for
// AzurermIothubSharedAccessPolicies.
type AzurermIothubSharedAccessPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzurermIothubSharedAccessPolicyLister
}

type azurermIothubSharedAccessPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAzurermIothubSharedAccessPolicyInformer constructs a new informer for AzurermIothubSharedAccessPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzurermIothubSharedAccessPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzurermIothubSharedAccessPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAzurermIothubSharedAccessPolicyInformer constructs a new informer for AzurermIothubSharedAccessPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzurermIothubSharedAccessPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzurermV1alpha1().AzurermIothubSharedAccessPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzurermV1alpha1().AzurermIothubSharedAccessPolicies().Watch(options)
			},
		},
		&azurermv1alpha1.AzurermIothubSharedAccessPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *azurermIothubSharedAccessPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzurermIothubSharedAccessPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azurermIothubSharedAccessPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&azurermv1alpha1.AzurermIothubSharedAccessPolicy{}, f.defaultInformer)
}

func (f *azurermIothubSharedAccessPolicyInformer) Lister() v1alpha1.AzurermIothubSharedAccessPolicyLister {
	return v1alpha1.NewAzurermIothubSharedAccessPolicyLister(f.Informer().GetIndexer())
}
