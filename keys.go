package beef

import "fmt"

func bucketNameForBF(name string) []byte {
	return []byte(fmt.Sprintf("bf:%v", name))
}
