package HTTPErrors

type Error struct {
    Message     string
    Code        int
}

// Takes a message and code and returns it as an Error struct
func NewError(msg string, code int) Error {
    return Error{
        Message: msg,
        Code:    code,
    }
}