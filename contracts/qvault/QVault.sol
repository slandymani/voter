// SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity 0.8.23;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./ITokenLock.sol";
import "./IQHolderRewardPool.sol";
import "../defi/token/IERC677.sol";
import "./ValidationRewardPools.sol";
import "../interfaces/IContractRegistry.sol";
import "../governance/validators/Validators.sol";
import "../governance/IParameters.sol";
import "../interfaces/IVotingWeightProxy.sol";
import "../common/Globals.sol";
import "../common/AddressStorageFactory.sol";
import "../common/CompoundRateKeeper.sol";
import "../common/CompoundRateKeeperFactory.sol";
import "./ATimeLockBase.sol";

/**
 * @title Q Vault
 * @notice Used to store funds
 */
contract QVault is ITokenLock, Initializable, IERC677, ATimeLockBase {
    struct ValidatorInfo {
        address validator;
        uint256 actualStake;
        uint256 normalizedStake;
        uint256 compoundRate;
        uint256 latestUpdateOfCompoundRate;
        uint256 idealStake;
        uint256 claimableReward;
    }

    struct DelegationAmount {
        uint256 actualStake;
        uint256 normalizedStake;
    }

    struct DelegationInfo {
        mapping(address => DelegationAmount) delegatedAmounts;
        AddressStorage delegatedTo;
        uint256 totalDelegatedStake;
    }

    struct BalanceDetails {
        uint256 currentBalance;
        uint256 normalizedBalance;
        uint256 compoundRate;
        uint256 lastUpdateOfCompoundRate;
        uint256 interestRate;
    }

    struct UserAccount {
        uint256 normalizedBalance;
        DelegationInfo delegationInfo;
        mapping(address => uint256) allowances;
    }

    IContractRegistry private registry;

    uint256 public aggregatedNormalizedBalance;

    CompoundRateKeeper public compoundRateKeeper;

    mapping(address => UserAccount) internal accounts;

    event UserDeposited(address indexed _user, uint256 _newDepositAmount, uint256 _newBalance);
    event DelegatorRewardClaimed(address indexed _delegator, uint256 _totalReward, uint256 _newBalance);
    event UserWithdrawn(address indexed _user, uint256 _withdrawnAmount, uint256 _newBalance);

    modifier onlyRewardPools() {
        require(
            msg.sender == address(_getValidatorsRewardPools()) || msg.sender == _getQHolderRewardPoolAddress(),
            "[QEC-017019]-Permission denied - only the reward pools contract has access."
        );
        _;
    }

    constructor() {}

    function initialize(address _registry) external virtual initializer {
        registry = IContractRegistry(_registry);
        compoundRateKeeper = CompoundRateKeeperFactory(registry.mustGetAddress(RKEY__CR_KEEPER_FACTORY)).create();
    }

    /**
     * @notice Updates compound rate
     * @return new compound rate
     */
    function updateCompoundRate() external returns (uint256) {
        uint256 _oldDenormalizedAggregatedBalance = compoundRateKeeper.denormalizeAmount(aggregatedNormalizedBalance);
        uint256 _newRate = compoundRateKeeper.update(
            IParameters(_getEPQFIParametersAddress()).getUint("governed.EPQFI.Q_rewardPoolInterest")
        );
        uint256 _newDenormalizedAggregatedBalance = compoundRateKeeper.denormalizeAmount(aggregatedNormalizedBalance);

        uint256 _accruedReward = _newDenormalizedAggregatedBalance - _oldDenormalizedAggregatedBalance;

        IQHolderRewardPool(_getQHolderRewardPoolAddress()).requestRewardTransfer(_accruedReward);

        return _newRate;
    }

    /**
     * @notice deposits funds
     * @return true if successful
     */
    function deposit() external payable returns (bool) {
        return depositTo(msg.sender);
    }

    /**
     * @notice withdraw is used to withdraw tokens from the user's balance from the QVault to his address
     *
     * @dev When the user's balance is zero, we get 017001 error.
     * When, due to the delegated stake, token withdraw is not possible, we receive error 017014.
     * If there are not enough unlocked tokens, then there will be an attempt to unlock the required amount of tokens.
     *
     * @param _amount The amount of tokens that the user wants to withdraw
     * the maximum that is on the balance will be displayed, otherwise it will return a refusal
     *
     * @return true if the withdrawal of tokens was successful
     */
    function withdraw(uint256 _amount) external returns (bool) {
        return withdrawTo(msg.sender, _amount);
    }

    /**
     * @notice lock is used to increase the user's current locked tokens
     *
     * @dev For token lock, the corresponding proxy contract method is used.
     * When trying to lock an amount more than is currently available, we get an error 017016
     *
     * @param _amount The amount of locked tokens
     */
    function lock(uint256 _amount) external override {
        IVotingWeightProxy _votingWeightProxy = _getVotingWeightProxy();

        VotingLockInfo memory _votingLockInfo = _getLockInfo();
        uint256 _totalLockedAmount = _votingLockInfo.lockedAmount + _votingLockInfo.pendingUnlockAmount;

        require(
            balanceOf(msg.sender) - _totalLockedAmount >= _amount,
            "[QEC-017016]-The lock amount must not exceed the available balance."
        );

        _votingWeightProxy.lock(msg.sender, _amount);

        _checkTokenLockInvariance();
    }

    /**
     * @notice announceUnlock is used to announce unlocking of locked tokens
     *
     * @dev To announce the unlocking of tokens, the corresponding proxy contract method is used
     *
     * @param _amount The amount of tokens for which unlocking will be announced
     */
    function announceUnlock(uint256 _amount) external override {
        _getVotingWeightProxy().announceUnlock(msg.sender, _amount, 0);

        _checkTokenLockInvariance();
    }

    /**
     * @notice unlock is used for the final unlocking of tokens announced for unlocking
     *
     * @dev To unlock tokens, the corresponding method of the proxy contract is used
     *
     * @param _amount Amount of tokens to be unlocked
     */
    function unlock(uint256 _amount) external override {
        IVotingWeightProxy _votingWeightProxy = _getVotingWeightProxy();

        _votingWeightProxy.unlock(msg.sender, _amount);

        _checkTokenLockInvariance();
    }

    /**
     * @notice Delegates amounts to transferred users
     * @param _delegatedTo array of addresses for delegation
     * @param _stakes array of amounts to be delegated
     */
    function delegateStake(address[] memory _delegatedTo, uint256[] memory _stakes) external {
        require(balanceOf(msg.sender) != 0, "[QEC-017014]-Insufficient funds to cover all delegations.");

        require(
            _delegatedTo.length == _stakes.length,
            "[QEC-017012]-The number of candidates and stakes should be the same, stake delegation failed."
        );

        DelegationInfo storage delegationInfo = accounts[msg.sender].delegationInfo;
        AddressStorage _newDelegations = _getOrCreateDelegatedToStorage(msg.sender);

        Validators _validators = Validators(registry.mustGetAddress(RKEY__VALIDATORS));
        ValidationRewardPools _pools = _getValidatorsRewardPools();

        for (uint256 i = 0; i < _delegatedTo.length; i++) {
            address _delegationAddr = _delegatedTo[i];
            if (delegationInfo.delegatedAmounts[_delegationAddr].actualStake < _stakes[i]) {
                _pools.updateValidatorsCompoundRate(_delegationAddr);
            }
        }

        _claimStakeDelegatorReward(msg.sender);

        for (uint256 i = 0; i < _delegatedTo.length; i++) {
            address _delegationAddr = _delegatedTo[i];
            uint256 _normalizedStake;
            uint256 _actualStake;

            // remove delegation
            if (_stakes[i] == 0) {
                uint256 _oldStake = delegationInfo.delegatedAmounts[_delegationAddr].actualStake;
                delegationInfo.totalDelegatedStake = delegationInfo.totalDelegatedStake - _oldStake;

                _pools.subAggregatedNormalizedStake(
                    _delegationAddr,
                    delegationInfo.delegatedAmounts[_delegationAddr].normalizedStake
                );

                _validators.refreshDelegatedStake(_delegationAddr);

                _newDelegations.mustRemove(_delegationAddr);
                delete delegationInfo.delegatedAmounts[_delegationAddr];

                continue;
            }

            // update stake
            if (delegationInfo.delegatedAmounts[_delegationAddr].actualStake != 0) {
                _normalizedStake = _pools.getNormalizedAmount(_delegationAddr, _stakes[i]);
                _pools.subAggregatedNormalizedStake(
                    _delegationAddr,
                    delegationInfo.delegatedAmounts[_delegationAddr].normalizedStake
                );
                _pools.addAggregatedNormalizedStake(_delegationAddr, _normalizedStake);
                delegationInfo.delegatedAmounts[_delegationAddr].normalizedStake = _normalizedStake;

                _actualStake = _pools.getDenormalizedAmount(_delegationAddr, _normalizedStake);
                uint256 _oldStake = delegationInfo.delegatedAmounts[_delegationAddr].actualStake;
                delegationInfo.delegatedAmounts[_delegationAddr].actualStake = _actualStake;

                delegationInfo.totalDelegatedStake = delegationInfo.totalDelegatedStake + _actualStake - _oldStake;

                _validators.refreshDelegatedStake(_delegationAddr);

                continue;
            }

            // create delegation
            _newDelegations.add(_delegationAddr);

            _normalizedStake = _pools.getNormalizedAmount(_delegationAddr, _stakes[i]);
            delegationInfo.delegatedAmounts[_delegationAddr].normalizedStake = _normalizedStake;
            _pools.addAggregatedNormalizedStake(_delegationAddr, _normalizedStake);

            _actualStake = _pools.getDenormalizedAmount(_delegationAddr, _normalizedStake);
            delegationInfo.delegatedAmounts[_delegationAddr].actualStake = _actualStake;
            delegationInfo.totalDelegatedStake += _actualStake;

            _validators.refreshDelegatedStake(_delegationAddr);
        }

        uint256 maxTargets = IParameters(_getEPQFIParametersAddress()).getUint(
            "governed.EPQFI.maximumDelegationTargets"
        );
        require(
            delegationInfo.delegatedTo.size() <= maxTargets,
            "[QEC-017013]-The limit of candidates for the delegation has been exceeded, stake delegation failed."
        );

        require(
            balanceOf(msg.sender) >= delegationInfo.totalDelegatedStake,
            "[QEC-017014]-Insufficient funds to cover all delegations."
        );
    }

    /**
     * @notice getLockInfo is used to get information about the user's locked tokens
     *
     * @dev We receive error 028001 if this contract is not the source of the token lock
     *
     * @return information about the current lock of tokens as an object of the VotingLockInfo structure
     */
    function getLockInfo() external view override returns (VotingLockInfo memory) {
        return _getLockInfo();
    }

    /**
     * @notice Retrieves list of delegations
     * @param _delegator the user who delegated part of the stakes to other users
     * @return ValidatorInfo type array
     */
    function getDelegationsList(address _delegator) external view returns (ValidatorInfo[] memory) {
        DelegationInfo storage delegationInfo = accounts[_delegator].delegationInfo;

        ValidationRewardPools _pools = _getValidatorsRewardPools();
        address[] memory _delegatedTo;

        if (delegationInfo.delegatedTo != AddressStorage(address(0)))
            _delegatedTo = delegationInfo.delegatedTo.getAddresses();
        ValidatorInfo[] memory result = new ValidatorInfo[](_delegatedTo.length);

        for (uint256 i = 0; i < _delegatedTo.length; i++) {
            result[i].validator = _delegatedTo[i];
            result[i].actualStake = delegationInfo.delegatedAmounts[_delegatedTo[i]].actualStake;
            result[i].normalizedStake = delegationInfo.delegatedAmounts[_delegatedTo[i]].normalizedStake;

            result[i].compoundRate = _pools.getCompoundRate(_delegatedTo[i]);
            result[i].latestUpdateOfCompoundRate = _pools.getLastUpdateOfCompoundRate(_delegatedTo[i]);

            result[i].idealStake = _pools.getDenormalizedAmount(result[i].validator, result[i].normalizedStake);
            result[i].claimableReward = result[i].idealStake - result[i].actualStake;
        }

        return result;
    }

    /**
     * @notice Retrieves user normalized balance
     * @param _userAddress address to get normalized balance for
     * @return normalized balance for user
     */
    function getNormalizedBalance(address _userAddress) external view returns (uint256) {
        return accounts[_userAddress].normalizedBalance;
    }

    /**
     * @notice Claims all delegated awards
     * @return true if successfully
     */
    function claimStakeDelegatorReward() public returns (bool) {
        return _claimStakeDelegatorReward(msg.sender);
    }

    /**
     * Increases contract balance
     */
    function depositFromPool() public payable {
        _checkBalanceInvariant();
    }

    /**
     * @inheritdoc IERC20
     * @dev Moves `amount` tokens from the caller's account to `recipient`.
     * @notice [ERC20 method] transfer tokens from sender to receiver
     * @return true if transfer is successful
     */
    function transfer(address receiver, uint256 amount) public override returns (bool) {
        _subFromBalance(msg.sender, amount);
        _addToBalance(receiver, amount);
        _checkBalanceInvariant();
        emit Transfer(msg.sender, receiver, amount);
        return true;
    }

    /**
     * @inheritdoc IERC20
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     * @notice [ERC20 method] allows transferring sender's tokens
     * @return true if approving is successful
     */
    function approve(address spender, uint256 amount) public override returns (bool) {
        accounts[msg.sender].allowances[spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    /**
     * @inheritdoc IERC20
     * @dev Moves `amount` tokens from `sender` to `recipient` using the
     * @notice [ERC20 method] transfer allowed tokens from owner to receiver
     * @return true if transfer is successful
     */
    function transferFrom(
        address owner,
        address receiver,
        uint256 amount
    ) public override returns (bool) {
        uint256 currentAllowance = accounts[owner].allowances[msg.sender];
        require(currentAllowance >= amount, "[QEC-017020]-The allowance is too low for this operation.");
        accounts[owner].allowances[msg.sender] -= amount;
        _subFromBalance(owner, amount);
        _addToBalance(receiver, amount);
        _checkBalanceInvariant();
        emit Transfer(owner, receiver, amount);
        return true;
    }

    /**
     * @dev Atomically increases the allowance granted to `spender` by the caller.
     * @param spender the user whos allowance will be increased by `addedValue`
     * @param addedValue amount of increase in the allowance
     * @return success of the operation
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
        accounts[msg.sender].allowances[spender] += addedValue;
        return true;
    }

    /**
     * @dev Atomically decreases the allowance granted to `spender` by the caller.
     * @param spender the user whos allowance will be decreased by `subtractedValue`
     * @param subtractedValue amount of decrease in the allowance
     * @return success of the operation
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     * - `spender` must have allowance for the caller of at least
     * `subtractedValue`.
     */
    function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool) {
        uint256 currentAllowance = accounts[msg.sender].allowances[spender];
        require(currentAllowance >= subtractedValue, "[QEC-017020]-The allowance is too low for this operation.");
        accounts[msg.sender].allowances[spender] -= subtractedValue;
        return true;
    }

    /**
     * @notice Transfer tokens to the recipient
     * @param _to Address of the person to whom the funds will be transferred
     * @param _value The amount of funds to be transferred
     * @param _data Date of the transfer
     * @return true if everything went well
     */
    function transferAndCall(
        address _to,
        uint256 _value,
        bytes memory _data
    ) public override returns (bool) {
        bool result = transfer(_to, _value);
        if (!result) return false;

        IERC677TransferReceiver receiver = IERC677TransferReceiver(_to);
        receiver.tokenFallback(msg.sender, _value, _data);

        emit Transfer(msg.sender, _to, _value, _data);
        return true;
    }

    /**
     * @notice Withdraw from msg.sender's balance to _recipient
     * @param _recipient Recipient of funds
     * @param _amount The amount of funds to be withdrawn
     * @return true if everything went well
     */
    function withdrawTo(address _recipient, uint256 _amount) public returns (bool) {
        _subFromBalance(msg.sender, _amount);

        (bool _success, ) = _recipient.call{value: _amount}("");
        require(_success, "[QEC-017004]-Transfer of the withdrawal amount failed.");

        _checkBalanceInvariant();

        emit UserWithdrawn(_recipient, _amount, balanceOf(_recipient));

        return true;
    }

    /**
     * @notice deposit Q on recipient's balance
     * @param _recipient Recipient of funds
     * @return true if everything went well
     */
    function depositTo(address _recipient) public payable returns (bool) {
        require(msg.value > 0, "[QEC-017000]-Deposit amount must not be zero.");

        _addToBalance(_recipient, msg.value);

        _checkBalanceInvariant();

        emit UserDeposited(_recipient, msg.value, balanceOf(_recipient));
        return true;
    }

    /**
     * @notice [Deprecated use balanceOf] etrieves user balance
     * @param _userAddress address to get balance for
     * @return balance for user
     */
    function getUserBalance(address _userAddress) public view returns (uint256) {
        return balanceOf(_userAddress);
    }

    /**
     * @notice Returns detailed information about the user's balance
     * @return BalanceDetails struct
     */
    function getBalanceDetails() public view returns (BalanceDetails memory) {
        BalanceDetails memory _balanceDetails;
        CompoundRateKeeper _compoundRateKeeper = compoundRateKeeper;

        _balanceDetails.compoundRate = _compoundRateKeeper.getCurrentRate();
        _balanceDetails.lastUpdateOfCompoundRate = _compoundRateKeeper.getLastUpdate();
        _balanceDetails.currentBalance = balanceOf(msg.sender);
        _balanceDetails.normalizedBalance = accounts[msg.sender].normalizedBalance;
        _balanceDetails.interestRate = IParameters(_getEPQFIParametersAddress()).getUint(
            "governed.EPQFI.Q_rewardPoolInterest"
        );

        return _balanceDetails;
    }

    /**
     * @notice [ERC20 method] Returns total amount of tokens
     * @return amount of tokens in existence
     */
    function totalSupply() public view override returns (uint256) {
        return address(this).balance;
    }

    /**
     * @notice [ERC20 method] Returns balance of tokenOwner;
     * @param _userAddress owner of the tokens whose balance we check
     * @return amount of tokens owned by `_userAddress`
     */
    function balanceOf(address _userAddress) public view override returns (uint256) {
        return compoundRateKeeper.denormalizeAmount(accounts[_userAddress].normalizedBalance);
    }

    /**
     * @inheritdoc IERC20
     * @dev Returns the remaining number of tokens that `spender` will be
     * @notice [ERC20 method]  check allowance
     * @param tokenOwner owner of the tokens
     * @param spender user who is allowed to spend a certain amount of `tokenOwner` tokens
     * @return current allowance of `spender` to spend  `tokenOwner` tokens
     */
    function allowance(address tokenOwner, address spender) public view override returns (uint256) {
        return accounts[tokenOwner].allowances[spender];
    }

    /**
     * @notice Returns the name of the token.
     * @return the name of the token.
     */
    function name() public pure returns (string memory) {
        return "Vault Q";
    }

    /**
     * @notice Returns the symbol of the token, usually a shorter version of the name.
     * @return the symbol of the token.
     */
    function symbol() public pure returns (string memory) {
        return "VQ";
    }

    /**
     * @dev Returns the number of decimals used to get its user representation.
     * @return the number of decimals used to get its user representation.
     */
    function decimals() public pure returns (uint8) {
        return 18;
    }

    function _onTimeLockedDeposit(address _account, uint256 _amount) internal override {
        _addToBalance(_account, _amount);
    }

    function _getOrCreateDelegatedToStorage(address _delegator) private returns (AddressStorage) {
        DelegationInfo storage delegationInfo = accounts[_delegator].delegationInfo;
        if (delegationInfo.delegatedTo == AddressStorage(address(0))) {
            address[] memory t;
            delegationInfo.delegatedTo = AddressStorageFactory(registry.mustGetAddress(RKEY__ADDRESS_STORAGE_FACTORY))
                .create(t);
        }
        return delegationInfo.delegatedTo;
    }

    function _claimStakeDelegatorReward(address _delegator) private returns (bool) {
        if (accounts[_delegator].delegationInfo.totalDelegatedStake == 0) return false;
        ValidationRewardPools _pools = _getValidatorsRewardPools();

        uint256 _totalReward = 0;
        DelegationInfo storage delegationInfo = accounts[_delegator].delegationInfo;
        address[] memory _delegatedTo = _getOrCreateDelegatedToStorage(_delegator).getAddresses();
        for (uint256 i = 0; i < _delegatedTo.length; i++) {
            address _validator = _delegatedTo[i];

            uint256 _claimableReward = _pools.getDenormalizedAmount(
                _validator,
                delegationInfo.delegatedAmounts[_validator].normalizedStake
            ) - delegationInfo.delegatedAmounts[_validator].actualStake;

            if (_claimableReward == 0) {
                continue;
            }

            _pools.requestRewardTransfer(_validator, _claimableReward);
            delegationInfo.delegatedAmounts[_validator].actualStake += _claimableReward;

            _totalReward += _claimableReward;
        }

        delegationInfo.totalDelegatedStake += _totalReward;
        _addToBalance(_delegator, _totalReward);

        emit DelegatorRewardClaimed(_delegator, _totalReward, balanceOf(_delegator));
        return true;
    }

    function _addToBalance(address _receiver, uint256 _amount) private {
        if (_amount == 0) return;

        uint256 targetBalance = balanceOf(_receiver) + _amount;
        uint256 newNormalizedBalance = compoundRateKeeper.normalizeAmount(targetBalance);

        UserAccount storage account = accounts[_receiver];
        if (newNormalizedBalance <= account.normalizedBalance) return;

        uint256 normalizedDiff = newNormalizedBalance - account.normalizedBalance;
        account.normalizedBalance = newNormalizedBalance;
        aggregatedNormalizedBalance += normalizedDiff;
    }

    function _subFromBalance(address _owner, uint256 _amount) private {
        uint256 _balanceOfOwner = balanceOf(_owner);
        require(_amount <= _balanceOfOwner, "[QEC-017021]-Insufficient balance for withdrawal.");
        if (_amount == 0) return;
        VotingLockInfo memory _votingLockInfo = _getLockInfo(_owner);

        uint256 _targetBalance = _balanceOfOwner - _amount;

        require(
            _targetBalance >= accounts[_owner].delegationInfo.totalDelegatedStake,
            "[QEC-017014]-Insufficient funds to cover all delegations."
        );

        require(
            _targetBalance >= getMinimumBalance(_owner, block.timestamp),
            "[QEC-017022]-Balance must not fall below current time locks."
        );

        uint256 _totalLockedAmount = _votingLockInfo.lockedAmount + _votingLockInfo.pendingUnlockAmount;
        uint256 _nonLockedAmount = _balanceOfOwner - _totalLockedAmount;

        if (_nonLockedAmount < _amount) {
            _getVotingWeightProxy().unlock(_owner, _amount - _nonLockedAmount);
        }

        uint256 _newNormalizedBalance = compoundRateKeeper.normalizeAmount(_targetBalance);
        uint256 _normalizedDiff = accounts[_owner].normalizedBalance - _newNormalizedBalance;

        accounts[_owner].normalizedBalance = _newNormalizedBalance;
        aggregatedNormalizedBalance -= _normalizedDiff;
    }

    function _getVotingWeightProxy() private view returns (IVotingWeightProxy) {
        return IVotingWeightProxy(registry.mustGetAddress(RKEY__VOTING_WEIGHT_PROXY));
    }

    function _getValidatorsRewardPools() private view returns (ValidationRewardPools) {
        return ValidationRewardPools(registry.mustGetAddress(RKEY__VALIDATION_REWARDS_POOL));
    }

    function _getLockInfo(address _account) private view returns (VotingLockInfo memory) {
        return _getVotingWeightProxy().getLockInfo(address(this), _account);
    }

    function _getLockInfo() private view returns (VotingLockInfo memory) {
        return _getLockInfo(msg.sender);
    }

    function _checkTokenLockInvariance() private view {
        VotingLockInfo memory _votingLockInfo = _getLockInfo();
        uint256 _totalLockedAmount = _votingLockInfo.lockedAmount + _votingLockInfo.pendingUnlockAmount;

        assert(balanceOf(msg.sender) >= _totalLockedAmount);
    }

    function _checkBalanceInvariant() private view {
        assert(address(this).balance >= compoundRateKeeper.denormalizeAmount(aggregatedNormalizedBalance));
    }

    function _getEPQFIParametersAddress() private view returns (address) {
        return registry.mustGetAddress(RKEY__EPQFI_PARAMETERS);
    }

    function _getQHolderRewardPoolAddress() private view returns (address) {
        return registry.mustGetAddress(RKEY__QHOLDER_REWARD_POOL);
    }
}
