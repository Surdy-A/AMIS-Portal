package models

import (
	"github.com/Surdy-A/amis_portal/internal/modules/user/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var LGAName = map[string]string{
	"OS01": "Atakunmosa East Central LCDA",
	"OS02": "Atakunmosa West",
	"OS03": "Atakunmosa West Central",
	"OS04": "Ayedaade",
	"OS05": "Ayedaade South LCDA",
	"OS06": "Ayedire",
	"OS07": "Ayedire South LCDA",
	"OS09": "Boluwaduro East LCDA",
	"OS10": "Boripe",
	"OS11": "Boripe North LCDA",
	"OS12": "Ede East LCDA",
	"OS13": "Ede North",
	"OS14": "Ede North Area Council",
	"OS15": "Ede South",
	"OS16": "Egbedore",
	"OS17": "Egbedore Area Council",
	"OS18": "Egbedore South LCDA",
	"OS19": "Ejigbo",
	"OS20": "Ejigbo South LCDA",
	"OS21": "Ejigbo West LCDA",
	"OS22": "Ife Central",
	"OS23": "Ife Central West LCDA",
	"OS24": "Ife East",
	"OS25": "Ife North",
	"OS26": "Ife North Area Council",
	"OS27": "Ife North Central LCDA",
	"OS28": "Ife North West LCDA",
	"OS29": "Ife Ooye LCDA",
	"OS30": "Ife South",
	"OS31": "Ife South West",
	"OS32": "Ifedayo",
	"OS33": "Ifedayo Area Council",
	"OS34": "Ifelodun",
	"OS35": "Ifelodun North Area Council",
	"OS36": "Ifelodun North LCDA",
	"OS37": "Ila",
	"OS38": "Ila Central LCDA",
	"OS39": "Ilesa East",
	"OS40": "Ilesa North East LCDA",
	"OS41": "Ilesa West",
	"OS42": "Ilesa West Central LCDA",
	"OS43": "Irepodun",
	"OS44": "Irepodun South LCDA",
	"OS45": "Irewole",
	"OS46": "Irewole North East LCDA",
	"OS47": "Isokan",
	"OS48": "Isokan South LCDA",
	"OS49": "Iwo",
	"OS50": "Iwo East LCDA",
	"OS51": "Iwo West LCDA",
	"OS52": "Obokun",
	"OS53": "Obokun East LCDA",
	"OS54": "Odo Otin",
	"OS55": "Odo Otin North LCDA",
	"OS56": "Odo Otin South LCDA",
	"OS57": "Ola Oluwa",
	"OS58": "Ola Oluwa South East LCDA",
	"OS59": "Olorunda",
	"OS60": "Olorunda Area Council",
	"OS61": "Olorunda North LCDA",
	"OS62": "Oriade",
	"OS63": "Oriade South LCDA",
	"OS64": "Orolu",
	"OS65": "Orolu Area Council",
	"OS66": "Osogbo",
	"OS67": "Osogbo South LCDA",
	"OS68": "Osogbo West LCDA",
}

var StateGovStatusList = map[int]string{
	1: "Approved",
	2: "Under Processing",
	3: "Not Yet Applied",
}

var FederalGovStatusList = map[int]string{
	1: "Registered",
	2: "Not Registered",
}

var Ownerships = map[int]string{
	1: "Sole Proprietorship",
	2: "Partnership",
	3: "Mission",
}

var Sex = map[int]string{
	1: "Male",
	2: "Female",
}

var Class = map[int]string{
	3: "Grade 3",
	4: "Grade 4",
	5: "Grade 5",
	6: "Grade 6",
}

var SchoolTypeList = map[int]string{
	1: "Nursery",
	2: "Primary",
	3: "Secondary",
	4: "Islamiyyah",
}

var AcademicCompetitionCategory = map[int]string{
	1: "Primary",
	2: "JSS",
	3: "SSS",
}

var SchoolOperationTypeLists = map[int]string{
	1: "Day",
	2: "Boarding",
}

// var Class = map[int]string{
// 	3:  "Grade 3",
// 	4:  "Grade 4",
// 	5: "Grade 5",
// 	6: "Grade 6",
// }

// #School Name and Code Choice
// SCHOOL_NAME__AND_CODE = (
// 	"51241" 'ISLAMIC MODEL SCHOOL. IGANGAN - 51241",
// 	"61321" 'THE LIGHT NUR & PRY SCHOOL, ORILE- OWU - 61321",
// 	"61322" 'EPITOME MONTESSORI N/P SCH. ODE-OMU - 61322",
// 	"42501" 'FOMWAN N /P SCHOOL IREE - 42501",
// )

// LG_NAME__AND_CODE = (
// 	"51" 'ATAKUNMOSA WEST - 51",
// 	"61" 'AYEDAADE - 61",
// 	"42" 'BORIPE - 42",
// 	"11" 'EDE NORTH - 11",
// )

type Proprietor struct {
	gorm.Model
	Name          string `json:"name"`
	Qualification string `json:"qualification"`
	Phone         string `json:"phone"`
	Occupation    string `json:"occupation"`
	Address       string `json:"address"`
	SchoolID      uint
}

type School struct {
	gorm.Model
	UserID                  uint
	User                    models.User
	SchoolName              string         `form:"school_name" json:"school_name" gorm:"varchar:250"`
	Address                 string         `form:"address" json:"address" binding:"required,min=3,max=100"`
	Email                   string         `form:"email"  json:"email" binding:"required,min=3,max=100"`
	Phone                   string         `form:"phone" json:"phone" binding:"required,min=3,max=100"`
	Website                 string         `form:"website" json:"website"`
	About                   string         `form:"about" json:"about" binding:"required,min=3,max=100"`
	StateGovStatus          string         `form:"state_gov_status" json:"state_gov_status"`
	FederalGovernmentStatus string         `form:"federal_government_status" json:"federal_government_status"`
	CacRegNumber            string         `form:"cac_reg_number" json:"cac_reg_number" binding:"min=3,max=100"`
	AssociationDetails      string         `form:"association_details" json:"association_details" binding:"min=3,max=100"`
	SchoolCode              string         `form:"school_code" json:"school_code" binding:"min=3,max=100"`
	LGACode                 string         `form:"lga_code" json:"lga_code" binding:"min=3,max=100"`
	Ownership               string         `form:"ownership" json:"ownership"`
	LGA                     string         `form:"lga" json:"lga"`
	Logo                    string         `form:"logo" json:"logo" binding:"required,min=3,max=100"`
	Paid                    bool           `form:"other_association" json:"paid"`
	SchoolTypes             pq.StringArray `form:"school_types" json:"school_types" gorm:"type:text[]"`
	SchoolOperationTypes    pq.StringArray `form:"school_operation_types" json:"school_operation_types" gorm:"type:text[]"`
}
