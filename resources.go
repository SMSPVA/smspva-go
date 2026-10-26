// Code generated from sdk/contracts.json. DO NOT EDIT.
package smspva

import "context"

type Activation struct{ client *Client }

// GetBalance calls GET /activation/balance.
func (r *Activation) GetBalance(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getBalance", nil, options...)
}

// GetUserInfo calls GET /activation/userinfo.
func (r *Activation) GetUserInfo(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getUserInfo", nil, options...)
}

type ActivationGetOperatorsParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
}

// GetOperators calls GET /activation/operators/{country}.
func (r *Activation) GetOperators(ctx context.Context, params ActivationGetOperatorsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getOperators", params, options...)
}

type ActivationGetConversionsParams struct {
	Service *string "json:\"service,omitempty\"" // Required.
}

// GetConversions calls GET /activation/conversions/{service}.
func (r *Activation) GetConversions(ctx context.Context, params ActivationGetConversionsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getConversions", params, options...)
}

// GetMessageTypeSettings calls GET /activation/countryservicesettings.
func (r *Activation) GetMessageTypeSettings(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getMessageTypeSettings", nil, options...)
}

// GetAllNumberCounts calls GET /activation/countnumbers.
func (r *Activation) GetAllNumberCounts(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getAllNumberCounts", nil, options...)
}

type ActivationGetCountryNumberCountsParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
}

// GetCountryNumberCounts calls GET /activation/countnumbers/{country}.
func (r *Activation) GetCountryNumberCounts(ctx context.Context, params ActivationGetCountryNumberCountsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getCountryNumberCounts", params, options...)
}

type ActivationGetServicePriceParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Service *string "json:\"service,omitempty\"" // Required.
	Voice   *int64  "json:\"voice,omitempty\""   // Optional.
}

// GetServicePrice calls GET /activation/serviceprice/{country}/{service}.
func (r *Activation) GetServicePrice(ctx context.Context, params ActivationGetServicePriceParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getServicePrice", params, options...)
}

type ActivationGetCountryOperatorPricesParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Voice   *int64  "json:\"voice,omitempty\""   // Optional.
}

// GetCountryOperatorPrices calls GET /activation/serviceprice/{country}.
func (r *Activation) GetCountryOperatorPrices(ctx context.Context, params ActivationGetCountryOperatorPricesParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getCountryOperatorPrices", params, options...)
}

type ActivationGetServicePricesParams struct {
	Service *string "json:\"service,omitempty\"" // Required.
	Voice   *int64  "json:\"voice,omitempty\""   // Optional.
}

// GetServicePrices calls GET /activation/serviceprices/{service}.
func (r *Activation) GetServicePrices(ctx context.Context, params ActivationGetServicePricesParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getServicePrices", params, options...)
}

type ActivationGetAllServicePricesParams struct {
	Voice *int64 "json:\"voice,omitempty\"" // Optional.
}

// GetAllServicePrices calls GET /activation/servicesprices.
func (r *Activation) GetAllServicePrices(ctx context.Context, params ActivationGetAllServicePricesParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getAllServicePrices", params, options...)
}

type ActivationGetNumberParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Service *string "json:\"service,omitempty\"" // Required.
	Voice   *int64  "json:\"voice,omitempty\""   // Optional.
}

// GetNumber calls GET /activation/number/{country}/{service}.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Activation) GetNumber(ctx context.Context, params ActivationGetNumberParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getNumber", params, options...)
}

type ActivationGetNumberByOperatorParams struct {
	Country  *string "json:\"country,omitempty\""  // Required.
	Service  *string "json:\"service,omitempty\""  // Required.
	Operator *string "json:\"operator,omitempty\"" // Required.
	Voice    *int64  "json:\"voice,omitempty\""    // Optional.
}

// GetNumberByOperator calls GET /activation/number/{country}/{service}/{operator}.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Activation) GetNumberByOperator(ctx context.Context, params ActivationGetNumberByOperatorParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getNumberByOperator", params, options...)
}

type ActivationReceiveSmsParams struct {
	PartnerKey *string "json:\"partnerkey,omitempty\"" // Optional.
	OrderID    *int64  "json:\"orderid,omitempty\""    // Required.
}

// ReceiveSms calls GET /activation/sms/{orderid}.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Activation) ReceiveSms(ctx context.Context, params ActivationReceiveSmsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.receiveSms", params, options...)
}

// GetOrders calls GET /activation/orders.
func (r *Activation) GetOrders(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getOrders", nil, options...)
}

type ActivationGetNumberStatusParams struct {
	Number  *string "json:\"number,omitempty\""  // Required.
	Service *string "json:\"service,omitempty\"" // Required.
}

