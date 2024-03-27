// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/pkg/sbom/ent/documenttype"
	"github.com/bom-squad/protobom/pkg/sbom/ent/metadata"
	"github.com/bom-squad/protobom/pkg/sbom/ent/predicate"
)

// DocumentTypeUpdate is the builder for updating DocumentType entities.
type DocumentTypeUpdate struct {
	config
	hooks    []Hook
	mutation *DocumentTypeMutation
}

// Where appends a list predicates to the DocumentTypeUpdate builder.
func (dtu *DocumentTypeUpdate) Where(ps ...predicate.DocumentType) *DocumentTypeUpdate {
	dtu.mutation.Where(ps...)
	return dtu
}

// SetType sets the "type" field.
func (dtu *DocumentTypeUpdate) SetType(d documenttype.Type) *DocumentTypeUpdate {
	dtu.mutation.SetType(d)
	return dtu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dtu *DocumentTypeUpdate) SetNillableType(d *documenttype.Type) *DocumentTypeUpdate {
	if d != nil {
		dtu.SetType(*d)
	}
	return dtu
}

// ClearType clears the value of the "type" field.
func (dtu *DocumentTypeUpdate) ClearType() *DocumentTypeUpdate {
	dtu.mutation.ClearType()
	return dtu
}

// SetName sets the "name" field.
func (dtu *DocumentTypeUpdate) SetName(s string) *DocumentTypeUpdate {
	dtu.mutation.SetName(s)
	return dtu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dtu *DocumentTypeUpdate) SetNillableName(s *string) *DocumentTypeUpdate {
	if s != nil {
		dtu.SetName(*s)
	}
	return dtu
}

// ClearName clears the value of the "name" field.
func (dtu *DocumentTypeUpdate) ClearName() *DocumentTypeUpdate {
	dtu.mutation.ClearName()
	return dtu
}

// SetDescription sets the "description" field.
func (dtu *DocumentTypeUpdate) SetDescription(s string) *DocumentTypeUpdate {
	dtu.mutation.SetDescription(s)
	return dtu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dtu *DocumentTypeUpdate) SetNillableDescription(s *string) *DocumentTypeUpdate {
	if s != nil {
		dtu.SetDescription(*s)
	}
	return dtu
}

