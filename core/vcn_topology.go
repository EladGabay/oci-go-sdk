// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v39/common"
)

// VcnTopology Defines the representation of a virtual network topology for a VCN.
type VcnTopology struct {

	// Lists entities comprising the virtual network topology.
	Entities []interface{} `mandatory:"true" json:"entities"`

	// Lists relationships between entities in the virtual network topology.
	Relationships []TopologyEntityRelationship `mandatory:"true" json:"relationships"`

	// Records when the virtual network topology was created, in RFC3339 (https://tools.ietf.org/html/rfc3339) format for date and time.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for which the topology is generated.
	VcnId *string `mandatory:"false" json:"vcnId"`
}

//GetEntities returns Entities
func (m VcnTopology) GetEntities() []interface{} {
	return m.Entities
}

//GetRelationships returns Relationships
func (m VcnTopology) GetRelationships() []TopologyEntityRelationship {
	return m.Relationships
}

//GetTimeCreated returns TimeCreated
func (m VcnTopology) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m VcnTopology) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m VcnTopology) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVcnTopology VcnTopology
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVcnTopology
	}{
		"VCN",
		(MarshalTypeVcnTopology)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *VcnTopology) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		VcnId         *string                      `json:"vcnId"`
		Entities      []interface{}                `json:"entities"`
		Relationships []topologyentityrelationship `json:"relationships"`
		TimeCreated   *common.SDKTime              `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.VcnId = model.VcnId

	m.Entities = make([]interface{}, len(model.Entities))
	for i, n := range model.Entities {
		m.Entities[i] = n
	}

	m.Relationships = make([]TopologyEntityRelationship, len(model.Relationships))
	for i, n := range model.Relationships {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Relationships[i] = nn.(TopologyEntityRelationship)
		} else {
			m.Relationships[i] = nil
		}
	}

	m.TimeCreated = model.TimeCreated

	return
}
