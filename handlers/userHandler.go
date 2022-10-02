package handlers

import (
	Db "SvGorm/db"
	"SvGorm/models"
	"SvGorm/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func getUserById(r *http.Request) models.User2 {
	user := models.User2{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	Db.DbGorm.First(&user, id)
	return user
}

func GetUsersV2(w http.ResponseWriter, r *http.Request) {
	users := []models.User2{}
	DataUsers := []models.UserPublic{}
	data := models.UserPublic{}
	Db.DbGorm.Find(&users)

	for _, v := range users {
		data.Id = v.Id
		data.UserName = v.UserName
		data.Email = v.Email
		data.Created_at = v.Created_at.Format("2006-01-02")
		DataUsers = append(DataUsers, data)
	}
	if len(users) > 0 {
		utils.SendResponse(w, utils.RespOk{
			Message: "Usuarios encontrados",
			Data:    DataUsers,
		})
	} else {
		utils.SendResponseVoid(w)
	}
}

func GetUserV2(w http.ResponseWriter, r *http.Request) {
	user := getUserById(r)
	if user.Id > 0 {
		utils.SendResponse(w, utils.RespOk{
			Message: "Usuario encontrado",
			Data:    user,
		})
	} else {
		utils.SendResponseVoid(w)
	}
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	bodyJson := json.NewDecoder(r.Body)
	user := models.User2{}
	user.Created_at = time.Now()

	if err := bodyJson.Decode(&user); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	} else {
		Db.DbGorm.Save(&user)
		utils.CreatedResponse(w, utils.RespOk{
			Message: "Usuario modificado",
			Data:    user,
		})
	}
}

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		utils.BadResponse(w, utils.RespBad{
// 			Message:    "Los id no se pudo procesar",
// 			StatusCode: http.StatusBadRequest,
// 		})
// 	}
// 	models.DeleteUser(id)
// 	utils.SendResponse(w, utils.RespOk{
// 		Message: "Usuario modificado",
// 		Data:    models.NewUser("", "", ""),
// 	})
// }
