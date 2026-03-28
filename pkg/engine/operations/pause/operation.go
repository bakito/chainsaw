package pause

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/engine/operations"
	"github.com/kyverno/chainsaw/pkg/engine/operations/internal"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	"github.com/kyverno/chainsaw/pkg/logging"
)

type operation struct {
	onPause func()
}

func New(onPause func()) operations.Operation {
	return &operation{}
}

func (o *operation) Exec(ctx context.Context, _ apis.Bindings) (_ outputs.Outputs, _err error) {
	defer func() {
		internal.LogEnd(ctx, logging.Pause, nil, _err)
	}()
	internal.LogStart(ctx, logging.Pause, nil)
	return nil, o.execute()
}

func (o *operation) execute() error {
	o.onPause()
	return nil
}
