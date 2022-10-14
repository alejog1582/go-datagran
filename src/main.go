package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const xthreads = 5

var item = 1
var finish = false

var db = OpenDbConnection()

func main() {
	//Notificacion mattermost Inicio Proceso
	data_inicio := Payload{
		Text: "Inicio Consumo de API Datagran",
	}
	payloadBytes_inicio, err := json.Marshal(data_inicio)
	if err != nil {
		// handle err
	}
	body_mattermost_inicio := bytes.NewReader(payloadBytes_inicio)

	req_body_mattermost_inicio, err := http.NewRequest("POST", os.Getenv("URL_MATTERMOST"), body_mattermost_inicio)
	if err != nil {
		// handle err
	}
	req_body_mattermost_inicio.Header.Set("Content-Type", "application/json")

	resp_mattermost_inicio, err := http.DefaultClient.Do(req_body_mattermost_inicio)
	if err != nil {
		// handle err
	}
	defer resp_mattermost_inicio.Body.Close()

	//Se borra la taba de base de datos
	//db.Where("rp_fuente = ?", "freq").Delete(&Datagran_modelo_1{})
	//db.Where("rp_fuente = ?", "rp").Delete(&Datagran_modelo_1{})

	var ch = make(chan int) // This number 50 can be anything as long as it's larger than xthreads
	var wg sync.WaitGroup

	// This starts xthreads number of goroutines that wait for something to do
	wg.Add(xthreads)
	for i := 0; i < xthreads; i++ {
		go func() {
			for {
				iterator, ok := <-ch
				if !ok { // if there is nothing to do and the channel has been closed then end the goroutine
					return
				}
				response := getData(iterator)
				if len(response.Output.Rows) == 0 {
					return
				}

				if len(response.Output.Rows) < 100 {
					finish = true
					var total_row_db int
					db.Raw("SELECT COUNT(*) FROM datagran_modelo_1").Scan(&total_row_db)
					mensaje_total_row := fmt.Sprintf("El consumo de la API ha finalizado con %d registros ingresados", total_row_db)

					//Notificacion mattermost Inicio Proceso
					data_fin := Payload{
						Text: mensaje_total_row,
					}
					payloadBytes_fin, err := json.Marshal(data_fin)
					if err != nil {
						// handle err
					}
					body_mattermost_fin := bytes.NewReader(payloadBytes_fin)

					req_body_mattermost_fin, err := http.NewRequest("POST", os.Getenv("URL_MATTERMOST"), body_mattermost_fin)
					if err != nil {
						// handle err
					}
					req_body_mattermost_fin.Header.Set("Content-Type", "application/json")

					resp_mattermost_fin, err := http.DefaultClient.Do(req_body_mattermost_fin)
					if err != nil {
						// handle err
					}
					defer resp_mattermost_fin.Body.Close()
					return
				}
			}
		}()
	}

	// Now the jobs can be added to the channel, which is used as a queue
	for {
		if finish {
			break
		}

		ch <- item
		item += 100
	}

	for i := 0; i < xthreads; i++ {
		wg.Done()
	}

	close(ch) // This tells the goroutines there's nothing else to do
	wg.Wait() // Wait for the threads to finish
}

func getData(iterator int) Response {
	url := "https://api.v2.datagran.io/v2/rest_api_export_operator/62014f420dec1817c67869d9?table=operator_6201760fbf8bc8a1435787d0__sql_output&project_id=61f8095cff5ddfeaeaf2bd42&offset=" + strconv.Itoa(iterator)
	fmt.Println(url)
	t := time.Now()
	fmt.Println("Fecha de llamado a la api: ", t)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return Response{}
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", "91164I3bU9dff57OS71Va3ff-fahf504pB6571d57zE8d3afc8c195bN8e510f807fl74as64GX35kU9paCU80z0c27XbdeP5509D0H2O5g")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Response{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON: " + err.Error())
	}
	for i := 0; i < len(result.Output.Rows); i++ {
		Tienda_id_str := fmt.Sprintf("%v", result.Output.Rows[i][0])
		Codigo_producto_srt := fmt.Sprintf("%v", result.Output.Rows[i][1])
		Rating_prediction_srt := fmt.Sprintf("%v", result.Output.Rows[i][2])
		Rp_fuente_srt := fmt.Sprintf("%v", result.Output.Rows[i][3])
		row_datagran := Datagran_modelo_1{Tienda_id_str, Codigo_producto_srt, Rating_prediction_srt, Rp_fuente_srt}
		fmt.Println(row_datagran)
		save_bd := db.Create(&row_datagran) // pass pointer of data to Create */
		fmt.Println(save_bd)
	}
	return result
}

func OpenDbConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error loading db")
	}

	/*sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Error loading db", err)
	}*/

	/*res, err := sqlDB.Exec("SET GLOBAL sql_mode = '';")
	if err != nil {
		log.Fatal("Error loading db", err)
		log.Fatal(res)
		os.Exit(0)
	}*/

	return db
}

type Response struct {
	Output Output
}

type Output struct {
	Cols []Cols
	Rows [][]interface{}
}

type Cols struct {
	Type string
	Name string
}

type Datagran_modelo_1 struct {
	Tienda_id         string
	Codigo_producto   string
	Rating_prediction string
	Rp_fuente         string
}

type Payload struct {
	Text string `json:"text"`
}
