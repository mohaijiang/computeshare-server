// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeimage"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ComputeImageUpdate is the builder for updating ComputeImage entities.
type ComputeImageUpdate struct {
	config
	hooks    []Hook
	mutation *ComputeImageMutation
}

// Where appends a list predicates to the ComputeImageUpdate builder.
func (ciu *ComputeImageUpdate) Where(ps ...predicate.ComputeImage) *ComputeImageUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetName sets the "name" field.
func (ciu *ComputeImageUpdate) SetName(s string) *ComputeImageUpdate {
	ciu.mutation.SetName(s)
	return ciu
}

// SetImage sets the "image" field.
func (ciu *ComputeImageUpdate) SetImage(s string) *ComputeImageUpdate {
	ciu.mutation.SetImage(s)
	return ciu
}

// SetTag sets the "tag" field.
func (ciu *ComputeImageUpdate) SetTag(s string) *ComputeImageUpdate {
	ciu.mutation.SetTag(s)
	return ciu
}

// SetPort sets the "port" field.
func (ciu *ComputeImageUpdate) SetPort(i int32) *ComputeImageUpdate {
	ciu.mutation.ResetPort()
	ciu.mutation.SetPort(i)
	return ciu
}

// AddPort adds i to the "port" field.
func (ciu *ComputeImageUpdate) AddPort(i int32) *ComputeImageUpdate {
	ciu.mutation.AddPort(i)
	return ciu
}

// SetCommand sets the "command" field.
func (ciu *ComputeImageUpdate) SetCommand(s string) *ComputeImageUpdate {
	ciu.mutation.SetCommand(s)
	return ciu
}

// Mutation returns the ComputeImageMutation object of the builder.
func (ciu *ComputeImageUpdate) Mutation() *ComputeImageMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *ComputeImageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ciu.sqlSave, ciu.mutation, ciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *ComputeImageUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *ComputeImageUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *ComputeImageUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *ComputeImageUpdate) check() error {
	if v, ok := ciu.mutation.Name(); ok {
		if err := computeimage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.name": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Image(); ok {
		if err := computeimage.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.image": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Tag(); ok {
		if err := computeimage.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.tag": %w`, err)}
		}
	}
	return nil
}

func (ciu *ComputeImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ciu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(computeimage.Table, computeimage.Columns, sqlgraph.NewFieldSpec(computeimage.FieldID, field.TypeInt32))
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.Name(); ok {
		_spec.SetField(computeimage.FieldName, field.TypeString, value)
	}
	if value, ok := ciu.mutation.Image(); ok {
		_spec.SetField(computeimage.FieldImage, field.TypeString, value)
	}
	if value, ok := ciu.mutation.Tag(); ok {
		_spec.SetField(computeimage.FieldTag, field.TypeString, value)
	}
	if value, ok := ciu.mutation.Port(); ok {
		_spec.SetField(computeimage.FieldPort, field.TypeInt32, value)
	}
	if value, ok := ciu.mutation.AddedPort(); ok {
		_spec.AddField(computeimage.FieldPort, field.TypeInt32, value)
	}
	if value, ok := ciu.mutation.Command(); ok {
		_spec.SetField(computeimage.FieldCommand, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computeimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ciu.mutation.done = true
	return n, nil
}

// ComputeImageUpdateOne is the builder for updating a single ComputeImage entity.
type ComputeImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ComputeImageMutation
}

// SetName sets the "name" field.
func (ciuo *ComputeImageUpdateOne) SetName(s string) *ComputeImageUpdateOne {
	ciuo.mutation.SetName(s)
	return ciuo
}

// SetImage sets the "image" field.
func (ciuo *ComputeImageUpdateOne) SetImage(s string) *ComputeImageUpdateOne {
	ciuo.mutation.SetImage(s)
	return ciuo
}

// SetTag sets the "tag" field.
func (ciuo *ComputeImageUpdateOne) SetTag(s string) *ComputeImageUpdateOne {
	ciuo.mutation.SetTag(s)
	return ciuo
}

// SetPort sets the "port" field.
func (ciuo *ComputeImageUpdateOne) SetPort(i int32) *ComputeImageUpdateOne {
	ciuo.mutation.ResetPort()
	ciuo.mutation.SetPort(i)
	return ciuo
}

// AddPort adds i to the "port" field.
func (ciuo *ComputeImageUpdateOne) AddPort(i int32) *ComputeImageUpdateOne {
	ciuo.mutation.AddPort(i)
	return ciuo
}

// SetCommand sets the "command" field.
func (ciuo *ComputeImageUpdateOne) SetCommand(s string) *ComputeImageUpdateOne {
	ciuo.mutation.SetCommand(s)
	return ciuo
}

// Mutation returns the ComputeImageMutation object of the builder.
func (ciuo *ComputeImageUpdateOne) Mutation() *ComputeImageMutation {
	return ciuo.mutation
}

// Where appends a list predicates to the ComputeImageUpdate builder.
func (ciuo *ComputeImageUpdateOne) Where(ps ...predicate.ComputeImage) *ComputeImageUpdateOne {
	ciuo.mutation.Where(ps...)
	return ciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *ComputeImageUpdateOne) Select(field string, fields ...string) *ComputeImageUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated ComputeImage entity.
func (ciuo *ComputeImageUpdateOne) Save(ctx context.Context) (*ComputeImage, error) {
	return withHooks(ctx, ciuo.sqlSave, ciuo.mutation, ciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *ComputeImageUpdateOne) SaveX(ctx context.Context) *ComputeImage {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *ComputeImageUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *ComputeImageUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *ComputeImageUpdateOne) check() error {
	if v, ok := ciuo.mutation.Name(); ok {
		if err := computeimage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.name": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Image(); ok {
		if err := computeimage.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.image": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Tag(); ok {
		if err := computeimage.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf(`ent: validator failed for field "ComputeImage.tag": %w`, err)}
		}
	}
	return nil
}

func (ciuo *ComputeImageUpdateOne) sqlSave(ctx context.Context) (_node *ComputeImage, err error) {
	if err := ciuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(computeimage.Table, computeimage.Columns, sqlgraph.NewFieldSpec(computeimage.FieldID, field.TypeInt32))
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ComputeImage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, computeimage.FieldID)
		for _, f := range fields {
			if !computeimage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != computeimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.Name(); ok {
		_spec.SetField(computeimage.FieldName, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.Image(); ok {
		_spec.SetField(computeimage.FieldImage, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.Tag(); ok {
		_spec.SetField(computeimage.FieldTag, field.TypeString, value)
	}
	if value, ok := ciuo.mutation.Port(); ok {
		_spec.SetField(computeimage.FieldPort, field.TypeInt32, value)
	}
	if value, ok := ciuo.mutation.AddedPort(); ok {
		_spec.AddField(computeimage.FieldPort, field.TypeInt32, value)
	}
	if value, ok := ciuo.mutation.Command(); ok {
		_spec.SetField(computeimage.FieldCommand, field.TypeString, value)
	}
	_node = &ComputeImage{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computeimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ciuo.mutation.done = true
	return _node, nil
}
