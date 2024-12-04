package jsonHelper

func ConvertToBytes(source any) ([]byte, error) {
	return json.Marshal(source)
}
