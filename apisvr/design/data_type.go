package design

type DataType string

const (
	DataTypeResultType DataType = "result_type"
	DataTypePayload    DataType = "payload"
)

func (dt DataType) IsRT() bool {
	return dt == DataTypeResultType
}

func (dt DataType) IsPayload() bool {
	return dt == DataTypePayload
}
