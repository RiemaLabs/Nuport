// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AttestationProof is an auto generated low-level Go binding around an user-defined struct.
type AttestationProof struct {
	TupleRootNonce *big.Int
	Tuple          DataRootTuple
	Proof          BinaryMerkleProof
}

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// Namespace is an auto generated low-level Go binding around an user-defined struct.
type Namespace struct {
	Version [1]byte
	Id      [28]byte
}

// NamespaceMerkleMultiproof is an auto generated low-level Go binding around an user-defined struct.
type NamespaceMerkleMultiproof struct {
	BeginKey  *big.Int
	EndKey    *big.Int
	SideNodes []NamespaceNode
}

// NamespaceNode is an auto generated low-level Go binding around an user-defined struct.
type NamespaceNode struct {
	Min    Namespace
	Max    Namespace
	Digest []byte
}

// SharesProof is an auto generated low-level Go binding around an user-defined struct.
type SharesProof struct {
	Data             [][]byte
	ShareProofs      []NamespaceMerkleMultiproof
	Namespace        Namespace
	RowRoots         []NamespaceNode
	RowProofs        []BinaryMerkleProof
	AttestationProof AttestationProof
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"binaryMerkleTree_verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumBinaryMerkleTree.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDAOracle\",\"name\":\"_bridge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"}],\"name\":\"getEncode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"}],\"name\":\"verifyMultiRowRootsToDataRootTupleRootProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumDAVerifier.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDAOracle\",\"name\":\"_bridge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"}],\"name\":\"verifySharesToDataRootTupleRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumDAVerifier.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"}],\"name\":\"verifySharesToDataRootTupleRootProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumDAVerifier.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506116d68061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c8063126bc58314610059578063196c47a3146100825780635ed2a7e2146100a35780636d910d69146100b65780637e75e3d7146100c9575b5f80fd5b61006c61006736600461107f565b6100ea565b6040516100799190611123565b60405180910390f35b610095610090366004611135565b610117565b604051610079929190611182565b6100956100b136600461107f565b610141565b6100956100c4366004611135565b610159565b6100dc6100d73660046111a8565b61017a565b6040516100799291906111e9565b606082826040516020016100ff9291906113f6565b60405160208183030381529060405290505b92915050565b5f80610138836060015184608001518560a0015160200151602001516101b6565b91509150915091565b5f8061014d84846102bb565b915091505b9250929050565b5f80610138836060015184608001518560a00151602001516020015161031c565b5f8061014d838560a00151604001518660a00151602001516040516020016101a29190611502565b60405160208183030381529060405261034f565b5f8084518451146101cc57505f905060066102b3565b5f5b84518110156102aa575f6101fd8783815181106101ed576101ed611519565b60200260200101515f0151610440565b61022388848151811061021257610212611519565b602002602001015160200151610440565b88848151811061023557610235611519565b6020026020010151604001516040516020016102539392919061152d565b60405160208183030381529060405290505f6102898688858151811061027b5761027b611519565b60200260200101518461034f565b509050806102a0575f6002945094505050506102b3565b50506001016101ce565b5060015f915091505b935093915050565b5f805f806102d786866060015187608001518860a00151610487565b91509150816102ec575f935091506101529050565b5f8061030d876060015188608001518960a00151602001516020015161031c565b90999098509650505050505050565b5f805f8061032b8787876101b6565b9150915081610340575f935091506102b39050565b506001965f9650945050505050565b5f806001846040015111610375578351511561037057505f905060016102b3565b61039a565b61038784602001518560400151610538565b8451511461039a57505f905060016102b3565b83604001518460200151106103b457505f905060026102b3565b5f6103be846105cb565b8551519091505f036103ee5784604001516001036103e257851491505f90506102b3565b5f8092509250506102b3565b5f8061040787602001518860400151858a5f015161063c565b90925090505f81600581111561041f5761041f61116e565b14610431575f945092506102b3915050565b50909514955f95509350505050565b80516020808301516040516001600160f81b03199093169183019190915263ffffffff191660218201525f90603d0160405160208183030381529060405261011190611562565b805160208201516040808401519051631f3302a960e01b81525f9384936001600160a01b038a1693631f3302a9936104c493929190600401611595565b602060405180830381865afa1580156104df573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061050391906115cd565b61051257505f9050600461052f565b5f8061052787878760200151602001516101b6565b909450925050505b94509492505050565b5f6001821161054857505f610111565b610551826107b5565b61055d90610100611600565b90505f61056b600183611600565b6001901b905060018161057e9190611600565b841161058a5750610111565b8060010361059c576001915050610111565b6105b86105a98286611600565b6105b38386611600565b610538565b6105c3906001611613565b915050610111565b5f60025f60f81b836040516020016105e4929190611626565b60408051601f19818403018152908290526105fe91611656565b602060405180830381855afa158015610619573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906101119190611671565b5f80845f036106505750829050600361052f565b846001036106745782511561066a5750829050600461052f565b508290505f61052f565b82515f036106875750829050600561052f565b5f610691866107e1565b90505f6106ac855f600188516106a79190611600565b61081f565b90505f8289101561072c575f6106c48a858a8661063c565b925090505f8260058111156106db576106db61116e565b146106ee5750869450925061052f915050565b61071e818860018a516107019190611600565b8151811061071157610711611519565b60200260200101516109a8565b5f955095505050505061052f565b5f61074b61073a858c611600565b610744868c611600565b8a8661063c565b925090505f8260058111156107625761076261116e565b146107755750869450925061052f915050565b6107a587600189516107879190611600565b8151811061079757610797611519565b6020026020010151826109a8565b9a5f9a5098505050505050505050565b5f5b81816001901b10156107d5576107ce600182611613565b90506107b7565b61011181610100611600565b5f60018210156107ef575f80fd5b5f6107f983610a23565b90505f610807600183611600565b6001901b90508381036108185760011c5b9392505050565b6060818311156108895760405162461bcd60e51b815260206004820152602a60248201527f496e76616c69642072616e67653a205f626567696e2069732067726561746572604482015269081d1a185b8817d95b9960b21b60648201526084015b60405180910390fd5b83518311806108985750835182115b156108fd5760405162461bcd60e51b815260206004820152602f60248201527f496e76616c69642072616e67653a205f626567696e206f72205f656e6420617260448201526e65206f7574206f6620626f756e647360881b6064820152608401610880565b5f6109088484611600565b6001600160401b0381111561091f5761091f610a45565b604051908082528060200260200182016040528015610948578160200160208202803683370190505b509050835b8381101561099f5785818151811061096757610967611519565b602002602001015182868361097c9190611600565b8151811061098c5761098c611519565b602090810291909101015260010161094d565b50949350505050565b604051600160f81b602082015260218101839052604181018290525f9060029060610160408051601f19818403018152908290526109e591611656565b602060405180830381855afa158015610a00573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906108189190611671565b5f805b82156101115780610a3681611688565b915050600183901c9250610a26565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715610a7b57610a7b610a45565b60405290565b604051606081016001600160401b0381118282101715610a7b57610a7b610a45565b60405160c081016001600160401b0381118282101715610a7b57610a7b610a45565b604051601f8201601f191681016001600160401b0381118282101715610aed57610aed610a45565b604052919050565b5f6001600160401b03821115610b0d57610b0d610a45565b5060051b60200190565b5f82601f830112610b26575f80fd5b81356001600160401b03811115610b3f57610b3f610a45565b610b52601f8201601f1916602001610ac5565b818152846020838601011115610b66575f80fd5b816020850160208301375f918101602001919091529392505050565b5f82601f830112610b91575f80fd5b81356020610ba6610ba183610af5565b610ac5565b82815260059290921b84018101918181019086841115610bc4575f80fd5b8286015b84811015610c015780356001600160401b03811115610be5575f80fd5b610bf38986838b0101610b17565b845250918301918301610bc8565b509695505050505050565b5f60408284031215610c1c575f80fd5b610c24610a59565b905081356001600160f81b031981168114610c3d575f80fd5b8152602082013563ffffffff1981168114610c56575f80fd5b602082015292915050565b5f82601f830112610c70575f80fd5b81356020610c80610ba183610af5565b82815260059290921b84018101918181019086841115610c9e575f80fd5b8286015b84811015610c015780356001600160401b0380821115610cc0575f80fd5b9088019060a0828b03601f1901811315610cd8575f80fd5b610ce0610a81565b610cec8c898601610c0c565b8152610cfb8c60608601610c0c565b81890152908301359082821115610d10575f80fd5b610d1e8c8984870101610b17565b60408201528652505050918301918301610ca2565b5f82601f830112610d42575f80fd5b81356020610d52610ba183610af5565b82815260059290921b84018101918181019086841115610d70575f80fd5b8286015b84811015610c015780356001600160401b0380821115610d92575f80fd5b908801906060828b03601f1901811315610daa575f80fd5b610db2610a81565b838801358152604080850135828a0152918401359183831115610dd3575f80fd5b610de18d8a85880101610c61565b908201528652505050918301918301610d74565b5f60608284031215610e05575f80fd5b610e0d610a81565b905081356001600160401b03811115610e24575f80fd5b8201601f81018413610e34575f80fd5b80356020610e44610ba183610af5565b82815260059290921b83018101918181019087841115610e62575f80fd5b938201935b83851015610e8057843582529382019390820190610e67565b808652505080850135818501525050506040820135604082015292915050565b5f82601f830112610eaf575f80fd5b81356020610ebf610ba183610af5565b82815260059290921b84018101918181019086841115610edd575f80fd5b8286015b84811015610c015780356001600160401b03811115610efe575f80fd5b610f0c8986838b0101610df5565b845250918301918301610ee1565b5f8183036080811215610f2b575f80fd5b610f33610a81565b8335815291506040601f1982011215610f4a575f80fd5b50610f53610a59565b602083810135825260408401358282015282015260608201356001600160401b03811115610f7f575f80fd5b610f8b84828501610df5565b60408301525092915050565b5f60e08284031215610fa7575f80fd5b610faf610aa3565b905081356001600160401b0380821115610fc7575f80fd5b610fd385838601610b82565b83526020840135915080821115610fe8575f80fd5b610ff485838601610d33565b60208401526110068560408601610c0c565b6040840152608084013591508082111561101e575f80fd5b61102a85838601610c61565b606084015260a0840135915080821115611042575f80fd5b61104e85838601610ea0565b608084015260c0840135915080821115611066575f80fd5b5061107384828501610f1a565b60a08301525092915050565b5f8060408385031215611090575f80fd5b82356001600160a01b03811681146110a6575f80fd5b915060208301356001600160401b038111156110c0575f80fd5b6110cc85828601610f97565b9150509250929050565b5f5b838110156110f05781810151838201526020016110d8565b50505f910152565b5f815180845261110f8160208601602086016110d6565b601f01601f19169290920160200192915050565b602081525f61081860208301846110f8565b5f60208284031215611145575f80fd5b81356001600160401b0381111561115a575f80fd5b61116684828501610f97565b949350505050565b634e487b7160e01b5f52602160045260245ffd5b821515815260408101600b831061119b5761119b61116e565b8260208301529392505050565b5f80604083850312156111b9575f80fd5b82356001600160401b038111156111ce575f80fd5b6111da85828601610f97565b95602094909401359450505050565b8215158152604081016006831061119b5761119b61116e565b80516001600160f81b031916825260209081015163ffffffff1916910152565b5f82825180855260208086019550808260051b8401018186015f5b8481101561129d57601f19868403018952815160a061125d858351611202565b85820151604061126f81880183611202565b8084015193505050806080860152611289818601836110f8565b9a86019a945050509083019060010161123d565b5090979650505050505050565b5f82825180855260208086019550808260051b8401018186015f5b8481101561129d57858303601f1901895281518051845284810151858501526040908101516060918501829052906112ff81860183611222565b9a86019a94505050908301906001016112c5565b8051606080845281519084018190525f91602091908201906080860190845b8181101561134e57835183529284019291840191600101611332565b5050602085015160208701526040850151604087015280935050505092915050565b5f8282518085526020808601955060208260051b840101602086015f5b8481101561129d57601f198684030189526113a9838351611313565b9884019892509083019060010161138d565b805182525f60208201516113dc602085018280518252602090810151910152565b506040820151608060608501526111666080850182611313565b6001600160a01b038316815260406020808301829052835160e092840192909252815161012084018190525f92610140600583901b86018101929184019190860190855b818110156114695761013f198886030183526114578585516110f8565b9450928501929185019160010161143a565b5050505060208501519150603f198085830301606086015261148b82846112aa565b9250604086015191506114a16080860183611202565b60608601519150808584030160c08601526114bc8383611222565b925060808601519150808584030160e08601526114d98383611370565b925060a086015191508085840301610100860152506114f882826113bb565b9695505050505050565b815181526020808301519082015260408101610111565b634e487b7160e01b5f52603260045260245ffd5b62ffffff1984811682528316601d82015281515f9061155381603a8501602087016110d6565b91909101603a01949350505050565b8051602082015162ffffff19808216929190601d83101561158d57808184601d0360031b1b83161693505b505050919050565b8381526115af602082018480518252602090810151910152565b608060608201525f6115c46080830184611313565b95945050505050565b5f602082840312156115dd575f80fd5b81518015158114610818575f80fd5b634e487b7160e01b5f52601160045260245ffd5b81810381811115610111576101116115ec565b80820180821115610111576101116115ec565b6001600160f81b03198316815281515f906116488160018501602087016110d6565b919091016001019392505050565b5f82516116678184602087016110d6565b9190910192915050565b5f60208284031215611681575f80fd5b5051919050565b5f60018201611699576116996115ec565b506001019056fea2646970667358221220804036eed49176f7bede5a9054e567aca4078a19e682036c7567a272b2971eec64736f6c63430008160033",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// BindingsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingsMetaData.Bin instead.
var BindingsBin = BindingsMetaData.Bin

