package server

import (
    "fmt"
    "encoding/xml"  
    "bytes"
    "strconv"
    "io"
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


type Replay map[string]interface{}

func (r Replay) String(key string) string {
    if str, ok := r[key]; ok {
        return str.(string)
    }
    return ""
}


func (r Replay) Int64(key string) int64 {
    if val, ok := r[key]; ok {
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

func (r Replay) ToUserName() string {
    return r.String("ToUserName")
}
func (r Replay) FromUserName() string {
    return r.String("FromUserName")
}
func (r Replay) CreateTime() int64 {
    return r.Int64("CreateTime")
}
func (r Replay) MsgType() string {
    return r.String("MsgType")
}
func (r Replay) FuncFlag() int64 {
    return r.Int64("FuncFlag")
}

func (r Replay) SetToUserName(val string) Replay {
    r["ToUserName"] = val
    return r
}
func (r Replay) SetFromUserName(val string) Replay {
    r["FromUserName"] = val
    return r
}
func (r Replay) SetCreateTime(val int64) Replay {
    r["CreateTime"] = val
    return r
}
func (r Replay) SetMsgType(val string) Replay {
    r["MsgType"] = val
    return r
}
func (r Replay) SetFuncFlag(val int64) Replay {
    r["FuncFlag"] = val
    return r
}

func (r Replay) Content() string {
    return r.String("Content")
}
func (r Replay) SetContent(val string) Replay {
    r["Content"] = val
    return r
}

func MapToXmlString(m map[string]interface{}) string {
    buf := &bytes.Buffer{}
    for k, v := range m {

        if v != nil {
            switch v.(type) {
            case int:
                io.WriteString(buf, fmt.Sprintf("<%s>", k))
                io.WriteString(buf, fmt.Sprintf("%d", v))
                io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
            case int64:
                io.WriteString(buf, fmt.Sprintf("<%s>", k))
                io.WriteString(buf, fmt.Sprintf("%d", v))
                io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
            case string:
                io.WriteString(buf, fmt.Sprintf("<%s>", k))
                io.WriteString(buf, "<![CDATA["+v.(string)+"]]>")
                io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
            case map[string]interface{}:
                io.WriteString(buf, fmt.Sprintf("<%s>", k))
                io.WriteString(buf, MapToXmlString(v.(map[string]interface{})))
                io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
            case []interface{}:
                for _, t := range v.([]interface{}) {
                    switch t.(type) {
                    case map[string]interface{}:
                        io.WriteString(buf, fmt.Sprintf("<%s>", k))
                        io.WriteString(buf, MapToXmlString(t.(map[string]interface{})))
                        io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
                    }
                }
            }
        }

    }
    return buf.String()
}

