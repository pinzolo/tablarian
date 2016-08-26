# sales_order_header

General sales order information.

## 列一覧

| PK |           列名            |          型          | サイズ | NULL |       初期値       |                                                   コメント                                                    |
|----|---------------------------|----------------------|--------|------|--------------------|---------------------------------------------------------------------------------------------------------------|
|  1 | sales_order_id            | serial               |        | NO   |                    | Primary key.                                                                                                  |
|    | revision_number           | smallint             |        | NO   |                  0 | Incremental number to track changes to the sales order over time.                                             |
|    | order_date                | timestamp            |      6 | NO   | now()              | Dates the sales order was created.                                                                            |
|    | due_date                  | timestamp            |      6 | NO   |                    | Date the order is due to the customer.                                                                        |
|    | ship_date                 | timestamp            |      6 |      |                    | Date the order was shipped to the customer.                                                                   |
|    | status                    | smallint             |        | NO   |                  1 | Order current status. 1 = In process; 2 = Approved; 3 = Backordered; 4 = Rejected; 5 = Shipped; 6 = Cancelled |
|    | online_order_flag         | public.Flag          |        | NO   | true               | 0 = Order placed by sales person. 1 = Order placed online by customer.                                        |
|    | purchase_order_number     | public.OrderNumber   |     25 |      |                    | Customer purchase order number reference.                                                                     |
|    | account_number            | public.AccountNumber |     15 |      |                    | Financial accounting number reference.                                                                        |
|    | customer_id               | integer              |        | NO   |                    | Customer identification number. Foreign key to customer.business_entity_id.                                   |
|    | sales_person_id           | integer              |        |      |                    | Sales person who created the sales order. Foreign key to sales_person.business_entity_id.                     |
|    | territory_id              | integer              |        |      |                    | Territory in which the sale was made. Foreign key to sales_territory.sales_territory_id.                      |
|    | bill_to_address_id        | integer              |        | NO   |                    | Customer billing address. Foreign key to address.address_id.                                                  |
|    | ship_to_address_id        | integer              |        | NO   |                    | Customer shipping address. Foreign key to address.address_id.                                                 |
|    | ship_method_id            | integer              |        | NO   |                    | Shipping method. Foreign key to ship_method.ship_method_id.                                                   |
|    | credit_card_id            | integer              |        |      |                    | Credit card identification number. Foreign key to credit_card.credit_card_id.                                 |
|    | credit_card_approval_code | varchar              |     15 |      |                    | Approval code provided by the credit card company.                                                            |
|    | currency_rate_id          | integer              |        |      |                    | Currency exchange rate used. Foreign key to currency_rate.currency_rate_id.                                   |
|    | sub_total                 | numeric              |        | NO   |               0.00 | Sales subtotal. Computed as SUM(sales_order_detail.line_total)for the appropriate sales_order_id.             |
|    | tax_amt                   | numeric              |        | NO   |               0.00 | Tax amount.                                                                                                   |
|    | freight                   | numeric              |        | NO   |               0.00 | Shipping cost.                                                                                                |
|    | total_due                 | numeric              |        |      |                    | Total due from customer. Computed as subtotal + tax_amt + freight.                                            |
|    | comment                   | varchar              |    128 |      |                    | Sales representative comments.                                                                                |
|    | rowguid                   | uuid                 |        | NO   | uuid_generate_v1() |                                                                                                               |
|    | modified_date             | timestamp            |      6 | NO   | now()              |                                                                                                               |

## インデックス

|                 名前                 |       列       | ユニーク |
|--------------------------------------|----------------|----------|
| pk_sales_order_header_sales_order_id | sales_order_id | YES      |

## 制約

|             製薬名              | KIND  |                      CONTENT                       |
|---------------------------------|-------|----------------------------------------------------|
| ck_sales_order_header_due_date  | CHECK | (due_date >= order_date)                           |
| ck_sales_order_header_freight   | CHECK | (freight >= 0.00)                                  |
| ck_sales_order_header_ship_date | CHECK | ((ship_date >= order_date) OR (ship_date IS NULL)) |
| ck_sales_order_header_status    | CHECK | ((status >= 0) AND (status <= 8))                  |
| ck_sales_order_header_sub_total | CHECK | (sub_total >= 0.00)                                |
| ck_sales_order_header_tax_amt   | CHECK | (tax_amt >= 0.00)                                  |

## 参照キー

|                        参照名                        |         列         |      参照テーブル      |       参照列       |
|------------------------------------------------------|--------------------|------------------------|--------------------|
| fk_sales_order_header_address_bill_to_address_id     | bill_to_address_id | person.address         | address_id         |
| fk_sales_order_header_address_ship_to_address_id     | ship_to_address_id | person.address         | address_id         |
| fk_sales_order_header_credit_card_credit_card_id     | credit_card_id     | credit_card            | credit_card_id     |
| fk_sales_order_header_currency_rate_currency_rate_id | currency_rate_id   | currency_rate          | currency_rate_id   |
| fk_sales_order_header_customer_customer_id           | customer_id        | customer               | customer_id        |
| fk_sales_order_header_sales_person_sales_person_id   | sales_person_id    | sales_person           | business_entity_id |
| fk_sales_order_header_sales_territory_territory_id   | territory_id       | sales_territory        | territory_id       |
| fk_sales_order_header_ship_method_ship_method_id     | ship_method_id     | purchasing.ship_method | ship_method_id     |

## 被参照キー

|                             参照名                              |         参照元テーブル          |    参照元列    |    被参照列    |
|-----------------------------------------------------------------|---------------------------------|----------------|----------------|
| fk_sales_order_detail_sales_order_header_sales_order_id         | sales_order_detail              | sales_order_id | sales_order_id |
| fk_sales_order_header_sales_reason_sales_order_header_sales_ord | sales_order_header_sales_reason | sales_order_id | sales_order_id |
