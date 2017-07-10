package http


import(
        "net/http"
        "fmt"
        "encoding/json"
        "log"
        "contactform/populate"
        "contactform/config"
)


type FormData struct {
        FullName string
        CurrentAddr string
        UKAddr string
        Providers []struct {
                ProvidersName string
                ProvidersPlanNo string
        }
        DateOfBirth string
        NIN string
}


func ReqHandler(rw http.ResponseWriter, request *http.Request){
        decoder := json.NewDecoder(request.Body)
        var json_res FormData
        err := decoder.Decode(&json_res)
        if err != nil {
                fmt.Println("\n Crashing here, err != nil \n")
                log.Fatalln(err)
        }
        fmt.Println(json_res.Member)
        populate.FillTempl(json_res.Member, json_res.CurrentAddr, json_res.UKAddr, json_res.Email)
}


func Serve(){
        tlscfg := config.GetTlsConfig("config/")
        log.Print("Creating server with cert file: " + tlscfg.GetCert() + "and key file: " + tlscfg.GetKey())
        http.HandleFunc("/services/FormData", ReqHandler)
                log.Print("listening on port: " + tlscfg.GetPort())
        err := http.ListenAndServeTLS(tlscfg.GetPort(), tlscfg.GetCert(), tlscfg.GetKey(), nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}


