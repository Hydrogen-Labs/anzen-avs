// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/src/libraries/BN254.sol";

interface IAnzenTaskManager {
    // EVENTS

    event NewOraclePullTaskCreated(uint32 indexed taskIndex, OraclePullTask oraclePullTask);

    event OraclePullTaskResponded(
        OraclePullTaskResponse oraclePullTaskResponse, TaskResponseMetadata taskResponseMetadata
    );

    event TaskCompleted(uint32 indexed taskIndex);

    event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger);

    event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger);

    struct OraclePullTask {
        uint32 oracleIndex;
        int256 proposedSafetyFactor;
        uint32 taskCreatedBlock;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
    }

    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct OraclePullTaskResponse {
        uint32 referenceTaskIndex;
        int256 safetyFactor;
    }

    // Extra information related to taskResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than TaskResponse
    // This metadata is needed by the challenger, so we emit it in the TaskResponded event
    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS

    // NOTE: this function creates new oracle pull task.
    function createNewOraclePullTask(
        uint32 oracleIndex,
        int256 proposedSafetyFactor,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external;

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    // // NOTE: this function raises challenge to existing tasks.
    // function raiseAndResolveChallenge(
    //     Task calldata task,
    //     TaskResponse calldata taskResponse,
    //     TaskResponseMetadata calldata taskResponseMetadata,
    //     BN254.G1Point[] memory pubkeysOfNonSigningOperators
    // ) external;

    /// @notice Returns the TASK_RESPONSE_WINDOW_BLOCK
    function getTaskResponseWindowBlock() external view returns (uint32);

    function updateAggregator(address newAggregator) external;

    function updateTaskGenerator(address newTaskGenerator) external;
}
