package errors

type Error struct {
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Debug       string `protobuf:"bytes,3,opt,name=debug" json:"debug,omitempty"`
	Hash        string `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	Class       int32  `protobuf:"varint,6,opt,name=class" json:"class,omitempty"`
	Stack       string `protobuf:"bytes,7,opt,name=stack" json:"stack,omitempty"`
	Created     int64  `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Code        string `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	Base        *Error  `protobuf:"bytes,10,opt,name=base" json:"base,omitempty"`
	RequestId   string `protobuf:"bytes,12,opt,name=request_id" json:"request_id,omitempty"`
}
