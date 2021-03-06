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

package util

import (
	"time"
)

// Providing 0 duration to an informer indicates that resync should be delayed as long as possible
const (
	NoResyncPeriod time.Duration = 0 * time.Second

	NamespaceKind = "Namespace"
	ServiceKind   = "Service"

	// The following fields are used to interact with unstructured
	// resources.

	// Common fields
	SpecField = "spec"

	// Placement fields
	ClusterNamesField = "clusterNames"

	// Override fields
	ClusterNameField = "clusterName"
	OverridesField   = "overrides"
)

type ReconciliationStatus int

const (
	StatusAllOK ReconciliationStatus = iota
	StatusNeedsRecheck
	StatusError
	StatusNotSynced
)
