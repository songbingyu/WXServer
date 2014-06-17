//百度音乐搜索

package util

import (
    "strings"
    "github.com/golang/glog"
    "encoding/xml"
    "fmt"
     "github.com/qiniu/iconv"
    //"xml"
    //"github.com/djimenez/iconv-go"
    //"github.com/donnie4w/dom4g"

)

const musicUrl ="http://box.zhangmen.baidu.com/x?op=12&count=1&"

func Substr(str string, start, length int) string {
    rs := []rune(str)
    rl := len(rs)
    end := 0
        
    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length
    
    if start > end {
        start, end = end, start
    }
    
    if start < 0 {
        start = 0
    }
  if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }
    
    return string(rs[start:end])
}


type music struct {

    Encode  string  `xml:"encode"`
    Decode  string `xml:"decode"`
    Type    string `xml:"type"`
    Lrcid   string `xml:"lrcid"`
    Flag    string `xml:"flag"`

}

type p2p struct {

    XMLName xml.Name       `xml:"p2p"`
    Hash       string     `xml:"hash"`
    Url     string        `xml:"url"` 
    Type    string         `xml:"type"`
    Size    int64          `xml:"size"`
    Bitrate int64          `xml:"bitrace"`

}

type MusicUrl    struct {
    Url     music  `xml:"url"` 
    Durl    music  `xml:"durl,omitempty"` 
    P2p     p2p     
}


type MusicInfo struct {

    XMLName     xml.Name  `xml:"result"` 
    Count   string `xml:"count"`
   //MusicUrls   [] MusicUrl  `xml:",innerxml"`
    Url          music `xml:"url"`
    Durl         music  `xml:"durl"`
    P2p          p2p  `xml:p2p` 
 }


func  SearchMusic ( title string, author string  ) ( e error,  MusicUrl string, HQMusicUrl string, ThumbMediaId string )  {

    url := musicUrl + "title=" + title + "$$"+ author + "$$$$"
   fmt.Println( url ) 
    url = strings.Trim( url, " " )
    
    glog.Info( url )
 
    res, _ := get_info( url )
   
    //fmt.Println( string( res) )

    var musicInfo  MusicInfo
   

	cd, err := iconv.Open( "utf-8", "gb2312" )
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()

    data := cd.ConvString( string(res ) )

    /*_,_, ok := iconv.Convert(  res , data, "gb2312", "utf-8")
    _,_, ok := iconv.Convert(  res , data )
    if ok != nil {
       e = ok 
        return
    }*/
   
     
     
test := Substr(string( data ),strings.Index( string(data), "<result" ), len(string(data))  ) 
test1 := Substr(string( test ), 0 , strings.Index( string( test), "/durl>") +6  ) 
    test1  = test1 + "</result>"
    fmt.Println( string(data) )
    fmt.Println("------------------------------------------" )
    fmt.Println( test )
    fmt.Println("------------------------------------------" )
    fmt.Println( test1 )
    

    // err := xml.Unmarshal( [] byte ( test ), &musicInfo)
    err = xml.Unmarshal( [] byte ( test1 ), &musicInfo)
    if err != nil {
       
        e = err 
        fmt.Println( e ,musicInfo )
        return
    }

    fmt.Println(" begin parse" )
    
    


    fmt.Println( musicInfo ) 
    fmt.Println( musicInfo.Url.Encode  )
    fmt.Println( musicInfo.Url.Decode )
    fmt.Println("---------------------------------------")
    if musicInfo.Count == "0" {
    
        return
    }
    musicUrl := Substr( musicInfo.Url.Encode, 0, strings.LastIndex(musicInfo.Url.Encode,"/" ) +1 ) + musicInfo.Url.Decode  
    hqMusicUrl := Substr( musicInfo.Durl.Encode, 0,strings.LastIndex(musicInfo.Durl.Encode,"/" ) +1 ) + musicInfo.Durl.Decode  
    e =nil 
    fmt.Println( MusicUrl )
    fmt.Println( HQMusicUrl )
    cd2, err2 := iconv.Open( "gb2312", "utf8" )
	if err2 != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd2.Close()

    //绝壁一大坑
    MusicUrl = "<![CDATA["+ musicUrl +"]]>"

    HQMusicUrl ="<![CDATA["+ hqMusicUrl +"]]>"



    return  

}


