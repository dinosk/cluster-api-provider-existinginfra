package resource

import "github.com/twelho/capi-existinginfra/pkg/plan"

// base can be embedded into a struct to provide a default implementation of
// plan.Resource.
type base struct {
}

var _ plan.Resource = plan.RegisterResource(&base{})

// State implements plan.Resource.
func (b *base) State() plan.State {
	return plan.EmptyState
}

// QueryState implements plan.Resource.
func (b *base) QueryState(runner plan.Runner) (plan.State, error) {
	return plan.EmptyState, nil
}

// Apply implements plan.Resource.
func (b *base) Apply(runner plan.Runner, diff plan.Diff) (bool, error) {
	return true, nil
}

// Undo implements plan.Resource.
func (b *base) Undo(runner plan.Runner, current plan.State) error {
	return nil
}
