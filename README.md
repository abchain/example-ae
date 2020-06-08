# A full example for using hyperledger framework: the "atomenergy" project

This project illustrating the integration of chaincode modules and REST API provided by framework

### The SDK providing REST API

Under the ./sdk directory, you can build a completely "local" system to test your chaincode without 
running it on a real fabric platform (like fabric 1.4 or ya-fabric 0.9x). And it will be easy to 
distribute the same REST service part with platform-spec supporting by just add an shadow import 
of the corresponding adapter.
