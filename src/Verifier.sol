// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "./lib/verifier/DAVerifier.sol";
import "./lib/tree/binary/BinaryMerkleTree.sol";
import "./IDAOracle.sol";

contract Verifier {
    function verifySharesToDataRootTupleRoot(IDAOracle _bridge, SharesProof memory _sharesProof)
        external
        view
        returns (bool, DAVerifier.ErrorCodes)
    {
        return DAVerifier.verifySharesToDataRootTupleRoot(_bridge, _sharesProof);
    }

    function binaryMerkleTree_verify(SharesProof memory _sharesProof, bytes32 root)
        external
        pure
        returns (bool, BinaryMerkleTree.ErrorCodes)
    {
        return BinaryMerkleTree.verify(root, _sharesProof.attestationProof.proof, abi.encode(_sharesProof.attestationProof.tuple));
    }

    function verifyMultiRowRootsToDataRootTupleRootProof(SharesProof memory _sharesProof)
        external
        pure
        returns (bool, DAVerifier.ErrorCodes)
    {
        return DAVerifier.verifyMultiRowRootsToDataRootTupleRootProof(_sharesProof.rowRoots, _sharesProof.rowProofs, _sharesProof.attestationProof.tuple.dataRoot);
    }

    function verifySharesToDataRootTupleRootProof(SharesProof memory _sharesProof)
        external
        pure
        returns (bool, DAVerifier.ErrorCodes)
    {
        return DAVerifier.verifySharesToDataRootTupleRootProof(
            _sharesProof.rowRoots,
            _sharesProof.rowProofs,
            _sharesProof.attestationProof.tuple.dataRoot
        );
    }
}
