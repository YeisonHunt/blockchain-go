package contract

type SmartContract interface {
    Execute(input []byte) []byte
}

type SimpleContract struct {
    State map[string]string
}

func (sc *SimpleContract) Execute(input []byte) []byte {
    // Example: simple key-value storage
    key := string(input[:len(input)/2])
    value := string(input[len(input)/2:])
    sc.State[key] = value
    return []byte("Executed")
}