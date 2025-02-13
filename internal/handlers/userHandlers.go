package handlers

import (
	"PetProject/internal/userService"
	"PetProject/internal/web/users"
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id

	err := h.Service.DeleteUserByID(userID)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil

}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Password: &usr.Password,
			Email:    &usr.Email,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {

	userRequest := request.Body

	userToCreate := userService.User{
		Password: *userRequest.Password,
		Email:    *userRequest.Email,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Password: &createdUser.Password,
		Email:    &createdUser.Email,
	}

	return response, nil
}

func (h *UserHandler) PutUsersId(ctx context.Context, request users.PutUsersIdRequestObject) (users.PutUsersIdResponseObject, error) {
	userID := request.Id
	userRequest := request.Body

	userToUpdate := userService.User{
		Password: *userRequest.Password,
		Email:    *userRequest.Email,
	}

	updatedUser, err := h.Service.UpdateUserByID(userID, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PutUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Password: &updatedUser.Password,
		Email:    &updatedUser.Email,
	}

	return response, nil

}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUserByID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
