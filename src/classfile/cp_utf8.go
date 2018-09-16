package classfile

type ConstantUtf8Info struct {
	str string
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader)  {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}
