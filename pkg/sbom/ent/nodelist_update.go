// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/pkg/sbom/ent/document"
	"github.com/bom-squad/protobom/pkg/sbom/ent/edge"
	"github.com/bom-squad/protobom/pkg/sbom/ent/node"
	"github.com/bom-squad/protobom/pkg/sbom/ent/nodelist"
	"github.com/bom-squad/protobom/pkg/sbom/ent/predicate"
)

// NodeListUpdate is the builder for updating NodeList entities.
type NodeListUpdate struct {
	config
	hooks    []Hook
	mutation *NodeListMutation
}

// Where appends a list predicates to the NodeListUpdate builder.
func (nlu *NodeListUpdate) Where(ps ...predicate.NodeList) *NodeListUpdate {
	nlu.mutation.Where(ps...)
	return nlu
}

// SetRootElements sets the "root_elements" field.
func (nlu *NodeListUpdate) SetRootElements(s string) *NodeListUpdate {
	nlu.mutation.SetRootElements(s)
	return nlu
}

// SetNillableRootElements sets the "root_elements" field if the given value is not nil.
func (nlu *NodeListUpdate) SetNillableRootElements(s *string) *NodeListUpdate {
	if s != nil {
		nlu.SetRootElements(*s)
	}
	return nlu
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (nlu *NodeListUpdate) AddNodeIDs(ids ...string) *NodeListUpdate {
	nlu.mutation.AddNodeIDs(ids...)
	return nlu
}

// AddNodes adds the "nodes" edges to the Node entity.
func (nlu *NodeListUpdate) AddNodes(n ...*Node) *NodeListUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nlu.AddNodeIDs(ids...)
}

// AddEdgeIDs adds the "edges" edge to the Edge entity by IDs.
func (nlu *NodeListUpdate) AddEdgeIDs(ids ...int) *NodeListUpdate {
	nlu.mutation.AddEdgeIDs(ids...)
	return nlu
}

// AddEdges adds the "edges" edges to the Edge entity.
func (nlu *NodeListUpdate) AddEdges(e ...*Edge) *NodeListUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nlu.AddEdgeIDs(ids...)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (nlu *NodeListUpdate) SetDocumentID(id int) *NodeListUpdate {
	nlu.mutation.SetDocumentID(id)
	return nlu
}

// SetDocument sets the "document" edge to the Document entity.
func (nlu *NodeListUpdate) SetDocument(d *Document) *NodeListUpdate {
	return nlu.SetDocumentID(d.ID)
}

// Mutation returns the NodeListMutation object of the builder.
func (nlu *NodeListUpdate) Mutation() *NodeListMutation {
	return nlu.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (nlu *NodeListUpdate) ClearNodes() *NodeListUpdate {
	nlu.mutation.ClearNodes()
	return nlu
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (nlu *NodeListUpdate) RemoveNodeIDs(ids ...string) *NodeListUpdate {
	nlu.mutation.RemoveNodeIDs(ids...)
	return nlu
}

// RemoveNodes removes "nodes" edges to Node entities.
func (nlu *NodeListUpdate) RemoveNodes(n ...*Node) *NodeListUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nlu.RemoveNodeIDs(ids...)
}

// ClearEdges clears all "edges" edges to the Edge entity.
func (nlu *NodeListUpdate) ClearEdges() *NodeListUpdate {
	nlu.mutation.ClearEdges()
	return nlu
}

// RemoveEdgeIDs removes the "edges" edge to Edge entities by IDs.
func (nlu *NodeListUpdate) RemoveEdgeIDs(ids ...int) *NodeListUpdate {
	nlu.mutation.RemoveEdgeIDs(ids...)
	return nlu
}

// RemoveEdges removes "edges" edges to Edge entities.
func (nlu *NodeListUpdate) RemoveEdges(e ...*Edge) *NodeListUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nlu.RemoveEdgeIDs(ids...)
}

// ClearDocument clears the "document" edge to the Document entity.
func (nlu *NodeListUpdate) ClearDocument() *NodeListUpdate {
	nlu.mutation.ClearDocument()
	return nlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nlu *NodeListUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nlu.sqlSave, nlu.mutation, nlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nlu *NodeListUpdate) SaveX(ctx context.Context) int {
	affected, err := nlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nlu *NodeListUpdate) Exec(ctx context.Context) error {
	_, err := nlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nlu *NodeListUpdate) ExecX(ctx context.Context) {
	if err := nlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nlu *NodeListUpdate) check() error {
	if _, ok := nlu.mutation.DocumentID(); nlu.mutation.DocumentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NodeList.document"`)
	}
	return nil
}

