package operators

import (
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/codeobjects/expression"
	"github.com/thanhhuy5902/flux_language/exception"
)

type DivisionOperator struct {
	*codeobjects.BaseStatement
	Left  *expression.NumericExpression
	Right *expression.NumericExpression
}

func (d *DivisionOperator) SetLeftExpr(leftExpr *expression.NumericExpression) {
	d.Left = leftExpr
}

func (d *DivisionOperator) SetRightExpr(rightExpr *expression.NumericExpression) {
	d.Right = rightExpr
}

func (d *DivisionOperator) GetLeftExpr() *expression.NumericExpression {
	return d.Left
}

func (d *DivisionOperator) GetRightExpr() *expression.NumericExpression {
	return d.Right
}

func (d *DivisionOperator) Generate(ctx *codeobjects.GenerateContext) string {
	return d.Left.Generate(ctx) + " / " + d.Right.Generate(ctx)
}

func (d *DivisionOperator) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException{
	leftCtx := ctx.Clone()
	rightCtx := ctx.Clone()
	if d.Left != nil {
		err := d.Left.Execute(leftCtx)
		if err != nil {
			return err
		}
	} else {
		leftCtx.NumericValue = 0
	}

	if d.Right != nil {
		err := d.Right.Execute(rightCtx)
		if err != nil {
			return err
		}
	} else {
		rightCtx.NumericValue = 0
	}

	ctx.NumericValue = leftCtx.NumericValue / rightCtx.NumericValue
	return nil
}

func NewDivision(line int, startPos int, endPos int, left *expression.NumericExpression, right *expression.NumericExpression) *DivisionOperator {
	return &DivisionOperator{BaseStatement: &codeobjects.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, Left: left, Right: right}
}
