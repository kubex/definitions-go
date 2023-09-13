package app

type Category string

const (
	CategoryGeneral   Category = "general"
	CategoryProduct   Category = "product"
	CategoryLeads     Category = "leads"
	CategorySales     Category = "sales"
	CategoryBilling   Category = "billing"
	CategoryDelivery  Category = "delivery"
	CategoryRetain    Category = "retain"
	CategoryRecover   Category = "recover"
	CategoryUtilities Category = "utilities"
	CategorySettings  Category = "settings"
	CategorySecurity  Category = "security"
)

func (c Category) Name() string {
	switch c {
	case CategoryGeneral:
		return "General"
	case CategoryProduct:
		return "Product"
	case CategoryLeads:
		return "Leads"
	case CategorySales:
		return "Sales"
	case CategoryBilling:
		return "Billing"
	case CategoryDelivery:
		return "Delivery"
	case CategoryRetain:
		return "Retain"
	case CategoryRecover:
		return "Recover"
	case CategoryUtilities:
		return "Utilities"
	case CategorySettings:
		return "Settings"
	case CategorySecurity:
		return "Security"
	}
	return "Uncategorized"
}
