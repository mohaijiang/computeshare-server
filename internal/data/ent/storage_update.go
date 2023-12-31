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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/storage"
)

// StorageUpdate is the builder for updating Storage entities.
type StorageUpdate struct {
	config
	hooks    []Hook
	mutation *StorageMutation
}

// Where appends a list predicates to the StorageUpdate builder.
func (su *StorageUpdate) Where(ps ...predicate.Storage) *StorageUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetOwner sets the "owner" field.
func (su *StorageUpdate) SetOwner(s string) *StorageUpdate {
	su.mutation.SetOwner(s)
	return su
}

// SetType sets the "type" field.
func (su *StorageUpdate) SetType(i int32) *StorageUpdate {
	su.mutation.ResetType()
	su.mutation.SetType(i)
	return su
}

// SetNillableType sets the "type" field if the given value is not nil.
func (su *StorageUpdate) SetNillableType(i *int32) *StorageUpdate {
	if i != nil {
		su.SetType(*i)
	}
	return su
}

// AddType adds i to the "type" field.
func (su *StorageUpdate) AddType(i int32) *StorageUpdate {
	su.mutation.AddType(i)
	return su
}

// SetName sets the "name" field.
func (su *StorageUpdate) SetName(s string) *StorageUpdate {
	su.mutation.SetName(s)
	return su
}

// SetCid sets the "cid" field.
func (su *StorageUpdate) SetCid(s string) *StorageUpdate {
	su.mutation.SetCid(s)
	return su
}

// SetSize sets the "size" field.
func (su *StorageUpdate) SetSize(i int32) *StorageUpdate {
	su.mutation.ResetSize()
	su.mutation.SetSize(i)
	return su
}

// AddSize adds i to the "size" field.
func (su *StorageUpdate) AddSize(i int32) *StorageUpdate {
	su.mutation.AddSize(i)
	return su
}

// SetLastModify sets the "last_modify" field.
func (su *StorageUpdate) SetLastModify(t time.Time) *StorageUpdate {
	su.mutation.SetLastModify(t)
	return su
}

// SetNillableLastModify sets the "last_modify" field if the given value is not nil.
func (su *StorageUpdate) SetNillableLastModify(t *time.Time) *StorageUpdate {
	if t != nil {
		su.SetLastModify(*t)
	}
	return su
}

// SetParentID sets the "parent_id" field.
func (su *StorageUpdate) SetParentID(s string) *StorageUpdate {
	su.mutation.SetParentID(s)
	return su
}

