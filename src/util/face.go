package util

import (
    "net/http"
    "io/ioutil"
    "github.com/golang/glog"
    "encoding/json"
    "fmt"
)


const apiurl="https://apicn.faceplusplus.com"
const apikey="49573589b027c35e55bcb4d2a96d798f"
const apisecret="omZD3wEnwy8nX1KdAJAf10LnVI75stfT"


type FaceInfo  struct{
     Face []struct{
         Attribute struct{
             Age struct{
                 Range float64
                 Value float64
             }
             Gender struct{
                 Confidence float64
                 Value string
             }
             Glass struct {
                Confidence float64
                Value string
             }

             Race struct{
                 Confidence float64
                 Value string
             }
            Smiling struct {
                Value  float64
            }
            
         }
         Face_id string
         Position struct{
             Center struct{
                 X float64
                 Y float64
             }
             Eye_left struct{
                 X float64
                 Y float64
             }
             Eye_right struct{
                 X float64
                 Y float64
             }
             Height float64
             Mouth_left struct{
                 X float64
                 Y float64
             }
             Mouth_right struct{
                 X float64
                 Y float64
             }
             Nose struct{
                 X float64
                 Y float64
             }
             Width float64
         }
         Tag string
     }
     Img_height int
     Img_id string
     Img_width int
     Session_id string
     url string
 }


func get_info ( url string ) ( b []byte ,err  error  ) {
   
   r, e  := http.Get( url ) 
   if   e!= nil {
        err = e
        return 
   }
  
   fmt.Println(" GetResult") 
   data, e:= ioutil.ReadAll( r.Body )
   if e != nil {
        err = e
        return 
   }
    
   fmt.Println("Get data ")
   r.Body.Close()
   
   return data, nil    

}

func DetectionDetect( picurl string ) FaceInfo  {

     url := apiurl+"/v2/detection/detect?url="+picurl+"&api_secret="+apisecret+"&api_key="+apikey+"&attribute=glass,gender,age,race,smiling"

     glog.Info( url )
     res,_ := get_info( url )
     var f FaceInfo
     json.Unmarshal( res ,&f )
     return f
} 

