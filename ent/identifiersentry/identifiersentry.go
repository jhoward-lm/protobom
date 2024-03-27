// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------

package identifiersentry

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the identifiersentry type in the database.
	Label = "identifiers_entry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSoftwareIdentifierType holds the string denoting the software_identifier_type field in the database.
	FieldSoftwareIdentifierType = "software_identifier_type"
	// FieldSoftwareIdentifierValue holds the string denoting the software_identifier_value field in the database.
	FieldSoftwareIdentifierValue = "software_identifier_value"
	// EdgeNodes holds the string denoting the nodes edge name in mutations.
	EdgeNodes = "nodes"
	// Table holds the table name of the identifiersentry in the database.
	Table = "identifiers_entries"
	// NodesTable is the table that holds the nodes relation/edge. The primary key declared below.
	NodesTable = "node_identifiers"
	// NodesInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodesInverseTable = "nodes"
)

// Columns holds all SQL columns for identifiersentry fields.
var Columns = []string{
	FieldID,
	FieldSoftwareIdentifierType,
	FieldSoftwareIdentifierValue,
}

var (
	// NodesPrimaryKey and NodesColumn2 are the table columns denoting the
	// primary key for the nodes relation (M2M).
	NodesPrimaryKey = []string{"node_id", "identifiers_entry_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// SoftwareIdentifierType defines the type for the "software_identifier_type" enum field.
type SoftwareIdentifierType string

// SoftwareIdentifierType values.
const (
	SoftwareIdentifierTypeUNKNOWN_IDENTIFIER_TYPE SoftwareIdentifierType = "UNKNOWN_IDENTIFIER_TYPE"
	SoftwareIdentifierTypePURL                    SoftwareIdentifierType = "PURL"
	SoftwareIdentifierTypeCPE22                   SoftwareIdentifierType = "CPE22"
	SoftwareIdentifierTypeCPE23                   SoftwareIdentifierType = "CPE23"
	SoftwareIdentifierTypeGITOID                  SoftwareIdentifierType = "GITOID"
)

func (sit SoftwareIdentifierType) String() string {
	return string(sit)
}

// SoftwareIdentifierTypeValidator is a validator for the "software_identifier_type" field enum values. It is called by the builders before save.
func SoftwareIdentifierTypeValidator(sit SoftwareIdentifierType) error {
	switch sit {
	case SoftwareIdentifierTypeUNKNOWN_IDENTIFIER_TYPE, SoftwareIdentifierTypePURL, SoftwareIdentifierTypeCPE22, SoftwareIdentifierTypeCPE23, SoftwareIdentifierTypeGITOID:
		return nil
	default:
		return fmt.Errorf("identifiersentry: invalid enum value for software_identifier_type field: %q", sit)
	}
}

// OrderOption defines the ordering options for the IdentifiersEntry queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySoftwareIdentifierType orders the results by the software_identifier_type field.
func BySoftwareIdentifierType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSoftwareIdentifierType, opts...).ToFunc()
}

// BySoftwareIdentifierValue orders the results by the software_identifier_value field.
func BySoftwareIdentifierValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSoftwareIdentifierValue, opts...).ToFunc()
}

// ByNodesCount orders the results by nodes count.
func ByNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodesStep(), opts...)
	}
}

// ByNodes orders the results by nodes terms.
func ByNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, NodesTable, NodesPrimaryKey...),
	)
}