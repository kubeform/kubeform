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
	googlev1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/google/v1alpha1"
)

// GoogleStorageObjectAccessControlInformer provides access to a shared informer and lister for
// GoogleStorageObjectAccessControls.
type GoogleStorageObjectAccessControlInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GoogleStorageObjectAccessControlLister
}

type googleStorageObjectAccessControlInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewGoogleStorageObjectAccessControlInformer constructs a new informer for GoogleStorageObjectAccessControl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGoogleStorageObjectAccessControlInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGoogleStorageObjectAccessControlInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredGoogleStorageObjectAccessControlInformer constructs a new informer for GoogleStorageObjectAccessControl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGoogleStorageObjectAccessControlInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().GoogleStorageObjectAccessControls().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().GoogleStorageObjectAccessControls().Watch(options)
			},
		},
		&googlev1alpha1.GoogleStorageObjectAccessControl{},
		resyncPeriod,
		indexers,
	)
}

func (f *googleStorageObjectAccessControlInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGoogleStorageObjectAccessControlInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *googleStorageObjectAccessControlInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&googlev1alpha1.GoogleStorageObjectAccessControl{}, f.defaultInformer)
}

func (f *googleStorageObjectAccessControlInformer) Lister() v1alpha1.GoogleStorageObjectAccessControlLister {
	return v1alpha1.NewGoogleStorageObjectAccessControlLister(f.Informer().GetIndexer())
}