// DeployBindings deploys a new Ethereum contract, binding an instance of Bindings to it.
func DeployBindings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bindings, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// BinaryMerkleTreeVerify is a free data retrieval call binding the contract method 0x7e75e3d7.
//
// Solidity: function binaryMerkleTree_verify((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 root) pure returns(bool, uint8)
func (_Bindings *BindingsCaller) BinaryMerkleTreeVerify(opts *bind.CallOpts, _sharesProof SharesProof, root [32]byte) (bool, uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "binaryMerkleTree_verify", _sharesProof, root)

	if err != nil {
		return *new(bool), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// BinaryMerkleTreeVerify is a free data retrieval call binding the contract method 0x7e75e3d7.
//
// Solidity: function binaryMerkleTree_verify((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 root) pure returns(bool, uint8)
func (_Bindings *BindingsSession) BinaryMerkleTreeVerify(_sharesProof SharesProof, root [32]byte) (bool, uint8, error) {
	return _Bindings.Contract.BinaryMerkleTreeVerify(&_Bindings.CallOpts, _sharesProof, root)
}

// BinaryMerkleTreeVerify is a free data retrieval call binding the contract method 0x7e75e3d7.
//
// Solidity: function binaryMerkleTree_verify((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 root) pure returns(bool, uint8)
func (_Bindings *BindingsCallerSession) BinaryMerkleTreeVerify(_sharesProof SharesProof, root [32]byte) (bool, uint8, error) {
	return _Bindings.Contract.BinaryMerkleTreeVerify(&_Bindings.CallOpts, _sharesProof, root)
}

// GetEncode is a free data retrieval call binding the contract method 0x126bc583.
//
// Solidity: function getEncode(address _bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bytes)
func (_Bindings *BindingsCaller) GetEncode(opts *bind.CallOpts, _bridge common.Address, _sharesProof SharesProof) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getEncode", _bridge, _sharesProof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncode is a free data retrieval call binding the contract method 0x126bc583.
//
// Solidity: function getEncode(address _bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bytes)
func (_Bindings *BindingsSession) GetEncode(_bridge common.Address, _sharesProof SharesProof) ([]byte, error) {
	return _Bindings.Contract.GetEncode(&_Bindings.CallOpts, _bridge, _sharesProof)
}

// GetEncode is a free data retrieval call binding the contract method 0x126bc583.
//
// Solidity: function getEncode(address _bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bytes)
func (_Bindings *BindingsCallerSession) GetEncode(_bridge common.Address, _sharesProof SharesProof) ([]byte, error) {
	return _Bindings.Contract.GetEncode(&_Bindings.CallOpts, _bridge, _sharesProof)
}

// VerifyMultiRowRootsToDataRootTupleRootProof is a free data retrieval call binding the contract method 0x196c47a3.
//
// Solidity: function verifyMultiRowRootsToDataRootTupleRootProof((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bool, uint8)
func (_Bindings *BindingsCaller) VerifyMultiRowRootsToDataRootTupleRootProof(opts *bind.CallOpts, _sharesProof SharesProof) (bool, uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "verifyMultiRowRootsToDataRootTupleRootProof", _sharesProof)

	if err != nil {
		return *new(bool), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// VerifyMultiRowRootsToDataRootTupleRootProof is a free data retrieval call binding the contract method 0x196c47a3.
//
// Solidity: function verifyMultiRowRootsToDataRootTupleRootProof((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bool, uint8)
func (_Bindings *BindingsSession) VerifyMultiRowRootsToDataRootTupleRootProof(_sharesProof SharesProof) (bool, uint8, error) {
	return _Bindings.Contract.VerifyMultiRowRootsToDataRootTupleRootProof(&_Bindings.CallOpts, _sharesProof)
}

// VerifyMultiRowRootsToDataRootTupleRootProof is a free data retrieval call binding the contract method 0x196c47a3.
//
// Solidity: function verifyMultiRowRootsToDataRootTupleRootProof((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bool, uint8)
func (_Bindings *BindingsCallerSession) VerifyMultiRowRootsToDataRootTupleRootProof(_sharesProof SharesProof) (bool, uint8, error) {
	return _Bindings.Contract.VerifyMultiRowRootsToDataRootTupleRootProof(&_Bindings.CallOpts, _sharesProof)
}

// VerifySharesToDataRootTupleRoot is a free data retrieval call binding the contract method 0x5ed2a7e2.
//
// Solidity: function verifySharesToDataRootTupleRoot(address _bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) view returns(bool, uint8)
func (_Bindings *BindingsCaller) VerifySharesToDataRootTupleRoot(opts *bind.CallOpts, _bridge common.Address, _sharesProof SharesProof) (bool, uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "verifySharesToDataRootTupleRoot", _bridge, _sharesProof)

	if err != nil {
		return *new(bool), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// VerifySharesToDataRootTupleRoot is a free data retrieval call binding the contract method 0x5ed2a7e2.
//
// Solidity: function verifySharesToDataRootTupleRoot(address _bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) view returns(bool, uint8)
func (_Bindings *BindingsSession) VerifySharesToDataRootTupleRoot(_bridge common.Address, _sharesProof SharesProof) (bool, uint8, error) {
	return _Bindings.Contract.VerifySharesToDataRootTupleRoot(&_Bindings.CallOpts, _bridge, _sharesProof)
}

// VerifySharesToDataRootTupleRoot is a free data retrieval call binding the contract method 0x5ed2a7e2.
//
// Solidity: function verifySharesToDataRootTupleRoot(address _bridge, (bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) view returns(bool, uint8)
func (_Bindings *BindingsCallerSession) VerifySharesToDataRootTupleRoot(_bridge common.Address, _sharesProof SharesProof) (bool, uint8, error) {
	return _Bindings.Contract.VerifySharesToDataRootTupleRoot(&_Bindings.CallOpts, _bridge, _sharesProof)
}

// VerifySharesToDataRootTupleRootProof is a free data retrieval call binding the contract method 0x6d910d69.
//
// Solidity: function verifySharesToDataRootTupleRootProof((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bool, uint8)
func (_Bindings *BindingsCaller) VerifySharesToDataRootTupleRootProof(opts *bind.CallOpts, _sharesProof SharesProof) (bool, uint8, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "verifySharesToDataRootTupleRootProof", _sharesProof)

	if err != nil {
		return *new(bool), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// VerifySharesToDataRootTupleRootProof is a free data retrieval call binding the contract method 0x6d910d69.
//
// Solidity: function verifySharesToDataRootTupleRootProof((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bool, uint8)
func (_Bindings *BindingsSession) VerifySharesToDataRootTupleRootProof(_sharesProof SharesProof) (bool, uint8, error) {
	return _Bindings.Contract.VerifySharesToDataRootTupleRootProof(&_Bindings.CallOpts, _sharesProof)
}

// VerifySharesToDataRootTupleRootProof is a free data retrieval call binding the contract method 0x6d910d69.
//
// Solidity: function verifySharesToDataRootTupleRootProof((bytes[],(uint256,uint256,((bytes1,bytes28),(bytes1,bytes28),bytes)[])[],(bytes1,bytes28),((bytes1,bytes28),(bytes1,bytes28),bytes)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof) pure returns(bool, uint8)
func (_Bindings *BindingsCallerSession) VerifySharesToDataRootTupleRootProof(_sharesProof SharesProof) (bool, uint8, error) {
	return _Bindings.Contract.VerifySharesToDataRootTupleRootProof(&_Bindings.CallOpts, _sharesProof)
}
