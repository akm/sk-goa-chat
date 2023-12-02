package design

import (
	. "goa.design/goa/v3/dsl"
)

// APIに指定した引数は ***srvc を定義するファイルののパッケージ名として使われる
var _ = API("chatapi", func() {
	Title("Chat API")
	Version("0.1")
	Description("Chat API")

	// Serverに指定した引数は goa example で cmd 以下に作られる ディレクトリ名として使われる
	Server("apisvr", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})

	HTTP(func() {
		Path("/")
	})
})
