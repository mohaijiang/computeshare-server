// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycle"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// CycleUpdate is the builder for updating Cycle entities.
type CycleUpdate struct {
	config
	hooks    []Hook
	mutation *CycleMutation
}

// Where appends a list predicates to the CycleUpdate builder.
func (cu *CycleUpdate) Where(ps ...predicate.Cycle) *CycleUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetFkUserID sets the "fk_user_id" field.
func (cu *CycleUpdate) SetFkUserID(u uuid.UUID) *CycleUpdate {
	cu.mutation.SetFkUserID(u)
	return cu
}

// SetCycle sets the "cycle" field.
func (cu *CycleUpdate) SetCycle(f float64) *CycleUpdate {
	cu.mutation.ResetCycle()
	cu.mutation.SetCycle(f)
	return cu
}

// AddCycle adds f to the "cycle" field.
func (cu *CycleUpdate) AddCycle(f float64) *CycleUpdate {
	cu.mutation.AddCycle(f)
	return cu
}

// SetCreateTime sets the "create_time" field.
func (cu *CycleUpdate) SetCreateTime(t time.Time) *CycleUpdate {
	cu.mutation.SetCreateTime(t)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *CycleUpdate) SetUpdateTime(t time.Time) *CycleUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// Mutation returns the CycleMutation object of the builder.
func (cu *CycleUpdate) Mutation() *CycleMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CycleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CycleUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CycleUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CycleUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CycleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cycle.Table, cycle.Columns, sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.FkUserID(); ok {
		_spec.SetField(cycle.FieldFkUserID, field.TypeUUID, value)
	}
	if value, ok := cu.mutation.Cycle(); ok {
		_spec.SetField(cycle.FieldCycle, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedCycle(); ok {
		_spec.AddField(cycle.FieldCycle, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.CreateTime(); ok {
		_spec.SetField(cycle.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.SetField(cycle.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cycle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CycleUpdateOne is the builder for updating a single Cycle entity.
type CycleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CycleMutation
}

// SetFkUserID sets the "fk_user_id" field.
func (cuo *CycleUpdateOne) SetFkUserID(u uuid.UUID) *CycleUpdateOne {
	cuo.mutation.SetFkUserID(u)
	return cuo
}

// SetCycle sets the "cycle" field.
func (cuo *CycleUpdateOne) SetCycle(f float64) *CycleUpdateOne {
	cuo.mutation.ResetCycle()
	cuo.mutation.SetCycle(f)
	return cuo
}

// AddCycle adds f to the "cycle" field.
func (cuo *CycleUpdateOne) AddCycle(f float64) *CycleUpdateOne {
	cuo.mutation.AddCycle(f)
	return cuo
}

// SetCreateTime sets the "create_time" field.
func (cuo *CycleUpdateOne) SetCreateTime(t time.Time) *CycleUpdateOne {
	cuo.mutation.SetCreateTime(t)
	return cuo
}

// SetUpdateTime sets the "update_time" field.
func (cuo *CycleUpdateOne) SetUpdateTime(t time.Time) *CycleUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// Mutation returns the CycleMutation object of the builder.
func (cuo *CycleUpdateOne) Mutation() *CycleMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CycleUpdate builder.
func (cuo *CycleUpdateOne) Where(ps ...predicate.Cycle) *CycleUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CycleUpdateOne) Select(field string, fields ...string) *CycleUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cycle entity.
func (cuo *CycleUpdateOne) Save(ctx context.Context) (*Cycle, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CycleUpdateOne) SaveX(ctx context.Context) *Cycle {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CycleUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CycleUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CycleUpdateOne) sqlSave(ctx context.Context) (_node *Cycle, err error) {
	_spec := sqlgraph.NewUpdateSpec(cycle.Table, cycle.Columns, sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cycle.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cycle.FieldID)
		for _, f := range fields {
			if !cycle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cycle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.FkUserID(); ok {
		_spec.SetField(cycle.FieldFkUserID, field.TypeUUID, value)
	}
	if value, ok := cuo.mutation.Cycle(); ok {
		_spec.SetField(cycle.FieldCycle, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedCycle(); ok {
		_spec.AddField(cycle.FieldCycle, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.CreateTime(); ok {
		_spec.SetField(cycle.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.SetField(cycle.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &Cycle{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cycle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}