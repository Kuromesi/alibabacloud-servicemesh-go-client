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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "istio.io/api/alibabacloudservicemesh/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ASMCircuitBreakerLister helps list ASMCircuitBreakers.
// All objects returned here must be treated as read-only.
type ASMCircuitBreakerLister interface {
	// List lists all ASMCircuitBreakers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ASMCircuitBreaker, err error)
	// ASMCircuitBreakers returns an object that can list and get ASMCircuitBreakers.
	ASMCircuitBreakers(namespace string) ASMCircuitBreakerNamespaceLister
	ASMCircuitBreakerListerExpansion
}

// aSMCircuitBreakerLister implements the ASMCircuitBreakerLister interface.
type aSMCircuitBreakerLister struct {
	indexer cache.Indexer
}

// NewASMCircuitBreakerLister returns a new ASMCircuitBreakerLister.
func NewASMCircuitBreakerLister(indexer cache.Indexer) ASMCircuitBreakerLister {
	return &aSMCircuitBreakerLister{indexer: indexer}
}

// List lists all ASMCircuitBreakers in the indexer.
func (s *aSMCircuitBreakerLister) List(selector labels.Selector) (ret []*v1.ASMCircuitBreaker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ASMCircuitBreaker))
	})
	return ret, err
}

// ASMCircuitBreakers returns an object that can list and get ASMCircuitBreakers.
func (s *aSMCircuitBreakerLister) ASMCircuitBreakers(namespace string) ASMCircuitBreakerNamespaceLister {
	return aSMCircuitBreakerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ASMCircuitBreakerNamespaceLister helps list and get ASMCircuitBreakers.
// All objects returned here must be treated as read-only.
type ASMCircuitBreakerNamespaceLister interface {
	// List lists all ASMCircuitBreakers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ASMCircuitBreaker, err error)
	// Get retrieves the ASMCircuitBreaker from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ASMCircuitBreaker, error)
	ASMCircuitBreakerNamespaceListerExpansion
}

// aSMCircuitBreakerNamespaceLister implements the ASMCircuitBreakerNamespaceLister
// interface.
type aSMCircuitBreakerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ASMCircuitBreakers in the indexer for a given namespace.
func (s aSMCircuitBreakerNamespaceLister) List(selector labels.Selector) (ret []*v1.ASMCircuitBreaker, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ASMCircuitBreaker))
	})
	return ret, err
}

// Get retrieves the ASMCircuitBreaker from the indexer for a given namespace and name.
func (s aSMCircuitBreakerNamespaceLister) Get(name string) (*v1.ASMCircuitBreaker, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("asmcircuitbreaker"), name)
	}
	return obj.(*v1.ASMCircuitBreaker), nil
}
