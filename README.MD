## Getting Started

ensure you have setup go in your local machine.
 [https://go.dev/doc/install](https://go.dev/doc/install)

### Prerequisites
Clone the repo
  ```sh
  git clone https://github.com/tkivite/myGoBlockchain.git
  ```

### Usage

 ```sh
go run main.go getbalance -address ADDRESS - get balance for ADDRESS
go run main.go createblockchain -address ADDRESS creates a blockchain and rewards the mining fee
go run main.go printchain - Prints the blocks in the chain
go run main.go send -from FROM -to TO -amount AMOUNT - Send amount of coins from one address to another
go run main.go createwallet - Creates a new wallet
go run main.go listaddresses - Lists the addresses in the wallet file
  ```

