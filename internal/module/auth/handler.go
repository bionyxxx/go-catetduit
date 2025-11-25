package auth

import (
	"catetduit/internal/config"
	"catetduit/internal/helper"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   *Service
	validator *validator.Validate
}

func NewHandler(service *Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := helper.ResponseBadRequest(w, "Invalid request payload", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	if err := h.validator.Struct(req); err != nil {
		errDetails := helper.FormatValidationErrors(err)
		err := helper.ResponseUnprocessableEntity(w, "Validation failed", errDetails)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	loginResp, err := h.service.Authenticate(req.Email, req.Password)

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			err := helper.ResponseUnauthorized(w, "Invalid email or password")
			if err != nil {
				fmt.Println("Error sending response:", err)
			}
			return
		}
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Login successful", loginResp)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := helper.ResponseBadRequest(w, "Invalid request payload", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	if err := h.validator.Struct(req); err != nil {
		errDetails := helper.FormatValidationErrors(err)
		err := helper.ResponseUnprocessableEntity(w, "Validation failed", errDetails)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	refreshResp, err := h.service.RefreshToken(req.RefreshToken)

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			err := helper.ResponseUnauthorized(w, "Invalid refresh token")
			if err != nil {
				fmt.Println("Error sending response:", err)
			}
			return
		}
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Token refreshed successfully", refreshResp)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := helper.ResponseBadRequest(w, "Invalid request payload", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	if err := h.validator.Struct(req); err != nil {
		errDetails := helper.FormatValidationErrors(err)
		err := helper.ResponseUnprocessableEntity(w, "Validation failed", errDetails)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	userResp, err := h.service.Register(&req)

	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err = helper.ResponseCreated(w, "Registration successful", userResp)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}

}

func (h *Handler) GoogleLogin(oauth2Config *config.OAuth2Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oauth2Config.GoogleConfig.AuthCodeURL(oauth2Config.StateString)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

//// GoogleCallback
//func (h *Handler) GoogleCallback(oauth2Config *helper.OAuth2Config) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if r.FormValue("state") != oauth2Config.StateString {
//			err := helper.ResponseUnauthorized(w, "Invalid OAuth2 state")
//			if err != nil {
//				fmt.Println("Error sending response:", err)
//			}
//			return
//		}
//
//		token, err := oauth2Config.Config.Exchange(r.Context(), r.FormValue("code"), oauth2Config.AuthCodeOptions...)
//		if err != nil {
//			err := helper.ResponseInternalServerError(w, "Failed to exchange token", err.Error())
//			if err != nil {
//				fmt.Println("Error sending response:", err)
//			}
//			return
//		}
//
//		client := oauth2Config.Config.Client(r.Context(), token)
//		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
//		if err != nil || resp.StatusCode != http.StatusOK {
//			err := helper.ResponseInternalServerError(w, "Failed to get user info from Google", "")
//			if err != nil {
//				fmt.Println("Error sending response:", err)
//			}
//			return
//		}
//		defer resp.Body.Close()
//
//		var googleUser GoogleUserInfo
//		if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
//			err := helper.ResponseInternalServerError(w, "Failed to decode Google user info", err.Error())
//			if err != nil {
//				fmt.Println("Error sending response:", err)
//			}
//			return
//		}
//
//		authResp, err := h.service.GoogleAuthenticate(&googleUser)
//		if err != nil {
//			err := helper.ResponseInternalServerError(w, "Authentication failed", err.Error())
//			if err != nil {
//				fmt.Println("Error sending response:", err)
//			}
//			return
//		}
//
//		err = helper.ResponseOKWithData(w, "Login with Google successful", authResp)
//		if err != nil {
//			fmt.Println("Error sending response:", err)
//		}
//	}
//}
