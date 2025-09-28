# Arkh Blockchain Testnet Setup Script
# This script sets up a complete testnet for testing IBC functionality

Write-Host "🚀 Setting up Arkh Blockchain Testnet..." -ForegroundColor Green

# Step 1: Initialize the blockchain
Write-Host "📝 Step 1: Initializing blockchain..." -ForegroundColor Yellow
./arkhd init testnet --chain-id arkh-testnet-1

# Step 2: Create test accounts
Write-Host "👤 Step 2: Creating test accounts..." -ForegroundColor Yellow
./arkhd keys add alice --keyring-backend test
./arkhd keys add bob --keyring-backend test
./arkhd keys add validator --keyring-backend test

# Step 3: Add accounts to genesis
Write-Host "💰 Step 3: Adding accounts to genesis..." -ForegroundColor Yellow
./arkhd add-genesis-account alice 1000000000uarkh --keyring-backend test
./arkhd add-genesis-account bob 1000000000uarkh --keyring-backend test
./arkhd add-genesis-account validator 1000000000uarkh --keyring-backend test

# Step 4: Create validator
Write-Host "🏛️ Step 4: Creating validator..." -ForegroundColor Yellow
./arkhd gentx validator 1000000uarkh --chain-id arkh-testnet-1 --keyring-backend test

# Step 5: Collect genesis transactions
Write-Host "📋 Step 5: Collecting genesis transactions..." -ForegroundColor Yellow
./arkhd collect-gentxs

# Step 6: Validate genesis
Write-Host "✅ Step 6: Validating genesis..." -ForegroundColor Yellow
./arkhd validate-genesis

Write-Host "🎉 Testnet setup complete!" -ForegroundColor Green
Write-Host "To start the testnet, run: ./arkhd start" -ForegroundColor Cyan

