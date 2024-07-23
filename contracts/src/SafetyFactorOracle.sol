// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "./static/Structs.sol";
import {ISafetyFactorOracle} from "./interfaces/ISafetyFactorOracle.sol";
import {SafetyFactorOracleStorage} from "./SafetyFactorOracleStorage.sol";

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

contract SafetyFactorOracle is SafetyFactorOracleStorage, Initializable {
    // AnzenTaskManager uses BLS signatures to sign the safety factor updates
    modifier onlyAnzenTaskManager() {
        require(msg.sender == anzenTaskManager, "onlyAnzenTaskManager: not from anzen task manager");
        _;
    }

    modifier onlyAnzenTaskManagerOrFallBackIfDisputePending() {
        require(
            (msg.sender == anzenTaskManager && !pendingDispute)
                || (msg.sender == fallBackSafetyFactorPoster && pendingDispute),
            "onlyAnzenTaskManagerOrFallBackIfDispute: not from anzen task manager or fall back safety factor poster"
        );
        _;
    }

    modifier onlyAnzenGov() {
        require(msg.sender == anzenGov, "onlyAnzenGov: not from anzen gov");
        _;
    }

    function initialize(address _anzenTaskManager, address _anzenGov, address _fallBackSafetyFactorPoster)
        public
        initializer
    {
        anzenTaskManager = _anzenTaskManager;
        anzenGov = _anzenGov;
        fallBackSafetyFactorPoster = _fallBackSafetyFactorPoster;
    }

    function addProtocol(uint32 _protocolIdToAdd, address _reservesManager) external onlyAnzenGov {
        require(!activeProtocols[_protocolIdToAdd], "Protocol is already active");
        activeProtocols[_protocolIdToAdd] = true;
        protocolToReservesManager[_protocolIdToAdd] = _reservesManager;

        emit ProtocolAdded(_protocolIdToAdd, _reservesManager);
    }

    function removeProtocol(uint32 _protocolIdToRemove) external onlyAnzenGov {
        require(activeProtocols[_protocolIdToRemove], "Protocol is not active");
        activeProtocols[_protocolIdToRemove] = false;
        protocolToReservesManager[_protocolIdToRemove] = address(0);

        emit ProtocolRemoved(_protocolIdToRemove);
    }

    function setSafetyFactor(uint32 _protocolId, int256 _safetyFactor)
        external
        onlyAnzenTaskManagerOrFallBackIfDisputePending
    {
        require(activeProtocols[_protocolId], "Protocol is not active");
        safetyFactorSnapshots[_protocolId] = SafetyFactorSnapshot(_safetyFactor, block.timestamp);

        emit SFUpdated(_protocolId, _safetyFactor);

        // TODO: Add logic to pass the safety factor to the  AVSReservesManager
    }

    function getSafetyFactor(uint32 _protocolId) external view returns (SafetyFactorSnapshot memory safetyFactor) {
        require(activeProtocols[_protocolId], "Protocol is not active");
        return safetyFactorSnapshots[_protocolId];
    }

    function setDisputeStatus(bool _status) external {
        // TODO: set pendingDispute to true if an $EIGEN dispute is pending
        // TODO: for testing purposes, we will can modify the status of the dispute
        pendingDispute = _status;
    }

    function getDisputeStatus() external view returns (bool) {
        return pendingDispute;
    }
}
