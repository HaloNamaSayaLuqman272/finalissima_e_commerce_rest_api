package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient struct {
	// ini adalah konfigurasi struct untuk melakukan request ke API pihak ketiga
	// atau eksternal
	BaseURL string
	// "BaseURL" adalah alamat dasar API tujuan, seperti https://midtrans.com/
	Timeout int
	// "Timeout" batas waktu maksimal sebelum suatu request dianggap error/ gagal
	APIKey string
	// "APIKey" token/ kunci otentifikasi yg dibutuhkan untuk mengakses API eksternal
	Endpoint string
	// "Endpoint" jalur spesifik dari API yg akan dipanggil, seperti BaseURL/v1/charge
	Method string
	// "Method" metode HTTP yg digunakan, seperti "GET", "PUT", dst
}

func InitHTTPClient(baseUrl string, timeout int, apiKey string) HTTPClient {
	// kita mendeklarasikan sebuah fungsi biasa dan memberi nama "InitHTTPClient"
	// memiliki parameter input "baseUrl", "timeout", dan "apiKey" dan
	// mengembalikan tiga field pada struct "HTTPClient"
	return HTTPClient{
		BaseURL: baseUrl,
		Timeout: timeout,
		APIKey:  apiKey,
	}
}

func (h *HTTPClient) SendJSON(endpoint, method string, payload map[string]any) (string, error) {
	// kita mendeklarasikan sebuah fungsi dan memberi nama "SendJSON"
	// memiliki variabel penerima "h" bertipe data pointer ke "HTTPClient"
	// dan memiliki parameter input "endpoint", "method", dan "payload"
	// parameter input "payload" bertipe data "map", Key-nya betipe "string",
	// Value-nya bertipe "any"
	// "payload" ini kedepannya akan dilakukan proses "Unmarshall"
	// fungsi "SendJSON" ini akan mengembalikan bentuk "string" dan "error"
	// apabila ada masalah
	var requestBody map[string]any
	// kita membuat variabel dengan nama "requestBody" dan bentuk datanya
	// akan berupa "map[string]any"
	var requestPayload *bytes.Buffer
	// kita juga membuat variabel yg bernama "requestPayload" dan bentuk datanya
	// berupa pointer ke "bytes.Buffer"
	// "bytes.Buffer" adalah salah satu package bawaan Go yg memiliki fungsi
	// sebagai buffer/ wadah sementara untuk menampung data byte, yg bisa
	// ditulis satu per satu, kemudian dibaca ulang sebagai satu kesatuan

	if payload != nil {
		requestBody = payload
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return "", err
		}
		requestPayload = bytes.NewBuffer(jsonData)
		// hasil dari "jsonData" akan berbentuk "bytes.NewBuffer" dan akan
		// dibungkus memakai variabel "requestPayload"
	}

	req, err := http.NewRequest(method, h.BaseURL+endpoint, requestPayload)
	// ini adalah bentuk HTTP yg direncanakan
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	// "Content-Type" adalah header standar HTTP yg berfungsi menjelaskan
	// tipe/ format data yg dikirim
	// "application/json" adalah nilai dari header dan mengatakan bahwa
	// format datanya adalah JSON
	// code ini bertujuan untuk memberi tahu server tujuan bahwa data yg
	// dikirim di body request berformat JSON, supaya server bisa memproses
	// mem-parsing data itu dengan benar sesuai formatnya, seandainya
	// code ini tidak ada, ada resiko server salah menginterpretasi atau menolak
	// data yg dikirim
	req.Header.Set("Authorization", "Key: "+h.APIKey)
	// "Authorization" adalah nama header standar HTTP yg berfungsi membawa
	// kredensial/ token otentifikasi, supaya server tujuan tahu bahwa request ini
	// datang dari pihak yg terverifikasi
	// "Key: "+h.APIKey adalah penggabungan "Key" sebagai teks literal
	// "+" operator penggabungan
	// "h.APIKey" mengakses field "APIKey" dari struct "h" dan "h" adalah
	// variabel penerima dari pointer ke struct "HTTPClient"
	// contoh hasilnya "Authorization: Key abc1234xyz"

	client := &http.Client{Timeout: time.Duration(h.Timeout) * time.Second}
	// code ini menyiapkan alat pengirim "http.Client" yg dipakai buat mengirim
	// request ke API eksternal, sekaligus mengatur batas waktu maksimal client
	// menunggu balasan
	resp, err := client.Do(req)
	// code ini adalah membatasi waktu proses "req"
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	// "defer" adalah kata kunci khusus Go yg memiliki arti jalankan perintah
	// ini nanti, tepat sebelum fungsi ini selesai "return" apa pun jalur keluarnya
	// code ini untuk menjamin koneksi "HTTP response" selalu ditutup dengan
	// benar setelah selesai dipakai, secara otomatis pada akhir fungsi

	bodyBytes, err := io.ReadAll(resp.Body)
	// "io.ReadAll" dipakai untuk mengambil seluruh isi data dari sebuah
	// stream, seperti "resp.Body" sekaligus dan menjadikan nya menjadi
	// satu blok data utuh []byte yg siap diproses atau dipakai
	if err != nil {
		return "", err
	}

	isFailed := resp.StatusCode < 200 || resp.StatusCode >= 300
	// code ini adalah kondisi untuk menggambarkan jika terjadi masalah
	if isFailed {
		return string(bodyBytes), fmt.Errorf("request failed: %s", resp.Status)
	}

	return string(bodyBytes), nil
}

func (h *HTTPClient) SendFormEncoded(endpoint, method string, payload map[string]string) (string, error) {
	var requestBody map[string]string
	var requestPayload *bytes.Buffer

	if payload != nil {
		requestBody = payload
		values := url.Values{}
		// code ini adalah wadah kosong yg dipakai untuk menyusun parameter
		// query string URL secara terstruktur dan aman, akan ter-encode dengan benar

		for k, v := range requestBody {
			values.Set(k, v)
		}

		requestPayload = bytes.NewBufferString(values.Encode())
		// code ini bermaksud mengambil data yg tersusun di "values", mengubahnya
		// menjadi string query yg sudah aman/ ter-encode lewat ".Encode",
		// lalu bungkus string itu ke dalam "bytes.Buffer"
	}

	req, err := http.NewRequest(method, h.BaseURL+endpoint, requestPayload)
	// code ini bermaksud sebagai langkah membangun objek request HTTP yg
	// lengkap, namun bagian ini request belum terkirim
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Key", h.APIKey)

	client := &http.Client{Timeout: time.Duration(h.Timeout) * time.Second}
	resp, err := client.Do(req)
	// code ini adalah momen eksekusi nyata dari seluruh proses HTTP request
	// yg sudah disiapkan sebelumnya
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	isFailed := resp.StatusCode < 200 || resp.StatusCode >= 300
	// code ini adalah kondisi untuk menggambarkan jika terjadi masalah
	if isFailed {
		return string(bodyBytes), fmt.Errorf("request failed: %s", resp.Status)
	}

	return string(bodyBytes), nil
}
