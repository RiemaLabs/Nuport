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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"binaryMerkleTree_verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumBinaryMerkleTree.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"}],\"name\":\"verifyMultiRowRootsToDataRootTupleRootProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumDAVerifier.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDAOracle\",\"name\":\"_bridge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"}],\"name\":\"verifySharesToDataRootTupleRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumDAVerifier.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"namespace\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"digest\",\"type\":\"bytes\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"}],\"name\":\"verifySharesToDataRootTupleRootProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"enumDAVerifier.ErrorCodes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506113698061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c8063196c47a31461004e5780635ed2a7e2146100785780636d910d691461008b5780637e75e3d71461009e575b5f80fd5b61006161005c36600461102d565b6100bf565b60405161006f929190611072565b60405180910390f35b610061610086366004611098565b6100e9565b61006161009936600461102d565b610101565b6100b16100ac3660046110ef565b610122565b60405161006f929190611130565b5f806100e0836060015184608001518560a00151602001516020015161015e565b91509150915091565b5f806100f58484610263565b915091505b9250929050565b5f806100e0836060015184608001518560a0015160200151602001516102c4565b5f806100f5838560a00151604001518660a001516020015160405160200161014a9190611149565b6040516020818303038152906040526102f7565b5f80845184511461017457505f9050600661025b565b5f5b8451811015610252575f6101a587838151811061019557610195611160565b60200260200101515f01516103e8565b6101cb8884815181106101ba576101ba611160565b6020026020010151602001516103e8565b8884815181106101dd576101dd611160565b6020026020010151604001516040516020016101fb939291906111a1565b60405160208183030381529060405290505f6102318688858151811061022357610223611160565b6020026020010151846102f7565b50905080610248575f60029450945050505061025b565b5050600101610176565b5060015f915091505b935093915050565b5f805f8061027f86866060015187608001518860a00151610435565b9150915081610294575f935091506100fa9050565b5f806102b5876060015188608001518960a0015160200151602001516102c4565b90999098509650505050505050565b5f805f806102d387878761015e565b91509150816102e8575f9350915061025b9050565b506001965f9650945050505050565b5f80600184604001511161031d578351511561031857505f9050600161025b565b610342565b61032f846020015185604001516104e6565b8451511461034257505f9050600161025b565b836040015184602001511061035c57505f9050600261025b565b5f61036684610579565b8551519091505f0361039657846040015160010361038a57851491505f905061025b565b5f80925092505061025b565b5f806103af87602001518860400151858a5f01516105ea565b90925090505f8160058111156103c7576103c761105e565b146103d9575f9450925061025b915050565b50909514955f95509350505050565b80516020808301516040516001600160f81b03199093169183019190915263ffffffff191660218201525f90603d0160405160208183030381529060405261042f906111c9565b92915050565b805160208201516040808401519051631f3302a960e01b81525f9384936001600160a01b038a1693631f3302a993610472939291906004016111fc565b602060405180830381865afa15801561048d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b19190611285565b6104c057505f905060046104dd565b5f806104d5878787602001516020015161015e565b909450925050505b94509492505050565b5f600182116104f657505f61042f565b6104ff82610763565b61050b906101006112b8565b90505f6105196001836112b8565b6001901b905060018161052c91906112b8565b8411610538575061042f565b8060010361054a57600191505061042f565b61056661055782866112b8565b61056183866112b8565b6104e6565b6105719060016112cb565b949350505050565b5f60025f60f81b836040516020016105929291906112de565b60408051601f19818403018152908290526105ac916112f9565b602060405180830381855afa1580156105c7573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061042f9190611304565b5f80845f036105fe575082905060036104dd565b8460010361062257825115610618575082905060046104dd565b508290505f6104dd565b82515f03610635575082905060056104dd565b5f61063f8661078f565b90505f61065a855f6001885161065591906112b8565b6107cd565b90505f828910156106da575f6106728a858a866105ea565b925090505f8260058111156106895761068961105e565b1461069c575086945092506104dd915050565b6106cc818860018a516106af91906112b8565b815181106106bf576106bf611160565b6020026020010151610956565b5f95509550505050506104dd565b5f6106f96106e8858c6112b8565b6106f2868c6112b8565b8a866105ea565b925090505f8260058111156107105761071061105e565b14610723575086945092506104dd915050565b610753876001895161073591906112b8565b8151811061074557610745611160565b602002602001015182610956565b9a5f9a5098505050505050505050565b5f5b81816001901b10156107835761077c6001826112cb565b9050610765565b61042f816101006112b8565b5f600182101561079d575f80fd5b5f6107a7836109d1565b90505f6107b56001836112b8565b6001901b90508381036107c65760011c5b9392505050565b6060818311156108375760405162461bcd60e51b815260206004820152602a60248201527f496e76616c69642072616e67653a205f626567696e2069732067726561746572604482015269081d1a185b8817d95b9960b21b60648201526084015b60405180910390fd5b83518311806108465750835182115b156108ab5760405162461bcd60e51b815260206004820152602f60248201527f496e76616c69642072616e67653a205f626567696e206f72205f656e6420617260448201526e65206f7574206f6620626f756e647360881b606482015260840161082e565b5f6108b684846112b8565b6001600160401b038111156108cd576108cd6109f3565b6040519080825280602002602001820160405280156108f6578160200160208202803683370190505b509050835b8381101561094d5785818151811061091557610915611160565b602002602001015182868361092a91906112b8565b8151811061093a5761093a611160565b60209081029190910101526001016108fb565b50949350505050565b604051600160f81b602082015260218101839052604181018290525f9060029060610160408051601f1981840301815290829052610993916112f9565b602060405180830381855afa1580156109ae573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906107c69190611304565b5f805b821561042f57806109e48161131b565b915050600183901c92506109d4565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715610a2957610a296109f3565b60405290565b604051606081016001600160401b0381118282101715610a2957610a296109f3565b60405160c081016001600160401b0381118282101715610a2957610a296109f3565b604051601f8201601f191681016001600160401b0381118282101715610a9b57610a9b6109f3565b604052919050565b5f6001600160401b03821115610abb57610abb6109f3565b5060051b60200190565b5f82601f830112610ad4575f80fd5b81356001600160401b03811115610aed57610aed6109f3565b610b00601f8201601f1916602001610a73565b818152846020838601011115610b14575f80fd5b816020850160208301375f918101602001919091529392505050565b5f82601f830112610b3f575f80fd5b81356020610b54610b4f83610aa3565b610a73565b82815260059290921b84018101918181019086841115610b72575f80fd5b8286015b84811015610baf5780356001600160401b03811115610b93575f80fd5b610ba18986838b0101610ac5565b845250918301918301610b76565b509695505050505050565b5f60408284031215610bca575f80fd5b610bd2610a07565b905081356001600160f81b031981168114610beb575f80fd5b8152602082013563ffffffff1981168114610c04575f80fd5b602082015292915050565b5f82601f830112610c1e575f80fd5b81356020610c2e610b4f83610aa3565b82815260059290921b84018101918181019086841115610c4c575f80fd5b8286015b84811015610baf5780356001600160401b0380821115610c6e575f80fd5b9088019060a0828b03601f1901811315610c86575f80fd5b610c8e610a2f565b610c9a8c898601610bba565b8152610ca98c60608601610bba565b81890152908301359082821115610cbe575f80fd5b610ccc8c8984870101610ac5565b60408201528652505050918301918301610c50565b5f82601f830112610cf0575f80fd5b81356020610d00610b4f83610aa3565b82815260059290921b84018101918181019086841115610d1e575f80fd5b8286015b84811015610baf5780356001600160401b0380821115610d40575f80fd5b908801906060828b03601f1901811315610d58575f80fd5b610d60610a2f565b838801358152604080850135828a0152918401359183831115610d81575f80fd5b610d8f8d8a85880101610c0f565b908201528652505050918301918301610d22565b5f60608284031215610db3575f80fd5b610dbb610a2f565b905081356001600160401b03811115610dd2575f80fd5b8201601f81018413610de2575f80fd5b80356020610df2610b4f83610aa3565b82815260059290921b83018101918181019087841115610e10575f80fd5b938201935b83851015610e2e57843582529382019390820190610e15565b808652505080850135818501525050506040820135604082015292915050565b5f82601f830112610e5d575f80fd5b81356020610e6d610b4f83610aa3565b82815260059290921b84018101918181019086841115610e8b575f80fd5b8286015b84811015610baf5780356001600160401b03811115610eac575f80fd5b610eba8986838b0101610da3565b845250918301918301610e8f565b5f8183036080811215610ed9575f80fd5b610ee1610a2f565b8335815291506040601f1982011215610ef8575f80fd5b50610f01610a07565b602083810135825260408401358282015282015260608201356001600160401b03811115610f2d575f80fd5b610f3984828501610da3565b60408301525092915050565b5f60e08284031215610f55575f80fd5b610f5d610a51565b905081356001600160401b0380821115610f75575f80fd5b610f8185838601610b30565b83526020840135915080821115610f96575f80fd5b610fa285838601610ce1565b6020840152610fb48560408601610bba565b60408401526080840135915080821115610fcc575f80fd5b610fd885838601610c0f565b606084015260a0840135915080821115610ff0575f80fd5b610ffc85838601610e4e565b608084015260c0840135915080821115611014575f80fd5b5061102184828501610ec8565b60a08301525092915050565b5f6020828403121561103d575f80fd5b81356001600160401b03811115611052575f80fd5b61057184828501610f45565b634e487b7160e01b5f52602160045260245ffd5b821515815260408101600b831061108b5761108b61105e565b8260208301529392505050565b5f80604083850312156110a9575f80fd5b82356001600160a01b03811681146110bf575f80fd5b915060208301356001600160401b038111156110d9575f80fd5b6110e585828601610f45565b9150509250929050565b5f8060408385031215611100575f80fd5b82356001600160401b03811115611115575f80fd5b61112185828601610f45565b95602094909401359450505050565b8215158152604081016006831061108b5761108b61105e565b81518152602080830151908201526040810161042f565b634e487b7160e01b5f52603260045260245ffd5b5f81515f5b818110156111935760208185018101518683015201611179565b505f93019283525090919050565b62ffffff1984811682528316601d8201525f6111c0603a830184611174565b95945050505050565b8051602082015162ffffff19808216929190601d8310156111f457808184601d0360031b1b83161693505b505050919050565b8381525f6020611219602084018680518252602090810151910152565b60806060848101829052855191850152805160e085018190526020909101905f906101008601905b808310156112615783518252928401926001929092019190840190611241565b50602087015160a0870152604087015160c087015280945050505050949350505050565b5f60208284031215611295575f80fd5b815180151581146107c6575f80fd5b634e487b7160e01b5f52601160045260245ffd5b8181038181111561042f5761042f6112a4565b8082018082111561042f5761042f6112a4565b6001600160f81b0319831681525f6105716001830184611174565b5f6107c68284611174565b5f60208284031215611314575f80fd5b5051919050565b5f6001820161132c5761132c6112a4565b506001019056fea2646970667358221220458d6f25140af83d83590e6a5ab5fdc53b4b1faa32f899a792f41a97f8acaf5e64736f6c63430008160033",
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
