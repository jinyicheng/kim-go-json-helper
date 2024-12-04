package jsonHelper

func Unmarshal(source string, target any) error {
	return json.Unmarshal([]byte(source), &target)
}
