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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	googlev1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SourcerepoRepositoryIamBindingInformer provides access to a shared informer and lister for
// SourcerepoRepositoryIamBindings.
type SourcerepoRepositoryIamBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SourcerepoRepositoryIamBindingLister
}

type sourcerepoRepositoryIamBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSourcerepoRepositoryIamBindingInformer constructs a new informer for SourcerepoRepositoryIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSourcerepoRepositoryIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSourcerepoRepositoryIamBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSourcerepoRepositoryIamBindingInformer constructs a new informer for SourcerepoRepositoryIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSourcerepoRepositoryIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().SourcerepoRepositoryIamBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().SourcerepoRepositoryIamBindings(namespace).Watch(options)
			},
		},
		&googlev1alpha1.SourcerepoRepositoryIamBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *sourcerepoRepositoryIamBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSourcerepoRepositoryIamBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sourcerepoRepositoryIamBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&googlev1alpha1.SourcerepoRepositoryIamBinding{}, f.defaultInformer)
}

func (f *sourcerepoRepositoryIamBindingInformer) Lister() v1alpha1.SourcerepoRepositoryIamBindingLister {
	return v1alpha1.NewSourcerepoRepositoryIamBindingLister(f.Informer().GetIndexer())
}
