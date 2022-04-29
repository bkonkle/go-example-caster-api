// Code generated by entc, DO NOT EDIT.

package ent

import (
	"caster/ent/predicate"
	"caster/ent/show"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShowDelete is the builder for deleting a Show entity.
type ShowDelete struct {
	config
	hooks    []Hook
	mutation *ShowMutation
}

// Where appends a list predicates to the ShowDelete builder.
func (sd *ShowDelete) Where(ps ...predicate.Show) *ShowDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ShowDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sd.hooks) == 0 {
		affected, err = sd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sd.mutation = mutation
			affected, err = sd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sd.hooks) - 1; i >= 0; i-- {
			if sd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ShowDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ShowDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: show.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: show.FieldID,
			},
		},
	}
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// ShowDeleteOne is the builder for deleting a single Show entity.
type ShowDeleteOne struct {
	sd *ShowDelete
}

// Exec executes the deletion query.
func (sdo *ShowDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{show.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ShowDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}