package jsonHelper

func ConvertToString(source any) (string, error) {
	if bytes, err := ConvertToBytes(source); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
