// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "../static/Structs.sol";

interface IAVSReservesManager {
    // Event declaration
    event TokenFlowUpdated(uint256 newTokenFlow);

    event TokensTransferredToPaymentMaster(uint256 totalTokenTransfered);

    function initialize(
        SafetyFactorConfig memory _safetyFactorConfig,
        address _safetyFactorOracle,
        address _avsGov,
        address _anzenGov,
        uint32 _protocolId,
        address[] memory _rewardTokens,
        uint256[] memory _initial_tokenFlowsPerSecond,
        uint256 _performanceFeeBPS
    ) external;

    function updateFlow() external;

    function overrideTokensPerSecond(address[] memory _tokenAddresses, uint256[] memory _newTokensPerSecond) external;

    function updateSafetyFactorParams(SafetyFactorConfig memory newSafetyFactorConfig) external;

    // function transferToPaymentManager() external;

    // Use this function to get the amount of tokens that can be claimed by the AVS

    /**
     * Getter Functions
     */
    function getSafetyFactorConfig() external view returns (SafetyFactorConfig memory);
}
