// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycletransaction"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// CycleTransactionDelete is the builder for deleting a CycleTransaction entity.
type CycleTransactionDelete struct {
	config
	hooks    []Hook
	mutation *CycleTransactionMutation
}

// Where appends a list predicates to the CycleTransactionDelete builder.
func (ctd *CycleTransactionDelete) Where(ps ...predicate.CycleTransaction) *CycleTransactionDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *CycleTransactionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ctd.sqlExec, ctd.mutation, ctd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *CycleTransactionDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *CycleTransactionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cycletransaction.Table, sqlgraph.NewFieldSpec(cycletransaction.FieldID, field.TypeUUID))
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ctd.mutation.done = true
	return affected, err
}

// CycleTransactionDeleteOne is the builder for deleting a single CycleTransaction entity.
type CycleTransactionDeleteOne struct {
	ctd *CycleTransactionDelete
}

// Where appends a list predicates to the CycleTransactionDelete builder.
func (ctdo *CycleTransactionDeleteOne) Where(ps ...predicate.CycleTransaction) *CycleTransactionDeleteOne {
	ctdo.ctd.mutation.Where(ps...)
	return ctdo
}

// Exec executes the deletion query.
func (ctdo *CycleTransactionDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cycletransaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *CycleTransactionDeleteOne) ExecX(ctx context.Context) {
	if err := ctdo.Exec(ctx); err != nil {
		panic(err)
	}
}