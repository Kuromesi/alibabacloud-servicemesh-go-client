// Copyright 2025 Alibaba Cloud Service Mesh
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	clientset "istio.io/client-go/asm/pkg/clientset"
	internalinterfaces "istio.io/client-go/asm/pkg/informers/externalversions/internalinterfaces"
	v1 "istio.io/client-go/asm/pkg/listers/alibabacloudservicemesh/v1"
	"context"
	time "time"

	alibabacloudservicemeshv1 "istio.io/api/alibabacloudservicemesh/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ASMEgressTrafficPolicyInformer provides access to a shared informer and lister for
// ASMEgressTrafficPolicies.
type ASMEgressTrafficPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ASMEgressTrafficPolicyLister
}

type aSMEgressTrafficPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewASMEgressTrafficPolicyInformer constructs a new informer for ASMEgressTrafficPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewASMEgressTrafficPolicyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredASMEgressTrafficPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredASMEgressTrafficPolicyInformer constructs a new informer for ASMEgressTrafficPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredASMEgressTrafficPolicyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IstioV1().ASMEgressTrafficPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IstioV1().ASMEgressTrafficPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&alibabacloudservicemeshv1.ASMEgressTrafficPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *aSMEgressTrafficPolicyInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredASMEgressTrafficPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *aSMEgressTrafficPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&alibabacloudservicemeshv1.ASMEgressTrafficPolicy{}, f.defaultInformer)
}

func (f *aSMEgressTrafficPolicyInformer) Lister() v1.ASMEgressTrafficPolicyLister {
	return v1.NewASMEgressTrafficPolicyLister(f.Informer().GetIndexer())
}
