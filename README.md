# Nuport-contracts

<!-- markdownlint-disable MD013 MD041 -->

Nuport is a Nubit -> EVM message relay.
It is based on Umee's Gravity Bridge implementation, [Peggo](https://github.com/umee-network/peggo).
**This project is under active development and should not be used in production**.

## Table of Contents

- [Building From Source](#building-from-source)
- [Send a message from Nubit to an EVM chain](#send-a-message-from-nubit-to-an-evm-chain)
- [How it works](#how-it-works)

## Building From Source

### Dependencies

Initialize git submodules, needed for Forge dependencies:

```sh
git submodule init
git submodule update
```

To regenerate the Go ABI wrappers with `make gen`, you need the `abigen` tool.
Building requires [Go 1.19+](https://golang.org/dl/).
Install `abigen` with:

```sh
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
make devtools
```

### Build and Test Contracts

Build with:

```sh
forge build
```

Test with:

```sh
forge test
```

### Format

Format Solidity with:

```sh
forge fmt
```

### Regenerate Go Wrappers

Go wrappers can be regenerated with:

```sh
make
```

## Send a message from Nubit to an EVM chain

A message can be included on Nubit by using the Nubit validator.
Instructions [here](https://github.com/RiemaLabs/nubit-validator).

## How it works

Nuport allows Nubit block header data roots to be relayed in one direction, from Nubit to an EVM chain.
It does not support bridging assets such as fungible or non-fungible tokens directly, and cannot send messages from the EVM chain back to Nubit.

It works by relying on a set of signers to attest to some event on Nubit: the Nubit validator set.
Nuport contract keeps track of the Nubit validator set by updating its view of the validator set with `updateValidatorSet()`.
More than 2/3 of the voting power of the current view of the validator set must sign off on new relayed events, submitted with `submitDataRootTupleRoot()`.
Each event is a batch of `DataRootTuple`s, with each tuple representing a single data root (i.e. block header).
Relayed tuples are in the same order as Nubit block headers.

### Events and messages relayed

 **Validator sets**:
 The relayer informs the Nuport contract who are the current validators and their power.
 This results in an execution of the `updateValidatorSet` function.

 **Batches**:
 The relayer informs the Nuport contract of new data root tuple roots.
 This results in an execution of the `submitDataRootTupleRoot` function.
