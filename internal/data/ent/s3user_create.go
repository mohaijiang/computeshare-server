// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3user"
)

// S3UserCreate is the builder for creating a S3User entity.
type S3UserCreate struct {
	config
	mutation *S3UserMutation
	hooks    []Hook
}

// SetFkUserID sets the "fk_user_id" field.
func (sc *S3UserCreate) SetFkUserID(u uuid.UUID) *S3UserCreate {
	sc.mutation.SetFkUserID(u)
	return sc
}

// SetType sets the "type" field.
func (sc *S3UserCreate) SetType(i int8) *S3UserCreate {
	sc.mutation.SetType(i)
	return sc
}

// SetAccessKey sets the "access_key" field.
func (sc *S3UserCreate) SetAccessKey(s string) *S3UserCreate {
	sc.mutation.SetAccessKey(s)
	return sc
}

// SetSecretKey sets the "secret_key" field.
func (sc *S3UserCreate) SetSecretKey(s string) *S3UserCreate {
	sc.mutation.SetSecretKey(s)
	return sc
}

// SetCreateTime sets the "create_time" field.
func (sc *S3UserCreate) SetCreateTime(t time.Time) *S3UserCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *S3UserCreate) SetNillableCreateTime(t *time.Time) *S3UserCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *S3UserCreate) SetUpdateTime(t time.Time) *S3UserCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *S3UserCreate) SetNillableUpdateTime(t *time.Time) *S3UserCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *S3UserCreate) SetID(u uuid.UUID) *S3UserCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *S3UserCreate) SetNillableID(u *uuid.UUID) *S3UserCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the S3UserMutation object of the builder.
func (sc *S3UserCreate) Mutation() *S3UserMutation {
	return sc.mutation
}

// Save creates the S3User in the database.
func (sc *S3UserCreate) Save(ctx context.Context) (*S3User, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *S3UserCreate) SaveX(ctx context.Context) *S3User {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *S3UserCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *S3UserCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *S3UserCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := s3user.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := s3user.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := s3user.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *S3UserCreate) check() error {
	if _, ok := sc.mutation.FkUserID(); !ok {
		return &ValidationError{Name: "fk_user_id", err: errors.New(`ent: missing required field "S3User.fk_user_id"`)}
	}
	if _, ok := sc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "S3User.type"`)}
	}
	if _, ok := sc.mutation.AccessKey(); !ok {
		return &ValidationError{Name: "access_key", err: errors.New(`ent: missing required field "S3User.access_key"`)}
	}
	if v, ok := sc.mutation.AccessKey(); ok {
		if err := s3user.AccessKeyValidator(v); err != nil {
			return &ValidationError{Name: "access_key", err: fmt.Errorf(`ent: validator failed for field "S3User.access_key": %w`, err)}
		}
	}
	if _, ok := sc.mutation.SecretKey(); !ok {
		return &ValidationError{Name: "secret_key", err: errors.New(`ent: missing required field "S3User.secret_key"`)}
	}
	if v, ok := sc.mutation.SecretKey(); ok {
		if err := s3user.SecretKeyValidator(v); err != nil {
			return &ValidationError{Name: "secret_key", err: fmt.Errorf(`ent: validator failed for field "S3User.secret_key": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "S3User.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "S3User.update_time"`)}
	}
	return nil
}

func (sc *S3UserCreate) sqlSave(ctx context.Context) (*S3User, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *S3UserCreate) createSpec() (*S3User, *sqlgraph.CreateSpec) {
	var (
		_node = &S3User{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(s3user.Table, sqlgraph.NewFieldSpec(s3user.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.FkUserID(); ok {
		_spec.SetField(s3user.FieldFkUserID, field.TypeUUID, value)
		_node.FkUserID = value
	}
	if value, ok := sc.mutation.GetType(); ok {
		_spec.SetField(s3user.FieldType, field.TypeInt8, value)
		_node.Type = value
	}
	if value, ok := sc.mutation.AccessKey(); ok {
		_spec.SetField(s3user.FieldAccessKey, field.TypeString, value)
		_node.AccessKey = value
	}
	if value, ok := sc.mutation.SecretKey(); ok {
		_spec.SetField(s3user.FieldSecretKey, field.TypeString, value)
		_node.SecretKey = value
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(s3user.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(s3user.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// S3UserCreateBulk is the builder for creating many S3User entities in bulk.
type S3UserCreateBulk struct {
	config
	builders []*S3UserCreate
}

// Save creates the S3User entities in the database.
func (scb *S3UserCreateBulk) Save(ctx context.Context) ([]*S3User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*S3User, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*S3UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *S3UserCreateBulk) SaveX(ctx context.Context) []*S3User {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *S3UserCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *S3UserCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
