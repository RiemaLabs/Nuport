// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

/// @notice A tuple of data root with metadata. Each data root is associated
///  with a Nubit block height.
struct DataRootTuple {
    // Nubit block height the data root was included in.
    // Genesis block is height = 0.
    // First queryable block is height = 1.
    uint256 height;
    // Data root.
    bytes32 dataRoot;
}
