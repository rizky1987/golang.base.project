package services

import (
	"example/config/env"
	"example/http/helpers"
	"example/http/interfaces"
	"fmt"

	sapCore "example/utils/sap_core"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartHandler struct {
	DB             *gorm.DB
	Helper         helpers.HTTPHelper
	Config         env.Config
	CartRepository interfaces.CartInterface
}

func (_h *CartHandler) GetAllCart(ctx echo.Context) error {
	var (
		err error
	)

	// begin connect to SAP

	// isi parameter yang diperlukan untuk request ke SAP
	reqParam := sapCore.GetDataProductByPrincipalRequestToSAP{
		StartNo:   0,
		EndNo:     1000,
		Principal: "0001",
	}

	records, errorMessage, errorFileLocation, errorFileLine := sapCore.GetProductByPrincipal(reqParam)
	if errorMessage != "" {

		// ini hanya contoh saja kedapnnya sesuaikan dengan kebutuhan
		return _h.Helper.SendBadRequest(ctx, err.Error(), errorFileLocation, errorFileLine)

	}

	// jika sudah dapat "recods" kita bisa mengolah datanya sesuia dengan kebutuhan
	fmt.Println(records)

	// end connect to SAP
	return _h.Helper.SendAllDataSuccess(ctx, "Cart", records)
}
