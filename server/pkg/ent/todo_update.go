// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/pkg/ent/label"
	"server/pkg/ent/predicate"
	"server/pkg/ent/todo"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TodoUpdate) SetTitle(s string) *TodoUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableTitle(s *string) *TodoUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TodoUpdate) SetDescription(s string) *TodoUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableDescription(s *string) *TodoUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// SetStartAt sets the "start_at" field.
func (tu *TodoUpdate) SetStartAt(t time.Time) *TodoUpdate {
	tu.mutation.SetStartAt(t)
	return tu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableStartAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetStartAt(*t)
	}
	return tu
}

// SetEndAt sets the "end_at" field.
func (tu *TodoUpdate) SetEndAt(t time.Time) *TodoUpdate {
	tu.mutation.SetEndAt(t)
	return tu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableEndAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetEndAt(*t)
	}
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TodoUpdate) SetPriority(t todo.Priority) *TodoUpdate {
	tu.mutation.SetPriority(t)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TodoUpdate) SetNillablePriority(t *todo.Priority) *TodoUpdate {
	if t != nil {
		tu.SetPriority(*t)
	}
	return tu
}

// SetStatus sets the "status" field.
func (tu *TodoUpdate) SetStatus(t todo.Status) *TodoUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableStatus(t *todo.Status) *TodoUpdate {
	if t != nil {
		tu.SetStatus(*t)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TodoUpdate) SetCreatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableCreatedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TodoUpdate) SetUpdatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
func (tu *TodoUpdate) AddLabelIDs(ids ...int) *TodoUpdate {
	tu.mutation.AddLabelIDs(ids...)
	return tu
}

// AddLabels adds the "labels" edges to the Label entity.
func (tu *TodoUpdate) AddLabels(l ...*Label) *TodoUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return tu.AddLabelIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// ClearLabels clears all "labels" edges to the Label entity.
func (tu *TodoUpdate) ClearLabels() *TodoUpdate {
	tu.mutation.ClearLabels()
	return tu
}

// RemoveLabelIDs removes the "labels" edge to Label entities by IDs.
func (tu *TodoUpdate) RemoveLabelIDs(ids ...int) *TodoUpdate {
	tu.mutation.RemoveLabelIDs(ids...)
	return tu
}

// RemoveLabels removes "labels" edges to Label entities.
func (tu *TodoUpdate) RemoveLabels(l ...*Label) *TodoUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return tu.RemoveLabelIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TodoUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := todo.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if v, ok := tu.mutation.Priority(); ok {
		if err := todo.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Todo.priority": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(todo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(todo.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.StartAt(); ok {
		_spec.SetField(todo.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.EndAt(); ok {
		_spec.SetField(todo.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.LabelsTable,
			Columns: todo.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedLabelsIDs(); len(nodes) > 0 && !tu.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.LabelsTable,
			Columns: todo.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.LabelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.LabelsTable,
			Columns: todo.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetTitle sets the "title" field.
func (tuo *TodoUpdateOne) SetTitle(s string) *TodoUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableTitle(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TodoUpdateOne) SetDescription(s string) *TodoUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableDescription(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// SetStartAt sets the "start_at" field.
func (tuo *TodoUpdateOne) SetStartAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetStartAt(t)
	return tuo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableStartAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetStartAt(*t)
	}
	return tuo
}

// SetEndAt sets the "end_at" field.
func (tuo *TodoUpdateOne) SetEndAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetEndAt(t)
	return tuo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableEndAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetEndAt(*t)
	}
	return tuo
}

// SetPriority sets the "priority" field.
func (tuo *TodoUpdateOne) SetPriority(t todo.Priority) *TodoUpdateOne {
	tuo.mutation.SetPriority(t)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillablePriority(t *todo.Priority) *TodoUpdateOne {
	if t != nil {
		tuo.SetPriority(*t)
	}
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TodoUpdateOne) SetStatus(t todo.Status) *TodoUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableStatus(t *todo.Status) *TodoUpdateOne {
	if t != nil {
		tuo.SetStatus(*t)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TodoUpdateOne) SetCreatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableCreatedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TodoUpdateOne) SetUpdatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
func (tuo *TodoUpdateOne) AddLabelIDs(ids ...int) *TodoUpdateOne {
	tuo.mutation.AddLabelIDs(ids...)
	return tuo
}

// AddLabels adds the "labels" edges to the Label entity.
func (tuo *TodoUpdateOne) AddLabels(l ...*Label) *TodoUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return tuo.AddLabelIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// ClearLabels clears all "labels" edges to the Label entity.
func (tuo *TodoUpdateOne) ClearLabels() *TodoUpdateOne {
	tuo.mutation.ClearLabels()
	return tuo
}

// RemoveLabelIDs removes the "labels" edge to Label entities by IDs.
func (tuo *TodoUpdateOne) RemoveLabelIDs(ids ...int) *TodoUpdateOne {
	tuo.mutation.RemoveLabelIDs(ids...)
	return tuo
}

// RemoveLabels removes "labels" edges to Label entities.
func (tuo *TodoUpdateOne) RemoveLabels(l ...*Label) *TodoUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return tuo.RemoveLabelIDs(ids...)
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TodoUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := todo.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if v, ok := tuo.mutation.Priority(); ok {
		if err := todo.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Todo.priority": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(todo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(todo.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.StartAt(); ok {
		_spec.SetField(todo.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.EndAt(); ok {
		_spec.SetField(todo.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.LabelsTable,
			Columns: todo.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedLabelsIDs(); len(nodes) > 0 && !tuo.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.LabelsTable,
			Columns: todo.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.LabelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.LabelsTable,
			Columns: todo.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
