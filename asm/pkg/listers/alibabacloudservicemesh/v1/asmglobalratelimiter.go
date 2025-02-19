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

// ASMGlobalRateLimiterLister helps list ASMGlobalRateLimiters.
// All objects returned here must be treated as read-only.
type ASMGlobalRateLimiterLister interface {
	// List lists all ASMGlobalRateLimiters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ASMGlobalRateLimiter, err error)
	// ASMGlobalRateLimiters returns an object that can list and get ASMGlobalRateLimiters.
	ASMGlobalRateLimiters(namespace string) ASMGlobalRateLimiterNamespaceLister
	ASMGlobalRateLimiterListerExpansion
}

// aSMGlobalRateLimiterLister implements the ASMGlobalRateLimiterLister interface.
type aSMGlobalRateLimiterLister struct {
	indexer cache.Indexer
}

// NewASMGlobalRateLimiterLister returns a new ASMGlobalRateLimiterLister.
func NewASMGlobalRateLimiterLister(indexer cache.Indexer) ASMGlobalRateLimiterLister {
	return &aSMGlobalRateLimiterLister{indexer: indexer}
}

// List lists all ASMGlobalRateLimiters in the indexer.
func (s *aSMGlobalRateLimiterLister) List(selector labels.Selector) (ret []*v1.ASMGlobalRateLimiter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ASMGlobalRateLimiter))
	})
	return ret, err
}

// ASMGlobalRateLimiters returns an object that can list and get ASMGlobalRateLimiters.
func (s *aSMGlobalRateLimiterLister) ASMGlobalRateLimiters(namespace string) ASMGlobalRateLimiterNamespaceLister {
	return aSMGlobalRateLimiterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ASMGlobalRateLimiterNamespaceLister helps list and get ASMGlobalRateLimiters.
// All objects returned here must be treated as read-only.
type ASMGlobalRateLimiterNamespaceLister interface {
	// List lists all ASMGlobalRateLimiters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ASMGlobalRateLimiter, err error)
	// Get retrieves the ASMGlobalRateLimiter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ASMGlobalRateLimiter, error)
	ASMGlobalRateLimiterNamespaceListerExpansion
}

// aSMGlobalRateLimiterNamespaceLister implements the ASMGlobalRateLimiterNamespaceLister
// interface.
type aSMGlobalRateLimiterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ASMGlobalRateLimiters in the indexer for a given namespace.
func (s aSMGlobalRateLimiterNamespaceLister) List(selector labels.Selector) (ret []*v1.ASMGlobalRateLimiter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ASMGlobalRateLimiter))
	})
	return ret, err
}

// Get retrieves the ASMGlobalRateLimiter from the indexer for a given namespace and name.
func (s aSMGlobalRateLimiterNamespaceLister) Get(name string) (*v1.ASMGlobalRateLimiter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("asmglobalratelimiter"), name)
	}
	return obj.(*v1.ASMGlobalRateLimiter), nil
}
