## StoreType
Attribute Name | Data Type
------------ | :-------------:
id | Int!
name | String!
category | [ [StoreCategory](#storecategory) ]
store | [ [Store](#store) ]

## StoreCategory
Attribute Name | Data Type
------------ | :-------------:
id | Int!
name | String!
type | [StoreType](#storetype)!
store | [ [Store](#store) ]!

## Store
Attribute Name | Data Type | Comment
 ------------  | :-------: | :-----:
id | Int!
name | String!
address | String!
city | String!
state | String!
zip | String!
email | String!
description | String!
website | String
twitter | String
facebook | String
instagram | String
image | String
outlet | [ [Outlet](#outlet) ]!
type | [StoreType](#storetype)!
category | [StoreCategory](#storecategory)!
employee | [ [Employee](#employee) ]!
item_category | [ [ItemCategory](#itemcategory) ]! | Not Implemented
item_modifier | [ [Modifier](#modifier) ]! | Not Implemented
item | [ [Item](#item) ]! | Not Implemented
customer | [ [Customer](#customer) ]! | Not Implemented
discount | [ [Discount](#discount) ]! | Not Implemented
tax | [ [Tax](#tax) ]! | Not Implemented

## Outlet
Attribute Name | Data Type | Comment
 ------------  | :-------: | :-----:
id | Int!
name | String!
address | String!
phone_number | String!
city | String!
state | String!
status | Boolean!
zip | String!
store | [Store](#store)!
employee | [ [Employee](#employee) ]!
item_category | [ [ItemCategory](#itemcategory) ]! | Not Implemented
item_modifier | [ [Modifier](#modifier) ]! | Not Implemented
item | [ [Item](#item) ]! | Not Implemented
discount | [ [Discount](#discount) ]! | Not Implemented
tax | [ [Tax](#tax) ]! | Not Implemented

## Employee
Attribute Name | Data Type | Comment
 ------------  | :-------: | :-----:
id | Int!
outlet | [Outlet](#outlet)!
store | [Store](#store)!
first_name | String!
last_name | String!
phone_number | String!
email | String!
password | String!
confirmed | Boolean!
active | Boolean!
transaction | [ [Transaction](#transaction) ]! | Not Implemented

## ItemCategory
Attribute Name | Data Type
------------ | :-------------:
id | Int!
name | String!
outlet | [Outlet](#outlet)!
store | [Store](#store)!
item | [ [Item](#item) ]

## Modifier
Attribute Name | Data Type
------------ | :-------------:
id | Int!
name | String!
outlet | [Outlet](#outlet)!
store | [Store](#store)!
item | [ [Item](#item) ]
option | [ [Option](#option) ]

## Item
Attribute Name | Data Type
------------ | :-------------:
id | Int!
outlet | [Outlet](#outlet)!
store | [Store](#store)!
name | String!
category | [ItemCategory](#itemcategory)
description | String!
modifier | [Modifier](#modifier)
image | String
variant | [ [Variant](#variant) ]!
transaction_detail | [ [TransactionDetail](#transactiondetail) ]!

## Variant
Attribute Name | Data Type
------------ | :-------------:
id | Int!
item | [Item](#item)!
price | Int!
sku | String!
stock | Int!
track_stock | Boolean!
alert_at | Int!
alert | Int!
cost | Int!
track_cogs | Boolean!
transaction_detail | [ [TransactionDetail](#transactiondetail) ]!

## Customer
Attribute Name | Data Type
------------ | :-------------:
id | Int!
store | [Store](#store)!
phone_number | String
name | String!
email | String
birthday | Date
address | String
city | String
state | String
zip | String
transaction | [ [Transaction](#transaction) ]!

## Discount
Attribute Name | Data Type
------------ | :-------------:
id | Int!
outlet | [Outlet](#outlet)!
store | [Store](#store)!
amount | Int!
type | String!

## Tax
Attribute Name | Data Type
------------ | :-------------:
id | Int!
outlet | [Outlet](#outlet)!
store | [Store](#store)!
amount | Int!

## Transaction
Attribute Name | Data Type
------------ | :-------------:
id | Int!
receipt_number | String
employee | [Employee](#employee)!
customer | [Customer](#customer)!
comment | String
status | String!
discount | [Discount](#discount)
tax | [Tax](#tax)
updated_at | Date!
created_at | Date!
detail | [ [TransactionDetail](#transactiondetail) ]!

## TransactionDetail
Attribute Name | Data Type
------------ | :-------------:
id | Int!
transaction | [Transaction](#transaction)!
variant | [Variant](#variant)
item | [Item](#item)
quantity | Int
comment | String
price | Int
option | String
payment | [Payment](#payment)

## Option
Attribute Name | Data Type
------------ | :-------------:
id | Int!
name | String
price | Int
modifier | [Modifier](#modifier)

## TransactionPayment
Attribute Name | Data Type
------------ | :-------------:
transaction | [Transaction](#transaction)!
type | String
cash | Int
comment | String
