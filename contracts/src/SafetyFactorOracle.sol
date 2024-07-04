// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "./static/Structs.sol";
import {ISafetyFactorOracle} from "./interfaces/ISafetyFactorOracle.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Pausable} from "@eigenlayer/contracts/permissions/Pausable.sol";

contract SafetyFactorOracle is ISafetyFactorOracle, Ownable, Pausable {
    mapping(uint32 => SafetyFactorSnapshot) public safetyFactorSnapshots; // Safety Factor snapshots for each protocol
    mapping(uint32 => address) public protocolToReservesManager; // Mapping of protocol to AVSReservesManager
    mapping(uint32 => bool) public activeProtocols; // Mapping of protocol to active status

    address anzenTaskManager;
    address anzenGov;

    constructor(address _anzenTaskManager, address _anzenGov) Ownable() {
        anzenTaskManager = _anzenTaskManager;
        anzenGov = _anzenGov;
    }

    // AnzenTaskManager uses BLS signatures to sign the safety factor updates
    modifier onlyAnzenTaskManager() {
        require(msg.sender == anzenTaskManager, "onlyAnzenTaskManager: not from anzen task manager");
        _;
    }

    modifier onlyAnzenGov() {
        require(msg.sender == anzenGov, "onlyAnzenGov: not from anzen gov");
        _;
    }

    function addProtocol(uint32 _protocolIdToAdd, address _reservesManager) external onlyAnzenGov {
        activeProtocols[_protocolIdToAdd] = true;
        protocolToReservesManager[_protocolIdToAdd] = _reservesManager;
    }

    function removeProtocol(uint32 _protocolIdToRemove) external onlyAnzenGov {
        activeProtocols[_protocolIdToRemove] = false;
        protocolToReservesManager[_protocolIdToRemove] = address(0);
    }

    function setSafetyFactor(uint32 _protocolId, int256 _safetyFactor) external onlyAnzenTaskManager {
        require(activeProtocols[_protocolId], "Protocol is not active");
        safetyFactorSnapshots[_protocolId] = SafetyFactorSnapshot(_safetyFactor, block.timestamp);

        emit SFUpdated(_safetyFactor, _protocolId);

        // TODO: Add logic to pass the safety factor to the  AVSReservesManager
    }

    function getSafetyFactor(uint32 _protocolId) external view returns (SafetyFactorSnapshot memory safetyFactor) {
        return safetyFactorSnapshots[_protocolId];
    }
}
