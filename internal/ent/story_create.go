// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/weavle/internal/ent/sentence"
	"github.com/DanielTitkov/weavle/internal/ent/story"
	"github.com/google/uuid"
)

// StoryCreate is the builder for creating a Story entity.
type StoryCreate struct {
	config
	mutation *StoryMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (sc *StoryCreate) SetStatus(s story.Status) *StoryCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *StoryCreate) SetNillableStatus(s *story.Status) *StoryCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StoryCreate) SetCreatedAt(t time.Time) *StoryCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StoryCreate) SetNillableCreatedAt(t *time.Time) *StoryCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StoryCreate) SetUpdatedAt(t time.Time) *StoryCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetID sets the "id" field.
func (sc *StoryCreate) SetID(u uuid.UUID) *StoryCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StoryCreate) SetNillableID(u *uuid.UUID) *StoryCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddSentenceIDs adds the "sentences" edge to the Sentence entity by IDs.
func (sc *StoryCreate) AddSentenceIDs(ids ...uuid.UUID) *StoryCreate {
	sc.mutation.AddSentenceIDs(ids...)
	return sc
}

// AddSentences adds the "sentences" edges to the Sentence entity.
func (sc *StoryCreate) AddSentences(s ...*Sentence) *StoryCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddSentenceIDs(ids...)
}

// Mutation returns the StoryMutation object of the builder.
func (sc *StoryCreate) Mutation() *StoryMutation {
	return sc.mutation
}

// Save creates the Story in the database.
func (sc *StoryCreate) Save(ctx context.Context) (*Story, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StoryCreate) SaveX(ctx context.Context) *Story {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StoryCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StoryCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StoryCreate) defaults() {
	if _, ok := sc.mutation.Status(); !ok {
		v := story.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := story.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := story.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StoryCreate) check() error {
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Story.status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := story.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Story.status": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Story.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Story.updated_at"`)}
	}
	return nil
}

func (sc *StoryCreate) sqlSave(ctx context.Context) (*Story, error) {
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

func (sc *StoryCreate) createSpec() (*Story, *sqlgraph.CreateSpec) {
	var (
		_node = &Story{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(story.Table, sqlgraph.NewFieldSpec(story.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(story.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(story.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(story.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := sc.mutation.SentencesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StoryCreateBulk is the builder for creating many Story entities in bulk.
type StoryCreateBulk struct {
	config
	builders []*StoryCreate
}

// Save creates the Story entities in the database.
func (scb *StoryCreateBulk) Save(ctx context.Context) ([]*Story, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Story, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StoryMutation)
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
func (scb *StoryCreateBulk) SaveX(ctx context.Context) []*Story {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StoryCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StoryCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
