# portifyx-market-data-service

Market Data microservice caching CoinGecko live pricing, Bitkub/Binance ticker caches, and foreign exchange rates (USD/THB).

## Database Isolation

Exclusively queries and manages MongoDB database `portifyx_marketdata` and Redis TDC.

## Running Locally

```bash
make run
# or
go run main.go api
```