// Mutation returns the StorageMutation object of the builder.
func (su *StorageUpdate) Mutation() *StorageMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StorageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StorageUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StorageUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StorageUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StorageUpdate) check() error {
	if v, ok := su.mutation.Owner(); ok {
		if err := storage.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "Storage.owner": %w`, err)}
		}
	}
	if v, ok := su.mutation.Name(); ok {
		if err := storage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Storage.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.Cid(); ok {
		if err := storage.CidValidator(v); err != nil {
			return &ValidationError{Name: "cid", err: fmt.Errorf(`ent: validator failed for field "Storage.cid": %w`, err)}
		}
	}
	if v, ok := su.mutation.ParentID(); ok {
		if err := storage.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf(`ent: validator failed for field "Storage.parent_id": %w`, err)}
		}
	}
	return nil
}

func (su *StorageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(storage.Table, storage.Columns, sqlgraph.NewFieldSpec(storage.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Owner(); ok {
		_spec.SetField(storage.FieldOwner, field.TypeString, value)
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.SetField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedType(); ok {
		_spec.AddField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(storage.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Cid(); ok {
		_spec.SetField(storage.FieldCid, field.TypeString, value)
	}
	if value, ok := su.mutation.Size(); ok {
		_spec.SetField(storage.FieldSize, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedSize(); ok {
		_spec.AddField(storage.FieldSize, field.TypeInt32, value)
	}
	if value, ok := su.mutation.LastModify(); ok {
		_spec.SetField(storage.FieldLastModify, field.TypeTime, value)
	}
	if value, ok := su.mutation.ParentID(); ok {
		_spec.SetField(storage.FieldParentID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StorageUpdateOne is the builder for updating a single Storage entity.
type StorageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StorageMutation
}

// SetOwner sets the "owner" field.
func (suo *StorageUpdateOne) SetOwner(s string) *StorageUpdateOne {
	suo.mutation.SetOwner(s)
	return suo
}

// SetType sets the "type" field.
func (suo *StorageUpdateOne) SetType(i int32) *StorageUpdateOne {
	suo.mutation.ResetType()
	suo.mutation.SetType(i)
	return suo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableType(i *int32) *StorageUpdateOne {
	if i != nil {
		suo.SetType(*i)
	}
	return suo
}

// AddType adds i to the "type" field.
func (suo *StorageUpdateOne) AddType(i int32) *StorageUpdateOne {
	suo.mutation.AddType(i)
	return suo
}

// SetName sets the "name" field.
func (suo *StorageUpdateOne) SetName(s string) *StorageUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetCid sets the "cid" field.
func (suo *StorageUpdateOne) SetCid(s string) *StorageUpdateOne {
	suo.mutation.SetCid(s)
	return suo
}

// SetSize sets the "size" field.
func (suo *StorageUpdateOne) SetSize(i int32) *StorageUpdateOne {
	suo.mutation.ResetSize()
	suo.mutation.SetSize(i)
	return suo
}

// AddSize adds i to the "size" field.
func (suo *StorageUpdateOne) AddSize(i int32) *StorageUpdateOne {
	suo.mutation.AddSize(i)
	return suo
}

// SetLastModify sets the "last_modify" field.
func (suo *StorageUpdateOne) SetLastModify(t time.Time) *StorageUpdateOne {
	suo.mutation.SetLastModify(t)
	return suo
}

// SetNillableLastModify sets the "last_modify" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableLastModify(t *time.Time) *StorageUpdateOne {
	if t != nil {
		suo.SetLastModify(*t)
	}
	return suo
}

// SetParentID sets the "parent_id" field.
func (suo *StorageUpdateOne) SetParentID(s string) *StorageUpdateOne {
	suo.mutation.SetParentID(s)
	return suo
}

// Mutation returns the StorageMutation object of the builder.
func (suo *StorageUpdateOne) Mutation() *StorageMutation {
	return suo.mutation
}

// Where appends a list predicates to the StorageUpdate builder.
func (suo *StorageUpdateOne) Where(ps ...predicate.Storage) *StorageUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StorageUpdateOne) Select(field string, fields ...string) *StorageUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Storage entity.
func (suo *StorageUpdateOne) Save(ctx context.Context) (*Storage, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StorageUpdateOne) SaveX(ctx context.Context) *Storage {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StorageUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StorageUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StorageUpdateOne) check() error {
	if v, ok := suo.mutation.Owner(); ok {
		if err := storage.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "Storage.owner": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Name(); ok {
		if err := storage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Storage.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Cid(); ok {
		if err := storage.CidValidator(v); err != nil {
			return &ValidationError{Name: "cid", err: fmt.Errorf(`ent: validator failed for field "Storage.cid": %w`, err)}
		}
	}
	if v, ok := suo.mutation.ParentID(); ok {
		if err := storage.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf(`ent: validator failed for field "Storage.parent_id": %w`, err)}
		}
	}
	return nil
}

func (suo *StorageUpdateOne) sqlSave(ctx context.Context) (_node *Storage, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(storage.Table, storage.Columns, sqlgraph.NewFieldSpec(storage.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Storage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storage.FieldID)
		for _, f := range fields {
			if !storage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Owner(); ok {
		_spec.SetField(storage.FieldOwner, field.TypeString, value)
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.SetField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedType(); ok {
		_spec.AddField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(storage.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Cid(); ok {
		_spec.SetField(storage.FieldCid, field.TypeString, value)
	}
	if value, ok := suo.mutation.Size(); ok {
		_spec.SetField(storage.FieldSize, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedSize(); ok {
		_spec.AddField(storage.FieldSize, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.LastModify(); ok {
		_spec.SetField(storage.FieldLastModify, field.TypeTime, value)
	}
	if value, ok := suo.mutation.ParentID(); ok {
		_spec.SetField(storage.FieldParentID, field.TypeString, value)
	}
	_node = &Storage{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
