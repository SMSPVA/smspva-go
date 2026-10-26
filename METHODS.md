# Go SDK methods

Generated from the shared SDK contracts. All methods accept a context and optional request options. Request fields are pointers: nil omits a field. `Ptr(value)` creates a pointer.

## Activation

| Method | HTTP |
| --- | --- |
| `client.Activation.GetBalance` | `GET /activation/balance` |
| `client.Activation.GetUserInfo` | `GET /activation/userinfo` |
| `client.Activation.GetOperators` | `GET /activation/operators/{country}` |
| `client.Activation.GetConversions` | `GET /activation/conversions/{service}` |
| `client.Activation.GetMessageTypeSettings` | `GET /activation/countryservicesettings` |
| `client.Activation.GetAllNumberCounts` | `GET /activation/countnumbers` |
| `client.Activation.GetCountryNumberCounts` | `GET /activation/countnumbers/{country}` |
| `client.Activation.GetServicePrice` | `GET /activation/serviceprice/{country}/{service}` |
| `client.Activation.GetCountryOperatorPrices` | `GET /activation/serviceprice/{country}` |
| `client.Activation.GetServicePrices` | `GET /activation/serviceprices/{service}` |
| `client.Activation.GetAllServicePrices` | `GET /activation/servicesprices` |
| `client.Activation.GetNumber` | `GET /activation/number/{country}/{service}` |
| `client.Activation.GetNumberByOperator` | `GET /activation/number/{country}/{service}/{operator}` |
| `client.Activation.ReceiveSms` | `GET /activation/sms/{orderid}` |
| `client.Activation.GetOrders` | `GET /activation/orders` |
| `client.Activation.GetNumberStatus` | `GET /activation/numberstatus/{number}/{service}` |
| `client.Activation.ClearSms` | `PUT /activation/clearsms/{orderid}` |
| `client.Activation.StopSms` | `PUT /activation/stopsms/{orderid}` |
| `client.Activation.CancelOrder` | `PUT /activation/cancelorder/{orderid}` |
| `client.Activation.BlockNumber` | `PUT /activation/blocknumber/{orderid}` |

## Rental

| Method | HTTP |
| --- | --- |
| `client.Rental.GetCountries` | `GET /api/rent.php?method=getcountries` |
| `client.Rental.GetDefaultServices` | `GET /api/rent.php?method=get_default_services` |
| `client.Rental.GetCountriesByService` | `GET /api/rent.php?method=get_country_by_service` |
| `client.Rental.GetData` | `GET /api/rent.php?method=getdata` |
| `client.Rental.GetDataWithProviders` | `GET /api/rent.php?method=getdataWithProviders` |
| `client.Rental.GetCount` | `GET /api/rent.php?method=getcount` |
| `client.Rental.GetCountWithProviders` | `GET /api/rent.php?method=getCountWithProviders` |
| `client.Rental.GetMultiServiceCount` | `GET /api/rent.php?method=get_count_multi` |
| `client.Rental.GetCanadaPaypalDays` | `GET /api/rent.php?method=get_days_canada_paypal` |
| `client.Rental.GetMinimumPaypalDays` | `GET /api/rent.php?method=get_min_paypal_days` |
| `client.Rental.Create` | `GET /api/rent.php?method=create` |
| `client.Rental.CreateMulti` | `GET /api/rent.php?method=create_multi` |
| `client.Rental.AddServiceToOrder` | `GET /api/rent.php?method=add_service_to_order` |
| `client.Rental.GetOrders` | `GET /api/rent.php?method=orders` |
| `client.Rental.Activate` | `GET /api/rent.php?method=activate` |
| `client.Rental.GetSms` | `GET /api/rent.php?method=sms` |
| `client.Rental.Prolong` | `GET /api/rent.php?method=prolong` |
| `client.Rental.ProlongByDays` | `GET /api/rent.php?method=prolong_max` |
| `client.Rental.DeleteOrder` | `GET /api/rent.php?method=delete` |
| `client.Rental.GetHistory` | `GET /api/rent.php?method=get_rent_history` |
| `client.Rental.GetRestoreCost` | `GET /api/rent.php?method=restore_user_precalc` |
| `client.Rental.Restore` | `GET /api/rent.php?method=restore_user` |

## Email

| Method | HTTP |
| --- | --- |
| `client.Email.GetProviders` | `GET /api/rent_email.php?method=get_providers` |
| `client.Email.GetData` | `GET /api/rent_email.php?method=get_data` |
| `client.Email.GetCount` | `GET /api/rent_email.php?method=get_count` |
| `client.Email.CreateOrder` | `GET /api/rent_email.php?method=create_order` |
| `client.Email.GetOrders` | `GET /api/rent_email.php?method=get_orders` |
| `client.Email.GetMails` | `GET /api/rent_email.php?method=get_mails` |
| `client.Email.GetMailBody` | `GET /api/rent_email.php?method=get_mail_body` |
| `client.Email.BuyAlias` | `GET /api/rent_email.php?method=buy_alias` |
| `client.Email.GetMailsByAlias` | `GET /api/rent_email.php?method=get_mails_by_alias` |
| `client.Email.ProlongOrder` | `GET /api/rent_email.php?method=prolong_order` |
| `client.Email.CancelOrder` | `GET /api/rent_email.php?method=cancel_order` |
| `client.Email.GetArchivedOrders` | `GET /api/rent_email.php?method=get_archived_orders` |
| `client.Email.CalcRestoreOrder` | `GET /api/rent_email.php?method=calc_restore_order` |
| `client.Email.RestoreOrder` | `GET /api/rent_email.php?method=restore_order` |
