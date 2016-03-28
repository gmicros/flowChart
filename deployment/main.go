package main

import (
    "io/ioutil"
    "strings"
    "net/http"
    "html/template"
)

type page struct {
    Contents string
    YesLink string
    NoLink string
}

func handler(w http.ResponseWriter, r *http.Request) {        
    var p page;
    var t *template.Template;
    if r.URL.Path[1:] == "" {
        p = pages["main"];
    } else { 
        p = pages[r.URL.Path[1:]];
    }
    if(p.YesLink == "") {
        t, _ = template.ParseFiles("tempNoBtn.html")
    } else {
        t, _ = template.ParseFiles("temp.html")
    }
    t.Execute(w,p);    
}

var pages map[string]page;

func main() {
    
    pages = getPages("pages.csv")
    
    http.HandleFunc("/", handler)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("http/css"))))
    //http.ListenAndServe(":8080", nil)    

    port := os.Getenv("PORT")
    if port == "" {
	port = "8080"
    }
	http.ListenAndServe(":"+port, nil)
}

}

func getPages(filename string) map[string]page{
    pages := make (map[string]page);
    dat, err := ioutil.ReadFile(filename);    
    check(err);
    file := string(dat);
    lines := strings.Split(file, "\n");
    for i := 0; i < len(lines); i++ {
        words := strings.Split(lines[i], ",");
        if len(words) == 2 {
            pages[strings.TrimSpace(words[0])] = page{strings.TrimSpace(words[1]), "", ""};    
        }
        if len(words) == 4 {
            pages[words[0]] = page{strings.TrimSpace(words[1]), strings.TrimSpace(words[2]), strings.TrimSpace(words[3])};
        }        
    }
    return pages;
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
