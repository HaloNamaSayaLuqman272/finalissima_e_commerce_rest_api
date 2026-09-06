package dtos

// kita memerlukan "dtos" "Data Transfer Object" sebuah package berisi
// struct-struct khusus mengatur bentuk/ format data yg dikirm lewat API,
// memisahkan data internal eksternal "model database" dari data yg memang
// boleh/ perlu dilihat client, dan juga menstandarkan format response di
// seluruh endpoint aplikasi
type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	// "omitempty" adalah jika data kosong, field ini tidak akan muncul pada
	// hasil JSON
}
