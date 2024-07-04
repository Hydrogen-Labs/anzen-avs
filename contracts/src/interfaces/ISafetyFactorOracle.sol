// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "../static/Structs.sol";

interface ISafetyFactorOracle {
    // Events
    event SFUpdated(int256 newSF, uint32 protocolId);

    // Setter functions
    function setSafetyFactor(uint32 protocolId, int256 safetyFactor) external;

    // Getter functions
    function getSafetyFactor(uint32 protocolId) external view returns (SafetyFactorSnapshot memory safetyFactor);
}
