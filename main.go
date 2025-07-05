package main

import (
	//"fmt"
	//"log"
	"net/http"
	"os"
	// "log"
)

func main(){
	// fileServer := http.FileServer(http.Dir("./statics"))
	// mux.Handle("/", fileServer)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/experience", experience)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		_, pattern := mux.Handler(r)
		if pattern == "" {
			handle404(w)
			return
		}
		mux.ServeHTTP(w, r)
	})

	http.ListenAndServe(":9090", handler)

}
func home (w http.ResponseWriter , r *http.Request){
   if r.Method!="GET"{
      handle405(w)
	return
   }
   data ,err:= os.ReadFile("statics/index.html")
   if err!=nil{
	handle505(w)
	 return
   }
   w.Write(data)

}
func about (w http.ResponseWriter , r *http.Request){
	if r.Method!="GET"{
	 handle405(w)
		return
	}
	data ,err := os.ReadFile("statics/about.html")
	if err!=nil{
		handle505(w)
    return
	}
	w.Write(data)
}
func contact (w http.ResponseWriter , r *http.Request){
	if r.Method!="GET"{
		handle405(w)
		return
	}
	data ,err:= os.ReadFile("statics/contact.html")
	if err!=nil{
		handle505(w)
		return
	}
	w.Write(data)
}
func experience (w http.ResponseWriter , r *http.Request){
	if r.Method!="GET"{
		handle405(w)
		return
	}
	data ,err:= os.ReadFile("statics/experience.html")
	if err!=nil{
		handle505(w)
		return
	}
	w.Write(data)
}
func handle404(w http.ResponseWriter){
	data, err :=os.ReadFile("statics/404.html")
	if err!=nil {
	http.Error(w,"Not Found " ,http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNotFound)
		w.Write(data)
}
func handle405(w http.ResponseWriter){
	data, err :=os.ReadFile("statics/405.html")
	if err!=nil {
	http.Error(w,"Not Found " ,http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNotFound)
		w.Write(data)
}
func handle505(w http.ResponseWriter){
	// استخدمت هنا ال. ـ. عشان انا مش عايزه القيمه االتانيه مش مهمه وعندي الفانكشن بترج قيميتين فانا مش محتاجه قيمه فيهم 

	data, _ :=os.ReadFile("statics/505.html")

	w.WriteHeader(http.StatusNotFound)
		w.Write(data)
}