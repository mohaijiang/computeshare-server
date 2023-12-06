// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/mohaijiang/computeshare-server/internal/data/ent"
)

// The AgentFunc type is an adapter to allow the use of ordinary
// function as Agent mutator.
type AgentFunc func(context.Context, *ent.AgentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AgentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AgentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AgentMutation", m)
}

// The ComputeImageFunc type is an adapter to allow the use of ordinary
// function as ComputeImage mutator.
type ComputeImageFunc func(context.Context, *ent.ComputeImageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ComputeImageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ComputeImageMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ComputeImageMutation", m)
}

// The ComputeInstanceFunc type is an adapter to allow the use of ordinary
// function as ComputeInstance mutator.
type ComputeInstanceFunc func(context.Context, *ent.ComputeInstanceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ComputeInstanceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ComputeInstanceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ComputeInstanceMutation", m)
}

// The ComputeSpecFunc type is an adapter to allow the use of ordinary
// function as ComputeSpec mutator.
type ComputeSpecFunc func(context.Context, *ent.ComputeSpecMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ComputeSpecFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ComputeSpecMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ComputeSpecMutation", m)
}

// The DomainBindingFunc type is an adapter to allow the use of ordinary
// function as DomainBinding mutator.
type DomainBindingFunc func(context.Context, *ent.DomainBindingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DomainBindingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DomainBindingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DomainBindingMutation", m)
}

// The EmployeeFunc type is an adapter to allow the use of ordinary
// function as Employee mutator.
type EmployeeFunc func(context.Context, *ent.EmployeeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmployeeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmployeeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmployeeMutation", m)
}

// The GatewayFunc type is an adapter to allow the use of ordinary
// function as Gateway mutator.
type GatewayFunc func(context.Context, *ent.GatewayMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GatewayFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GatewayMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GatewayMutation", m)
}

// The GatewayPortFunc type is an adapter to allow the use of ordinary
// function as GatewayPort mutator.
type GatewayPortFunc func(context.Context, *ent.GatewayPortMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GatewayPortFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GatewayPortMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GatewayPortMutation", m)
}

// The NetworkMappingFunc type is an adapter to allow the use of ordinary
// function as NetworkMapping mutator.
type NetworkMappingFunc func(context.Context, *ent.NetworkMappingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NetworkMappingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NetworkMappingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NetworkMappingMutation", m)
}

// The S3BucketFunc type is an adapter to allow the use of ordinary
// function as S3Bucket mutator.
type S3BucketFunc func(context.Context, *ent.S3BucketMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f S3BucketFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.S3BucketMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.S3BucketMutation", m)
}

// The S3UserFunc type is an adapter to allow the use of ordinary
// function as S3User mutator.
type S3UserFunc func(context.Context, *ent.S3UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f S3UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.S3UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.S3UserMutation", m)
}

// The ScriptFunc type is an adapter to allow the use of ordinary
// function as Script mutator.
type ScriptFunc func(context.Context, *ent.ScriptMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScriptFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ScriptMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScriptMutation", m)
}

// The ScriptExecutionRecordFunc type is an adapter to allow the use of ordinary
// function as ScriptExecutionRecord mutator.
type ScriptExecutionRecordFunc func(context.Context, *ent.ScriptExecutionRecordMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScriptExecutionRecordFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ScriptExecutionRecordMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScriptExecutionRecordMutation", m)
}

// The StorageFunc type is an adapter to allow the use of ordinary
// function as Storage mutator.
type StorageFunc func(context.Context, *ent.StorageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StorageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.StorageMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StorageMutation", m)
}

// The StorageProviderFunc type is an adapter to allow the use of ordinary
// function as StorageProvider mutator.
type StorageProviderFunc func(context.Context, *ent.StorageProviderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StorageProviderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.StorageProviderMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StorageProviderMutation", m)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *ent.TaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TaskMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
