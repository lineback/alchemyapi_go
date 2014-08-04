# alchemyapi_go #

A sdk for AlchemyAPI using Golang


## AlchemyAPI ##


More information at: http://www.alchemyapi.com



## API Key ##

To use AlchemyAPI, you'll need to obtain an API key and attach that key to all requests. If you do not already have a key, please visit: http://www.alchemyapi.com/api/register.html



## Getting Started with the Golang SDK ##

To get started and run the example, simply:

   go get github.com/lineback/alchemyapi_go
   cd $GOPATH/src/github.com/lineback/alchemyapi_go/alchemyInit/
   go install
   cd ../alchemyAPI/
   alchemyInit YOUR_API_KEY
   go test 
   cd ../AlchemyExamples
   go install 
   alchemyInit YOUR_API_KEY
   alchemyExamples


Just replace YOUR_API_KEY with your 40 character API key from AlchemyAPI, and you should be good to go.