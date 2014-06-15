//main
package main
 
import (
    "flag"
    "server"
)



func main() {
    
    flag.Parse()
    app := server.NewApp( )  
    
    app.Start()
    app.Stop();
}


