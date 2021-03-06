/*
Copyright 2018 The Kubernetes Authors.

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

package common

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/kubernetes-sigs/federation-v2/pkg/apis/core/typeconfig"
	"github.com/kubernetes-sigs/federation-v2/pkg/controller/util"
	"github.com/kubernetes-sigs/federation-v2/pkg/kubefed2/federate"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func NewTestObjects(typeConfig typeconfig.Interface, namespace string, clusterNames []string) (template, placement, override *unstructured.Unstructured, err error) {
	template, err = NewTestTemplate(typeConfig, namespace)
	if err != nil {
		return nil, nil, nil, err
	}

	placement, err = GetPlacementTestObject(typeConfig, namespace, clusterNames)
	if err != nil {
		return nil, nil, nil, err
	}

	if typeConfig.GetOverride() != nil {
		filenameTemplate := fixtureFilenameTemplate(typeConfig)
		overrideFilename := fmt.Sprintf(filenameTemplate, "override")
		override = &unstructured.Unstructured{}
		err := federate.DecodeYAMLFromFile(overrideFilename, override)
		if err != nil {
			return nil, nil, nil, err
		}
		err = UpdateOverrideObject(typeConfig, namespace, clusterNames, override)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return template, placement, override, nil
}

func NewTestTemplate(typeConfig typeconfig.Interface, namespace string) (*unstructured.Unstructured, error) {
	filenameTemplate := fixtureFilenameTemplate(typeConfig)

	templateFilename := fmt.Sprintf(filenameTemplate, "template")
	template := &unstructured.Unstructured{}
	err := federate.DecodeYAMLFromFile(templateFilename, template)
	if err != nil {
		return nil, err
	}
	if typeConfig.GetNamespaced() {
		template.SetNamespace(namespace)
	}
	template.SetName("")
	template.SetGenerateName("test-crud-")

	return template, nil
}

func fixtureFilenameTemplate(typeConfig typeconfig.Interface) string {
	path := fixturePath()
	return filepath.Join(path, fmt.Sprintf("%s-%%s.yaml", strings.ToLower(typeConfig.GetTarget().Kind)))
}

func fixturePath() string {
	// Get the directory of the current executable
	_, filename, _, _ := runtime.Caller(0)
	commonPath := filepath.Dir(filename)
	return filepath.Join(commonPath, "fixtures")
}

func GetPlacementTestObject(typeConfig typeconfig.Interface, namespace string, clusterNames []string) (*unstructured.Unstructured, error) {
	path := fixturePath()
	placementFilename := filepath.Join(path, "placement.yaml")
	placement := &unstructured.Unstructured{}
	err := federate.DecodeYAMLFromFile(placementFilename, placement)
	if err != nil {
		return nil, err
	}
	// Usually placement scope matches resource scope, but
	// FederatedNamespacePlacement is namespaced to allow
	// namespace-scoped tenants to define placement for their
	// namespace.
	if typeConfig.GetNamespaced() || typeConfig.GetTemplate().Kind == util.NamespaceKind {
		placement.SetNamespace(namespace)
	}
	placement.SetName("")
	placement.SetGenerateName("test-crud-")
	placementAPIResource := typeConfig.GetPlacement()
	placement.SetKind(placementAPIResource.Kind)
	placement.SetAPIVersion(fmt.Sprintf("%s/%s", placementAPIResource.Group, placementAPIResource.Version))
	err = util.SetClusterNames(placement, clusterNames)
	if err != nil {
		return nil, err
	}
	return placement, nil
}

// UpdateOverrideObject sets the namespace and applies the given
// cluster names to the override resource provided.
func UpdateOverrideObject(typeConfig typeconfig.Interface, namespace string, clusterNames []string, override *unstructured.Unstructured) error {
	if typeConfig.GetNamespaced() {
		override.SetNamespace(namespace)
	}
	overridesSlice, ok, err := unstructured.NestedSlice(override.Object, "spec", "overrides")
	if err != nil {
		return fmt.Errorf("Error retrieving overrides for %q: %v", typeConfig.GetTemplate().Kind, err)
	}
	var targetOverrides map[string]interface{}
	if ok {
		targetOverrides = overridesSlice[0].(map[string]interface{})
	} else {
		targetOverrides = map[string]interface{}{}
	}
	targetOverrides[util.ClusterNameField] = clusterNames[0]
	overridesSlice[0] = targetOverrides
	err = unstructured.SetNestedSlice(override.Object, overridesSlice, "spec", "overrides")
	if err != nil {
		return fmt.Errorf("Error setting overrides for %q: %v", typeConfig.GetTemplate().Kind, err)
	}
	return nil
}
