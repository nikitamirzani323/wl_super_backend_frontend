package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/wl_super_backend_frontend/entities"
)

type responsemaster struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Listcurr interface{} `json:"listcurr"`
	Listbank interface{} `json:"listbank"`
	Record   interface{} `json:"record"`
}

func Masterhome(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(entities.Home)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsemaster{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/master")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsemaster)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":   result.Status,
			"message":  result.Message,
			"record":   result.Record,
			"listcurr": result.Listcurr,
			"listbank": result.Listbank,
			"time":     time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func MasterSave(c *fiber.Ctx) error {
	type payload_mastersave struct {
		Page              string `json:"page"`
		Sdata             string `json:"sdata" `
		Master_id         string `json:"master_id" `
		Master_idcurr     string `json:"master_idcurr" `
		Master_name       string `json:"master_name" `
		Master_owner      string `json:"master_owner" `
		Master_phone1     string `json:"master_phone1" `
		Master_phone2     string `json:"master_phone2" `
		Master_email      string `json:"master_email" `
		Master_note       string `json:"master_note" `
		Master_bank_id    string `json:"master_bank_id" `
		Master_bank_norek string `json:"master_bank_norek" `
		Master_bank_name  string `json:"master_bank_name" `
		Master_status     string `json:"master_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_mastersave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":   hostname,
			"page":              client.Page,
			"sdata":             client.Sdata,
			"master_id":         client.Master_id,
			"master_idcurr":     client.Master_idcurr,
			"master_name":       client.Master_name,
			"master_owner":      client.Master_owner,
			"master_phone1":     client.Master_phone1,
			"master_phone2":     client.Master_phone2,
			"master_email":      client.Master_email,
			"master_note":       client.Master_note,
			"master_bank_id":    client.Master_bank_id,
			"master_bank_norek": client.Master_bank_norek,
			"master_bank_name":  client.Master_bank_name,
			"master_status":     client.Master_status,
		}).
		Post(PATH + "api/mastersave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func MasteradminSave(c *fiber.Ctx) error {
	type payload_masteradminsave struct {
		Page                 string `json:"page"`
		Sdata                string `json:"sdata" `
		Masteradmin_id       int    `json:"masteradmin_id" `
		Masteradmin_idmaster string `json:"masteradmin_idmaster" `
		Masteradmin_tipe     string `json:"masteradmin_tipe" `
		Masteradmin_username string `json:"masteradmin_username" `
		Masteradmin_password string `json:"masteradmin_password" `
		Masteradmin_name     string `json:"masteradmin_name" `
		Masteradmin_phone1   string `json:"masteradmin_phone1" `
		Masteradmin_phone2   string `json:"masteradmin_phone2" `
		Masteradmin_status   string `json:"masteradmin_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_masteradminsave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":      hostname,
			"page":                 client.Page,
			"sdata":                client.Sdata,
			"masteradmin_id":       client.Masteradmin_id,
			"masteradmin_idmaster": client.Masteradmin_idmaster,
			"masteradmin_tipe":     client.Masteradmin_tipe,
			"masteradmin_username": client.Masteradmin_username,
			"masteradmin_password": client.Masteradmin_password,
			"masteradmin_name":     client.Masteradmin_name,
			"masteradmin_phone1":   client.Masteradmin_phone1,
			"masteradmin_phone2":   client.Masteradmin_phone2,
			"masteradmin_status":   client.Masteradmin_status,
		}).
		Post(PATH + "api/masteradminsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
