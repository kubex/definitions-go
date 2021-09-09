package app

type Category string

const (
	CategoryBooks            Category = "books"
	CategoryMedical          Category = "medical"
	CategoryDeveloperTools   Category = "developer-tools"
	CategoryFinance          Category = "finance"
	CategoryEducation        Category = "education"
	CategoryNews             Category = "news"
	CategoryEntertainment    Category = "entertainment"
	CategoryProductivity     Category = "productivity"
	CategoryReference        Category = "reference"
	CategoryGames            Category = "games"
	CategorySocialNetworking Category = "social-networking"
	CategoryLifestyle        Category = "lifestyle"
	CategoryUtilities        Category = "utilities"
	CategoryOther            Category = "other"
	CategoryAccountingTax    Category = "accounting-and-tax"
	CategorySales            Category = "sales"
	CategoryCustomerService  Category = "customer-service"
	CategorySupply           Category = "supply"
	CategoryHealthWellness   Category = "health-and-wellness"
	CategoryInsurance        Category = "insurance"
	CategoryWebServices      Category = "web-services"
	CategoryLegal            Category = "legal"
	CategoryTravel           Category = "travel"
	CategoryMarketing        Category = "marketing"
	CategoryTransport        Category = "transport"
	CategoryEvents           Category = "events"
	CategoryProduction       Category = "production"
	CategoryPurchasing       Category = "purchasing"
	CategoryHumanResources   Category = "human-resources"
	CategoryOperations       Category = "operations"
	CategoryRecruitment      Category = "recruitment"
	CategoryAffiliates       Category = "affiliates"
)

func (c Category) Name() string {
	switch c {
	case CategoryBooks:
		return "Books"
	case CategoryMedical:
		return "Medical"
	case CategoryDeveloperTools:
		return "Developer Tools"
	case CategoryFinance:
		return "Finance"
	case CategoryEducation:
		return "Education"
	case CategoryNews:
		return "News"
	case CategoryEntertainment:
		return "Entertainment"
	case CategoryProductivity:
		return "Productivity"
	case CategoryReference:
		return "Reference"
	case CategoryGames:
		return "Games"
	case CategorySocialNetworking:
		return "Social Networking"
	case CategoryLifestyle:
		return "Lifestyle"
	case CategoryUtilities:
		return "Utilities"
	case CategoryOther:
		return "Other"
	case CategoryAccountingTax:
		return "Accounting And Tax"
	case CategorySales:
		return "Sales"
	case CategoryCustomerService:
		return "Customer Service"
	case CategorySupply:
		return "Supply"
	case CategoryHealthWellness:
		return "Health And Wellness"
	case CategoryInsurance:
		return "Insurance"
	case CategoryWebServices:
		return "Web Services"
	case CategoryLegal:
		return "Legal"
	case CategoryTravel:
		return "Travel"
	case CategoryMarketing:
		return "Marketing"
	case CategoryTransport:
		return "Transport"
	case CategoryEvents:
		return "Events"
	case CategoryProduction:
		return "Production"
	case CategoryPurchasing:
		return "Purchasing"
	case CategoryHumanResources:
		return "Human Resources"
	case CategoryOperations:
		return "Operations"
	case CategoryRecruitment:
		return "Recruitment"
	case CategoryAffiliates:
		return "Affiliates"
	}
	return "Uncategorized"
}
