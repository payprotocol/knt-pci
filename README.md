# Payprotocol Token Chaincode

Payprotocol Token (PCI)

#

## API

method __`func`__ [arg1, _arg2_, ... ] {trs1, _trs2_, ... }
- method : __query__ or __invoke__
- func : function name
- [arg] : mandatory argument
- [_arg_] : optional argument
- {trs} : mandatory transient
- {_trs_} : optional transient

#

> query __`burn`__ [total_supply, genesis_account_balance, max_amount]
- Get burnable amount
- [total_supply] : current total supply
- [genesis_account_balance] : current balance amount of the genesis account
- [max_amount] : amount to burn, but the actual return value less or equal than it, (if max_amount <= 0, no limit)

> query __`mint`__ [total_supply, genesis_account_balance, max_amount]
- Get mintable amount
- [total_supply] : current total supply
- [genesis_account_balance] : current balance amount of the genesis account
- [max_amount] : amount to mint, but the actual return value less or equal than it, (if max_amount <= 0, no limit)

> query __`token`__
- Get initial meta data of the token
