package main

import (
	"fmt"
	"github.com/majestrate/i2p-tools/sam3"
)

const yoursam = "127.0.0.1:7656" // sam bridge

func main() {

	//akey := sam3.I2PKeys{"IpZoI5s~6KDQIEh5LxQ3lKa3--cvduBND-oti8FrV5pFrJsJZkhaoa9iI1xSVTgDaUrlo~wu3n3JCyfoKjEmrc8KKmoH6N7XlS18sCqQDWG-hsMC9m9scD6qfKA~pkI7Zha4opeEISzvpxOlCQw3yqslXAztAfIoWo5oEMFgxVaDblWjV-IrNF2UZFVttVVRR2osEhg-b4hk5Gfcc~XWtQd~pYoNkRdKCGK6tqQ8AoPiXPYsAM28ig-ZNMBbUEHE3oDJxXzBhbNFmNBxwSe7zy6Y3uxWPPdI2zNmqQZMSXSvK1WAJrxU0qvl3PUyDztDB6teAt-uo6DNJ2N5xwKVUAoq2Bopcgqi4zSFWLWjc6tc6IrCKde2h~KtSAn30oUBeqYj6xLbq4JceG-7451ZWQr3Ld2YvHKFWb79YN9UDHFq2mHQflXpI5T1RoQFG2yU7h6dA6HMOz8mEfNKyuHSVmDlRk3zs39aW6CAOvbs3AB4mpOEZK-~XIgPyEjFlMxEAAAA",
	//	"IpZoI5s~6KDQIEh5LxQ3lKa3--cvduBND-oti8FrV5pFrJsJZkhaoa9iI1xSVTgDaUrlo~wu3n3JCyfoKjEmrc8KKmoH6N7XlS18sCqQDWG-hsMC9m9scD6qfKA~pkI7Zha4opeEISzvpxOlCQw3yqslXAztAfIoWo5oEMFgxVaDblWjV-IrNF2UZFVttVVRR2osEhg-b4hk5Gfcc~XWtQd~pYoNkRdKCGK6tqQ8AoPiXPYsAM28ig-ZNMBbUEHE3oDJxXzBhbNFmNBxwSe7zy6Y3uxWPPdI2zNmqQZMSXSvK1WAJrxU0qvl3PUyDztDB6teAt-uo6DNJ2N5xwKVUAoq2Bopcgqi4zSFWLWjc6tc6IrCKde2h~KtSAn30oUBeqYj6xLbq4JceG-7451ZWQr3Ld2YvHKFWb79YN9UDHFq2mHQflXpI5T1RoQFG2yU7h6dA6HMOz8mEfNKyuHSVmDlRk3zs39aW6CAOvbs3AB4mpOEZK-~XIgPyEjFlMxEAAAAXuEA8Cwa6lt~9FbuJmUdsAzVcCUOyQ7OG2Te4pS1lTGFlQBf4yjVbbnRIabkeXy1Zv-WeR3xDX~ijay1J~4ecnmESK00rx03~kzV~sQcZexDYmN011RtxJQPzn8bWqiOBzcA9vvLOJZpJ7dxesfh6eM4etk19q8HjczR95YrD3uus39EnQHKcOL~9Y4dJVFXJcMcMpJ-I4n5AJnC1XpmtOexpHQqFQWlUGooY3
	key, _ := sam3.LoadKeysIncompat(akey)

	fmt.Println(key.String())

}