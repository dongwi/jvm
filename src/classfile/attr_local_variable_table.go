package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*localVariableTableEntry
}

type localVariableTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (self *LocalVariableTableAttribute)readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*localVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &localVariableTableEntry{
			startPc:reader.readUint16(),
			lineNumber:reader.readUint16(),
		}
	}
}
