package operators

import (
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/codeobjects/expression"
	"github.com/thanhhuy5902/flux_language/exception"
)

type Add struct {
	*codeobjects.BaseStatement
	Left  *expression.NumericExpression
	Right *expression.NumericExpression
}

func (a *Add) SetLeftExpr(leftExpr *expression.NumericExpression) {
	a.Left = leftExpr
}

func (a *Add) SetRightExpr(rightExpr *expression.NumericExpression) {
	a.Right = rightExpr
}

func (a *Add) GetLeftExpr() *expression.NumericExpression {
	return a.Left
}

func (a *Add) GetRightExpr() *expression.NumericExpression {
	return a.Right
}

func (a *Add) Generate(ctx *codeobjects.GenerateContext) string {
	return a.Left.Generate(ctx) + " + " + a.Right.Generate(ctx)
}

func (a *Add) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	leftCtx := ctx.Clone()
	rightCtx := ctx.Clone()
	if a.Left != nil {
		err := a.Left.Execute(leftCtx)
		if err != nil {
			return err
		}
	} else {
		leftCtx.NumericValue = 0
	}

	if a.Right != nil {
		err := a.Right.Execute(rightCtx)
		if err != nil {
			return err
		}
	} else {
		rightCtx.NumericValue = 0
	}

	ctx.NumericValue = leftCtx.NumericValue + rightCtx.NumericValue
	return nil
}

func NewAdd(line int, startPos int, endPos int, left *expression.NumericExpression, right *expression.NumericExpression) *Add {
	return &Add{BaseStatement: &codeobjects.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, Left: left, Right: right}
}
