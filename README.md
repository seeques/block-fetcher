# block-fetcher

A CLI tool to fetch Ethereum block data, transactions, ERC-20 transfer events, and decode ERC-20 method calls.

## Installation
```bash
go install github.com/seeques/block-fetcher@latest
```

Or build from source:
```bash
git clone https://github.com/seeques/block-fetcher.git
cd block-fetcher
go build -o block-fetcher
```

## Usage

By default, connects to `http://localhost:8545` (Anvil/local node).

Use `--rpc` flag for remote nodes:
```bash
block-fetcher  --rpc https://eth-mainnet.g.alchemy.com/v2/YOUR_KEY
```

### Commands

| Command | Description |
|---------|-------------|
| `txs` | Fetch transactions from a block |
| `receipt` | Get transaction receipt |
| `events` | Fetch ERC-20 transfer events |
| `selectors` | Decode ERC-20 method call from tx data |

### txs

Fetch transactions from a block:
```bash
block-fetcher txs                    # latest block
block-fetcher txs --block 12345      # specific block
block-fetcher txs --block 12345 -r   # include receipts
```

### receipt

Get a transaction receipt:
```bash
block-fetcher receipt 0xYOUR_TX_HASH
```

### events

Fetch ERC-20 transfer events:
```bash
block-fetcher events --address 0xERC20_ADDR              # latest block
block-fetcher events --address 0xERC20_ADDR --block 123  # specific block
```

Example output:
```
Block Number: 23946951
Index: 0
Event name: Transfer
From: 0xba1738b7dF19509ed812AF60bA1D023FfC0650B4
To: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Value: 300000000000000000000000
```

### selectors

Decode ERC-20 method call from transaction input data:
```bash
block-fetcher selectors --data a9059cbb000000000000000000000000...
```

> Note: `--data` should not include the `0x` prefix

Example output:
```
Method: transfer
To: 0x6BDaa62de230e76CF56662DFA8aa4159362BDdd1
Value: 700711029142889
```
