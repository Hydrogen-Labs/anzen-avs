// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "../static/Structs.sol";

interface ISafetyFactorOracle {
    // Events
    event SFUpdated(uint32 protocolId, int256 newSF);

    event ProtocolAdded(uint32 protocolId, address reservesManager);

    event ProtocolRemoved(uint32 protocolId);

    event DisputeStatusSet(bool status);

    /**
     *
     *                         Public Functions
     *
     */
    function setDisputeStatus(bool status) external;
    // TODO: remove input parameter when $EIGEN dispute is implemented

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

    function getDisputeStatus() external view returns (bool);
}
