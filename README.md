# Block fetcher
A CLI tool to fetch Ethereum block data, transactions, ERC-20 transfer events, and ERC20 method calls from transaction data
## Installation
```bash
go install github.com/seeques/block-fetcher@latest
```

## Usage

### Global flag
```bash
block-fetcher any_command --rpc https://eth-mainnet.g.alchemy.com/v2/YOUR_KEY
```

### Fetch transactions from a block
```bash
# Latest block (local Anvil)
block-fetcher txs

# Specific block
block-fetcher txs --block 12345

# Show receipt as if enabled
block-fetcher txs --block 1 -r true
```

### Get transaction receipt only
```bash
block-fetcher receipt 0xYOUR_TX_HASH
```

### Fetch ERC-20 transfer events
```bash
# Latest block by default
block-fetcher events --address 0xERC20_ADDR

# Specific block
block-fetcher events --address 0xERC20_ADDR --block 12345
```

**Example from anvil fork of ethereum**:
```bash
# Input
block-fetcher events --address $DAI
# Output
Block Number: 23946951
Index: 0
Event name: Transfer
From: 0xba1738b7dF19509ed812AF60bA1D023FfC0650B4
To: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Value: 300000000000000000000000
```

### Fetch ERC-20 method call from transaction data
```bash
# RAW_BYTES must not include 0x
block-fetcher selectors --data RAW_BYTES
```

**Example from anvil fork of ethereum**:
```bash
# Input
block-fetcher selectors --data a9059cbb0000000000000000000000006bdaa62de230e76cf56662dfa8aa4159362bddd100000000000000000000000000000000000000000000000000027d4afffb7569
# Output
Method: transfer
To: 0x6BDaa62de230e76CF56662DFA8aa4159362BDdd1
Value: 700711029142889
```