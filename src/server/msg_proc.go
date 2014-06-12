package server

import (
    //"errors"
    "net/http"
    "time"
)

const (
    MSG_TEXT     = "text"
    MSG_IMAGE    = "image"
    MSG_LOCATION = "location"
    MSG_LINK     = "link"
    MSG_EVENT    = "event"
    MSG_VOICE    = "voice"
    MSG_VIDEO    = "video"
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
             return nil
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
   
    proc.send_text(w,  m.FromUserName(), m.ToUserName(), string( "欢迎加入万卷书屋。。") ); 
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

func ( proc MsgProc ) send_text( w http.ResponseWriter, ToUserName string , FromUserName string, Content string ) {

  resp := Replay{}
  resp.SetToUserName( ToUserName )
  resp.SetFromUserName( FromUserName )
  resp.SetCreateTime( time.Now().Unix())
  resp.SetMsgType( MSG_TEXT )
  resp.SetContent( Content )
  proc.send_msg( w, resp )
}

func ( proc MsgProc ) send_msg( w http.ResponseWriter, resp Replay ) {

    w.Write([]byte("<xml>"))
    body := MapToXmlString( resp )
    w.Write([]byte( body ))
    w.Write([]byte("</xml>"))   
} 
