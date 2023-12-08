package callback

import (
	"context"

	"github.com/goexl/powerjob/internal/internal/core"
	"github.com/goexl/powerjob/internal/param"
)

type Post func(ctx context.Context, core *param.Core, path string, req core.Application, rsp core.SuccessChecker) error
