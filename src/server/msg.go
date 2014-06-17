package server

import (
    //"fmt"
    "encoding/xml"  
    //"bytes"
    "strconv"
    //"io"
    //"github.com/golang/glog"
)



type Msg  map[string] interface{}

type BaseMsg struct {

    XMLName xml.Name `xml:"xml"`
    field []string `xml:"any"`
}


func ( m Msg ) String(key string) string {
    
    if str, ok := m[key]; ok {
        return str.(string)
    }
    return ""
}

func (m Msg ) Int64(key string) int64 {
     
    if val, ok := m[key]; ok {
        switch val.(type) {
        case string:
             i, _ := strconv.ParseInt(val.(string), 0, 64)
            return i
        case int:
            return int64(val.(int))
        case int64:
            return val.(int64)
        }
    }
    return 0
}

func ( m Msg ) ToUserName() string {
  
      return m.String("ToUserName")
}

func ( m Msg ) FromUserName() string {
        
      return m.String("FromUserName")
}

func (m Msg ) CreateTime() int64 {
    
      return m.Int64("CreateTime")
}
func ( m Msg ) MsgType() string {
    
      return m.String("MsgType")
}
func ( m  Msg ) MsgId() int64 {
    
      return m.Int64("MsgId")
}

func ( m Msg ) Content() string {
    
     return m.String("Content")
}

func ( m Msg ) PicUrl() string {
    
    return m.String("PicUrl")
}

func ( m Msg ) MediaId() string {
   
    return m.String("MediaId")

}

func ( m Msg) Format() string {

    return m.String("Format");
}

func ( m Msg ) ThumbMediaId() string {

    return m.String("ThumbMediaId");
}

func (m Msg ) Location_X() string {
    
    return m.String("Location_X")
}

func (m Msg ) Location_Y() string {
    
    return m.String("Location_Y")
}

func (m Msg ) Scale() int64 {
    
    return m.Int64("Scale")
}

func (m Msg ) Label() string {
    
    return m.String("Label")
}

func (m Msg ) Title() string {
    
    return m.String("Title")
}

func ( m Msg ) Description() string {
   
    return m.String("Description")
}

func ( m Msg ) Url() string {
    
    return m.String("Url")
}

func ( m Msg ) Event() string {
    
    return m.String("Event")
}

func ( m Msg ) EventKey() string {
   
     return m.String("EventKey")
}

func ( m Msg ) Ticket() string {
   
    return m.String("Ticket"); 
}

func ( m Msg ) Latitude() string {
   
    return m.String("Latitude");
}

func ( m Msg ) Longitude() string {
   
     return m.String("Longitude");
}

func ( m Msg ) Precision() string {
    
    return m.String("Precision");
}


type respBase struct {

    ToUserName      string  
    FromUserName    string
    CreateTime      int64
    MsgType         string
     
}




type RespText   struct {

    XMLName     xml.Name    `xml:"xml"`
    respBase      
    Content     string 
}

type MusicInfo struct {

  XMLName       xml.Name    `xml:"Music"`
  Title         string
  Description   string
  MusicUrl      string
  HQMusicUrl    string
 // ThumbMediaId  string

}


type RespMusic struct {

    XMLName     xml.Name    `xml:"xml"`
    respBase    
    Music       MusicInfo 

}
type item  struct {

  XMLName       xml.Name    `xml:"item"`
  Title         string
  Description   string
  PicUrl        string
  Url           string
 
}
type RespTextAndPic struct {

    XMLName         xml.Name    `xml:"xml"`
    respBase    
    ArticleCount    int64       `xml:",omitempty"`
    Articles        []item     `xml:"Articles>item,omitempty"` 
}


