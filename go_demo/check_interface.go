package go_demo



type Buffer interface {
	Write(b []byte) error
}


type ImpBuffer struct {

}

//func (*ImpBuffer) Writee(b []byte) error {
//	return nil
//}

func (*ImpBuffer) Write(b []byte) error {
	return nil
}

// 直到运行时才会报出这个错误，如果想编译的时候就进行报错 ⬇️
// Cannot use '(*ImpBuffer)(nil)' (type *ImpBuffer) as the type Buffer Type does not implement 'Buffer

var _ Buffer = (*ImpBuffer)(nil)
