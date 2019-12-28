/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.5.3;

import "./RollupUtils.sol";

import "../challenge/IChallengeFactory.sol";

import "../arch/Protocol.sol";


contract Staking {

    // VM already initialized"
    string constant INIT_TWICE = "INIT_TWICE";
    // Challenge factory must be nonzero
    string constant INIT_NONZERO = "INIT_NONZERO";

    // Invalid staker
    string constant INV_STAKER = "INV_STAKER";

    // must supply stake value
    string constant STK_AMT = "STK_AMT";
    // Staker already exists
    string constant ALRDY_STAKED = "ALRDY_STAKED";

    // Challenge can only be resolved by spawned contract
    string constant RES_CHAL_SENDER = "RES_CHAL_SENDER";

    // Stakers must have a conflict over pending top
    string constant PND_CHAL_TYPE = "PND_CHAL_TYPE";

    // Stakers must have a conflict over messages
    string constant MSGS_CHAL_TYPE = "MSGS_CHAL_TYPE";

    // Stakers must have a conflict over execution
    string constant EXEC_CHAL_TYPE = "EXEC_CHAL_TYPE";

    // staker1 staked after deadline
    string constant STK1_DEADLINE = "STK1_DEADLINE";
    // staker2 staked after deadline
    string constant STK2_DEADLINE = "STK2_DEADLINE";
    // staker1 already in a challenge
    string constant STK1_IN_CHAL = "STK1_IN_CHAL";
    // staker2 already in a challenge
    string constant STK2_IN_CHAL = "STK1_IN_CHAL";
    // Child types must be ordered
    string constant TYPE_ORDER = "TYPE_ORDER";
    // Invalid child type
    string constant INVLD_CHLD_TYPE = "INVLD_CHLD_TYPE";
    // Invalid staker1 proof
    string constant STK1_PROOF = "STK1_PROOF";
    // Invalid staker2 proof
    string constant STK2_PROOF = "STK2_PROOF";

    uint internal constant VALID_CHILD_TYPE = 0;
    uint internal constant INVALID_PENDING_TOP_CHILD_TYPE = 1;
    uint internal constant INVALID_MESSAGES_CHILD_TYPE = 2;
    uint internal constant INVALID_EXECUTION_CHILD_TYPE = 3;
    uint internal constant MAX_CHILD_TYPE = 3;

    IChallengeFactory public challengeFactory;

    struct Staker {
        bytes32 location;
        uint128 creationTime;
        bool inChallenge;
    }

    uint128 private stakeRequirement;
    mapping(address => Staker) private stakers;
    uint private stakerCount;

    event RollupStakeCreated(
        address staker,
        bytes32 nodeHash,
        uint    blockNumber
    );

    event RollupChallengeStarted(
        address asserter,
        address challenger,
        uint    challengeType,
        address challengeContract
    );

    event RollupChallengeCompleted(
        address challengeContract,
        address winner,
        address loser
    );

    struct ChallengeData {
        address payable[2] stakerAddresses;
        bytes32 node;
        uint disputableDeadline;
        uint[2] stakerPositions;
        bytes32[2] vmProtoHashes;
        bytes32[] proof1;
        bytes32[] proof2;
    }

    struct StartExecutionChallengeData {
        bytes32 beforeHash;
        bytes32 beforeInbox;
        uint64[2] timeBounds;
        bytes32 pendingAssertion;
        bytes32 beforePendingTop;
        bytes32 importedMessageSlice;
        uint32 importedMessageCount;
        bytes32 assertionHash;
    }

    // fields
    //  node
    //  beforeHash
    //  beforeInbox
    //  pendingAssertionHash
    //  beforePendingTop
    //  importedMessageSlice
    //  assertionHash

    function startExecutionChallenge(
        bytes32[7] calldata _fields,
        address payable[2] calldata stakerAddresses,
        uint disputableDeadline,
        uint[2] calldata stakerPositions,
        bytes32[2] calldata vmProtoHashes,
        bytes32[] calldata proof1,
        bytes32[] calldata proof2,
        uint64[2] calldata _timeBounds,
        uint32 _importedMessageCount
    )
        external
    {
        return _startExecutionChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            StartExecutionChallengeData(
                _fields[1],
                _fields[2],
                _timeBounds,
                _fields[3],
                _fields[4],
                _fields[5],
                _importedMessageCount,
                _fields[6]
            )
        );
    }

    // fields
    //  node
    //  preconditionHash
    //  afterPendingTop
    //  currentPending
    //  importedAssertion
    //  assertionHash

    function startPendingTopChallenge(
        bytes32[6] calldata _fields,
        address payable[2] calldata stakerAddresses,
        uint disputableDeadline,
        uint[2] calldata stakerPositions,
        bytes32[2] calldata vmProtoHashes,
        bytes32[] calldata proof1,
        bytes32[] calldata proof2
    )
        external
    {
        return _startPendingTopChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            _fields[1],
            _fields[2],
            _fields[3],
            _fields[4],
            _fields[5]
        );
    }

    // fields
    //  node
    //  preconditionHash
    //  afterPendingTop
    //  currentPending
    //  beforePendingTop
    //  importedMessageSlice
    //  assertionHash

    function startMessagesChallenge(
        bytes32[7] memory _fields,
        address payable[2] memory stakerAddresses,
        uint disputableDeadline,
        uint[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2,
        uint32 _importedMessageCount
    )
        public
    {
        return _startMessagesChallenge(
            ChallengeData(
                stakerAddresses,
                _fields[0],
                disputableDeadline,
                stakerPositions,
                vmProtoHashes,
                proof1,
                proof2
            ),
            _fields[1],
            _fields[2],
            _fields[3],
            _fields[4],
            _fields[5],
            _importedMessageCount,
            _fields[6]
        );
    }

    function resolveChallenge(address payable winner, address loser) external {
        address sender = msg.sender;
        bytes32 codehash;
        assembly { codehash := extcodehash(sender) }
        address challengeContract = challengeFactory.generateCloneAddress(address(winner), loser, codehash);
        require(challengeContract == msg.sender, RES_CHAL_SENDER);
        Staker storage winningStaker = getValidStaker(address(winner));
        winner.transfer(stakeRequirement / 2);
        winningStaker.inChallenge = false;
        deleteStaker(loser);

        emit RollupChallengeCompleted(msg.sender, address(winner), loser);
    }

    function init(
        uint128 _stakeRequirement,
        address _challengeFactoryAddress
    )
        internal
    {
        require(address(challengeFactory) == address(0), INIT_TWICE);
        require(_challengeFactoryAddress != address(0), INIT_NONZERO);

        challengeFactory = IChallengeFactory(_challengeFactoryAddress);

        // VM parameters
        stakeRequirement = _stakeRequirement;
    }

    function getValidStaker(address _stakerAddress) internal view returns (Staker storage) {
        Staker storage staker = stakers[_stakerAddress];
        require(staker.location != 0x00, INV_STAKER);
        return staker;
    }

    function createStake(
        bytes32 location
    )
        internal
    {
        require(msg.value == stakeRequirement, STK_AMT);
        require(stakers[msg.sender].location != 0x00, ALRDY_STAKED);
        stakers[msg.sender] = Staker(
            location,
            uint128(block.number),
            false
        );
        stakerCount++;

        emit RollupStakeCreated(msg.sender, location, block.number);
    }

    function deleteStakerWithPayout(address payable _stakerAddress) internal {
        deleteStaker(_stakerAddress);
        _stakerAddress.transfer(stakeRequirement);
    }

    function deleteStaker(address _stakerAddress) private {
        delete stakers[_stakerAddress];
        stakerCount--;
    }

    function getStakerCount() internal view returns(uint) {
        return stakerCount;
    }

    function _startPendingTopChallenge(
        ChallengeData memory _challenge,
        bytes32 _preconditionHash,
        bytes32 _afterPendingTop,
        bytes32 _currentPending,
        bytes32 _importedAssertion,
        bytes32 _assertionHash
    )
        private
    {
        require(_challenge.stakerPositions[1] == INVALID_PENDING_TOP_CHILD_TYPE, PND_CHAL_TYPE);

        verifyConflict(
            _challenge,
            RollupUtils.disputableNodeHash(
                _preconditionHash,
                RollupUtils.pendingAssertionHash(
                    _afterPendingTop,
                    _currentPending
                ),
                _importedAssertion,
                _assertionHash
            )
        );

        address newChallengeAddr = challengeFactory.createPendingTopChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            _currentPending,
            _afterPendingTop
        );

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function _startMessagesChallenge(
        ChallengeData memory _challenge,
        bytes32 _preconditionHash,
        bytes32 _afterPendingTop,
        bytes32 _currentPending,
        bytes32 _beforePendingTop,
        bytes32 _importedMessageSlice,
        uint32 _importedMessageCount,
        bytes32 _assertionHash
    )
        private
    {
        require(_challenge.stakerPositions[1] == INVALID_MESSAGES_CHILD_TYPE, MSGS_CHAL_TYPE);

        verifyConflict(
            _challenge,
            RollupUtils.disputableNodeHash(
                _preconditionHash,
                RollupUtils.pendingAssertionHash(
                    _afterPendingTop,
                    _currentPending
                ),
                RollupUtils.importedAssertionHash(
                    _beforePendingTop,
                    _importedMessageCount,
                    _importedMessageSlice
                ),
                _assertionHash
            )
        );

        address newChallengeAddr = challengeFactory.createMessagesChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            _beforePendingTop,
            _afterPendingTop,
            _importedMessageSlice,
            _importedMessageCount
        );


        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function _startExecutionChallenge(
        ChallengeData memory _challenge,
        StartExecutionChallengeData memory data
    )
        private
    {
        require(_challenge.stakerPositions[1] == INVALID_EXECUTION_CHILD_TYPE, EXEC_CHAL_TYPE);

        verifyConflict(
            _challenge,
            RollupUtils.disputableNodeHash(
                Protocol.generatePreconditionHash(
                    data.beforeHash,
                    data.timeBounds,
                    data.beforeInbox
                ),
                data.pendingAssertion,
                RollupUtils.importedAssertionHash(
                    data.beforePendingTop,
                    data.importedMessageCount,
                    data.importedMessageSlice
                ),
                data.assertionHash
            )
        );

        address newChallengeAddr = challengeFactory.createExecutionChallenge(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            0, // Challenge period
            data.beforeHash,
            Protocol.addMessagesToInbox(data.beforeInbox, data.importedMessageSlice),
            data.timeBounds,
            data.assertionHash
        );

        emit RollupChallengeStarted(
            _challenge.stakerAddresses[0],
            _challenge.stakerAddresses[1],
            _challenge.stakerPositions[1],
            newChallengeAddr
        );
    }

    function verifyConflict(
        ChallengeData memory _challenge,
        bytes32 disputableNodeHash
    )
        private
    {
        Staker storage staker1 = getValidStaker(_challenge.stakerAddresses[0]);
        Staker storage staker2 = getValidStaker(_challenge.stakerAddresses[1]);

        require(staker1.creationTime < _challenge.disputableDeadline, STK1_DEADLINE);
        require(staker2.creationTime < _challenge.disputableDeadline, STK2_DEADLINE);
        require(!staker1.inChallenge, STK1_IN_CHAL);
        require(!staker2.inChallenge, STK2_IN_CHAL);
        require(_challenge.stakerPositions[0] < _challenge.stakerPositions[1], TYPE_ORDER);
        require(_challenge.stakerPositions[1] <= MAX_CHILD_TYPE, INVLD_CHLD_TYPE);
        require(
            RollupUtils.isPath(
                RollupUtils.childNodeHash(
                    _challenge.node,
                    _challenge.disputableDeadline,
                    disputableNodeHash,
                    _challenge.stakerPositions[0],
                    _challenge.vmProtoHashes[0]
                ),
                staker1.location,
                _challenge.proof1
            ),
            STK1_PROOF
        );
        require(
            RollupUtils.isPath(
                RollupUtils.childNodeHash(
                    _challenge.node,
                    _challenge.disputableDeadline,
                    disputableNodeHash,
                    _challenge.stakerPositions[1],
                    _challenge.vmProtoHashes[1]
                ),
                staker2.location,
                _challenge.proof2
            ),
            STK2_PROOF
        );

        staker1.inChallenge = true;
        staker2.inChallenge = true;
    }
}
