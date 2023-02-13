package testing

import (
	"encoding/json"
	"go_STAN/internal/db"
)

func GetTestOrders() ([]byte, []byte) {
	order1 := db.Order{
		OrderUID:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: db.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: db.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDT:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []db.Item{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				RID:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       "2021-11-26T06:22:19Z",
		OofShard:          "1",
	}

	order2 := db.Order{
		OrderUID:    "g900jan4b2b55b1test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: db.Delivery{
			Name:    "Tester Testingov",
			Phone:   "+9167654321",
			Zip:     "1179809",
			City:    "Bruh City",
			Address: "Ploshad Mira 0",
			Region:  "Sussy amogus",
			Email:   "tester@gmail.com",
		},
		Payment: db.Payment{
			Transaction:  "g900jan4b2b55b1test",
			RequestID:    "",
			Currency:     "RUB",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDT:    1637907727,
			Bank:         "sber",
			DeliveryCost: 25000,
			GoodsTotal:   322,
			CustomFee:    2,
		},
		Items: []db.Item{
			{
				ChrtID:      8434235,
				TrackNumber: "WBILMTESTONETRACK",
				Price:       555,
				RID:         "ab4300817b461ae0btest",
				Name:        "SusItem1",
				Sale:        30,
				Size:        "0",
				TotalPrice:  322,
				NmID:        2186252,
				Brand:       "Sabo Vivienne",
				Status:      202,
			},
			{
				ChrtID:      5531535,
				TrackNumber: "WBILMTESTTWOTRACK",
				Price:       777,
				RID:         "ab4322800b992ae0btest",
				Name:        "SusItem2",
				Sale:        40,
				Size:        "0",
				TotalPrice:  228,
				NmID:        1337252,
				Brand:       "Sabos Viviennetire",
				Status:      202,
			},
		},
		Locale:            "ru",
		InternalSignature: "",
		CustomerID:        "tester",
		DeliveryService:   "meest",
		Shardkey:          "8",
		SmID:              99,
		DateCreated:       "2022-01-26T06:22:19Z",
		OofShard:          "1",
	}

	jsonOrder1, _ := json.Marshal(order1)
	jsonOrder2, _ := json.Marshal(order2)

	return jsonOrder1, jsonOrder2
}
