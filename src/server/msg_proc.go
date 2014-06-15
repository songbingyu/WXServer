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
   
   // proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "欢迎加入万卷书屋。。") ); 
    
    art := [] item{}
    art1 := item{}
    art1.Title = "[golang] xml解析"
    art1.Description = "golang解析xml真是好用,特别是struct属性的tag让程序简单了许多,其他变成语言需要特殊类型的在golang里直接使用tag舒服"
    art1.PicUrl = "http://a.hiphotos.baidu.com/image/pic/item/e7cd7b899e510fb37ff3361fdb33c895d0430c4b.jpg"
    art1.Url    = "http://www.dotcoo.com/golang-xml-reader"
    
    art2 := item{}
    art2.Title = "[golang] long-polling 长轮训"
    art2.Description = "ＴＣＰ　轮训。。。"
    art2.PicUrl = "http://att.newsmth.net/nForum/att/Nanjing/230042/3739/large"
    art2.Url    = "http://www.dotcoo.com/golang-long-polling"
    
    art3 := item{}
    art3.Title = "师兄可厉害了"
    art3.Description = "要期末考试了让师兄帮我说说功课"
    art3.PicUrl = "http://att.newsmth.net/nForum/att/SCU/58151/1348/large"
    art3.Url    = "http://www.newsmth.net/nForum/#!article/SCU/58151"
   art = append( art ,art1 )   
   art = append( art ,art2 )   
   art = append( art ,art3 )   
   proc.send_pic_and_text( w,m.FromUserName(), m.ToUserName(),art )
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


