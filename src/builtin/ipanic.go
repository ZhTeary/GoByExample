package builtin

import "os"

func Ipanic() {
	panic("a problem")

	_, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
	}
}
