package jsonHelper

func Mapping(source any, target any) error {
	var converted string
	var err error
	if converted, err = ConvertToString(source); err != nil {
		return err
	} else {
		if err = Unmarshal(converted, &target); err != nil {
			return err
		}
		return nil
	}
}
