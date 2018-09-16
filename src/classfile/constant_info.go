package classfile

const (
	CONSTANT_Class = 7
	CONSTANT_String = 8
	CONSTANT_Fieldref = 9
	CONSTANT_Methodref = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_Interger = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_NameAndType = 12
	CONSTANT_Utf8 = 1
	CONSTANT_MethodHande = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokdDynamic = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	return c
}
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Interger : return &ConstantIntergerInfo{}
	case CONSTANT_Float : return &ConstantFloatInfo{}
	case CONSTANT_Long : return &ConstantLongInfo{}
	case CONSTANT_Double : return &ConstantDoubleInfo{}
	case CONSTANT_Utf8 : return &ConstantUtf8Info{}
	case CONSTANT_String : return &ConstantStringInfo{cp:cp}
	case CONSTANT_Class : return &ConstantClassInfo{cp:cp}
	case CONSTANT_Fieldref : return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp:cp}}
	case CONSTANT_Methodref : return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp:cp}}
	case CONSTANT_NameAndType : return &ConstantNameAndTypeInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}