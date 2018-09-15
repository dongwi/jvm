package classpath

import (
	"os"
	"strings"
)

const pathListSeparetor  = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparetor) {
		return  newCompositeEntry(path)
	}
	if strings.Contains(path, "*") {
		return newWildCardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
			return newZipEntry(path)
	}
	return newDirEntry(path)
}



