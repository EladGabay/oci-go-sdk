// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v39/common"
)

// ConnectDescriptor Connect Descriptor details.
type ConnectDescriptor struct {

	// Host of the connect descriptor.
	Host *string `mandatory:"false" json:"host"`

	// Port of the connect descriptor.
	Port *int `mandatory:"false" json:"port"`

	// Database service name.
	DatabaseServiceName *string `mandatory:"false" json:"databaseServiceName"`

	// Connect string.
	ConnectString *string `mandatory:"false" json:"connectString"`
}

func (m ConnectDescriptor) String() string {
	return common.PointerString(m)
}
