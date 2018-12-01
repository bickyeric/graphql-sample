## Store
Attribute Name | Data Type
------------ | :-------------:
id | Int
name | String
address | String
city | String
state | String
zip | String
email | String
description | String
website | String
twitter | String
facebook | String
instagram | String
image | String
outlet | [Outlet](#outlet)
type | [StoreType](#storetype)
category | [Category](#category)

## StoreType
Attribute Name | Data Type
------------ | :-------------:
id | Int
name | String
category | [StoreCategory](#storecategory)
store | [Store](#store)

## Outlet
Attribute Name | Data Type
------------ | :-------------:
id | Int
name | String
address | String
phone_number | String
city | String
state | String
status | Boolean
zip | String
store | [Store](#store)

## StoreCategory
Attribute Name | Data Type
------------ | :-------------:
id | Int
name | String
type | [StoreType](#storetype)
