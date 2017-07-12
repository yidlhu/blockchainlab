export PATH=$PATH:$GOPATH/bin

cd $GOPATH/src/chaincodes/BCDEMO

echo 'clear vendor folders'
rm -Rf vendor
rm BCDEMO

echo 'building...'
govendor init
govendor add +external
govendor add github.com/hyperledger/fabric/peer
govendor add chaincodes/BCDEMO/db
govendor add chaincodes/BCDEMO/model
# govendor add chaincodes/assets
go build





