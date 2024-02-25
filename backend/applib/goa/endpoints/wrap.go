package goaendpoints

import (
	"fmt"
	"reflect"

	goa "goa.design/goa/v3/pkg"
)

// Wrap は endpoints で指定された struct の各フィールドのうち goa.Endpoint 型のフィールドに対して、
// エラーハンドラ eh を適用した関数を作成します。
// 戻り値は、元の endpoints と同じ型の struct で、 eh が適用された goa.Endpoint 型のフィールドを持ちます。
//
// この関数は reflect を使用しますが、サーバー起動時にのみ使用され、リクエスト処理時には使用されません。
// サーバー起動も著しく遅くなることはないので reflect を使っても問題ないと判断しました。
func Wrap[Endpoints any](endpoints Endpoints, wrapper func(goa.Endpoint) goa.Endpoint) Endpoints {
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
		var dstFunc goa.Endpoint = wrapper(fn)
		// dstField := dstVal.FieldByName(field.Name)
		dstField := dstVal.Field(i)
		dstField.Set(reflect.ValueOf(dstFunc))
	}

	return dstVal.Addr().Interface().(Endpoints)
}
