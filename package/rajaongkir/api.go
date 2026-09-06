package rajaongkir

import (
	"encoding/json"
	"finalissima_e_commerce_rest_api/package/clients"
	"finalissima_e_commerce_rest_api/package/constant"
	"finalissima_e_commerce_rest_api/package/utils"
	"net/http"
)

const BASE_URL = "https://rajaongkir.komerce.id/api/v1"

type Service interface {
	GetDeliveryFee(req GetFeeRequest) (float64, string, error)
}

type service struct {
	client clients.HTTPClient
}

func InitService() Service {
	return &service{
		client: clients.InitHTTPClient(BASE_URL, 10, utils.GetConfigurance(constant.RAJAONGKIR_API_KEY)),
	}
	// code ini untuk melakukan sambungan koneksi ke "Rajaongkir"
}

func (r *service) GetDeliveryFee(req GetFeeRequest) (float64, string, error) {
	var response FeeResponse

	payload := map[string]string{
		"origin":      req.Origin,
		"destination": req.Destination,
		//"weight":      strconv.Itoa(totalWeight),
		"courier": req.Courier,
	}

	res, err := r.client.SendFormEncoded(
		"calculate/district/domestic-cost",
		http.MethodPost,
		payload,
	)
	if err != nil {
		return 0, "", err
	}
	if err := json.Unmarshal([]byte(res), &response); err != nil {
		return 0, "", err
	}

	fee := float64(response.Data[0].Cost)
	etd := response.Data[0].Etd

	return fee, etd, nil
}
