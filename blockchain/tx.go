
package blockchain

type TxOutput struct {
    Value int

    PubKey string
}

//Important to note that each output is Indivisible.
//You cannot "make change" with any output.
//If the Value is 10, in order to give someone 5, we need to make two five coin outputs.

type TxInput struct {
    ID []byte
    Out int
    Sig string

}

func (in *TxInput) CanUnlock(data string) bool {
    return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
    return out.PubKey == data
}