// GetNumberStatus calls GET /activation/numberstatus/{number}/{service}.
func (r *Activation) GetNumberStatus(ctx context.Context, params ActivationGetNumberStatusParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.getNumberStatus", params, options...)
}

type ActivationClearSmsParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
}

// ClearSms calls PUT /activation/clearsms/{orderid}.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Activation) ClearSms(ctx context.Context, params ActivationClearSmsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.clearSms", params, options...)
}

type ActivationStopSmsParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
}

// StopSms calls PUT /activation/stopsms/{orderid}.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Activation) StopSms(ctx context.Context, params ActivationStopSmsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.stopSms", params, options...)
}

type ActivationCancelOrderParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
}

// CancelOrder calls PUT /activation/cancelorder/{orderid}.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Activation) CancelOrder(ctx context.Context, params ActivationCancelOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.cancelOrder", params, options...)
}

type ActivationBlockNumberParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
}

// BlockNumber calls PUT /activation/blocknumber/{orderid}.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Activation) BlockNumber(ctx context.Context, params ActivationBlockNumberParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "activation.blockNumber", params, options...)
}

type Rental struct{ client *Client }

// GetCountries calls GET /api/rent.php?method=getcountries.
func (r *Rental) GetCountries(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getCountries", nil, options...)
}

// GetDefaultServices calls GET /api/rent.php?method=get_default_services.
func (r *Rental) GetDefaultServices(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getDefaultServices", nil, options...)
}

type RentalGetCountriesByServiceParams struct {
	Service *string "json:\"service,omitempty\"" // Required.
	Dtype   *string "json:\"dtype,omitempty\""   // Required.
	Dcount  *int64  "json:\"dcount,omitempty\""  // Optional.
}

// GetCountriesByService calls GET /api/rent.php?method=get_country_by_service.
func (r *Rental) GetCountriesByService(ctx context.Context, params RentalGetCountriesByServiceParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getCountriesByService", params, options...)
}

type RentalGetDataParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Dtype   *string "json:\"dtype,omitempty\""   // Optional.
	Dcount  *int64  "json:\"dcount,omitempty\""  // Optional.
	Extend  *string "json:\"extend,omitempty\""  // Optional.
}

// GetData calls GET /api/rent.php?method=getdata.
func (r *Rental) GetData(ctx context.Context, params RentalGetDataParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getData", params, options...)
}

type RentalGetDataWithProvidersParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Dtype   *string "json:\"dtype,omitempty\""   // Optional.
	Dcount  *int64  "json:\"dcount,omitempty\""  // Optional.
	Extend  *string "json:\"extend,omitempty\""  // Optional.
}

// GetDataWithProviders calls GET /api/rent.php?method=getdataWithProviders.
func (r *Rental) GetDataWithProviders(ctx context.Context, params RentalGetDataWithProvidersParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getDataWithProviders", params, options...)
}

type RentalGetCountParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Dtype   *string "json:\"dtype,omitempty\""   // Optional.
	Dcount  *int64  "json:\"dcount,omitempty\""  // Optional.
}

// GetCount calls GET /api/rent.php?method=getcount.
func (r *Rental) GetCount(ctx context.Context, params RentalGetCountParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getCount", params, options...)
}

type RentalGetCountWithProvidersParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Dtype   *string "json:\"dtype,omitempty\""   // Optional.
	Dcount  *int64  "json:\"dcount,omitempty\""  // Optional.
}

// GetCountWithProviders calls GET /api/rent.php?method=getCountWithProviders.
func (r *Rental) GetCountWithProviders(ctx context.Context, params RentalGetCountWithProvidersParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getCountWithProviders", params, options...)
}

type RentalGetMultiServiceCountParams struct {
	Country  *string "json:\"country,omitempty\""  // Required.
	Services *string "json:\"services,omitempty\"" // Required.
	Dtype    *string "json:\"dtype,omitempty\""    // Required.
	Dcount   *int64  "json:\"dcount,omitempty\""   // Required.
}

// GetMultiServiceCount calls GET /api/rent.php?method=get_count_multi.
func (r *Rental) GetMultiServiceCount(ctx context.Context, params RentalGetMultiServiceCountParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getMultiServiceCount", params, options...)
}

// GetCanadaPaypalDays calls GET /api/rent.php?method=get_days_canada_paypal.
func (r *Rental) GetCanadaPaypalDays(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getCanadaPaypalDays", nil, options...)
}

type RentalGetMinimumPaypalDaysParams struct {
	Country *string "json:\"country,omitempty\"" // Required.
	Service *string "json:\"service,omitempty\"" // Required.
}

// GetMinimumPaypalDays calls GET /api/rent.php?method=get_min_paypal_days.
func (r *Rental) GetMinimumPaypalDays(ctx context.Context, params RentalGetMinimumPaypalDaysParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getMinimumPaypalDays", params, options...)
}