// ClearDescription clears the value of the "description" field.
func (dtu *DocumentTypeUpdate) ClearDescription() *DocumentTypeUpdate {
	dtu.mutation.ClearDescription()
	return dtu
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (dtu *DocumentTypeUpdate) SetMetadataID(id string) *DocumentTypeUpdate {
	dtu.mutation.SetMetadataID(id)
	return dtu
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (dtu *DocumentTypeUpdate) SetNillableMetadataID(id *string) *DocumentTypeUpdate {
	if id != nil {
		dtu = dtu.SetMetadataID(*id)
	}
	return dtu
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (dtu *DocumentTypeUpdate) SetMetadata(m *Metadata) *DocumentTypeUpdate {
	return dtu.SetMetadataID(m.ID)
}

// Mutation returns the DocumentTypeMutation object of the builder.
func (dtu *DocumentTypeUpdate) Mutation() *DocumentTypeMutation {
	return dtu.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (dtu *DocumentTypeUpdate) ClearMetadata() *DocumentTypeUpdate {
	dtu.mutation.ClearMetadata()
	return dtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dtu *DocumentTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dtu.sqlSave, dtu.mutation, dtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dtu *DocumentTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := dtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dtu *DocumentTypeUpdate) Exec(ctx context.Context) error {
	_, err := dtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtu *DocumentTypeUpdate) ExecX(ctx context.Context) {
	if err := dtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtu *DocumentTypeUpdate) check() error {
	if v, ok := dtu.mutation.GetType(); ok {
		if err := documenttype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DocumentType.type": %w`, err)}
		}
	}
	return nil
}

func (dtu *DocumentTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(documenttype.Table, documenttype.Columns, sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt))
	if ps := dtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtu.mutation.GetType(); ok {
		_spec.SetField(documenttype.FieldType, field.TypeEnum, value)
	}
	if dtu.mutation.TypeCleared() {
		_spec.ClearField(documenttype.FieldType, field.TypeEnum)
	}
	if value, ok := dtu.mutation.Name(); ok {
		_spec.SetField(documenttype.FieldName, field.TypeString, value)
	}
	if dtu.mutation.NameCleared() {
		_spec.ClearField(documenttype.FieldName, field.TypeString)
	}
	if value, ok := dtu.mutation.Description(); ok {
		_spec.SetField(documenttype.FieldDescription, field.TypeString, value)
	}
	if dtu.mutation.DescriptionCleared() {
		_spec.ClearField(documenttype.FieldDescription, field.TypeString)
	}
	if dtu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   documenttype.MetadataTable,
			Columns: []string{documenttype.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   documenttype.MetadataTable,
			Columns: []string{documenttype.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{documenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dtu.mutation.done = true
	return n, nil
}

// DocumentTypeUpdateOne is the builder for updating a single DocumentType entity.
type DocumentTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DocumentTypeMutation
}

// SetType sets the "type" field.
func (dtuo *DocumentTypeUpdateOne) SetType(d documenttype.Type) *DocumentTypeUpdateOne {
	dtuo.mutation.SetType(d)
	return dtuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dtuo *DocumentTypeUpdateOne) SetNillableType(d *documenttype.Type) *DocumentTypeUpdateOne {
	if d != nil {
		dtuo.SetType(*d)
	}
	return dtuo
}

// ClearType clears the value of the "type" field.
func (dtuo *DocumentTypeUpdateOne) ClearType() *DocumentTypeUpdateOne {
	dtuo.mutation.ClearType()
	return dtuo
}

// SetName sets the "name" field.
func (dtuo *DocumentTypeUpdateOne) SetName(s string) *DocumentTypeUpdateOne {
	dtuo.mutation.SetName(s)
	return dtuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dtuo *DocumentTypeUpdateOne) SetNillableName(s *string) *DocumentTypeUpdateOne {
	if s != nil {
		dtuo.SetName(*s)
	}
	return dtuo
}

// ClearName clears the value of the "name" field.
func (dtuo *DocumentTypeUpdateOne) ClearName() *DocumentTypeUpdateOne {
	dtuo.mutation.ClearName()
	return dtuo
}

// SetDescription sets the "description" field.
func (dtuo *DocumentTypeUpdateOne) SetDescription(s string) *DocumentTypeUpdateOne {
	dtuo.mutation.SetDescription(s)
	return dtuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dtuo *DocumentTypeUpdateOne) SetNillableDescription(s *string) *DocumentTypeUpdateOne {
	if s != nil {
		dtuo.SetDescription(*s)
	}
	return dtuo
}

// ClearDescription clears the value of the "description" field.
func (dtuo *DocumentTypeUpdateOne) ClearDescription() *DocumentTypeUpdateOne {
	dtuo.mutation.ClearDescription()
	return dtuo
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (dtuo *DocumentTypeUpdateOne) SetMetadataID(id string) *DocumentTypeUpdateOne {
	dtuo.mutation.SetMetadataID(id)
	return dtuo
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (dtuo *DocumentTypeUpdateOne) SetNillableMetadataID(id *string) *DocumentTypeUpdateOne {
	if id != nil {
		dtuo = dtuo.SetMetadataID(*id)
	}
	return dtuo
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (dtuo *DocumentTypeUpdateOne) SetMetadata(m *Metadata) *DocumentTypeUpdateOne {
	return dtuo.SetMetadataID(m.ID)
}

// Mutation returns the DocumentTypeMutation object of the builder.
func (dtuo *DocumentTypeUpdateOne) Mutation() *DocumentTypeMutation {
	return dtuo.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (dtuo *DocumentTypeUpdateOne) ClearMetadata() *DocumentTypeUpdateOne {
	dtuo.mutation.ClearMetadata()
	return dtuo
}

// Where appends a list predicates to the DocumentTypeUpdate builder.
func (dtuo *DocumentTypeUpdateOne) Where(ps ...predicate.DocumentType) *DocumentTypeUpdateOne {
	dtuo.mutation.Where(ps...)
	return dtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dtuo *DocumentTypeUpdateOne) Select(field string, fields ...string) *DocumentTypeUpdateOne {
	dtuo.fields = append([]string{field}, fields...)
	return dtuo
}

// Save executes the query and returns the updated DocumentType entity.
func (dtuo *DocumentTypeUpdateOne) Save(ctx context.Context) (*DocumentType, error) {
	return withHooks(ctx, dtuo.sqlSave, dtuo.mutation, dtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dtuo *DocumentTypeUpdateOne) SaveX(ctx context.Context) *DocumentType {
	node, err := dtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dtuo *DocumentTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := dtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtuo *DocumentTypeUpdateOne) ExecX(ctx context.Context) {
	if err := dtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtuo *DocumentTypeUpdateOne) check() error {
	if v, ok := dtuo.mutation.GetType(); ok {
		if err := documenttype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DocumentType.type": %w`, err)}
		}
	}
	return nil
}

func (dtuo *DocumentTypeUpdateOne) sqlSave(ctx context.Context) (_node *DocumentType, err error) {
	if err := dtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(documenttype.Table, documenttype.Columns, sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt))
	id, ok := dtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DocumentType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, documenttype.FieldID)
		for _, f := range fields {
			if !documenttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != documenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtuo.mutation.GetType(); ok {
		_spec.SetField(documenttype.FieldType, field.TypeEnum, value)
	}
	if dtuo.mutation.TypeCleared() {
		_spec.ClearField(documenttype.FieldType, field.TypeEnum)
	}
	if value, ok := dtuo.mutation.Name(); ok {
		_spec.SetField(documenttype.FieldName, field.TypeString, value)
	}
	if dtuo.mutation.NameCleared() {
		_spec.ClearField(documenttype.FieldName, field.TypeString)
	}
	if value, ok := dtuo.mutation.Description(); ok {
		_spec.SetField(documenttype.FieldDescription, field.TypeString, value)
	}
	if dtuo.mutation.DescriptionCleared() {
		_spec.ClearField(documenttype.FieldDescription, field.TypeString)
	}
	if dtuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   documenttype.MetadataTable,
			Columns: []string{documenttype.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtuo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   documenttype.MetadataTable,
			Columns: []string{documenttype.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DocumentType{config: dtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{documenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dtuo.mutation.done = true
	return _node, nil
}
