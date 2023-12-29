//SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity 0.8.23;

import "../IParameters.sol";
import "../IPanel.sol";
import "../IVoting.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../../interfaces/IContractRegistry.sol";
import "../../common/Globals.sol";
import "../../interfaces/IVotingWeightProxy.sol";

/**
 * @title Constitution Voting
 * @notice Used to vote for changes in constitution parameters
 */
contract ConstitutionVoting is IConstitutionVoting, Initializable {
    /**
     * @notice Enumerator that specifies type of voting
     */
    enum Classification {
        BASIC,
        FUNDAMENTAL,
        DETAILED
    }

    /**
     * @notice Structure that holds proposal parameters
     */
    struct ConstitutionProposal {
        BaseProposal base;
        mapping(uint256 => ParameterInfo) parameters;
        uint256 parametersSize;
        Classification classification;
        bytes32 newConstitutionHash;
        bytes32 currentConstitutionHash;
    }

    bytes32 public constitutionHash;

    uint256 public proposalCounter = 0;

    mapping(uint256 => ConstitutionProposal) public proposals;
    mapping(uint256 => mapping(address => bool)) public hasUserVoted;
    mapping(uint256 => mapping(address => bool)) public hasRootVetoed;

    IContractRegistry private registry;

    event UserVoted(uint256 indexed _proposalId, VotingOption _votingOption, uint256 _amount);
    event ProposalExecuted(uint256 indexed _proposalId, bytes32 _constitutionHash);

    modifier shouldExist(uint256 _id) {
        require(getStatus(_id) != ProposalStatus.NONE, "[QEC-001011]-The proposal does not exist.");
        _;
    }

    /**
     * @notice Restricts only root node can interact
     */
    modifier onlyRoot() {
        require(
            IRoots(_getRootNodesAddress()).isMember(msg.sender),
            "[QEC-001012]-Permission denied - only root nodes have access."
        );
        _;
    }

    constructor() {}

    /**
     * @notice Initializes the first parameters
     * @param _constitutionHash hash of constitution
     * @param _registry address of registry contract
     */
    function initialize(bytes32 _constitutionHash, address _registry) external initializer {
        constitutionHash = _constitutionHash;
        registry = IContractRegistry(_registry);
    }

    /**
     * @notice Creates address proposal
     * @param _remark Some message with details (may be link)
     * @param _proposalType type of proposal taken from Classification
     * @param _hashValue new constitution`s hash
     * @param _paramKey the key of the parameter to change
     * @param _paramValue new parameter value
     * @return id of new proposal
     */
    function createAddrProposal(
        string memory _remark,
        Classification _proposalType,
        bytes32 _hashValue,
        string memory _paramKey,
        address _paramValue
    ) external returns (uint256) {
        ParameterInfo memory _param;
        _param.paramKey = _paramKey;
        _param.paramType = ParameterType.ADDRESS;
        _param.addrValue = _paramValue;

        ParameterInfo[] memory _parametersArr = new ParameterInfo[](1);
        _parametersArr[0] = _param;

        return createProposal(_remark, _proposalType, _hashValue, _parametersArr);
    }

    /**
     * @notice Creates bool proposal
     * @param _remark Some message with details (may be link)
     * @param _proposalType type of proposal taken from Classification
     * @param _hashValue new constitution`s hash
     * @param _paramKey the key of the parameter to change
     * @param _paramValue new parameter value
     * @return id of new proposal
     */
    function createBoolProposal(
        string memory _remark,
        Classification _proposalType,
        bytes32 _hashValue,
        string memory _paramKey,
        bool _paramValue
    ) external returns (uint256) {
        ParameterInfo memory _param;
        _param.paramKey = _paramKey;
        _param.paramType = ParameterType.BOOL;
        _param.boolValue = _paramValue;

        ParameterInfo[] memory _parametersArr = new ParameterInfo[](1);
        _parametersArr[0] = _param;

        return createProposal(_remark, _proposalType, _hashValue, _parametersArr);
    }

    /**
     * @notice Creates bytes proposal
     * @param _remark Some message with details (may be link)
     * @param _proposalType type of proposal taken from Classification
     * @param _hashValue new constitution`s hash
     * @param _paramKey the key of the parameter to change
     * @param _paramValue new parameter value
     * @return id of new proposal
     */
    function createBytesProposal(
        string memory _remark,
        Classification _proposalType,
        bytes32 _hashValue,
        string memory _paramKey,
        bytes32 _paramValue
    ) external returns (uint256) {
        ParameterInfo memory _param;
        _param.paramKey = _paramKey;
        _param.paramType = ParameterType.BYTES32;
        _param.bytes32Value = _paramValue;

        ParameterInfo[] memory _parametersArr = new ParameterInfo[](1);
        _parametersArr[0] = _param;

        return createProposal(_remark, _proposalType, _hashValue, _parametersArr);
    }

    /**
     * @notice Creates string proposal
     * @param _remark Some message with details (may be link)
     * @param _proposalType type of proposal taken from Classification
     * @param _hashValue new constitution`s hash
     * @param _paramKey the key of the parameter to change
     * @param _paramValue new parameter value
     * @return id of new proposal
     */
    function createStrProposal(
        string memory _remark,
        Classification _proposalType,
        bytes32 _hashValue,
        string memory _paramKey,
        string memory _paramValue
    ) external returns (uint256) {
        ParameterInfo memory _param;
        _param.paramKey = _paramKey;
        _param.paramType = ParameterType.STRING;
        _param.strValue = _paramValue;

        ParameterInfo[] memory _parametersArr = new ParameterInfo[](1);
        _parametersArr[0] = _param;

        return createProposal(_remark, _proposalType, _hashValue, _parametersArr);
    }

    /**
     * @notice Creates uint proposal
     * @param _remark Some message with details (may be link)
     * @param _proposalType type of proposal taken from Classification
     * @param _hashValue new constitution`s hash
     * @param _paramKey the key of the parameter to change
     * @param _paramValue new parameter value
     * @return id of new proposal
     */
    function createUintProposal(
        string memory _remark,
        Classification _proposalType,
        bytes32 _hashValue,
        string memory _paramKey,
        uint32 _paramValue
    ) external returns (uint256) {
        ParameterInfo memory _param;
        _param.paramKey = _paramKey;
        _param.paramType = ParameterType.UINT;
        _param.uintValue = _paramValue;

        ParameterInfo[] memory _parametersArr = new ParameterInfo[](1);
        _parametersArr[0] = _param;

        return createProposal(_remark, _proposalType, _hashValue, _parametersArr);
    }

    /**
     * @notice Applies changes for specified proposal
     * @param _proposalId Proposal id
     */
    function execute(uint256 _proposalId) external override {
        require(
            getStatus(_proposalId) == ProposalStatus.PASSED,
            "[QEC-001004]-Proposal must be PASSED before excecuting."
        );

        proposals[_proposalId].base.executed = true;

        require(
            _applyProposedChanges(_proposalId),
            "[QEC-001005]-Failed to apply changes to the parameters, proposal execution failed."
        );

        constitutionHash = proposals[_proposalId].newConstitutionHash;

        emit ProposalExecuted(_proposalId, proposals[_proposalId].newConstitutionHash);
    }

    /**
     * @notice Give a vote for the specified proposal
     * @param _proposalId Proposal id
     */
    function voteFor(uint256 _proposalId) external override shouldExist(_proposalId) {
        _vote(_proposalId, VotingOption.FOR);
    }

    /**
     * @notice Give a vote against the specified proposal
     * @param _proposalId Proposal id
     */
    function voteAgainst(uint256 _proposalId) external override shouldExist(_proposalId) {
        _vote(_proposalId, VotingOption.AGAINST);
    }

    /**
     * @notice Root gives a veto for the specified proposal
     * @param _id Proposal id
     */
    function veto(uint256 _id) external override onlyRoot shouldExist(_id) {
        require(
            getStatus(_id) == ProposalStatus.ACCEPTED,
            "[QEC-001006]-Proposal must be ACCEPTED before casting a root node veto."
        );
        require(!hasRootVetoed[_id][msg.sender], "[QEC-001007]-The caller has already vetoed the proposal.");
        hasRootVetoed[_id][msg.sender] = true;

        ++proposals[_id].base.counters.vetosCount;
    }

    /**
     * @dev getter return constitution hash
     * @param _id id of the proposal, in which to take constitution
     * @return constitution hash
     */
    function getConstitutionHash(uint256 _id) external view returns (bytes32) {
        return proposals[_id].newConstitutionHash;
    }

    /**
     * @notice Gets an array of parameter information from a specific proposal
     * @param _proposalId Proposal id
     * @return array of parameter information structures
     */
    function getParametersArr(uint256 _proposalId)
        external
        view
        shouldExist(_proposalId)
        returns (ParameterInfo[] memory)
    {
        ConstitutionProposal storage prop = proposals[_proposalId];

        uint256 _arraySize = prop.parametersSize;
        ParameterInfo[] memory _parametersArr = new ParameterInfo[](_arraySize);

        for (uint256 i = 0; i < _arraySize; i++) {
            _parametersArr[i] = prop.parameters[i];
        }

        return _parametersArr;
    }

    /**
     * @notice Returns the base structure for the given id
     * @param _proposalId Proposal id
     * @return structure of type BaseProposal
     */
    function getProposal(uint256 _proposalId)
        external
        view
        override
        shouldExist(_proposalId)
        returns (BaseProposal memory)
    {
        return proposals[_proposalId].base;
    }

    /**
     * @notice Gets statistic for the proposal
     * @param _proposalId Proposal id
     * @return structure of type VotingStats
     */
    function getProposalStats(uint256 _proposalId)
        external
        view
        override
        shouldExist(_proposalId)
        returns (VotingStats memory)
    {
        ConstitutionProposal storage prop = proposals[_proposalId];
        VotingStats memory _stats;

        _stats.requiredQuorum = prop.base.params.requiredQuorum;
        _stats.requiredMajority = prop.base.params.requiredMajority;

        uint256 _totalWeight = prop.base.counters.weightFor + prop.base.counters.weightAgainst;
        uint256 _Qamount = getTotalQInExistence(block.number);

        _stats.currentQuorum = _calculatePercentage(_totalWeight, _Qamount);
        _stats.currentMajority = _calculatePercentage(prop.base.counters.weightFor, _totalWeight);
        _stats.currentVetoPercentage = _getVetosPercentageNumber(_proposalId);

        return _stats;
    }

    /**
     * @notice Gets count for vetos number
     * @param _proposalId Proposal id
     * @return number of vetos
     */
    function getVetosNumber(uint256 _proposalId) external view override shouldExist(_proposalId) returns (uint256) {
        return proposals[_proposalId].base.counters.vetosCount;
    }

    /**
     * @notice Gets percent for vetos number
     * @param _proposalId Proposal id
     * @return percent of vetos
     */
    function getVetosPercentage(uint256 _proposalId) external view override shouldExist(_proposalId) returns (uint256) {
        return _getVetosPercentageNumber(_proposalId);
    }

    /**
     * @notice Creates the proposal with several parameters of different types
     * @param _remark Some message with details (may be link)
     * @param _proposalType Type of proposal taken from Classification
     * @param _hashValue New constitution`s hash
     * @param _parametersArr An array of parameters that will be updated in case of a successful vote
     * @return id of new proposal
     */
    function createProposal(
        string memory _remark,
        Classification _proposalType,
        bytes32 _hashValue,
        ParameterInfo[] memory _parametersArr
    ) public returns (uint256) {
        uint256 _parametersArrLength = _parametersArr.length;

        ConstitutionProposal storage prop = proposals[proposalCounter];
        prop.parametersSize = _parametersArrLength;

        for (uint256 i = 0; i < _parametersArrLength; i++) {
            prop.parameters[i] = _parametersArr[i];
        }

        return _createProposal(_remark, _proposalType, _hashValue);
    }

    /**
     * @notice Returns information about the user's voting rights by the proposal id
     * @param _proposalId Proposal id
     * @return VotingWeightInfo struct
     */
    function getVotingWeightInfo(uint256 _proposalId)
        public
        view
        override
        shouldExist(_proposalId)
        returns (VotingWeightInfo memory)
    {
        IVotingWeightProxy _proxy = IVotingWeightProxy(_getVotingWeightProxyAddress());
        VotingWeightInfo memory _info;
        _info.base = _proxy.getBaseVotingWeightInfo(msg.sender, proposals[_proposalId].base.params.votingEndTime);
        _info.hasAlreadyVoted = hasUserVoted[_proposalId][msg.sender];
        _info.canVote = _info.base.ownWeight > 0 && !_info.hasAlreadyVoted;

        return _info;
    }

    /**
     * @notice Returns the status of the given proposal
     * @param _id Proposal id
     * @return status of proposal
     */
    function getStatus(uint256 _id) public view returns (ProposalStatus) {
        ConstitutionProposal storage prop = proposals[_id];
        if (prop.base.params.votingEndTime == 0) {
            return ProposalStatus.NONE;
        }

        if (prop.base.executed) {
            return ProposalStatus.EXECUTED;
        }

        if (block.timestamp < prop.base.params.votingEndTime) {
            return ProposalStatus.PENDING;
        }

        uint256 _totalWeight = prop.base.counters.weightFor + prop.base.counters.weightAgainst;
        uint256 _Qamount = getTotalQInExistence(block.number);
        uint256 _actualQuorum = _calculatePercentage(_totalWeight, _Qamount);
        if (_actualQuorum < prop.base.params.requiredQuorum) {
            return ProposalStatus.REJECTED;
        }

        uint256 _actualMajority = _calculatePercentage(prop.base.counters.weightFor, _totalWeight);
        if (_actualMajority <= prop.base.params.requiredMajority) {
            return ProposalStatus.REJECTED;
        }

        if (_getVetosPercentageNumber(_id) > getDecimal() / 2) {
            return ProposalStatus.REJECTED;
        }

        if (block.timestamp < prop.base.params.vetoEndTime) {
            return ProposalStatus.ACCEPTED;
        }

        if (constitutionHash != prop.currentConstitutionHash) {
            return ProposalStatus.OBSOLETE;
        }

        if (block.timestamp > prop.base.params.vetoEndTime + prop.base.params.proposalExecutionP) {
            return ProposalStatus.EXPIRED;
        }

        return ProposalStatus.PASSED;
    }

    /**
     * @dev Internal func to create all types of proposals
     */
    function _createProposal(
        string memory _remark,
        Classification _proposalType,
        bytes32 _hashValue
    ) private returns (uint256) {
        ConstitutionProposal storage prop = proposals[proposalCounter];
        prop.classification = _proposalType;
        prop.newConstitutionHash = _hashValue;
        prop.currentConstitutionHash = constitutionHash;
        prop.base.remark = _remark;

        uint256 _votePeriod = 0;
        uint256 _vetoPeriod = 0;
        uint256 _quorum = 0;
        uint256 _majority = 0;

        IParameters _constitution = IParameters(_getConstitutionParametersAddress());
        if (_proposalType == Classification.BASIC) {
            _votePeriod = _constitution.getUint("constitution.voting.basicQSectionVP");
            _vetoPeriod = _constitution.getUint("constitution.voting.basicQSectionRNVALP");
            _quorum = _constitution.getUint("constitution.voting.basicQSectionQRM");
            _majority = _constitution.getUint("constitution.voting.basicQSectionRMAJ");
        }
        if (_proposalType == Classification.FUNDAMENTAL) {
            _votePeriod = _constitution.getUint("constitution.voting.fundQSectionVP");
            _vetoPeriod = _constitution.getUint("constitution.voting.fundQSectionRNVALP");
            _quorum = _constitution.getUint("constitution.voting.fundQSectionQRM");
            _majority = _constitution.getUint("constitution.voting.fundQSectionRMAJ");
        }
        if (_proposalType == Classification.DETAILED) {
            _votePeriod = _constitution.getUint("constitution.voting.detailedQSectionVP");
            _vetoPeriod = _constitution.getUint("constitution.voting.detailedQSectionRNVALP");
            _quorum = _constitution.getUint("constitution.voting.detailedQSectionQRM");
            _majority = _constitution.getUint("constitution.voting.detailedQSectionRMAJ");
        }

        require(_votePeriod != 0, "[QEC-001000]-Invalid voting period parameter, proposal creation failed.");
        require(_vetoPeriod != 0, "[QEC-001001]-Invalid veto period parameter, proposal creation failed.");
        require(_quorum != 0, "[QEC-001002]-Invalid quorum parameter, proposal creation failed.");
        require(_majority != 0, "[QEC-001003]-Invalid required majority parameter, proposal creation failed.");

        prop.base.params.votingEndTime = block.timestamp + _votePeriod;
        prop.base.params.vetoEndTime = block.timestamp + _votePeriod + _vetoPeriod;
        prop.base.params.requiredQuorum = _quorum;
        prop.base.params.requiredMajority = _majority;
        prop.base.params.proposalExecutionP = _constitution.getUint("constitution.proposalExecutionP");

        emit ProposalCreated(proposalCounter, proposals[proposalCounter].base);

        return proposalCounter++;
    }

    /**
     * @dev Underlying function for execute
     */
    function _applyProposedChanges(uint256 _id) private returns (bool) {
        IMutableParameters constitParams = IMutableParameters(_getConstitutionParametersAddress());

        ConstitutionProposal storage prop = proposals[_id];
        ParameterInfo memory _param;
        ParameterType _type;
        string memory _key;

        for (uint256 i = 0; i < prop.parametersSize; i++) {
            _param = prop.parameters[i];
            _type = _param.paramType;
            _key = _param.paramKey;

            if (_type == ParameterType.ADDRESS) {
                constitParams.setAddr(_key, _param.addrValue);
            } else if (_type == ParameterType.UINT) {
                constitParams.setUint(_key, _param.uintValue);
            } else if (_type == ParameterType.STRING) {
                constitParams.setString(_key, _param.strValue);
            } else if (_type == ParameterType.BYTES32) {
                constitParams.setBytes32(_key, _param.bytes32Value);
            } else if (_type == ParameterType.BOOL) {
                constitParams.setBool(_key, _param.boolValue);
            }
        }

        return true;
    }

    /**
     * @dev Votes for the underlying proposal
     */
    function _vote(uint256 _id, VotingOption _votingOption) private returns (bool) {
        require(getStatus(_id) == ProposalStatus.PENDING, "[QEC-001008]-Voting is only possible on PENDING proposals.");
        require(!hasUserVoted[_id][msg.sender], "[QEC-001009]-The caller has already voted for the proposal.");

        hasUserVoted[_id][msg.sender] = true;

        IVotingWeightProxy proxy = IVotingWeightProxy(_getVotingWeightProxyAddress());
        uint256 _totalLockedQ = proxy.extendLocking(msg.sender, proposals[_id].base.params.votingEndTime);
        require(_totalLockedQ > 0, "[QEC-001013]-The total voting weight must be greater than zero.");

        if (_votingOption == VotingOption.FOR) {
            proposals[_id].base.counters.weightFor += _totalLockedQ;
        } else {
            proposals[_id].base.counters.weightAgainst += _totalLockedQ;
        }

        emit UserVoted(_id, _votingOption, _totalLockedQ);

        return true;
    }

    /**
     * @dev Internally counts the vetos percentage
     */
    function _getVetosPercentageNumber(uint256 _proposalId) private view returns (uint256) {
        return
            (proposals[_proposalId].base.counters.vetosCount * getDecimal()) / IRoots(_getRootNodesAddress()).getSize();
    }

    function _getRootNodesAddress() private view returns (address) {
        return registry.mustGetAddress(RKEY__ROOT_NODES);
    }

    function _getVotingWeightProxyAddress() private view returns (address) {
        return registry.mustGetAddress(RKEY__VOTING_WEIGHT_PROXY);
    }

    function _getConstitutionParametersAddress() private view returns (address) {
        return registry.mustGetAddress(RKEY__CONSTITUTION_PARAMETERS);
    }

    /**
     * @dev Calculates the percentage x/y
     */
    function _calculatePercentage(uint256 part, uint256 amount) private pure returns (uint256) {
        if (amount == 0) {
            return 0;
        }
        return (part * getDecimal()) / amount;
    }
}
