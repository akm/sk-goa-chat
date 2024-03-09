package design

import (
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func InResultType() bool {
	for _, exp := range eval.Context.Stack {
		if _, ok := exp.(*expr.ResultTypeExpr); ok {
			return true
		}
	}
	return false
}

var InRT = InResultType

func InPayload() bool {
	// ResultType と違って Type に渡された関数内から呼ばれているかどうかは eval.Context.Stack には反映されていない
	// なので、ResultType でなければ Payload だと判断してしまう
	return !InResultType()
}
