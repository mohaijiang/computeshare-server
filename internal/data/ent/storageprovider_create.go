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
	"github.com/mohaijiang/computeshare-server/internal/data/ent/storageprovider"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
)

// StorageProviderCreate is the builder for creating a StorageProvider entity.
type StorageProviderCreate struct {
	config
	mutation *StorageProviderMutation
	hooks    []Hook
}

// SetAgentID sets the "agent_id" field.
func (spc *StorageProviderCreate) SetAgentID(u uuid.UUID) *StorageProviderCreate {
	spc.mutation.SetAgentID(u)
	return spc
}

// SetStatus sets the "status" field.
func (spc *StorageProviderCreate) SetStatus(cps consts.StorageProviderStatus) *StorageProviderCreate {
	spc.mutation.SetStatus(cps)
	return spc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (spc *StorageProviderCreate) SetNillableStatus(cps *consts.StorageProviderStatus) *StorageProviderCreate {
	if cps != nil {
		spc.SetStatus(*cps)
	}
	return spc
}

// SetMasterServer sets the "master_server" field.
func (spc *StorageProviderCreate) SetMasterServer(s string) *StorageProviderCreate {
	spc.mutation.SetMasterServer(s)
	return spc
}

// SetPublicIP sets the "public_ip" field.
func (spc *StorageProviderCreate) SetPublicIP(s string) *StorageProviderCreate {
	spc.mutation.SetPublicIP(s)
	return spc
}

// SetPublicPort sets the "public_port" field.
func (spc *StorageProviderCreate) SetPublicPort(i int32) *StorageProviderCreate {
	spc.mutation.SetPublicPort(i)
	return spc
}

// SetGrpcPort sets the "grpc_port" field.
func (spc *StorageProviderCreate) SetGrpcPort(i int32) *StorageProviderCreate {
	spc.mutation.SetGrpcPort(i)
	return spc
}

// SetCreatedTime sets the "created_time" field.
func (spc *StorageProviderCreate) SetCreatedTime(t time.Time) *StorageProviderCreate {
	spc.mutation.SetCreatedTime(t)
	return spc
}

// SetID sets the "id" field.
func (spc *StorageProviderCreate) SetID(u uuid.UUID) *StorageProviderCreate {
	spc.mutation.SetID(u)
	return spc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (spc *StorageProviderCreate) SetNillableID(u *uuid.UUID) *StorageProviderCreate {
	if u != nil {
		spc.SetID(*u)
	}
	return spc
}

// Mutation returns the StorageProviderMutation object of the builder.
func (spc *StorageProviderCreate) Mutation() *StorageProviderMutation {
	return spc.mutation
}

// Save creates the StorageProvider in the database.
func (spc *StorageProviderCreate) Save(ctx context.Context) (*StorageProvider, error) {
	spc.defaults()
	return withHooks(ctx, spc.sqlSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *StorageProviderCreate) SaveX(ctx context.Context) *StorageProvider {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *StorageProviderCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *StorageProviderCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spc *StorageProviderCreate) defaults() {
	if _, ok := spc.mutation.Status(); !ok {
		v := storageprovider.DefaultStatus
		spc.mutation.SetStatus(v)
	}
	if _, ok := spc.mutation.ID(); !ok {
		v := storageprovider.DefaultID()
		spc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *StorageProviderCreate) check() error {
	if _, ok := spc.mutation.AgentID(); !ok {
		return &ValidationError{Name: "agent_id", err: errors.New(`ent: missing required field "StorageProvider.agent_id"`)}
	}
	if _, ok := spc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "StorageProvider.status"`)}
	}
	if _, ok := spc.mutation.MasterServer(); !ok {
		return &ValidationError{Name: "master_server", err: errors.New(`ent: missing required field "StorageProvider.master_server"`)}
	}
	if v, ok := spc.mutation.MasterServer(); ok {
		if err := storageprovider.MasterServerValidator(v); err != nil {
			return &ValidationError{Name: "master_server", err: fmt.Errorf(`ent: validator failed for field "StorageProvider.master_server": %w`, err)}
		}
	}
	if _, ok := spc.mutation.PublicIP(); !ok {
		return &ValidationError{Name: "public_ip", err: errors.New(`ent: missing required field "StorageProvider.public_ip"`)}
	}
	if v, ok := spc.mutation.PublicIP(); ok {
		if err := storageprovider.PublicIPValidator(v); err != nil {
			return &ValidationError{Name: "public_ip", err: fmt.Errorf(`ent: validator failed for field "StorageProvider.public_ip": %w`, err)}
		}
	}
	if _, ok := spc.mutation.PublicPort(); !ok {
		return &ValidationError{Name: "public_port", err: errors.New(`ent: missing required field "StorageProvider.public_port"`)}
	}
	if _, ok := spc.mutation.GrpcPort(); !ok {
		return &ValidationError{Name: "grpc_port", err: errors.New(`ent: missing required field "StorageProvider.grpc_port"`)}
	}
	if _, ok := spc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "StorageProvider.created_time"`)}
	}
	return nil
}

func (spc *StorageProviderCreate) sqlSave(ctx context.Context) (*StorageProvider, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
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
	spc.mutation.id = &_node.ID
	spc.mutation.done = true
	return _node, nil
}

func (spc *StorageProviderCreate) createSpec() (*StorageProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &StorageProvider{config: spc.config}
		_spec = sqlgraph.NewCreateSpec(storageprovider.Table, sqlgraph.NewFieldSpec(storageprovider.FieldID, field.TypeUUID))
	)
	if id, ok := spc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := spc.mutation.AgentID(); ok {
		_spec.SetField(storageprovider.FieldAgentID, field.TypeUUID, value)
		_node.AgentID = value
	}
	if value, ok := spc.mutation.Status(); ok {
		_spec.SetField(storageprovider.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := spc.mutation.MasterServer(); ok {
		_spec.SetField(storageprovider.FieldMasterServer, field.TypeString, value)
		_node.MasterServer = value
	}
	if value, ok := spc.mutation.PublicIP(); ok {
		_spec.SetField(storageprovider.FieldPublicIP, field.TypeString, value)
		_node.PublicIP = value
	}
	if value, ok := spc.mutation.PublicPort(); ok {
		_spec.SetField(storageprovider.FieldPublicPort, field.TypeInt32, value)
		_node.PublicPort = value
	}
	if value, ok := spc.mutation.GrpcPort(); ok {
		_spec.SetField(storageprovider.FieldGrpcPort, field.TypeInt32, value)
		_node.GrpcPort = value
	}
	if value, ok := spc.mutation.CreatedTime(); ok {
		_spec.SetField(storageprovider.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = value
	}
	return _node, _spec
}

// StorageProviderCreateBulk is the builder for creating many StorageProvider entities in bulk.
type StorageProviderCreateBulk struct {
	config
	builders []*StorageProviderCreate
}

// Save creates the StorageProvider entities in the database.
func (spcb *StorageProviderCreateBulk) Save(ctx context.Context) ([]*StorageProvider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*StorageProvider, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StorageProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *StorageProviderCreateBulk) SaveX(ctx context.Context) []*StorageProvider {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *StorageProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *StorageProviderCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}
