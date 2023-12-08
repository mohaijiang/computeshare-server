// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/storageprovider"
)

// StorageProviderDelete is the builder for deleting a StorageProvider entity.
type StorageProviderDelete struct {
	config
	hooks    []Hook
	mutation *StorageProviderMutation
}

// Where appends a list predicates to the StorageProviderDelete builder.
func (spd *StorageProviderDelete) Where(ps ...predicate.StorageProvider) *StorageProviderDelete {
	spd.mutation.Where(ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *StorageProviderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, spd.sqlExec, spd.mutation, spd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *StorageProviderDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *StorageProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(storageprovider.Table, sqlgraph.NewFieldSpec(storageprovider.FieldID, field.TypeUUID))
	if ps := spd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	spd.mutation.done = true
	return affected, err
}

// StorageProviderDeleteOne is the builder for deleting a single StorageProvider entity.
type StorageProviderDeleteOne struct {
	spd *StorageProviderDelete
}

// Where appends a list predicates to the StorageProviderDelete builder.
func (spdo *StorageProviderDeleteOne) Where(ps ...predicate.StorageProvider) *StorageProviderDeleteOne {
	spdo.spd.mutation.Where(ps...)
	return spdo
}

// Exec executes the deletion query.
func (spdo *StorageProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{storageprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *StorageProviderDeleteOne) ExecX(ctx context.Context) {
	if err := spdo.Exec(ctx); err != nil {
		panic(err)
	}
}