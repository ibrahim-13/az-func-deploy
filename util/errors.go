package util

func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
