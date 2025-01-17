// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	extensionsv1alpha1 "istio.io/client-go/pkg/apis/extensions/v1alpha1"
	versioned "istio.io/client-go/pkg/clientset/versioned"
	internalinterfaces "istio.io/client-go/pkg/informers/externalversions/internalinterfaces"
	v1alpha1 "istio.io/client-go/pkg/listers/extensions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceMetadataInformer provides access to a shared informer and lister for
// ServiceMetadatas.
type ServiceMetadataInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceMetadataLister
}

type serviceMetadataInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceMetadataInformer constructs a new informer for ServiceMetadata type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceMetadataInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceMetadataInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceMetadataInformer constructs a new informer for ServiceMetadata type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceMetadataInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1alpha1().ServiceMetadatas(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1alpha1().ServiceMetadatas(namespace).Watch(context.TODO(), options)
			},
		},
		&extensionsv1alpha1.ServiceMetadata{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceMetadataInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceMetadataInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceMetadataInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensionsv1alpha1.ServiceMetadata{}, f.defaultInformer)
}

func (f *serviceMetadataInformer) Lister() v1alpha1.ServiceMetadataLister {
	return v1alpha1.NewServiceMetadataLister(f.Informer().GetIndexer())
}