type RentalCreateParams struct {
	Dtype      *string "json:\"dtype,omitempty\""       // Required.
	Dcount     *int64  "json:\"dcount,omitempty\""      // Required.
	Country    *string "json:\"country,omitempty\""     // Required.
	Service    *string "json:\"service,omitempty\""     // Required.
	Provider   *string "json:\"provider,omitempty\""    // Optional.
	PartnerKey *string "json:\"partner_key,omitempty\"" // Optional.
}

// Create calls GET /api/rent.php?method=create.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) Create(ctx context.Context, params RentalCreateParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.create", params, options...)
}

type RentalCreateMultiParams struct {
	Dtype      *string "json:\"dtype,omitempty\""       // Required.
	Dcount     *int64  "json:\"dcount,omitempty\""      // Required.
	Country    *string "json:\"country,omitempty\""     // Required.
	Services   *string "json:\"services,omitempty\""    // Required.
	Provider   *string "json:\"provider,omitempty\""    // Optional.
	PartnerKey *string "json:\"partner_key,omitempty\"" // Optional.
}

// CreateMulti calls GET /api/rent.php?method=create_multi.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) CreateMulti(ctx context.Context, params RentalCreateMultiParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.createMulti", params, options...)
}

type RentalAddServiceToOrderParams struct {
	ID       *int64  "json:\"id,omitempty\""        // Required.
	Service  *string "json:\"service,omitempty\""   // Required.
	Pnumber  *string "json:\"pnumber,omitempty\""   // Required.
	RentDays *int64  "json:\"rent_days,omitempty\"" // Optional.
}

// AddServiceToOrder calls GET /api/rent.php?method=add_service_to_order.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) AddServiceToOrder(ctx context.Context, params RentalAddServiceToOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.addServiceToOrder", params, options...)
}

// GetOrders calls GET /api/rent.php?method=orders.
func (r *Rental) GetOrders(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getOrders", nil, options...)
}

type RentalActivateParams struct {
	ID *int64 "json:\"id,omitempty\"" // Required.
}

// Activate calls GET /api/rent.php?method=activate.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) Activate(ctx context.Context, params RentalActivateParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.activate", params, options...)
}

type RentalGetSmsParams struct {
	ID *int64 "json:\"id,omitempty\"" // Required.
}

// GetSms calls GET /api/rent.php?method=sms.
func (r *Rental) GetSms(ctx context.Context, params RentalGetSmsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getSms", params, options...)
}

type RentalProlongParams struct {
	ID     *int64  "json:\"id,omitempty\""     // Required.
	Dtype  *string "json:\"dtype,omitempty\""  // Required.
	Dcount *int64  "json:\"dcount,omitempty\"" // Required.
}

// Prolong calls GET /api/rent.php?method=prolong.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) Prolong(ctx context.Context, params RentalProlongParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.prolong", params, options...)
}

type RentalProlongByDaysParams struct {
	ID     *int64  "json:\"id,omitempty\""     // Required.
	Dtype  *string "json:\"dtype,omitempty\""  // Required.
	Dcount *int64  "json:\"dcount,omitempty\"" // Required.
}

// ProlongByDays calls GET /api/rent.php?method=prolong_max.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) ProlongByDays(ctx context.Context, params RentalProlongByDaysParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.prolongByDays", params, options...)
}

type RentalDeleteOrderParams struct {
	ID *int64 "json:\"id,omitempty\"" // Required.
}

// DeleteOrder calls GET /api/rent.php?method=delete.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) DeleteOrder(ctx context.Context, params RentalDeleteOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.deleteOrder", params, options...)
}

type RentalGetHistoryParams struct {
	Skip *int64 "json:\"skip,omitempty\"" // Optional.
	Take *int64 "json:\"take,omitempty\"" // Optional.
}

// GetHistory calls GET /api/rent.php?method=get_rent_history.
func (r *Rental) GetHistory(ctx context.Context, params RentalGetHistoryParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getHistory", params, options...)
}

type RentalGetRestoreCostParams struct {
	ID *int64 "json:\"id,omitempty\"" // Required.
}

// GetRestoreCost calls GET /api/rent.php?method=restore_user_precalc.
func (r *Rental) GetRestoreCost(ctx context.Context, params RentalGetRestoreCostParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.getRestoreCost", params, options...)
}

type RentalRestoreParams struct {
	ID *int64 "json:\"id,omitempty\"" // Required.
}

// Restore calls GET /api/rent.php?method=restore_user.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Rental) Restore(ctx context.Context, params RentalRestoreParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "rental.restore", params, options...)
}

type Email struct{ client *Client }

