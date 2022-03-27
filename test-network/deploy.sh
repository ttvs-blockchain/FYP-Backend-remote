./network.sh down

./network.sh up createChannel -c mychannel -ca

./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go

./network.sh up createChannel -c globalchannel -ca

./network.sh deployCC -c globalchannel -ccn basic-global -ccp ../asset-transfer-basic/chaincode-go-global/ -ccl go

cd ../asset-transfer-basic/application-go
rm wallet/appUser.id 
go run main.go