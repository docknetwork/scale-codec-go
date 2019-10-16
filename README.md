Golang SCALE Codec
=
This is an implementation of Scale-codec in go. 
The following implementations were used as a reference:
* [Python](https://github.com/polkascan/py-scale-codec)
* [Rust](https://github.com/paritytech/parity-scale-codec)


To know more about the role of this library check this [link](https://medium.com/polkadot-network/polkascan-development-update-1-8451c4fcfc2e).

And more info about CODEC types is [here](https://polkadot.js.org/api/types/#codec-types).

Installation
-

Do
```shell script
go get github.com/docknetwork/scale-codec-go/codec
```
or from cloned repo
```shell script
cd codec && go install
``` 

Examples
-
Parsing primitive types:
```go
package main

import (
	"fmt"
	"scale/codec"
)

func main() {
	offsetBytes, err := codec.NewBytes("0x02093d00")
	value, err := offsetBytes.ToCompactUInt32()
	fmt.Println(value, err)
    // 1000000 <nil>
}
```

Parsing bytes to existing structure:
```go
package main

import (
	"fmt"
	"scale/codec"
)


func main() {
	offsetBytes, err := codec.NewBytes("0x0c00")
	prefs, err := offsetBytes.ToValidatorPrefsLegacy()
	fmt.Println(prefs.UnstakeThreshold, err)
    // 3 <nil>  
}
```

Creating your own structure:
```go
package main

import (
	"fmt"
	"scale/codec"
)

type ValidatorPrefsLegacy struct {
	UnstakeThreshold codec.U32
	ValidatorPayment codec.Balance
}

func ToValidatorPrefsLegacy(sb *codec.OffsetBytes) (res ValidatorPrefsLegacy, err error) {
	unstakeThreshold, err := sb.ToCompactUInt32()
	if err != nil {
		return
	}
	validatorPayment, err := sb.ToCompactBalance()
	if err != nil {
		return
	}
	res.UnstakeThreshold = unstakeThreshold
	res.ValidatorPayment = validatorPayment
	return
}

func main() {
	offsetBytes, err := codec.NewBytes("0x0c00")
	prefs, err := ToValidatorPrefsLegacy(&offsetBytes)
	fmt.Println(prefs.UnstakeThreshold, err)
}
```