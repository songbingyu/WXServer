package server

import (
    "fmt"
    "crypto/sha1"
    "sort"
    "io"
)


var  token = "wanjuanshuwu";

func CheckSignature( signature string,  timestamp string, nonce string )  bool {
	
	tmps:=[]string{token,timestamp,nonce}
	sort.Strings(tmps)
	tmpStr:=tmps[0]+tmps[1]+tmps[2]
    tmp:=str2sha1(tmpStr)
   if tmp==signature {
   		return true
   }

   return false
}

func str2sha1(data string) string{
    t:=sha1.New()
    io.WriteString(t,data)
    return fmt.Sprintf("%x",t.Sum(nil))
}


