package server

import (
    //"errors"i
    "fmt"
    "strconv"
    "net/http"
    "time"
    "github.com/golang/glog"
    "encoding/xml"
    "util" 
    "strings"
)

const (
    MSG_TEXT     = "text"
    MSG_IMAGE    = "image"
    MSG_LOCATION = "location"
    MSG_LINK     = "link"
    MSG_EVENT    = "event"
    MSG_VOICE    = "voice"
    MSG_VIDEO    = "video" 
    MSG_NEWS     = "news"
    MSG_MUSIC    = "music"
)

const (
    EVENT_SUBSCRIBE   = "subscribe"
    EVENT_UNSUBSCRIBE = "unsubscribe"
    EVENT_LOCATION    = "LOCATION"
    EVENT_CLICK       = "CLICK"
    EVENT_VIEW        = "VIEW"
)



type MsgProc  struct{


} 

func ( proc MsgProc)  Proc( w http.ResponseWriter,  m * Msg ) error {

    msgType := m.MsgType()
    switch( msgType ) {
        case MSG_TEXT:
             i := proc.proc_text( w, m )
             return i
        case MSG_IMAGE:
             i := proc.proc_pic ( w, m )
             return i
        case MSG_EVENT:
             i := proc.proc_event( w,m )
             return i 
        case MSG_LOCATION:
             return nil
        case MSG_LINK:
             return nil
        case MSG_VOICE:
             return nil
        case MSG_VIDEO:
             return nil
    } 
    
   return nil
}


func ( proc MsgProc ) proc_text( w http.ResponseWriter,  m *Msg )  error {
  
    b := strings.Contains( m.Content(), "音乐@")
  //  fmt.Println(*m )
    if b == true  {

        var title string
        var author string
        content := strings.Trim( m.Content(), " " )
        keys := strings.Split( content, "@" )
        fmt.Println( m.Content(), keys ,  len(keys) ) 
        if len(keys) == 1 {
            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "请输入音乐名!") ); 
            return nil
        }
        
        title = keys[1]
        author = ""
        if len(keys)  >= 3 {
            author = keys[2] 
        }
        fmt.Println( title, author )
 
        ok, musicUrl, HDMusicUrl , _:= util.SearchMusic( title, author )
        if  ok!= nil  {
            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "你找的音乐不存在!") ); 

            return nil
        } 
        
        fmt.Println( " go......" )  
        resp := RespMusic{}
        resp.ToUserName = m.FromUserName()
        resp.FromUserName =m.ToUserName() 
        resp.CreateTime = time.Now().Unix()
        resp.MsgType =MSG_MUSIC
        resp.Music.MusicUrl = strings.Trim( musicUrl," " )
        resp.Music.HQMusicUrl = strings.Trim(  HDMusicUrl," ")
        resp.Music.Title = title
        resp.Music.Description = "百度歌曲"
        //resp.Music.ThumbMediaId="00"
        
        proc.send_music( w, resp )
        return nil 

    }

 
    b = strings.Contains( m.Content(), "天气@")
    if b == true  {

        var location  string
        content := strings.Trim( m.Content(), " " )
        keys := strings.Split( content, "@" )
        if len(keys) == 1 {
            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "请输入音乐名!") ); 
            return nil
        }
        
        location = keys[1]
        ok, weather:= util.SearchWeather( location )
        if ok != nil {
        
            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "没找到该地方。。。") ) 
            return nil
        }
        
        if  weather.Error != 0 {

            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "没找到该地方。。。") ) 
            return nil
            
        } 
   
        items := [] item {}
               for _,v := range weather.Results[0].Whether_data  {
                t   := item{}
                t.Title = v.Date +" "+ v.Weather + " "+v.Wind+" " +v.Temperature
                t.Description =""
                t.PicUrl = v.DayPictureUrl
                t.Url    = v.DayPictureUrl
                items    =append( items, t )        

        }

        proc.send_pic_and_text( w,m.FromUserName(), m.ToUserName(), items )
        return nil
    }       

    return nil
}

func ( proc MsgProc ) proc_event(w http.ResponseWriter,  m *Msg )  error {

    eventType := m.Event();
    switch ( eventType ) {
        case EVENT_SUBSCRIBE:
            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "好读书，爱学习。好学书屋，http://www.HaoXueShuWu.com。主营IT类、游戏类、经管类书籍和相关资料。欢迎您的光临-_") ); 
            return nil
        case EVENT_UNSUBSCRIBE:
            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "好读书，爱学习。好学书屋，http://www.HaoXueShuWu.com。主营IT类、游戏类、经管类书籍和相关资料。期待您的光临-_") ); 
            return nil
        case EVENT_LOCATION:
            proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "你的位置是") ); 
            return nil
    }
    return nil

}

func ( proc MsgProc ) proc_pic( w http.ResponseWriter,  m *Msg ) error {

    faceInfo :=util.DetectionDetect( m.PicUrl()  )
    glog.Info( faceInfo )
    var content string=""
    switch len( faceInfo.Face ) {
        case 0:
         content = "哥们，你长的太帅了，系统都识别不出来了？"
        case 1:
             attribute:=faceInfo.Face[0].Attribute
             age:=attribute.Age
             gender:=attribute.Gender
             var faceGender string
             if gender.Value=="Male"{
                 faceGender="男"
             }else{
                 faceGender="女"
             }
             faceAgeValue:=fmt.Sprintf("%d",int(age.Value))
             faceAgeRange:=fmt.Sprintf("%d",int(age.Range))
             race  := attribute.Race
             glass := attribute.Glass
             smil  := attribute.Smiling
             content="性别："+faceGender+"\n"+"年龄："+faceAgeValue+"(±"+faceAgeRange+")" + "\n" +"人种:" + race.Value + "\n" +"微笑程度:" + strconv.FormatFloat( smil.Value, 'f', -1, 64) + "\n"+"眼镜:" + glass.Value    
        default: 
            content =" 人太多了，搞基吗？" 
    } 
            
    proc.send_text(w,  m.FromUserName(), m.ToUserName(), content  ); 
    return nil

} 
func ( proc MsgProc ) send_text( w http.ResponseWriter, ToUserName string , FromUserName string, Content string ) {

  resp := RespText{}
  resp.ToUserName = ToUserName 
  resp.FromUserName = FromUserName 
  resp.CreateTime = time.Now().Unix()
  resp.MsgType= MSG_TEXT 
  resp.Content = Content
  data,_ := xml.Marshal( resp )  
  SendMsg( w, data )
}

func ( proc MsgProc ) send_music( w http.ResponseWriter,  respMusic RespMusic ) {

    //data,_ := xml.Marshal( respMusic ) 
    data,_ := xml.MarshalIndent( respMusic,"" ," " ) 
    fmt.Println( string( data) ) 
    SendMsg( w, data  )
}




func ( proc MsgProc ) send_pic_and_text(w http.ResponseWriter, ToUserName string , FromUserName string, art  []item  ) {


  resp := RespTextAndPic{}
  resp.ToUserName=ToUserName
  resp.FromUserName = FromUserName 
  resp.CreateTime = time.Now().Unix()
  resp.MsgType =MSG_NEWS 
  resp.ArticleCount=int64(len(art))
  resp.Articles = art
  data,_ := xml.Marshal( resp )  
  SendMsg( w, data )

}


func SendMsg( w http.ResponseWriter, data[] byte  ) {

    fmt.Fprintf(w, string(data))
    glog.Info("%s", string(data) ) 
} 