// GetProviders calls GET /api/rent_email.php?method=get_providers.
func (r *Email) GetProviders(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getProviders", nil, options...)
}

type EmailGetDataParams struct {
	Pcode *string "json:\"pcode,omitempty\"" // Optional.
}

// GetData calls GET /api/rent_email.php?method=get_data.
func (r *Email) GetData(ctx context.Context, params EmailGetDataParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getData", params, options...)
}

type EmailGetCountParams struct {
	Pcode *string "json:\"pcode,omitempty\"" // Optional.
}

// GetCount calls GET /api/rent_email.php?method=get_count.
func (r *Email) GetCount(ctx context.Context, params EmailGetCountParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getCount", params, options...)
}

type EmailCreateOrderParams struct {
	Scode  *string "json:\"scode,omitempty\""  // Required.
	Pcode  *string "json:\"pcode,omitempty\""  // Optional.
	Dtype  *string "json:\"dtype,omitempty\""  // Required.
	Dcount *int64  "json:\"dcount,omitempty\"" // Required.
}

// CreateOrder calls GET /api/rent_email.php?method=create_order.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Email) CreateOrder(ctx context.Context, params EmailCreateOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.createOrder", params, options...)
}

// GetOrders calls GET /api/rent_email.php?method=get_orders.
func (r *Email) GetOrders(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getOrders", nil, options...)
}

type EmailGetMailsParams struct {
	OrderID *int64  "json:\"orderid,omitempty\"" // Required.
	T       *int64  "json:\"t,omitempty\""       // Optional.
	S       *int64  "json:\"s,omitempty\""       // Optional.
	Sq      *string "json:\"sq,omitempty\""      // Optional.
}

// GetMails calls GET /api/rent_email.php?method=get_mails.
func (r *Email) GetMails(ctx context.Context, params EmailGetMailsParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getMails", params, options...)
}

type EmailGetMailBodyParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
	Emailid *int64 "json:\"emailid,omitempty\"" // Required.
}

// GetMailBody calls GET /api/rent_email.php?method=get_mail_body.
func (r *Email) GetMailBody(ctx context.Context, params EmailGetMailBodyParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getMailBody", params, options...)
}

type EmailBuyAliasParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
	Aliasid *int64 "json:\"aliasid,omitempty\"" // Required.
}

// BuyAlias calls GET /api/rent_email.php?method=buy_alias.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Email) BuyAlias(ctx context.Context, params EmailBuyAliasParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.buyAlias", params, options...)
}

type EmailGetMailsByAliasParams struct {
	OrderID *int64  "json:\"orderid,omitempty\"" // Required.
	Aliasid *int64  "json:\"aliasid,omitempty\"" // Required.
	T       *int64  "json:\"t,omitempty\""       // Optional.
	S       *int64  "json:\"s,omitempty\""       // Optional.
	Sq      *string "json:\"sq,omitempty\""      // Optional.
}

// GetMailsByAlias calls GET /api/rent_email.php?method=get_mails_by_alias.
func (r *Email) GetMailsByAlias(ctx context.Context, params EmailGetMailsByAliasParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getMailsByAlias", params, options...)
}

type EmailProlongOrderParams struct {
	OrderID *int64  "json:\"orderid,omitempty\"" // Required.
	Dtype   *string "json:\"dtype,omitempty\""   // Required.
	Dcount  *int64  "json:\"dcount,omitempty\""  // Required.
}

// ProlongOrder calls GET /api/rent_email.php?method=prolong_order.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Email) ProlongOrder(ctx context.Context, params EmailProlongOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.prolongOrder", params, options...)
}

type EmailCancelOrderParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
}

// CancelOrder calls GET /api/rent_email.php?method=cancel_order.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Email) CancelOrder(ctx context.Context, params EmailCancelOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.cancelOrder", params, options...)
}

// GetArchivedOrders calls GET /api/rent_email.php?method=get_archived_orders.
func (r *Email) GetArchivedOrders(ctx context.Context, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.getArchivedOrders", nil, options...)
}

type EmailCalcRestoreOrderParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
}

// CalcRestoreOrder calls GET /api/rent_email.php?method=calc_restore_order.
func (r *Email) CalcRestoreOrder(ctx context.Context, params EmailCalcRestoreOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.calcRestoreOrder", params, options...)
}

type EmailRestoreOrderParams struct {
	OrderID *int64 "json:\"orderid,omitempty\"" // Required.
}

// RestoreOrder calls GET /api/rent_email.php?method=restore_order.
// This operation changes state or may charge the account. It is never retried by the SDK.
func (r *Email) RestoreOrder(ctx context.Context, params EmailRestoreOrderParams, options ...RequestOption) (Response, error) {
	return r.client.call(ctx, "email.restoreOrder", params, options...)
}
