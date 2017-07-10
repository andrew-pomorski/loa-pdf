package populate

import(
        "bytes"
        "io/ioutil"
        "log"
        "fmt"
        "os/exec"
        "html/template"
        "contactform/mail"
        "time"
)



var TemplateFile, _ =  template.ParseFiles("../templates/template.html")

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



func FillTempl(FullName string, CurrentAddr string, UKAddr string, Providers []struct, DateOfBirth string, NIN string:w){
        fmt.Printf("FillTempl")
        buff := bytes.NewBufferString("")
        sent_at := time.Now()
        pdf_path := "../templates/loa.pdf"
        // Compile and allocate in buffer
        err := TemplateFile.Execute(buff, FormData{
              FullName: FullName,
			  UKAddr: UKAddr,
			  Providers: {
				Providers
			  },
			  DateOfBirth: DateOfBirth,
			  NIN: NIN
        })
        if err != nil {
                log.Fatalln(err)
        }
        err = ioutil.WriteFile("loa_letter.html", buff.Bytes(), 0666)
        if err != nil {
                log.Fatalln(err)
        }
        //convert compiled file to pdf
        err = exec.Command("wkhtmltopdf", "loa_letter.html", pdf_path).Run()
        mail.Send(sent_at, pdf_path, Email);
        if err == nil {
                fmt.Printf("[+ TEMPLATE] Save successful")
        } else {
                fmt.Printf("[- TEMPLATE] Error generating PDF %s", err)
        }
}
