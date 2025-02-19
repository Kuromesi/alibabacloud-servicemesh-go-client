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

// ASMSwimLaneGroupLister helps list ASMSwimLaneGroups.
// All objects returned here must be treated as read-only.
type ASMSwimLaneGroupLister interface {
	// List lists all ASMSwimLaneGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ASMSwimLaneGroup, err error)
	// Get retrieves the ASMSwimLaneGroup from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ASMSwimLaneGroup, error)
	ASMSwimLaneGroupListerExpansion
}

// aSMSwimLaneGroupLister implements the ASMSwimLaneGroupLister interface.
type aSMSwimLaneGroupLister struct {
	indexer cache.Indexer
}

// NewASMSwimLaneGroupLister returns a new ASMSwimLaneGroupLister.
func NewASMSwimLaneGroupLister(indexer cache.Indexer) ASMSwimLaneGroupLister {
	return &aSMSwimLaneGroupLister{indexer: indexer}
}

// List lists all ASMSwimLaneGroups in the indexer.
func (s *aSMSwimLaneGroupLister) List(selector labels.Selector) (ret []*v1.ASMSwimLaneGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ASMSwimLaneGroup))
	})
	return ret, err
}

// Get retrieves the ASMSwimLaneGroup from the index for a given name.
func (s *aSMSwimLaneGroupLister) Get(name string) (*v1.ASMSwimLaneGroup, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("asmswimlanegroup"), name)
	}
	return obj.(*v1.ASMSwimLaneGroup), nil
}
