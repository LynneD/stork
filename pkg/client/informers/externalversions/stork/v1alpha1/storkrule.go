/*
Copyright 2018 Openstorage.org

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

	stork_v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	versioned "github.com/libopenstorage/stork/pkg/client/clientset/versioned"
	internalinterfaces "github.com/libopenstorage/stork/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/libopenstorage/stork/pkg/client/listers/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorkRuleInformer provides access to a shared informer and lister for
// StorkRules.
type StorkRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorkRuleLister
}

type storkRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStorkRuleInformer constructs a new informer for StorkRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorkRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorkRuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStorkRuleInformer constructs a new informer for StorkRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorkRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorkV1alpha1().StorkRules(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorkV1alpha1().StorkRules(namespace).Watch(options)
			},
		},
		&stork_v1alpha1.StorkRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *storkRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorkRuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storkRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&stork_v1alpha1.StorkRule{}, f.defaultInformer)
}

func (f *storkRuleInformer) Lister() v1alpha1.StorkRuleLister {
	return v1alpha1.NewStorkRuleLister(f.Informer().GetIndexer())
}
