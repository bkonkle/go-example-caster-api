// Code generated by entc, DO NOT EDIT.

package ent

import (
	"caster/ent/episode"
	"caster/ent/predicate"
	"caster/ent/show"
	"caster/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShowUpdate is the builder for updating Show entities.
type ShowUpdate struct {
	config
	hooks    []Hook
	mutation *ShowMutation
}

// Where appends a list predicates to the ShowUpdate builder.
func (su *ShowUpdate) Where(ps ...predicate.Show) *ShowUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *ShowUpdate) SetCreatedAt(t time.Time) *ShowUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *ShowUpdate) SetNillableCreatedAt(t *time.Time) *ShowUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *ShowUpdate) SetUpdatedAt(t time.Time) *ShowUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *ShowUpdate) SetNillableUpdatedAt(t *time.Time) *ShowUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// SetTitle sets the "title" field.
func (su *ShowUpdate) SetTitle(s string) *ShowUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetSummary sets the "summary" field.
func (su *ShowUpdate) SetSummary(s string) *ShowUpdate {
	su.mutation.SetSummary(s)
	return su
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (su *ShowUpdate) SetNillableSummary(s *string) *ShowUpdate {
	if s != nil {
		su.SetSummary(*s)
	}
	return su
}

// ClearSummary clears the value of the "summary" field.
func (su *ShowUpdate) ClearSummary() *ShowUpdate {
	su.mutation.ClearSummary()
	return su
}

// SetPicture sets the "picture" field.
func (su *ShowUpdate) SetPicture(s string) *ShowUpdate {
	su.mutation.SetPicture(s)
	return su
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (su *ShowUpdate) SetNillablePicture(s *string) *ShowUpdate {
	if s != nil {
		su.SetPicture(*s)
	}
	return su
}

// ClearPicture clears the value of the "picture" field.
func (su *ShowUpdate) ClearPicture() *ShowUpdate {
	su.mutation.ClearPicture()
	return su
}

// SetContent sets the "content" field.
func (su *ShowUpdate) SetContent(u *utils.Content) *ShowUpdate {
	su.mutation.SetContent(u)
	return su
}

// ClearContent clears the value of the "content" field.
func (su *ShowUpdate) ClearContent() *ShowUpdate {
	su.mutation.ClearContent()
	return su
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (su *ShowUpdate) AddEpisodeIDs(ids ...string) *ShowUpdate {
	su.mutation.AddEpisodeIDs(ids...)
	return su
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (su *ShowUpdate) AddEpisodes(e ...*Episode) *ShowUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddEpisodeIDs(ids...)
}

// Mutation returns the ShowMutation object of the builder.
func (su *ShowUpdate) Mutation() *ShowMutation {
	return su.mutation
}

// ClearEpisodes clears all "episodes" edges to the Episode entity.
func (su *ShowUpdate) ClearEpisodes() *ShowUpdate {
	su.mutation.ClearEpisodes()
	return su
}

// RemoveEpisodeIDs removes the "episodes" edge to Episode entities by IDs.
func (su *ShowUpdate) RemoveEpisodeIDs(ids ...string) *ShowUpdate {
	su.mutation.RemoveEpisodeIDs(ids...)
	return su
}

// RemoveEpisodes removes "episodes" edges to Episode entities.
func (su *ShowUpdate) RemoveEpisodes(e ...*Episode) *ShowUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveEpisodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ShowUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShowUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShowUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShowUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ShowUpdate) check() error {
	if v, ok := su.mutation.Title(); ok {
		if err := show.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Show.title": %w`, err)}
		}
	}
	if v, ok := su.mutation.Summary(); ok {
		if err := show.SummaryValidator(v); err != nil {
			return &ValidationError{Name: "summary", err: fmt.Errorf(`ent: validator failed for field "Show.summary": %w`, err)}
		}
	}
	if v, ok := su.mutation.Picture(); ok {
		if err := show.PictureValidator(v); err != nil {
			return &ValidationError{Name: "picture", err: fmt.Errorf(`ent: validator failed for field "Show.picture": %w`, err)}
		}
	}
	return nil
}

func (su *ShowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   show.Table,
			Columns: show.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: show.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: show.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: show.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: show.FieldTitle,
		})
	}
	if value, ok := su.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: show.FieldSummary,
		})
	}
	if su.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: show.FieldSummary,
		})
	}
	if value, ok := su.mutation.Picture(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: show.FieldPicture,
		})
	}
	if su.mutation.PictureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: show.FieldPicture,
		})
	}
	if value, ok := su.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: show.FieldContent,
		})
	}
	if su.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: show.FieldContent,
		})
	}
	if su.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   show.EpisodesTable,
			Columns: []string{show.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: episode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedEpisodesIDs(); len(nodes) > 0 && !su.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   show.EpisodesTable,
			Columns: []string{show.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: episode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   show.EpisodesTable,
			Columns: []string{show.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: episode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{show.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ShowUpdateOne is the builder for updating a single Show entity.
type ShowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShowMutation
}

// SetCreatedAt sets the "created_at" field.
func (suo *ShowUpdateOne) SetCreatedAt(t time.Time) *ShowUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *ShowUpdateOne) SetNillableCreatedAt(t *time.Time) *ShowUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *ShowUpdateOne) SetUpdatedAt(t time.Time) *ShowUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *ShowUpdateOne) SetNillableUpdatedAt(t *time.Time) *ShowUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// SetTitle sets the "title" field.
func (suo *ShowUpdateOne) SetTitle(s string) *ShowUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetSummary sets the "summary" field.
func (suo *ShowUpdateOne) SetSummary(s string) *ShowUpdateOne {
	suo.mutation.SetSummary(s)
	return suo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (suo *ShowUpdateOne) SetNillableSummary(s *string) *ShowUpdateOne {
	if s != nil {
		suo.SetSummary(*s)
	}
	return suo
}

// ClearSummary clears the value of the "summary" field.
func (suo *ShowUpdateOne) ClearSummary() *ShowUpdateOne {
	suo.mutation.ClearSummary()
	return suo
}

// SetPicture sets the "picture" field.
func (suo *ShowUpdateOne) SetPicture(s string) *ShowUpdateOne {
	suo.mutation.SetPicture(s)
	return suo
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (suo *ShowUpdateOne) SetNillablePicture(s *string) *ShowUpdateOne {
	if s != nil {
		suo.SetPicture(*s)
	}
	return suo
}

// ClearPicture clears the value of the "picture" field.
func (suo *ShowUpdateOne) ClearPicture() *ShowUpdateOne {
	suo.mutation.ClearPicture()
	return suo
}

// SetContent sets the "content" field.
func (suo *ShowUpdateOne) SetContent(u *utils.Content) *ShowUpdateOne {
	suo.mutation.SetContent(u)
	return suo
}

// ClearContent clears the value of the "content" field.
func (suo *ShowUpdateOne) ClearContent() *ShowUpdateOne {
	suo.mutation.ClearContent()
	return suo
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (suo *ShowUpdateOne) AddEpisodeIDs(ids ...string) *ShowUpdateOne {
	suo.mutation.AddEpisodeIDs(ids...)
	return suo
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (suo *ShowUpdateOne) AddEpisodes(e ...*Episode) *ShowUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddEpisodeIDs(ids...)
}

// Mutation returns the ShowMutation object of the builder.
func (suo *ShowUpdateOne) Mutation() *ShowMutation {
	return suo.mutation
}

// ClearEpisodes clears all "episodes" edges to the Episode entity.
func (suo *ShowUpdateOne) ClearEpisodes() *ShowUpdateOne {
	suo.mutation.ClearEpisodes()
	return suo
}

// RemoveEpisodeIDs removes the "episodes" edge to Episode entities by IDs.
func (suo *ShowUpdateOne) RemoveEpisodeIDs(ids ...string) *ShowUpdateOne {
	suo.mutation.RemoveEpisodeIDs(ids...)
	return suo
}

// RemoveEpisodes removes "episodes" edges to Episode entities.
func (suo *ShowUpdateOne) RemoveEpisodes(e ...*Episode) *ShowUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveEpisodeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ShowUpdateOne) Select(field string, fields ...string) *ShowUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Show entity.
func (suo *ShowUpdateOne) Save(ctx context.Context) (*Show, error) {
	var (
		err  error
		node *Show
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShowUpdateOne) SaveX(ctx context.Context) *Show {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ShowUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShowUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ShowUpdateOne) check() error {
	if v, ok := suo.mutation.Title(); ok {
		if err := show.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Show.title": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Summary(); ok {
		if err := show.SummaryValidator(v); err != nil {
			return &ValidationError{Name: "summary", err: fmt.Errorf(`ent: validator failed for field "Show.summary": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Picture(); ok {
		if err := show.PictureValidator(v); err != nil {
			return &ValidationError{Name: "picture", err: fmt.Errorf(`ent: validator failed for field "Show.picture": %w`, err)}
		}
	}
	return nil
}

func (suo *ShowUpdateOne) sqlSave(ctx context.Context) (_node *Show, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   show.Table,
			Columns: show.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: show.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Show.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, show.FieldID)
		for _, f := range fields {
			if !show.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != show.FieldID {
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
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: show.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: show.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: show.FieldTitle,
		})
	}
	if value, ok := suo.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: show.FieldSummary,
		})
	}
	if suo.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: show.FieldSummary,
		})
	}
	if value, ok := suo.mutation.Picture(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: show.FieldPicture,
		})
	}
	if suo.mutation.PictureCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: show.FieldPicture,
		})
	}
	if value, ok := suo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: show.FieldContent,
		})
	}
	if suo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: show.FieldContent,
		})
	}
	if suo.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   show.EpisodesTable,
			Columns: []string{show.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: episode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedEpisodesIDs(); len(nodes) > 0 && !suo.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   show.EpisodesTable,
			Columns: []string{show.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: episode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   show.EpisodesTable,
			Columns: []string{show.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: episode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Show{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{show.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
