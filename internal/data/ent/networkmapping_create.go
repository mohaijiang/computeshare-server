// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/networkmapping"
)

// NetworkMappingCreate is the builder for creating a NetworkMapping entity.
type NetworkMappingCreate struct {
	config
	mutation *NetworkMappingMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (nmc *NetworkMappingCreate) SetName(s string) *NetworkMappingCreate {
	nmc.mutation.SetName(s)
	return nmc
}

// SetProtocol sets the "protocol" field.
func (nmc *NetworkMappingCreate) SetProtocol(s string) *NetworkMappingCreate {
	nmc.mutation.SetProtocol(s)
	return nmc
}

// SetNillableProtocol sets the "protocol" field if the given value is not nil.
func (nmc *NetworkMappingCreate) SetNillableProtocol(s *string) *NetworkMappingCreate {
	if s != nil {
		nmc.SetProtocol(*s)
	}
	return nmc
}

// SetFkGatewayID sets the "fk_gateway_id" field.
func (nmc *NetworkMappingCreate) SetFkGatewayID(u uuid.UUID) *NetworkMappingCreate {
	nmc.mutation.SetFkGatewayID(u)
	return nmc
}

// SetGatewayPort sets the "gateway_port" field.
func (nmc *NetworkMappingCreate) SetGatewayPort(i int32) *NetworkMappingCreate {
	nmc.mutation.SetGatewayPort(i)
	return nmc
}

// SetGatewayIP sets the "gateway_ip" field.
func (nmc *NetworkMappingCreate) SetGatewayIP(s string) *NetworkMappingCreate {
	nmc.mutation.SetGatewayIP(s)
	return nmc
}

// SetComputerPort sets the "computer_port" field.
func (nmc *NetworkMappingCreate) SetComputerPort(i int32) *NetworkMappingCreate {
	nmc.mutation.SetComputerPort(i)
	return nmc
}

// SetStatus sets the "status" field.
func (nmc *NetworkMappingCreate) SetStatus(i int) *NetworkMappingCreate {
	nmc.mutation.SetStatus(i)
	return nmc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nmc *NetworkMappingCreate) SetNillableStatus(i *int) *NetworkMappingCreate {
	if i != nil {
		nmc.SetStatus(*i)
	}
	return nmc
}

// SetFkComputerID sets the "fk_computer_id" field.
func (nmc *NetworkMappingCreate) SetFkComputerID(u uuid.UUID) *NetworkMappingCreate {
	nmc.mutation.SetFkComputerID(u)
	return nmc
}

// SetFkUserID sets the "fk_user_id" field.
func (nmc *NetworkMappingCreate) SetFkUserID(u uuid.UUID) *NetworkMappingCreate {
	nmc.mutation.SetFkUserID(u)
	return nmc
}

// SetID sets the "id" field.
func (nmc *NetworkMappingCreate) SetID(u uuid.UUID) *NetworkMappingCreate {
	nmc.mutation.SetID(u)
	return nmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nmc *NetworkMappingCreate) SetNillableID(u *uuid.UUID) *NetworkMappingCreate {
	if u != nil {
		nmc.SetID(*u)
	}
	return nmc
}

// Mutation returns the NetworkMappingMutation object of the builder.
func (nmc *NetworkMappingCreate) Mutation() *NetworkMappingMutation {
	return nmc.mutation
}

// Save creates the NetworkMapping in the database.
func (nmc *NetworkMappingCreate) Save(ctx context.Context) (*NetworkMapping, error) {
	nmc.defaults()
	return withHooks(ctx, nmc.sqlSave, nmc.mutation, nmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nmc *NetworkMappingCreate) SaveX(ctx context.Context) *NetworkMapping {
	v, err := nmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nmc *NetworkMappingCreate) Exec(ctx context.Context) error {
	_, err := nmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmc *NetworkMappingCreate) ExecX(ctx context.Context) {
	if err := nmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nmc *NetworkMappingCreate) defaults() {
	if _, ok := nmc.mutation.Protocol(); !ok {
		v := networkmapping.DefaultProtocol
		nmc.mutation.SetProtocol(v)
	}
	if _, ok := nmc.mutation.Status(); !ok {
		v := networkmapping.DefaultStatus
		nmc.mutation.SetStatus(v)
	}
	if _, ok := nmc.mutation.ID(); !ok {
		v := networkmapping.DefaultID()
		nmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nmc *NetworkMappingCreate) check() error {
	if _, ok := nmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "NetworkMapping.name"`)}
	}
	if v, ok := nmc.mutation.Name(); ok {
		if err := networkmapping.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "NetworkMapping.name": %w`, err)}
		}
	}
	if _, ok := nmc.mutation.Protocol(); !ok {
		return &ValidationError{Name: "protocol", err: errors.New(`ent: missing required field "NetworkMapping.protocol"`)}
	}
	if v, ok := nmc.mutation.Protocol(); ok {
		if err := networkmapping.ProtocolValidator(v); err != nil {
			return &ValidationError{Name: "protocol", err: fmt.Errorf(`ent: validator failed for field "NetworkMapping.protocol": %w`, err)}
		}
	}
	if _, ok := nmc.mutation.FkGatewayID(); !ok {
		return &ValidationError{Name: "fk_gateway_id", err: errors.New(`ent: missing required field "NetworkMapping.fk_gateway_id"`)}
	}
	if _, ok := nmc.mutation.GatewayPort(); !ok {
		return &ValidationError{Name: "gateway_port", err: errors.New(`ent: missing required field "NetworkMapping.gateway_port"`)}
	}
	if _, ok := nmc.mutation.GatewayIP(); !ok {
		return &ValidationError{Name: "gateway_ip", err: errors.New(`ent: missing required field "NetworkMapping.gateway_ip"`)}
	}
	if _, ok := nmc.mutation.ComputerPort(); !ok {
		return &ValidationError{Name: "computer_port", err: errors.New(`ent: missing required field "NetworkMapping.computer_port"`)}
	}
	if _, ok := nmc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "NetworkMapping.status"`)}
	}
	if _, ok := nmc.mutation.FkComputerID(); !ok {
		return &ValidationError{Name: "fk_computer_id", err: errors.New(`ent: missing required field "NetworkMapping.fk_computer_id"`)}
	}
	if _, ok := nmc.mutation.FkUserID(); !ok {
		return &ValidationError{Name: "fk_user_id", err: errors.New(`ent: missing required field "NetworkMapping.fk_user_id"`)}
	}
	return nil
}

func (nmc *NetworkMappingCreate) sqlSave(ctx context.Context) (*NetworkMapping, error) {
	if err := nmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nmc.driver, _spec); err != nil {
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
	nmc.mutation.id = &_node.ID
	nmc.mutation.done = true
	return _node, nil
}

func (nmc *NetworkMappingCreate) createSpec() (*NetworkMapping, *sqlgraph.CreateSpec) {
	var (
		_node = &NetworkMapping{config: nmc.config}
		_spec = sqlgraph.NewCreateSpec(networkmapping.Table, sqlgraph.NewFieldSpec(networkmapping.FieldID, field.TypeUUID))
	)
	if id, ok := nmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nmc.mutation.Name(); ok {
		_spec.SetField(networkmapping.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := nmc.mutation.Protocol(); ok {
		_spec.SetField(networkmapping.FieldProtocol, field.TypeString, value)
		_node.Protocol = value
	}
	if value, ok := nmc.mutation.FkGatewayID(); ok {
		_spec.SetField(networkmapping.FieldFkGatewayID, field.TypeUUID, value)
		_node.FkGatewayID = value
	}
	if value, ok := nmc.mutation.GatewayPort(); ok {
		_spec.SetField(networkmapping.FieldGatewayPort, field.TypeInt32, value)
		_node.GatewayPort = value
	}
	if value, ok := nmc.mutation.GatewayIP(); ok {
		_spec.SetField(networkmapping.FieldGatewayIP, field.TypeString, value)
		_node.GatewayIP = value
	}
	if value, ok := nmc.mutation.ComputerPort(); ok {
		_spec.SetField(networkmapping.FieldComputerPort, field.TypeInt32, value)
		_node.ComputerPort = value
	}
	if value, ok := nmc.mutation.Status(); ok {
		_spec.SetField(networkmapping.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := nmc.mutation.FkComputerID(); ok {
		_spec.SetField(networkmapping.FieldFkComputerID, field.TypeUUID, value)
		_node.FkComputerID = value
	}
	if value, ok := nmc.mutation.FkUserID(); ok {
		_spec.SetField(networkmapping.FieldFkUserID, field.TypeUUID, value)
		_node.FkUserID = value
	}
	return _node, _spec
}

// NetworkMappingCreateBulk is the builder for creating many NetworkMapping entities in bulk.
type NetworkMappingCreateBulk struct {
	config
	builders []*NetworkMappingCreate
}

// Save creates the NetworkMapping entities in the database.
func (nmcb *NetworkMappingCreateBulk) Save(ctx context.Context) ([]*NetworkMapping, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nmcb.builders))
	nodes := make([]*NetworkMapping, len(nmcb.builders))
	mutators := make([]Mutator, len(nmcb.builders))
	for i := range nmcb.builders {
		func(i int, root context.Context) {
			builder := nmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NetworkMappingMutation)
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
					_, err = mutators[i+1].Mutate(root, nmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nmcb *NetworkMappingCreateBulk) SaveX(ctx context.Context) []*NetworkMapping {
	v, err := nmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nmcb *NetworkMappingCreateBulk) Exec(ctx context.Context) error {
	_, err := nmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmcb *NetworkMappingCreateBulk) ExecX(ctx context.Context) {
	if err := nmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
