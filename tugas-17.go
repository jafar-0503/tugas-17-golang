package main

import "fmt"
import "encoding/json"
import "net/http"
import "database/sql"
import _"mysql-master"

type materi_golang struct{
  ID int
  Nama string
  Jurusan string
  Alamat string
}
//koneksi ke database
func koneksi()(*sql.DB, error){
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/materi_golang")
  if err != nil{
    return nil, err
  }
  return db, nil
}

var data =[]materi_golang{}

//ambil data dari function main
func main(){
    ambil_data()
    http.HandleFunc("/mhs", ambil_mhs)
    http.HandleFunc("/cari_mhs", cari_mhs)

    fmt.Println("Menjalankan Web Server localhost:8080")
    http.ListenAndServe(":8080",nil)
}
//get data from function ambil_mhs
func ambil_mhs(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST"{
    var result, err = json.Marshal(data)

    if err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Write(result)
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}
//get data from function cari_mhs
func cari_mhs(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST"{
    var nama = r.FormValue("Nama")
    var result []byte
    var err error

    for _, each := range data {
      if each.Nama == nama{
        result, err = json.Marshal(each)

        if err != nil{
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        w.Write(result)
        return
      }
    }
    http.Error(w, "Data Mahasiswa Tidak Tersedia", http.StatusBadRequest)
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}
//get data from function ambil_data
func ambil_data(){
  db, err := koneksi()

  if err != nil{
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  rows, err := db.Query("select * from tbl_mahasiswa")
  if err != nil{
    fmt.Println(err.Error())
    return
}
defer rows.Close()

for rows.Next(){
  var each = materi_golang{}
  var err = rows.Scan(&each.ID, &each.Nama, &each.Jurusan, &each.Alamat)

  if err != nil{
    fmt.Println(err.Error())
    return
  }
  data = append(data, each)
}
if err = rows.Err(); err != nil{
  fmt.Println(err.Error())
  return
}
}
