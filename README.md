# polygonscan-go-api
Wrapper around polygonscan's API to easily query it in go

---
## Installation

`go get github.com/VrncQuentin/polygonscan-go-api`

---
## Roadmap
Here's the current TODO list of the project.

### API Implementation
- [x] Account
  - [x] Balances
  - [x] Transaction histories
  - [x] Token Transfer Events
  - [x] Mined Blocks
- [ ] Contracts
- [x] Transactions
- [x] Blocks
- [ ] Logs
- [ ] GETH/Parity Proxy
- [x] Tokens
  - [x] ERC20 Supply Info (in stats module)
  - [x] ERC20 Balance (in account module)
- [x] Stats

- [ ] Tests

### Improvements

- [ ] Full Documentation
- [ ] Improve typing, currently the API only returns string: convert to more accurate types 
- [ ] Reduce code duplications
