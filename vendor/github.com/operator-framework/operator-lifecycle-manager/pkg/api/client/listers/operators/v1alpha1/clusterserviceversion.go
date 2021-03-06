/*
Copyright Red Hat, Inc.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterServiceVersionLister helps list ClusterServiceVersions.
type ClusterServiceVersionLister interface {
	// List lists all ClusterServiceVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterServiceVersion, err error)
	// ClusterServiceVersions returns an object that can list and get ClusterServiceVersions.
	ClusterServiceVersions(namespace string) ClusterServiceVersionNamespaceLister
	ClusterServiceVersionListerExpansion
}

// clusterServiceVersionLister implements the ClusterServiceVersionLister interface.
type clusterServiceVersionLister struct {
	indexer cache.Indexer
}

// NewClusterServiceVersionLister returns a new ClusterServiceVersionLister.
func NewClusterServiceVersionLister(indexer cache.Indexer) ClusterServiceVersionLister {
	return &clusterServiceVersionLister{indexer: indexer}
}

// List lists all ClusterServiceVersions in the indexer.
func (s *clusterServiceVersionLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterServiceVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterServiceVersion))
	})
	return ret, err
}

// ClusterServiceVersions returns an object that can list and get ClusterServiceVersions.
func (s *clusterServiceVersionLister) ClusterServiceVersions(namespace string) ClusterServiceVersionNamespaceLister {
	return clusterServiceVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterServiceVersionNamespaceLister helps list and get ClusterServiceVersions.
type ClusterServiceVersionNamespaceLister interface {
	// List lists all ClusterServiceVersions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterServiceVersion, err error)
	// Get retrieves the ClusterServiceVersion from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ClusterServiceVersion, error)
	ClusterServiceVersionNamespaceListerExpansion
}

// clusterServiceVersionNamespaceLister implements the ClusterServiceVersionNamespaceLister
// interface.
type clusterServiceVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterServiceVersions in the indexer for a given namespace.
func (s clusterServiceVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterServiceVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterServiceVersion))
	})
	return ret, err
}

// Get retrieves the ClusterServiceVersion from the indexer for a given namespace and name.
func (s clusterServiceVersionNamespaceLister) Get(name string) (*v1alpha1.ClusterServiceVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterserviceversion"), name)
	}
	return obj.(*v1alpha1.ClusterServiceVersion), nil
}
