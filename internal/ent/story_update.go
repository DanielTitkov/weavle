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
	"github.com/DanielTitkov/weavle/internal/ent/predicate"
	"github.com/DanielTitkov/weavle/internal/ent/sentence"
	"github.com/DanielTitkov/weavle/internal/ent/story"
	"github.com/google/uuid"
)

// StoryUpdate is the builder for updating Story entities.
type StoryUpdate struct {
	config
	hooks    []Hook
	mutation *StoryMutation
}

// Where appends a list predicates to the StoryUpdate builder.
func (su *StoryUpdate) Where(ps ...predicate.Story) *StoryUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetStatus sets the "status" field.
func (su *StoryUpdate) SetStatus(s story.Status) *StoryUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *StoryUpdate) SetNillableStatus(s *story.Status) *StoryUpdate {
	if s != nil {
		su.SetStatus(*s)
	}
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *StoryUpdate) SetCreatedAt(t time.Time) *StoryUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *StoryUpdate) SetNillableCreatedAt(t *time.Time) *StoryUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StoryUpdate) SetUpdatedAt(t time.Time) *StoryUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// AddSentenceIDs adds the "sentences" edge to the Sentence entity by IDs.
func (su *StoryUpdate) AddSentenceIDs(ids ...uuid.UUID) *StoryUpdate {
	su.mutation.AddSentenceIDs(ids...)
	return su
}

// AddSentences adds the "sentences" edges to the Sentence entity.
func (su *StoryUpdate) AddSentences(s ...*Sentence) *StoryUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSentenceIDs(ids...)
}

// Mutation returns the StoryMutation object of the builder.
func (su *StoryUpdate) Mutation() *StoryMutation {
	return su.mutation
}

// ClearSentences clears all "sentences" edges to the Sentence entity.
func (su *StoryUpdate) ClearSentences() *StoryUpdate {
	su.mutation.ClearSentences()
	return su
}

// RemoveSentenceIDs removes the "sentences" edge to Sentence entities by IDs.
func (su *StoryUpdate) RemoveSentenceIDs(ids ...uuid.UUID) *StoryUpdate {
	su.mutation.RemoveSentenceIDs(ids...)
	return su
}

// RemoveSentences removes "sentences" edges to Sentence entities.
func (su *StoryUpdate) RemoveSentences(s ...*Sentence) *StoryUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSentenceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StoryUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StoryUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StoryUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StoryUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StoryUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := story.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StoryUpdate) check() error {
	if v, ok := su.mutation.Status(); ok {
		if err := story.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Story.status": %w`, err)}
		}
	}
	return nil
}

func (su *StoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(story.Table, story.Columns, sqlgraph.NewFieldSpec(story.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(story.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(story.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(story.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.SentencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   story.SentencesTable,
			Columns: []string{story.SentencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentence.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSentencesIDs(); len(nodes) > 0 && !su.mutation.SentencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   story.SentencesTable,
			Columns: []string{story.SentencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SentencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   story.SentencesTable,
			Columns: []string{story.SentencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{story.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StoryUpdateOne is the builder for updating a single Story entity.
type StoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoryMutation
}

// SetStatus sets the "status" field.
func (suo *StoryUpdateOne) SetStatus(s story.Status) *StoryUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *StoryUpdateOne) SetNillableStatus(s *story.Status) *StoryUpdateOne {
	if s != nil {
		suo.SetStatus(*s)
	}
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *StoryUpdateOne) SetCreatedAt(t time.Time) *StoryUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *StoryUpdateOne) SetNillableCreatedAt(t *time.Time) *StoryUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StoryUpdateOne) SetUpdatedAt(t time.Time) *StoryUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// AddSentenceIDs adds the "sentences" edge to the Sentence entity by IDs.
func (suo *StoryUpdateOne) AddSentenceIDs(ids ...uuid.UUID) *StoryUpdateOne {
	suo.mutation.AddSentenceIDs(ids...)
	return suo
}

// AddSentences adds the "sentences" edges to the Sentence entity.
func (suo *StoryUpdateOne) AddSentences(s ...*Sentence) *StoryUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSentenceIDs(ids...)
}

// Mutation returns the StoryMutation object of the builder.
func (suo *StoryUpdateOne) Mutation() *StoryMutation {
	return suo.mutation
}

// ClearSentences clears all "sentences" edges to the Sentence entity.
func (suo *StoryUpdateOne) ClearSentences() *StoryUpdateOne {
	suo.mutation.ClearSentences()
	return suo
}

// RemoveSentenceIDs removes the "sentences" edge to Sentence entities by IDs.
func (suo *StoryUpdateOne) RemoveSentenceIDs(ids ...uuid.UUID) *StoryUpdateOne {
	suo.mutation.RemoveSentenceIDs(ids...)
	return suo
}

// RemoveSentences removes "sentences" edges to Sentence entities.
func (suo *StoryUpdateOne) RemoveSentences(s ...*Sentence) *StoryUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSentenceIDs(ids...)
}

// Where appends a list predicates to the StoryUpdate builder.
func (suo *StoryUpdateOne) Where(ps ...predicate.Story) *StoryUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StoryUpdateOne) Select(field string, fields ...string) *StoryUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Story entity.
func (suo *StoryUpdateOne) Save(ctx context.Context) (*Story, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StoryUpdateOne) SaveX(ctx context.Context) *Story {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StoryUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StoryUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StoryUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := story.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StoryUpdateOne) check() error {
	if v, ok := suo.mutation.Status(); ok {
		if err := story.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Story.status": %w`, err)}
		}
	}
	return nil
}

func (suo *StoryUpdateOne) sqlSave(ctx context.Context) (_node *Story, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(story.Table, story.Columns, sqlgraph.NewFieldSpec(story.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Story.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, story.FieldID)
		for _, f := range fields {
			if !story.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != story.FieldID {
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
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(story.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(story.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(story.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.SentencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   story.SentencesTable,
			Columns: []string{story.SentencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentence.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSentencesIDs(); len(nodes) > 0 && !suo.mutation.SentencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   story.SentencesTable,
			Columns: []string{story.SentencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SentencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   story.SentencesTable,
			Columns: []string{story.SentencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Story{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{story.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}