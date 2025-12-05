# Block fetcher
A CLI tool to fetch Ethereum block data, transactions, and ERC-20 transfer events (latter TBD)

## Installation
```bash
go install github.com/seeques/block-fetcher@latest
```

## Usage

### Fetch transactions from a block
```bash
# Latest block (local Anvil)
block-fetcher txs

# Specific block
block-fetcher txs --block 12345

# Custom RPC
block-fetcher txs --rpc https://eth-mainnet.g.alchemy.com/v2/YOUR_KEY --block 12345

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
block-fetcher --address 0xERC20_ADDR

# Specific block
block-fetcher --address 0xERC20_ADDR --block 12345