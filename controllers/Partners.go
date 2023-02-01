package controllers

import (
    "fmt"
	"net/http"
	"example/task/model"
    "example/task/database"
	"github.com/gin-gonic/gin"
)

func AddPartner(c *gin.Context) {
   var partner model.Partner

   if err := c.Bind(&partner); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status_code": 500,
            "api_version": "v1",
            "endpoint": "/AddPartner",
            "status": "Server Error !",
            "msg":    "Server Error !",
            "data":   nil,
        })
        c.Abort()
        return
    } 

    check := model.CheckPartner(partner.Wallet_address1, partner.Wallet_address2)

    if check > 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "status_code": 409,
            "api_version": "v1",
            "endpoint": "/AddPartner",
            "status": "Failure!",
            "msg":    "Conflict partner",
            "data": check, 
         })
    } else {
        database.Database.Create(&partner)
        c.JSON(http.StatusBadRequest, gin.H{
            "status_code": 200,
            "api_version": "v1",
            "endpoint": "/AddPartner",
            "status": "Success!",
            "msg":    "Success",
         })
    }
}

func GetPartner(c *gin.Context) {
    var partner model.Partner
    var user model.User

    if err := c.Bind(&partner); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status_code": 500,
            "api_version": "v1",
            "endpoint": "/AddPartner",
            "status": "Server Error !",
            "msg":    "Server Error !",
            "data":   nil,
        })
        c.Abort()
        return
    } 

    fmt.Println("**********", partner.Wallet_address1)

    // var cnt int

    // database.Database.Raw("SELECT U.* FROM partners PP LEFT JOIN users U ON U.wallet_address = PP.wallet_address2 WHERE PP.wallet_address1 = ?", partner.Wallet_address1).Scan(&cnt)

    err := database.Database.Raw("SELECT U.* FROM partners PP LEFT JOIN users U ON U.wallet_address = PP.wallet_address2 WHERE PP.wallet_address1 = ?", partner.Wallet_address1).Find(&user).Error

    // if cnt == 0 {
    //      c.JSON(http.StatusInternalServerError, gin.H{
    //         "status_code": 204,
    //         "api_version": "v1",
    //         "endpoint": "/GetPartner",
    //         "status": "No Content !",
    //         "msg":    "No Content !",
    //     })
    //     return
    // } 
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status_code": 204,
            "api_version": "v1",
            "endpoint": "/GetPartner",
            "status": "No Content !",
            "msg":    "No Content !",
        })
        return
    } 


    c.JSON(http.StatusInternalServerError, gin.H{
        "status_code": 200,
        "api_version": "v1",
        "endpoint": "/GetPartner",
        "status": "Success !",
        "msg":    "Success !",
        "data": user,
    })


}
