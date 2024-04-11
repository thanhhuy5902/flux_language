package expression

import (
	"fmt"
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/exception"
)

type NumericExpression struct {
	*codeobjects.BaseStatement
	LeftExpr  *NumericExpression
	Op        Operator
	RightExpr *NumericExpression
	Value     float64
	GetVar    *GetVar
}

func (n NumericExpression) Generate(ctx *codeobjects.GenerateContext) string {

	if n.LeftExpr == nil && n.RightExpr == nil {
		return fmt.Sprintf("%v", n.Value)
	}
	return n.Op.Generate(ctx)
}

func (n NumericExpression) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	if n.LeftExpr == nil && n.RightExpr == nil {
		ctx.NumericValue = n.Value
		return nil
	}
	n.Op.SetLeftExpr(n.LeftExpr)
	n.Op.SetRightExpr(n.RightExpr)
	return n.Op.Execute(ctx)
}

func NewNumericExpression(line int, startPos int, endPos int, leftExpr *NumericExpression, operator Operator, rightExpr *NumericExpression, value float64) *NumericExpression {
	return &NumericExpression{BaseStatement: &codeobjects.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, LeftExpr: leftExpr, Op: operator, RightExpr: rightExpr, Value: value}
}
