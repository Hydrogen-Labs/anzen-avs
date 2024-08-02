// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

import "./interfaces/ISafetyFactorOracle.sol";
import "./AnzenTaskManager.sol";
import "./AVSReservesManager.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract AVSReservesManagerFactoryStorage {
    uint32 public lastAVSReservesManagerId;

    mapping(uint32 => address) public avsReservesManagers;
    mapping(address => bool) public hasAVSReservesManager;

    event AVSReservesManagerCreated(
        address indexed avsReservesManager, uint32 avsReservesManagerId, address avsServiceManager
    );
}
