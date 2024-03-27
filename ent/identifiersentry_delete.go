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
package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/ent/identifiersentry"
	"github.com/bom-squad/protobom/ent/predicate"
)

// IdentifiersEntryDelete is the builder for deleting a IdentifiersEntry entity.
type IdentifiersEntryDelete struct {
	config
	hooks    []Hook
	mutation *IdentifiersEntryMutation
}

// Where appends a list predicates to the IdentifiersEntryDelete builder.
func (ied *IdentifiersEntryDelete) Where(ps ...predicate.IdentifiersEntry) *IdentifiersEntryDelete {
	ied.mutation.Where(ps...)
	return ied
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ied *IdentifiersEntryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ied.sqlExec, ied.mutation, ied.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ied *IdentifiersEntryDelete) ExecX(ctx context.Context) int {
	n, err := ied.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ied *IdentifiersEntryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(identifiersentry.Table, sqlgraph.NewFieldSpec(identifiersentry.FieldID, field.TypeInt))
	if ps := ied.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ied.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ied.mutation.done = true
	return affected, err
}

// IdentifiersEntryDeleteOne is the builder for deleting a single IdentifiersEntry entity.
type IdentifiersEntryDeleteOne struct {
	ied *IdentifiersEntryDelete
}

// Where appends a list predicates to the IdentifiersEntryDelete builder.
func (iedo *IdentifiersEntryDeleteOne) Where(ps ...predicate.IdentifiersEntry) *IdentifiersEntryDeleteOne {
	iedo.ied.mutation.Where(ps...)
	return iedo
}

// Exec executes the deletion query.
func (iedo *IdentifiersEntryDeleteOne) Exec(ctx context.Context) error {
	n, err := iedo.ied.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{identifiersentry.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iedo *IdentifiersEntryDeleteOne) ExecX(ctx context.Context) {
	if err := iedo.Exec(ctx); err != nil {
		panic(err)
	}
}