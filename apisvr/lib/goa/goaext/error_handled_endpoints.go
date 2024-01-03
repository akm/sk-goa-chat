package goaext

import (
	"context"
	"fmt"
	"reflect"

	goa "goa.design/goa/v3/pkg"
)

// ErrorHandledEndpoints は endpoints で指定された struct の各フィールドのうち
// goa.Endpoint 型のフィールドに対して、エラーハンドラ eh を適用した関数を作成します。
// 戻り値は、元の endpoints と同じ型の struct で、 eh が適用された goa.Endpoint 型のフィールドを持ちます。
func ErrorHandledEndpoints[T any](endpoints T, eh func(error) error) T {
	srcVal := reflect.ValueOf(endpoints).Elem() // ポインタ
	fmt.Printf("srcVal: [%T] %+v\n", srcVal.Interface(), srcVal.Interface())
	dstVal := reflect.New(srcVal.Type()).Elem() // 同じ型のstructを作成。New はポインタのValueを返すので Elem を使う
	fmt.Printf("dstVal: [%T] %+v\n", dstVal.Interface(), dstVal.Interface())

	fieldNum := srcVal.NumField()
	for i := 0; i < fieldNum; i++ {
		field := srcVal.Field(i)
		if field.Kind() != reflect.Func {
			continue
		}
		fv := field.Interface()
		fn, ok := fv.(goa.Endpoint)
		if !ok {
			continue
		}
		var dstFunc goa.Endpoint = func(ctx context.Context, req any) (any, error) {
			r, err := fn(ctx, req)
			if err != nil {
				err = eh(err)
			}
			return r, err
		}
		// dstField := dstVal.FieldByName(field.Name)
		dstField := dstVal.Field(i)
		dstField.Set(reflect.ValueOf(dstFunc))
	}

	return dstVal.Addr().Interface().(T)
}
