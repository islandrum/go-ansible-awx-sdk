package awx
import (
	"bytes"
	"encoding/json"
	"fmt"
)

// UserService implements awx Users apis.
type UserService struct {
	client *Client
}

// ListUsersResponse represents `ListUsers` endpoint response.
type ListUsersResponse struct {
	Pagination
	Results []*User `json:"results"`
}

const usersAPIEndpoint = "/api/v2/users/"

// ListUsers shows list of awx Users.
func (u *UserService) ListUsers(params map[string]string) ([]*User, *ListUsersResponse, error) {
	result := new(ListUsersResponse)
	resp, err := u.client.Requester.GetJSON(usersAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateUser creates an awx User.
func (u *UserService) CreateUser(data map[string]interface{}, params map[string]string) (*User, error) {
	mandatoryFields = []string{"username", "password", "first_name", "last_name", "email"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(User)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if User exists and return proper error

	resp, err := u.client.Requester.PostJSON(usersAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUser update an awx user.
func (u *UserService) UpdateUser(id int, data map[string]interface{}, params map[string]string) (*User, error) {
	result := new(User)
	endpoint := fmt.Sprintf("%s%d", usersAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Requester.PutJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteUser delete an awx User.
func (u *UserService) DeleteUser(id int) (*User, error) {
	result := new(User)
	endpoint := fmt.Sprintf("%s%d", usersAPIEndpoint, id)

	resp, err := u.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
