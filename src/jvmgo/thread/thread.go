package thread

type Thread struct {
    // todo
    pc      uint32
    stack   *Stack
}

type Stack struct {
    frames  []*Frame
}

type Frame struct {

}
