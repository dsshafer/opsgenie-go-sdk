package userv2

// UpdateUserResponse is response with status of updating user.
type UpdateUserResponse struct {
	Result string `json:"result"`
	*ResponseMeta
}
