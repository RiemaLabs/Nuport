# Nuport Fraud Proofs

## Nuport Intro

A Nuport rollup is a blockchain that uses Nubit for data availability but settles on any EVM chain. Nuport operates by having the Nubit validator set periodically sign over batched data commitments and validator set updates, which are relayed an EVM smart contract. The data commitments are stored in the EVM chain's state, and can be used to prove inclusion of any data historically posted to Nubit.

## Fraud Proofs

Fraud proofs can be used to inform light clients (including on-chain smart contract light clients) in the case of an invalid rollup state transition or unavailable rollup block data—specifically rollup block data that is claimed to be on Nubit but is not. They rely on rollup full nodes getting the data that was published to Nubit, and executing all the state transitions to verify the rollup state. If they discover an invalid state transition or unavailable rollup block, they emit a fraud proof with the necessary information to convince light clients that fraud happened. This allows for trust-minimized light clients, as the network only needs one honest full node to create the fraud proof and propagate it.

## Rollup Header

Rollups can adopt many approaches to prove that fraud occurred. One of which could be having the following fields in the rollup header:

- Rollup block state root
- A sequence of spans in Nubit: which references where the rollup data was published in the Nubit chain.

> [!NOTE]  
> The sequence of spans can be defined using the following: `height`, `start index`, and `length` in the Nubit block, in the case of a single Nubit block. However, it could be generalized to span over multiple blocks.

For the rest of the document, we will suppose that the sequence of spans only references one Nubit block.

## Proving Unavailable Data

By construction, the rollup block data **is the sequence of spans defined in the header**. Thus to prove that the rollup data is unavailable, it is necessary and sufficient to show that the sequence of spans doesn't belong to the Nubit block, i.e. the span is out of bounds.

We could prove that via creating a binary [Merkle proof](https://github.com/RiemaLabs/nubit-core/blob/0424f0816e5d27806b7aed8c63cba38e65f89525/crypto/merkle/proof.go#L19-L31) of any row/column to the Nubit data root. This proof will provide the `total` which is the number of rows/columns in the extended data square. This can be used to calculate the square size.

Then, we will use that information to check if the provided transaction index, in the header, is out of the square size bounds.

For the data root, we will use a binary Merkle proof to prove its inclusion in a data root tuple root that was committed to by the Nuport smart contract. More on this in [here](#1-data-root-inclusion-proof).

## Proving an Invalid State Transition

In order to prove an invalid transaction in the rollup, we need to prove the following:

- Prove that the transaction was posted to Nubit, and
- Prove that the transaction is invalid. This is left to the rollup to define.

The first part, proving that the transaction was posted to Nubit, can be done in three steps:

1. Prove that the data root tuple is committed to by the Nuport smart contract
2. Verify inclusion proof of the transaction to Nubit data root
3. Prove that the transaction is in the rollup sequence spans

### 1. Data root inclusion proof

To prove the data root is committed to by the Nuport smart contract, we will need to provide a Merkle proof of the data root tuple to a data root tuple root. This can be created using the [`data_root_inclusion_proof`](https://github.com/RiemaLabs/nubit-core/blob/0424f0816e5d27806b7aed8c63cba38e65f89525/rpc/client/http/http.go#L519-L538) query.

### 2. Transaction inclusion proof

To prove that a rollup transaction is part of the data root, we will need to provide two proofs: a namespace Merkle proof of the transaction to a row root. This could be done via proving the shares that contain the transaction to the row root using a namespace Merkle proof. And, a binary Merkle proof of the row root to the data root.

These proofs can be generated using the [`ProveShares`](https://github.com/RiemaLabs/nubit-core/blob/0424f0816e5d27806b7aed8c63cba38e65f89525/rpc/client/http/http.go#L553-L570) query.

### 3. Transaction part of the rollup sequence

To prove that a transaction is part of the rollup sequence of spans, we take the authenticated share proof and use the shares begin/end key to define the share position in the row.

Then, we use the row proof to get the row index in the extended Nubit square and get the index of the share in row major order:

```solidity
uint256 shareIndexInRow = shareProof.shareProofs[0].beginKey;
uint256 shareIndexInRowMajorOrder = shareIndexInRow + shareProof.rowProofs[0].numLeaves * shareProof.rowProofs[0].key;
```

Finally, we can compare the computed index with the rollup header sequence of spans, and be sure that the share/transaction is part of the rollup data.

Check the `RollupInclusionProofs.t.sol` for an example.
