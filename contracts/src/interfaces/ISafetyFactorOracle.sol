// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "../static/Structs.sol";

interface ISafetyFactorOracle {
    // Events
    event SFUpdated(int256 newSF, uint32 protocolId);

    /**
     *
     *                         Anzen Gov Functions
     *
     */
    function addProtocol(uint32 protocolIdToAdd, address reservesManager) external;

    function removeProtocol(uint32 protocolIdToRemove) external;

    /**
     *
     *                         Anzen Task Manager Functions
     *
     */
    function setSafetyFactor(uint32 protocolId, int256 safetyFactor) external;

    /**
     *
     *                         Getter Functions
     *
     */
    function getSafetyFactor(uint32 protocolId) external view returns (SafetyFactorSnapshot memory safetyFactor);
}
