package util

import  (
    "encoding/json"
    "fmt"
    // "github.com/qiniu/iconv"
)

type  WeatherData   struct {           
        
        Date                string      `json:"date"`
        DayPictureUrl       string      `json:"dayPictureUrl"`
        NightPictureUrl     string      `json:"nightPictureUrl"`
        Weather             string      `json:"weather"`
        Wind                string      `json:"wind"`
        Temperature         string      `json:"temperature"`
            
} 


type Result struct {

    CurrentCity      string                `json:"currentCity"`
    Whether_data     []WeatherData         `json:"weather_data"`
}

type Weather struct {

    Error       int64       `json:"error"`
    Status      string      `json:"status"`
    Date        string      `json:"date"`
    Results     []Result      `json:"results"`

}
const weather_url = "http://api.map.baidu.com/telematics/v3/weather?"


func SearchWeather ( location string )( err error,  weather Weather ) {

    url := weather_url +"location="+location +"&output=json&ak=A72e372de05e63c8740b2622d0ed8ab1"

    fmt.Println( url )
    
    res,_ := get_info( url )
    
    fmt.Println( string(res ) ) 
    ok := json.Unmarshal( [] byte ( res ), &weather )
    if ok != nil {
        fmt.Println( ok )
        return  
    }
    fmt.Println( weather ) 
    err = nil  
    return 

}

 
