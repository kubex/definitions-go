package app

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type Category string

const (
	CategoryAdministration Category = "administration"
	CategoryMarketing      Category = "marketing"
	CategoryDevelopment    Category = "development"
	CategorySupport        Category = "support"
	CategoryPayments       Category = "payments"
	CategoryCommunication  Category = "communication"
	CategoryChurn          Category = "churn"
	CategoryRenewal        Category = "renewal"
	CategoryFinance        Category = "finance"
	CategoryKnowledge      Category = "knowledge"
	CategoryOperations     Category = "operations"
	CategoryReporting      Category = "reporting"
	CategoryCustomers      Category = "customers"
	CategoryGeneral        Category = "general"
	CategoryProduct        Category = "product"
	CategoryLeads          Category = "leads"
	CategorySales          Category = "sales"
	CategoryBilling        Category = "billing"
	CategoryDelivery       Category = "delivery"
	CategoryRetain         Category = "retain"
	CategoryRecover        Category = "recover"
	CategoryUtilities      Category = "utilities"
	CategorySettings       Category = "settings"
	CategorySecurity       Category = "security"
	CategoryHR             Category = "human-resources"
)

func (c Category) Name() string {
	return cases.Title(language.English, cases.Compact).String(strings.ReplaceAll(string(c), "-", " "))
}
