package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCmdIndex(t *testing.T) {
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	cmdIndex.Run([]string{})
	expected := `
country_region_currency
credit_card                      Customer credit card information.
currency                         Lookup table containing standard ISO currencies.
currency_rate                    Currency exchange rates.
customer                         Current customer information. Also see the Person and Store tables.
person_credit_card               Cross-reference table mapping people to their credit card information in the credit_card table.
sales_order_detail               Individual products associated with a specific sales order. See sales_order_header.
sales_order_header               General sales order information.
sales_order_header_sales_reason  Cross-reference table mapping sales orders to sales reason codes.
sales_person                     Sales representative current information.
sales_person_quota_history       Sales performance tracking.
sales_reason                     Lookup table of customer purchase reasons.
sales_tax_rate                   Tax rate lookup table.
sales_territory                  Sales territory lookup table.
sales_territory_history          Sales representative transfers to other sales territories.
shopping_cart_item               Contains online customer orders until the order is submitted or cancelled.
special_offer                    Sale discounts lookup table.
special_offer_product            Cross-reference table mapping products to special offer discounts.
store                            Customers (resellers) of Adventure Works products.`
	actual := buf.String()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("\nactual:\n%v\nexpected:%v\n", actual, expected)
	}
}

func TestCmdIndexWithOtherConfig(t *testing.T) {
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-test")
	configFile = "test/tablarian-aw.config"
	cmdIndex.Run([]string{})
	expected := `
country_region_currency
credit_card                      Customer credit card information.
currency                         Lookup table containing standard ISO currencies.
currency_rate                    Currency exchange rates.
customer                         Current customer information. Also see the Person and Store tables.
person_credit_card               Cross-reference table mapping people to their credit card information in the credit_card table.
sales_order_detail               Individual products associated with a specific sales order. See sales_order_header.
sales_order_header               General sales order information.
sales_order_header_sales_reason  Cross-reference table mapping sales orders to sales reason codes.
sales_person                     Sales representative current information.
sales_person_quota_history       Sales performance tracking.
sales_reason                     Lookup table of customer purchase reasons.
sales_tax_rate                   Tax rate lookup table.
sales_territory                  Sales territory lookup table.
sales_territory_history          Sales representative transfers to other sales territories.
shopping_cart_item               Contains online customer orders until the order is submitted or cancelled.
special_offer                    Sale discounts lookup table.
special_offer_product            Cross-reference table mapping products to special offer discounts.
store                            Customers (resellers) of Adventure Works products.`
	actual := buf.String()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("\nactual:\n%v\nexpected:%v\n", actual, expected)
	}
}

func TestCmdIndexWithOtherConfigByAbsPath(t *testing.T) {
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	absPath, err := testConfigFilePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure config loading: %v", err)
	}
	configFile = "@" + absPath
	cmdIndex.Run([]string{})
	expected := `
country_region_currency
credit_card                      Customer credit card information.
currency                         Lookup table containing standard ISO currencies.
currency_rate                    Currency exchange rates.
customer                         Current customer information. Also see the Person and Store tables.
person_credit_card               Cross-reference table mapping people to their credit card information in the credit_card table.
sales_order_detail               Individual products associated with a specific sales order. See sales_order_header.
sales_order_header               General sales order information.
sales_order_header_sales_reason  Cross-reference table mapping sales orders to sales reason codes.
sales_person                     Sales representative current information.
sales_person_quota_history       Sales performance tracking.
sales_reason                     Lookup table of customer purchase reasons.
sales_tax_rate                   Tax rate lookup table.
sales_territory                  Sales territory lookup table.
sales_territory_history          Sales representative transfers to other sales territories.
shopping_cart_item               Contains online customer orders until the order is submitted or cancelled.
special_offer                    Sale discounts lookup table.
special_offer_product            Cross-reference table mapping products to special offer discounts.
store                            Customers (resellers) of Adventure Works products.`
	actual := buf.String()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("\nactual:\n%v\nexpected:%v\n", actual, expected)
	}
}

func TestCmdIndexWithNoCommentOption(t *testing.T) {
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	withoutTableComment = true
	cmdIndex.Run([]string{})
	expected := `
country_region_currency
credit_card
currency
currency_rate
customer
person_credit_card
sales_order_detail
sales_order_header
sales_order_header_sales_reason
sales_person
sales_person_quota_history
sales_reason
sales_tax_rate
sales_territory
sales_territory_history
shopping_cart_item
special_offer
special_offer_product
store`
	actual := buf.String()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("\nactual:\n%v\nexpected:%v\n", actual, expected)
	}
}

func TestCmdIndexWithInvalidJson(t *testing.T) {
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("invalid-json")
	stat := cmdIndex.Run([]string{})
	if stat == 0 {
		t.Error("Index command should not finish normally on invalid schema.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "unexpected end of JSON input"; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdIndexWithDbError(t *testing.T) {
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("db-error")
	stat := cmdIndex.Run([]string{})
	if stat == 0 {
		t.Error("Index command should not finish normally on invalid schema.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), `pq: role "foobar" does not exist`; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}
