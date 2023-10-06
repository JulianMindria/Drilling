package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID      int
	Name    string
	Picture string
	Price   int
}

func GetallProduct(c *gin.Context) {
	products := []Product{
		{
			ID:      1,
			Name:    "Beras",
			Picture: "https://ipanganan.com/images/banner1a.png",
			Price:   10000,
		},
		{
			ID:      2,
			Name:    "Teh",
			Picture: "https://cf.shopee.co.id/file/2c559cae36b8f0b9008a73cfe861927d",
			Price:   9000,
		},
		{
			ID:      3,
			Name:    "Gula",
			Picture: "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/11/6/556a6450-25f6-41d2-8cc3-07f35d4989e7.jpg",
			Price:   14000,
		},
		{
			ID:      4,
			Name:    "Kopi",
			Picture: "https://assets.promediateknologi.com/crop/0x0:0x0/750x500/photo/2022/05/11/2929532943.png",
			Price:   8000,
		},
		{
			ID:      5,
			Name:    "Telur",
			Picture: "https://cdn-2.tstatic.net/batam/foto/bank/images/telur-ayam_20180416_112847.jpg",
			Price:   12000,
		},
		{
			ID:      6,
			Name:    "Jeruk",
			Picture: "https://png.pngtree.com/png-clipart/20201208/original/pngtree-stack-cut-oranges-png-image_5529928.jpg",
			Price:   18000,
		},
		{
			ID:      7,
			Name:    "Apel",
			Picture: "https://png.pngtree.com/png-clipart/20190614/original/pngtree-photo-realistic-of-red-apple-full-editable-isolated-on-white-png-image_3718196.jpg",
			Price:   28000,
		},
	}

	c.JSON(http.StatusOK, products)
}
