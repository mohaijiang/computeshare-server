// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycleorder"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// CycleOrderDelete is the builder for deleting a CycleOrder entity.
type CycleOrderDelete struct {
	config
	hooks    []Hook
	mutation *CycleOrderMutation
}

// Where appends a list predicates to the CycleOrderDelete builder.
func (cod *CycleOrderDelete) Where(ps ...predicate.CycleOrder) *CycleOrderDelete {
	cod.mutation.Where(ps...)
	return cod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cod *CycleOrderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cod.sqlExec, cod.mutation, cod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cod *CycleOrderDelete) ExecX(ctx context.Context) int {
	n, err := cod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cod *CycleOrderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cycleorder.Table, sqlgraph.NewFieldSpec(cycleorder.FieldID, field.TypeUUID))
	if ps := cod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cod.mutation.done = true
	return affected, err
}

// CycleOrderDeleteOne is the builder for deleting a single CycleOrder entity.
type CycleOrderDeleteOne struct {
	cod *CycleOrderDelete
}

// Where appends a list predicates to the CycleOrderDelete builder.
func (codo *CycleOrderDeleteOne) Where(ps ...predicate.CycleOrder) *CycleOrderDeleteOne {
	codo.cod.mutation.Where(ps...)
	return codo
}

// Exec executes the deletion query.
func (codo *CycleOrderDeleteOne) Exec(ctx context.Context) error {
	n, err := codo.cod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cycleorder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (codo *CycleOrderDeleteOne) ExecX(ctx context.Context) {
	if err := codo.Exec(ctx); err != nil {
		panic(err)
	}
}
