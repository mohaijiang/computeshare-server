// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/script"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/scriptexecutionrecord"
)

// ScriptExecutionRecordCreate is the builder for creating a ScriptExecutionRecord entity.
type ScriptExecutionRecordCreate struct {
	config
	mutation *ScriptExecutionRecordMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (serc *ScriptExecutionRecordCreate) SetUserID(s string) *ScriptExecutionRecordCreate {
	serc.mutation.SetUserID(s)
	return serc
}

// SetFkScriptID sets the "fk_script_id" field.
func (serc *ScriptExecutionRecordCreate) SetFkScriptID(i int32) *ScriptExecutionRecordCreate {
	serc.mutation.SetFkScriptID(i)
	return serc
}

// SetScriptContent sets the "script_content" field.
func (serc *ScriptExecutionRecordCreate) SetScriptContent(s string) *ScriptExecutionRecordCreate {
	serc.mutation.SetScriptContent(s)
	return serc
}

// SetTaskNumber sets the "task_number" field.
func (serc *ScriptExecutionRecordCreate) SetTaskNumber(i int32) *ScriptExecutionRecordCreate {
	serc.mutation.SetTaskNumber(i)
	return serc
}

// SetScriptName sets the "script_name" field.
func (serc *ScriptExecutionRecordCreate) SetScriptName(s string) *ScriptExecutionRecordCreate {
	serc.mutation.SetScriptName(s)
	return serc
}

// SetFileAddress sets the "file_address" field.
func (serc *ScriptExecutionRecordCreate) SetFileAddress(s string) *ScriptExecutionRecordCreate {
	serc.mutation.SetFileAddress(s)
	return serc
}

// SetExecuteState sets the "execute_state" field.
func (serc *ScriptExecutionRecordCreate) SetExecuteState(i int32) *ScriptExecutionRecordCreate {
	serc.mutation.SetExecuteState(i)
	return serc
}

// SetExecuteResult sets the "execute_result" field.
func (serc *ScriptExecutionRecordCreate) SetExecuteResult(s string) *ScriptExecutionRecordCreate {
	serc.mutation.SetExecuteResult(s)
	return serc
}

// SetCreateTime sets the "create_time" field.
func (serc *ScriptExecutionRecordCreate) SetCreateTime(t time.Time) *ScriptExecutionRecordCreate {
	serc.mutation.SetCreateTime(t)
	return serc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (serc *ScriptExecutionRecordCreate) SetNillableCreateTime(t *time.Time) *ScriptExecutionRecordCreate {
	if t != nil {
		serc.SetCreateTime(*t)
	}
	return serc
}

// SetUpdateTime sets the "update_time" field.
func (serc *ScriptExecutionRecordCreate) SetUpdateTime(t time.Time) *ScriptExecutionRecordCreate {
	serc.mutation.SetUpdateTime(t)
	return serc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (serc *ScriptExecutionRecordCreate) SetNillableUpdateTime(t *time.Time) *ScriptExecutionRecordCreate {
	if t != nil {
		serc.SetUpdateTime(*t)
	}
	return serc
}

// SetID sets the "id" field.
func (serc *ScriptExecutionRecordCreate) SetID(i int32) *ScriptExecutionRecordCreate {
	serc.mutation.SetID(i)
	return serc
}

// SetScriptID sets the "script" edge to the Script entity by ID.
func (serc *ScriptExecutionRecordCreate) SetScriptID(id int32) *ScriptExecutionRecordCreate {
	serc.mutation.SetScriptID(id)
	return serc
}

// SetNillableScriptID sets the "script" edge to the Script entity by ID if the given value is not nil.
func (serc *ScriptExecutionRecordCreate) SetNillableScriptID(id *int32) *ScriptExecutionRecordCreate {
	if id != nil {
		serc = serc.SetScriptID(*id)
	}
	return serc
}

// SetScript sets the "script" edge to the Script entity.
func (serc *ScriptExecutionRecordCreate) SetScript(s *Script) *ScriptExecutionRecordCreate {
	return serc.SetScriptID(s.ID)
}

// Mutation returns the ScriptExecutionRecordMutation object of the builder.
func (serc *ScriptExecutionRecordCreate) Mutation() *ScriptExecutionRecordMutation {
	return serc.mutation
}

// Save creates the ScriptExecutionRecord in the database.
func (serc *ScriptExecutionRecordCreate) Save(ctx context.Context) (*ScriptExecutionRecord, error) {
	serc.defaults()
	return withHooks(ctx, serc.sqlSave, serc.mutation, serc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (serc *ScriptExecutionRecordCreate) SaveX(ctx context.Context) *ScriptExecutionRecord {
	v, err := serc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (serc *ScriptExecutionRecordCreate) Exec(ctx context.Context) error {
	_, err := serc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (serc *ScriptExecutionRecordCreate) ExecX(ctx context.Context) {
	if err := serc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (serc *ScriptExecutionRecordCreate) defaults() {
	if _, ok := serc.mutation.CreateTime(); !ok {
		v := scriptexecutionrecord.DefaultCreateTime
		serc.mutation.SetCreateTime(v)
	}
	if _, ok := serc.mutation.UpdateTime(); !ok {
		v := scriptexecutionrecord.DefaultUpdateTime
		serc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (serc *ScriptExecutionRecordCreate) check() error {
	if _, ok := serc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ScriptExecutionRecord.user_id"`)}
	}
	if _, ok := serc.mutation.FkScriptID(); !ok {
		return &ValidationError{Name: "fk_script_id", err: errors.New(`ent: missing required field "ScriptExecutionRecord.fk_script_id"`)}
	}
	if v, ok := serc.mutation.FkScriptID(); ok {
		if err := scriptexecutionrecord.FkScriptIDValidator(v); err != nil {
			return &ValidationError{Name: "fk_script_id", err: fmt.Errorf(`ent: validator failed for field "ScriptExecutionRecord.fk_script_id": %w`, err)}
		}
	}
	if _, ok := serc.mutation.ScriptContent(); !ok {
		return &ValidationError{Name: "script_content", err: errors.New(`ent: missing required field "ScriptExecutionRecord.script_content"`)}
	}
	if v, ok := serc.mutation.ScriptContent(); ok {
		if err := scriptexecutionrecord.ScriptContentValidator(v); err != nil {
			return &ValidationError{Name: "script_content", err: fmt.Errorf(`ent: validator failed for field "ScriptExecutionRecord.script_content": %w`, err)}
		}
	}
	if _, ok := serc.mutation.TaskNumber(); !ok {
		return &ValidationError{Name: "task_number", err: errors.New(`ent: missing required field "ScriptExecutionRecord.task_number"`)}
	}
	if v, ok := serc.mutation.TaskNumber(); ok {
		if err := scriptexecutionrecord.TaskNumberValidator(v); err != nil {
			return &ValidationError{Name: "task_number", err: fmt.Errorf(`ent: validator failed for field "ScriptExecutionRecord.task_number": %w`, err)}
		}
	}
	if _, ok := serc.mutation.ScriptName(); !ok {
		return &ValidationError{Name: "script_name", err: errors.New(`ent: missing required field "ScriptExecutionRecord.script_name"`)}
	}
	if v, ok := serc.mutation.ScriptName(); ok {
		if err := scriptexecutionrecord.ScriptNameValidator(v); err != nil {
			return &ValidationError{Name: "script_name", err: fmt.Errorf(`ent: validator failed for field "ScriptExecutionRecord.script_name": %w`, err)}
		}
	}
	if _, ok := serc.mutation.FileAddress(); !ok {
		return &ValidationError{Name: "file_address", err: errors.New(`ent: missing required field "ScriptExecutionRecord.file_address"`)}
	}
	if _, ok := serc.mutation.ExecuteState(); !ok {
		return &ValidationError{Name: "execute_state", err: errors.New(`ent: missing required field "ScriptExecutionRecord.execute_state"`)}
	}
	if _, ok := serc.mutation.ExecuteResult(); !ok {
		return &ValidationError{Name: "execute_result", err: errors.New(`ent: missing required field "ScriptExecutionRecord.execute_result"`)}
	}
	if _, ok := serc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "ScriptExecutionRecord.create_time"`)}
	}
	if _, ok := serc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "ScriptExecutionRecord.update_time"`)}
	}
	return nil
}

func (serc *ScriptExecutionRecordCreate) sqlSave(ctx context.Context) (*ScriptExecutionRecord, error) {
	if err := serc.check(); err != nil {
		return nil, err
	}
	_node, _spec := serc.createSpec()
	if err := sqlgraph.CreateNode(ctx, serc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	serc.mutation.id = &_node.ID
	serc.mutation.done = true
	return _node, nil
}

func (serc *ScriptExecutionRecordCreate) createSpec() (*ScriptExecutionRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &ScriptExecutionRecord{config: serc.config}
		_spec = sqlgraph.NewCreateSpec(scriptexecutionrecord.Table, sqlgraph.NewFieldSpec(scriptexecutionrecord.FieldID, field.TypeInt32))
	)
	if id, ok := serc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := serc.mutation.UserID(); ok {
		_spec.SetField(scriptexecutionrecord.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := serc.mutation.FkScriptID(); ok {
		_spec.SetField(scriptexecutionrecord.FieldFkScriptID, field.TypeInt32, value)
		_node.FkScriptID = value
	}
	if value, ok := serc.mutation.ScriptContent(); ok {
		_spec.SetField(scriptexecutionrecord.FieldScriptContent, field.TypeString, value)
		_node.ScriptContent = value
	}
	if value, ok := serc.mutation.TaskNumber(); ok {
		_spec.SetField(scriptexecutionrecord.FieldTaskNumber, field.TypeInt32, value)
		_node.TaskNumber = value
	}
	if value, ok := serc.mutation.ScriptName(); ok {
		_spec.SetField(scriptexecutionrecord.FieldScriptName, field.TypeString, value)
		_node.ScriptName = value
	}
	if value, ok := serc.mutation.FileAddress(); ok {
		_spec.SetField(scriptexecutionrecord.FieldFileAddress, field.TypeString, value)
		_node.FileAddress = value
	}
	if value, ok := serc.mutation.ExecuteState(); ok {
		_spec.SetField(scriptexecutionrecord.FieldExecuteState, field.TypeInt32, value)
		_node.ExecuteState = value
	}
	if value, ok := serc.mutation.ExecuteResult(); ok {
		_spec.SetField(scriptexecutionrecord.FieldExecuteResult, field.TypeString, value)
		_node.ExecuteResult = value
	}
	if value, ok := serc.mutation.CreateTime(); ok {
		_spec.SetField(scriptexecutionrecord.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := serc.mutation.UpdateTime(); ok {
		_spec.SetField(scriptexecutionrecord.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if nodes := serc.mutation.ScriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scriptexecutionrecord.ScriptTable,
			Columns: []string{scriptexecutionrecord.ScriptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(script.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.script_script_execution_records = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScriptExecutionRecordCreateBulk is the builder for creating many ScriptExecutionRecord entities in bulk.
type ScriptExecutionRecordCreateBulk struct {
	config
	builders []*ScriptExecutionRecordCreate
}

// Save creates the ScriptExecutionRecord entities in the database.
func (sercb *ScriptExecutionRecordCreateBulk) Save(ctx context.Context) ([]*ScriptExecutionRecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sercb.builders))
	nodes := make([]*ScriptExecutionRecord, len(sercb.builders))
	mutators := make([]Mutator, len(sercb.builders))
	for i := range sercb.builders {
		func(i int, root context.Context) {
			builder := sercb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScriptExecutionRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, sercb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sercb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, sercb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sercb *ScriptExecutionRecordCreateBulk) SaveX(ctx context.Context) []*ScriptExecutionRecord {
	v, err := sercb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sercb *ScriptExecutionRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := sercb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sercb *ScriptExecutionRecordCreateBulk) ExecX(ctx context.Context) {
	if err := sercb.Exec(ctx); err != nil {
		panic(err)
	}
}
