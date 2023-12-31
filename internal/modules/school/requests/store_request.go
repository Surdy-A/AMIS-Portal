package requets

type StoreRequest struct {
	SchoolName string `form:"school_name" json:"school_name" binding:"required,min=3,max=100"`
}
