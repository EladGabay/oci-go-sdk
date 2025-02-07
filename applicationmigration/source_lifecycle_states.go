// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"strings"
)

// SourceLifecycleStatesEnum Enum with underlying type: string
type SourceLifecycleStatesEnum string

// Set of constants representing the allowable values for SourceLifecycleStatesEnum
const (
	SourceLifecycleStatesCreating SourceLifecycleStatesEnum = "CREATING"
	SourceLifecycleStatesDeleting SourceLifecycleStatesEnum = "DELETING"
	SourceLifecycleStatesUpdating SourceLifecycleStatesEnum = "UPDATING"
	SourceLifecycleStatesActive   SourceLifecycleStatesEnum = "ACTIVE"
	SourceLifecycleStatesInactive SourceLifecycleStatesEnum = "INACTIVE"
	SourceLifecycleStatesDeleted  SourceLifecycleStatesEnum = "DELETED"
)

var mappingSourceLifecycleStatesEnum = map[string]SourceLifecycleStatesEnum{
	"CREATING": SourceLifecycleStatesCreating,
	"DELETING": SourceLifecycleStatesDeleting,
	"UPDATING": SourceLifecycleStatesUpdating,
	"ACTIVE":   SourceLifecycleStatesActive,
	"INACTIVE": SourceLifecycleStatesInactive,
	"DELETED":  SourceLifecycleStatesDeleted,
}

var mappingSourceLifecycleStatesEnumLowerCase = map[string]SourceLifecycleStatesEnum{
	"creating": SourceLifecycleStatesCreating,
	"deleting": SourceLifecycleStatesDeleting,
	"updating": SourceLifecycleStatesUpdating,
	"active":   SourceLifecycleStatesActive,
	"inactive": SourceLifecycleStatesInactive,
	"deleted":  SourceLifecycleStatesDeleted,
}

// GetSourceLifecycleStatesEnumValues Enumerates the set of values for SourceLifecycleStatesEnum
func GetSourceLifecycleStatesEnumValues() []SourceLifecycleStatesEnum {
	values := make([]SourceLifecycleStatesEnum, 0)
	for _, v := range mappingSourceLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceLifecycleStatesEnumStringValues Enumerates the set of values in String for SourceLifecycleStatesEnum
func GetSourceLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"DELETING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingSourceLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceLifecycleStatesEnum(val string) (SourceLifecycleStatesEnum, bool) {
	enum, ok := mappingSourceLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
