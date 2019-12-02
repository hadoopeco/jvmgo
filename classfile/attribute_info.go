package classfile

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/23 19:29

attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
var (
    _attrDeprecated = &DeprecatedAttribute{}
    _attrSynthetic  = &SyntheticAttribute{}
)


 type AttributeInfo interface {
       readInfo(reader *ClassReader)
 }


 func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo{
     attributesCount :=  reader.readUint16()
     attributes := make([]AttributeInfo, attributesCount)
     for i := range attributes {
       attributes[i] = readAttribute(reader, cp)
     }
     return  attributes
 }

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
     attributeNameIndex := reader.readUint16()
     attrName := cp.getUtf8(attributeNameIndex)
     attrlen := reader.readUint32()
     attrInfo := newAttributeInfo(attrName,attrlen,cp)
     attrInfo.readInfo(reader)

     return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
    switch attrName {
    //case "AnnotationDefault":
    case "BootstrapMethods":
        return &BootstrapMethodsAttribute{}
    case "Code":
        return &CodeAttribute{cp: cp}
    case "ConstantValue":
        return &ConstantValueAttribute{}
    case "Deprecated":
        return _attrDeprecated
    case "EnclosingMethod":
       return &EnclosingMethodAttribute{cp: cp}
    case "Exceptions":
        return &ExceptionsAttribute{}
    case "InnerClasses":
       return &InnerClassesAttribute{}
    case "LineNumberTable":
        return &LineNumberTableAttribute{}
    case "LocalVariableTable":
        return &LocalVariableTableAttribute{}
    case "LocalVariableTypeTable":
        return &LocalVariableTypeTableAttribute{}
        // case "MethodParameters":
        // case "RuntimeInvisibleAnnotations":
        // case "RuntimeInvisibleParameterAnnotations":
        // case "RuntimeInvisibleTypeAnnotations":
        // case "RuntimeVisibleAnnotations":
        // case "RuntimeVisibleParameterAnnotations":
        // case "RuntimeVisibleTypeAnnotations":
    //case "Signature":
    //    return &SignatureAttribute{cp: cp}
    case "SourceFile":
        return &SourceFileAttribute{cp: cp}
        // case "SourceDebugExtension":
        // case "StackMapTable":
    case "Synthetic":
        return _attrSynthetic
    default:
        return &UnparsedAttribute{attrName, attrLen, nil}
    }
}