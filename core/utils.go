package core

import (
	"math/big"

	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
)

// this hardcodes abi.encode() for anzentaskmanager.IAnzenTaskManagerTaskResponse
// unclear why abigen doesn't provide this out of the box...

func AbiEncodePullOracleTaskResponse(h *anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse) ([]byte, error) {

	// The order here has to match the field ordering of anzentaskmanager.IAnzenTaskManagerTaskResponse
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "safetyFactor",
			Type: "int256",
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	bytes, err := arguments.Pack(h)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetPullOracleTaskResponseDigest(h *anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse) ([32]byte, error) {
	encodeTaskResponseByte, err := AbiEncodePullOracleTaskResponse(h)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

// BINDING UTILS - conversion from contract structs to golang structs

// BN254.sol is a library, so bindings for G1 Points and G2 Points are only generated
// in every contract that imports that library. Thus the output here will need to be
// type casted if G1Point is needed to interface with another contract (eg: BLSPublicKeyCompendium.sol)
func ConvertToBN254G1Point(input *bls.G1Point) anzentaskmanager.BN254G1Point {
	output := anzentaskmanager.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) anzentaskmanager.BN254G2Point {
	output := anzentaskmanager.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
