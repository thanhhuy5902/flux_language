package operators

import (
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/codeobjects/expression"
	"github.com/thanhhuy5902/flux_language/exception"
)

type Multi struct {
	*codeobjects.BaseStatement
	Left  *expression.NumericExpression
	Right *expression.NumericExpression
}

func (m *Multi) SetLeftExpr(leftExpr *expression.NumericExpression) {
	m.Left = leftExpr
}

func (m *Multi) SetRightExpr(rightExpr *expression.NumericExpression) {
	m.Right = rightExpr
}

func (m *Multi) GetLeftExpr() *expression.NumericExpression {
	return m.Left
}

func (m *Multi) GetRightExpr() *expression.NumericExpression {
	return m.Right
}

func (m *Multi) Generate(ctx *codeobjects.GenerateContext) string {
	return m.Left.Generate(ctx) + " * " + m.Right.Generate(ctx)
}

func (m *Multi) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	leftCtx := ctx.Clone()
	rightCtx := ctx.Clone()
	if m.Left != nil {
		err := m.Left.Execute(leftCtx)
		if err != nil {
			return err
		}
	} else {
		leftCtx.NumericValue = 0
	}

	if m.Right != nil {
		err := m.Right.Execute(rightCtx)
		if err != nil {
			return err
		}
	} else {
		rightCtx.NumericValue = 0
	}

	ctx.NumericValue = leftCtx.NumericValue * rightCtx.NumericValue
	return nil
}

func NewMulti(line int, startPos int, endPos int, left *expression.NumericExpression, right *expression.NumericExpression) *Multi {
	return &Multi{BaseStatement: &codeobjects.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, Left: left, Right: right}
}
