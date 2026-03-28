package operations

import (
	"context"

	oppause "github.com/kyverno/chainsaw/pkg/engine/operations/pause"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
)

type pauseAction struct {
	onPause func()
}

func (o pauseAction) Execute(ctx context.Context, tc enginecontext.TestContext) (outputs.Outputs, error) {
	op := oppause.New(o.onPause)
	return op.Exec(ctx, tc.Bindings())
}

func pauseOperation(onPause func()) Operation {
	return pauseAction{
		onPause: onPause,
	}
}
