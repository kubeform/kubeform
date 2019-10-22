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

	awsv1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VpcIpv4CIDRBlockAssociationInformer provides access to a shared informer and lister for
// VpcIpv4CIDRBlockAssociations.
type VpcIpv4CIDRBlockAssociationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VpcIpv4CIDRBlockAssociationLister
}

type vpcIpv4CIDRBlockAssociationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVpcIpv4CIDRBlockAssociationInformer constructs a new informer for VpcIpv4CIDRBlockAssociation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVpcIpv4CIDRBlockAssociationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVpcIpv4CIDRBlockAssociationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVpcIpv4CIDRBlockAssociationInformer constructs a new informer for VpcIpv4CIDRBlockAssociation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVpcIpv4CIDRBlockAssociationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().VpcIpv4CIDRBlockAssociations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().VpcIpv4CIDRBlockAssociations(namespace).Watch(options)
			},
		},
		&awsv1alpha1.VpcIpv4CIDRBlockAssociation{},
		resyncPeriod,
		indexers,
	)
}

func (f *vpcIpv4CIDRBlockAssociationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVpcIpv4CIDRBlockAssociationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vpcIpv4CIDRBlockAssociationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.VpcIpv4CIDRBlockAssociation{}, f.defaultInformer)
}

func (f *vpcIpv4CIDRBlockAssociationInformer) Lister() v1alpha1.VpcIpv4CIDRBlockAssociationLister {
	return v1alpha1.NewVpcIpv4CIDRBlockAssociationLister(f.Informer().GetIndexer())
}
