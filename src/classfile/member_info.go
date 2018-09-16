package classfile

type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16//方法或字段的访问域
	nameIndex	uint16//字段名或方法名在常量池中的索引
	descriptorIndex uint16//字段或方法的描述符，在常量池中的索引
	attributes  []AttributeInfo//字段或方法的属性表
}

func readMembers(reader *ClassReader, cp ConstantPool) [] *MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp : cp,
		accessFlags:reader.readUint16(),
		nameIndex:reader.readUint16(),
		descriptorIndex:reader.readUint16(),
		attributes:readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlag() uint16 {
	return self.accessFlags
}
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}