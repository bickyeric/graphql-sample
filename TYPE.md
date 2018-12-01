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
Attribute Name | Data Type
------------ | :-------------:
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
item_category | [ [ItemCategory](#itemcategory) ]!
item_modifier | [ [Modifier](#modifier) ]!
item | [ [Item](#item) ]!

## Outlet
Attribute Name | Data Type
------------ | :-------------:
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
item_category | [ [ItemCategory](#itemcategory) ]!
item_modifier | [ [Modifier](#modifier) ]!
item | [ [Item](#item) ]!

## Employee
Attribute Name | Data Type
------------ | :-------------:
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
