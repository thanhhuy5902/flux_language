package expression

import (
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/exception"
)

type Operator interface {
	SetLeftExpr(leftExpr *NumericExpression)
	SetRightExpr(rightExpr *NumericExpression)
	GetLeftExpr() *NumericExpression
	GetRightExpr() *NumericExpression
	Generate(ctx *codeobjects.GenerateContext) string
	Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException
}
