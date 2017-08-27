# solidity-contract-governess
Solidity contract management using golang

Install geth from following link https://www.ethereum.org/cli

## Setup private ethereum net for the testing purpose.
```
$ geth --dev 
$ get attach 
$ personal.newAccount('testnet')
$ minor.start()
```

## Generate contract ABI using remix or any other tool.
```
$ abigen --abi token.abi --pkg main --type Token --out token.go
```
