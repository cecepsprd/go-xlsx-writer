package model

type Usergroup struct {
	ID            int    `json:"id"`
	UsergroupName string `db:"usergroup_name"`
	Note          string `json:"note"`
}

type User struct {
	ID            int     `db:"id"`
	UsergroupID   int     `db:"usergroup_id"`
	Username      string  `db:"username"`
	Password      string  `db:"password"`
	FullName      string  `db:"full_name"`
	IsActive      int     `db:"is_active"`
	CreatedBy     int     `db:"created_by"`
	Email         string  `db:"email"`
	Phone         string  `db:"phone"`
	IsSuperuser   int     `db:"is_superuser"`
	IsVerified    int     `db:"is_verified"`
	IsSales       int     `db:"is_sales"`
	AndroidID     string  `db:"android_id"`
	Saldo         float64 `db:"saldo"`
	CatalogProfit float64 `db:"catalog_profit"`
	IsFnb         int     `db:"is_fnb"`
}

type Poskora struct {
	ID              int     `db:"id"`
	UserID          int     `db:"user_id"`
	WarehouseID     int     `db:"warehouse_id"`
	Code            string  `db:"code"`
	IsActive        int     `db:"is_active"`
	MaxMember       int     `db:"max_member"`
	Type            string  `db:"type"`
	ShippingAddress string  `db:"shipping_address"`
	VillageID       int     `db:"village_id"`
	Latitude        float64 `db:"latitude"`
	Longitude       float64 `db:"longitude"`
	IsWow           int     `db:"is_wow"`
	CreatedBy       int     `db:"created_by"`
}

type KK struct {
	ID             int     `db:"id"`
	PoskoraID      int     `db:"poskora_id"`
	Name           string  `db:"name"`
	Address        string  `db:"address"`
	VillageID      int     `db:"village_id"`
	Phone          string  `db:"phone"`
	IsDeleted      int     `db:"is_deleted"`
	IsActive       int     `db:"is_active"`
	IsVerified     int     `db:"is_verified"`
	IsPoskora      int     `db:"is_poskora"`
	IsDefault      int     `db:"is_default"`
	SaldoBalance   float64 `db:"saldo_balance"`
	TotalOrder     int     `db:"total_order"`
	CreatedBy      int     `db:"created_by"`
	UpdatedBy      int     `db:"updated_by"`
	SurveyAnswered int     `db:"survey_answered"`
}
