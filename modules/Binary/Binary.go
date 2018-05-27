package Binary

import "reflect"

func GetExtensionName(data []byte) string {
	pngHeader := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	jpgHeader := []byte{255, 216}
	gifHeader := []byte{71, 73, 70}

	if reflect.DeepEqual(data[:len(pngHeader)], pngHeader) {
		return ".png"
	}
	if reflect.DeepEqual(data[:len(jpgHeader)], jpgHeader) {
		return ".jpg"
	}
	if reflect.DeepEqual(data[:len(gifHeader)], jpgHeader) {
		return ".gif"
	}
	return ".unknown"
}