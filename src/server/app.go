package server

import (
    "net/http"
    "io/ioutil" 
    "x2j"
    "github.com/golang/glog"
)


var app_ *App

func NewApp() *App {
    if app_ == nil {
        app_ = new ( App )
    }
    return app_;
}


type App struct {

  msg_proc MsgProc

}

func (app *App )  Start() {
    
    glog.Info("Server begin listen 80 ........");
    http.HandleFunc("/monitor", handler )
    http.ListenAndServe(":80", nil)

}

func (app *App)  Stop() {
    
    glog.Flush()
}



func /*(app *App)*/  handler(w http.ResponseWriter, req *http.Request) {
    
     if req.Method == "GET" { 
	    
            glog.Info(" receive GET " )
            //参数验证 
	        var signature = req.FormValue("signature");
	        var timestamp = req.FormValue("timestamp");
	        var nonce = req.FormValue("nonce");
	      	var echostr = req.FormValue("echostr")
	     	
	     	if CheckSignature( signature, timestamp, nonce) {

	     		glog.Info( "signature=%s, timestamp=%s, nonce=%s,echostr=%s" , 
	       				 							signature, timestamp,nonce,echostr  );
	     		byteArray := []byte(echostr)
	     		w.Write( byteArray )
	     	} 
         
    } else if req.Method == "POST" {
       
       glog.Info(" receive post" ) 
       app_.parse_req( w, req )
    }  

}
 
func (app *App )  parse_req( w http.ResponseWriter, req *http.Request ) {
    
    defer req.Body.Close()  
  
    body, err := ioutil.ReadAll( req.Body )  
    if err != nil {  
        glog.Error( "read msg error: %s",err )
        return  
    }  
  
    glog.Info( string(body) );  

    root, err := x2j.DocToMap(string(body))
    if err != nil {
        glog.Error(" bad xml: %s", body )
        return
    }

    msg := Msg(root["xml"].(map[string]interface{}))
   
    err = app.msg_proc.Proc( w ,&msg )
    if err != nil {
        glog.Error( " msg proc fail ... " );
    }

    return  
}

