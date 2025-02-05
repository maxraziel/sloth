// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PrometheusRuleLister helps list PrometheusRules.
// All objects returned here must be treated as read-only.
type PrometheusRuleLister interface {
	// List lists all PrometheusRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CustomPrometheusRules, err error)
	// PrometheusRules returns an object that can list and get PrometheusRules.
	PrometheusRules(namespace string) PrometheusRuleNamespaceLister
	PrometheusRuleListerExpansion
}

// prometheusRuleLister implements the PrometheusRuleLister interface.
type prometheusRuleLister struct {
	indexer cache.Indexer
}

// NewPrometheusRuleLister returns a new PrometheusRuleLister.
func NewPrometheusRuleLister(indexer cache.Indexer) PrometheusRuleLister {
	return &prometheusRuleLister{indexer: indexer}
}

// List lists all PrometheusRules in the indexer.
func (s *prometheusRuleLister) List(selector labels.Selector) (ret []*v1.CustomPrometheusRules, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CustomPrometheusRules))
	})
	return ret, err
}

// PrometheusRules returns an object that can list and get PrometheusRules.
func (s *prometheusRuleLister) PrometheusRules(namespace string) PrometheusRuleNamespaceLister {
	return prometheusRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrometheusRuleNamespaceLister helps list and get PrometheusRules.
// All objects returned here must be treated as read-only.
type PrometheusRuleNamespaceLister interface {
	// List lists all PrometheusRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CustomPrometheusRules, err error)
	// Get retrieves the PrometheusRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CustomPrometheusRules, error)
	PrometheusRuleNamespaceListerExpansion
}

// prometheusRuleNamespaceLister implements the PrometheusRuleNamespaceLister
// interface.
type prometheusRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrometheusRules in the indexer for a given namespace.
func (s prometheusRuleNamespaceLister) List(selector labels.Selector) (ret []*v1.CustomPrometheusRules, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CustomPrometheusRules))
	})
	return ret, err
}

// Get retrieves the PrometheusRule from the indexer for a given namespace and name.
func (s prometheusRuleNamespaceLister) Get(name string) (*v1.CustomPrometheusRules, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("prometheusrule"), name)
	}
	return obj.(*v1.CustomPrometheusRules), nil
}
