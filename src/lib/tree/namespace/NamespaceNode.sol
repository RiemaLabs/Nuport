// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "../Types.sol";

/// @notice Namespace Merkle Tree node.
struct NamespaceNode {
    // Minimum namespace.
    Namespace min;
    // Maximum namespace.
    Namespace max;
    // Node value.
    bytes digest;
}

/// @notice Compares two `NamespaceNode`s.
/// @param first First node.
/// @param second Second node.
/// @return `true` is equal, `false otherwise.
// solhint-disable-next-line func-visibility
function namespaceNodeEquals(NamespaceNode memory first, NamespaceNode memory second) pure returns (bool) {
    if (!first.min.equalTo(second.min)) {
        return false;
    }
    if (!first.max.equalTo(second.max)) {
        return false;
    }
    if (first.digest.length != second.digest.length) {
        return false;
    }
    for (uint i = 0; i < first.digest.length; i++) {
        if (first.digest[i] != second.digest[i]) {
            return false;
        }
    }
    return true;
    //return first.min.equalTo(second.min) && first.max.equalTo(second.max) && (first.digest == second.digest);
}