func (nlu *NodeListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nlu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(nodelist.Table, nodelist.Columns, sqlgraph.NewFieldSpec(nodelist.FieldID, field.TypeInt))
	if ps := nlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nlu.mutation.RootElements(); ok {
		_spec.SetField(nodelist.FieldRootElements, field.TypeString, value)
	}
	if nlu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nlu.mutation.RemovedNodesIDs(); len(nodes) > 0 && !nlu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nlu.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nlu.mutation.EdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.EdgesTable,
			Columns: []string{nodelist.EdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nlu.mutation.RemovedEdgesIDs(); len(nodes) > 0 && !nlu.mutation.EdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.EdgesTable,
			Columns: []string{nodelist.EdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nlu.mutation.EdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.EdgesTable,
			Columns: []string{nodelist.EdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nlu.mutation.DocumentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nodelist.DocumentTable,
			Columns: []string{nodelist.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nlu.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nodelist.DocumentTable,
			Columns: []string{nodelist.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodelist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nlu.mutation.done = true
	return n, nil
}

// NodeListUpdateOne is the builder for updating a single NodeList entity.
type NodeListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeListMutation
}

// SetRootElements sets the "root_elements" field.
func (nluo *NodeListUpdateOne) SetRootElements(s string) *NodeListUpdateOne {
	nluo.mutation.SetRootElements(s)
	return nluo
}

// SetNillableRootElements sets the "root_elements" field if the given value is not nil.
func (nluo *NodeListUpdateOne) SetNillableRootElements(s *string) *NodeListUpdateOne {
	if s != nil {
		nluo.SetRootElements(*s)
	}
	return nluo
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (nluo *NodeListUpdateOne) AddNodeIDs(ids ...string) *NodeListUpdateOne {
	nluo.mutation.AddNodeIDs(ids...)
	return nluo
}

// AddNodes adds the "nodes" edges to the Node entity.
func (nluo *NodeListUpdateOne) AddNodes(n ...*Node) *NodeListUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nluo.AddNodeIDs(ids...)
}

// AddEdgeIDs adds the "edges" edge to the Edge entity by IDs.
func (nluo *NodeListUpdateOne) AddEdgeIDs(ids ...int) *NodeListUpdateOne {
	nluo.mutation.AddEdgeIDs(ids...)
	return nluo
}

// AddEdges adds the "edges" edges to the Edge entity.
func (nluo *NodeListUpdateOne) AddEdges(e ...*Edge) *NodeListUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nluo.AddEdgeIDs(ids...)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (nluo *NodeListUpdateOne) SetDocumentID(id int) *NodeListUpdateOne {
	nluo.mutation.SetDocumentID(id)
	return nluo
}

// SetDocument sets the "document" edge to the Document entity.
func (nluo *NodeListUpdateOne) SetDocument(d *Document) *NodeListUpdateOne {
	return nluo.SetDocumentID(d.ID)
}

// Mutation returns the NodeListMutation object of the builder.
func (nluo *NodeListUpdateOne) Mutation() *NodeListMutation {
	return nluo.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (nluo *NodeListUpdateOne) ClearNodes() *NodeListUpdateOne {
	nluo.mutation.ClearNodes()
	return nluo
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (nluo *NodeListUpdateOne) RemoveNodeIDs(ids ...string) *NodeListUpdateOne {
	nluo.mutation.RemoveNodeIDs(ids...)
	return nluo
}

// RemoveNodes removes "nodes" edges to Node entities.
func (nluo *NodeListUpdateOne) RemoveNodes(n ...*Node) *NodeListUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nluo.RemoveNodeIDs(ids...)
}

// ClearEdges clears all "edges" edges to the Edge entity.
func (nluo *NodeListUpdateOne) ClearEdges() *NodeListUpdateOne {
	nluo.mutation.ClearEdges()
	return nluo
}

// RemoveEdgeIDs removes the "edges" edge to Edge entities by IDs.
func (nluo *NodeListUpdateOne) RemoveEdgeIDs(ids ...int) *NodeListUpdateOne {
	nluo.mutation.RemoveEdgeIDs(ids...)
	return nluo
}

// RemoveEdges removes "edges" edges to Edge entities.
func (nluo *NodeListUpdateOne) RemoveEdges(e ...*Edge) *NodeListUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nluo.RemoveEdgeIDs(ids...)
}

// ClearDocument clears the "document" edge to the Document entity.
func (nluo *NodeListUpdateOne) ClearDocument() *NodeListUpdateOne {
	nluo.mutation.ClearDocument()
	return nluo
}

// Where appends a list predicates to the NodeListUpdate builder.
func (nluo *NodeListUpdateOne) Where(ps ...predicate.NodeList) *NodeListUpdateOne {
	nluo.mutation.Where(ps...)
	return nluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nluo *NodeListUpdateOne) Select(field string, fields ...string) *NodeListUpdateOne {
	nluo.fields = append([]string{field}, fields...)
	return nluo
}

// Save executes the query and returns the updated NodeList entity.
func (nluo *NodeListUpdateOne) Save(ctx context.Context) (*NodeList, error) {
	return withHooks(ctx, nluo.sqlSave, nluo.mutation, nluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nluo *NodeListUpdateOne) SaveX(ctx context.Context) *NodeList {
	node, err := nluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nluo *NodeListUpdateOne) Exec(ctx context.Context) error {
	_, err := nluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nluo *NodeListUpdateOne) ExecX(ctx context.Context) {
	if err := nluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nluo *NodeListUpdateOne) check() error {
	if _, ok := nluo.mutation.DocumentID(); nluo.mutation.DocumentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NodeList.document"`)
	}
	return nil
}

func (nluo *NodeListUpdateOne) sqlSave(ctx context.Context) (_node *NodeList, err error) {
	if err := nluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(nodelist.Table, nodelist.Columns, sqlgraph.NewFieldSpec(nodelist.FieldID, field.TypeInt))
	id, ok := nluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NodeList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nodelist.FieldID)
		for _, f := range fields {
			if !nodelist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nodelist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nluo.mutation.RootElements(); ok {
		_spec.SetField(nodelist.FieldRootElements, field.TypeString, value)
	}
	if nluo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nluo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !nluo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nluo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nluo.mutation.EdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.EdgesTable,
			Columns: []string{nodelist.EdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nluo.mutation.RemovedEdgesIDs(); len(nodes) > 0 && !nluo.mutation.EdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.EdgesTable,
			Columns: []string{nodelist.EdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nluo.mutation.EdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.EdgesTable,
			Columns: []string{nodelist.EdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nluo.mutation.DocumentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nodelist.DocumentTable,
			Columns: []string{nodelist.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nluo.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   nodelist.DocumentTable,
			Columns: []string{nodelist.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NodeList{config: nluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodelist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nluo.mutation.done = true
	return _node, nil
}
