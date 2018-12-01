## Store
type **Store** {
id: Int
name: String
address: String
city: String
state: String
zip: String
email: String
description: String
website: String
twitter: String
facebook: String
instagram: String
image: String
outlet: **Outlet!**
type: **StoreType!**
category: **Category!**
}

> type **StoreType** {
> id: Int
> name: String
> category: **StoreCategory**
> store: **Store**
> }

> type **Outlet** {
> id: Int
> name: String
> address: String
> phone_number: String
> city: String
> state: String
> status: Boolean
> zip: String
> store: **Store!**
> }

> type **StoreCategory** {
> id: Int
> name: String
> type: **StoreType**
> }