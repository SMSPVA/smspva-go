// Code generated from sdk/contracts.json. DO NOT EDIT.
package smspva

import (
	"context"
	"encoding/json"
)

var namedCalls = map[string]func(context.Context, *Client, json.RawMessage) (Response, error){
	"activation.getBalance": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Activation.GetBalance(ctx)
	},
	"activation.getUserInfo": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Activation.GetUserInfo(ctx)
	},
	"activation.getOperators": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetOperatorsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetOperators(ctx, p)
	},
	"activation.getConversions": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetConversionsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetConversions(ctx, p)
	},
	"activation.getMessageTypeSettings": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Activation.GetMessageTypeSettings(ctx)
	},
	"activation.getAllNumberCounts": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Activation.GetAllNumberCounts(ctx)
	},
	"activation.getCountryNumberCounts": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetCountryNumberCountsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetCountryNumberCounts(ctx, p)
	},
	"activation.getServicePrice": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetServicePriceParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetServicePrice(ctx, p)
	},
	"activation.getCountryOperatorPrices": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetCountryOperatorPricesParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetCountryOperatorPrices(ctx, p)
	},
	"activation.getServicePrices": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetServicePricesParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetServicePrices(ctx, p)
	},
	"activation.getAllServicePrices": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetAllServicePricesParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetAllServicePrices(ctx, p)
	},
	"activation.getNumber": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetNumberParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetNumber(ctx, p)
	},
	"activation.getNumberByOperator": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetNumberByOperatorParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetNumberByOperator(ctx, p)
	},
	"activation.receiveSms": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationReceiveSmsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.ReceiveSms(ctx, p)
	},
	"activation.getOrders": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Activation.GetOrders(ctx)
	},
	"activation.getNumberStatus": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationGetNumberStatusParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.GetNumberStatus(ctx, p)
	},
	"activation.clearSms": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationClearSmsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.ClearSms(ctx, p)
	},
	"activation.stopSms": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationStopSmsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.StopSms(ctx, p)
	},
	"activation.cancelOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationCancelOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.CancelOrder(ctx, p)
	},
	"activation.blockNumber": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p ActivationBlockNumberParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Activation.BlockNumber(ctx, p)
	},
	"rental.getCountries": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Rental.GetCountries(ctx)
	},
	"rental.getDefaultServices": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Rental.GetDefaultServices(ctx)
	},
	"rental.getCountriesByService": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetCountriesByServiceParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetCountriesByService(ctx, p)
	},
	"rental.getData": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetDataParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetData(ctx, p)
	},
	"rental.getDataWithProviders": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetDataWithProvidersParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetDataWithProviders(ctx, p)
	},
	"rental.getCount": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetCountParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetCount(ctx, p)
	},
	"rental.getCountWithProviders": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetCountWithProvidersParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetCountWithProviders(ctx, p)
	},
	"rental.getMultiServiceCount": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetMultiServiceCountParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetMultiServiceCount(ctx, p)
	},
	"rental.getCanadaPaypalDays": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Rental.GetCanadaPaypalDays(ctx)
	},
	"rental.getMinimumPaypalDays": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetMinimumPaypalDaysParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetMinimumPaypalDays(ctx, p)
	},
	"rental.create": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalCreateParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.Create(ctx, p)
	},
	"rental.createMulti": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalCreateMultiParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.CreateMulti(ctx, p)
	},
	"rental.addServiceToOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalAddServiceToOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.AddServiceToOrder(ctx, p)
	},
	"rental.getOrders": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Rental.GetOrders(ctx)
	},
	"rental.activate": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalActivateParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.Activate(ctx, p)
	},
	"rental.getSms": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetSmsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetSms(ctx, p)
	},
	"rental.prolong": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalProlongParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.Prolong(ctx, p)
	},
	"rental.prolongByDays": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalProlongByDaysParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.ProlongByDays(ctx, p)
	},
	"rental.deleteOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalDeleteOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.DeleteOrder(ctx, p)
	},
	"rental.getHistory": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetHistoryParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetHistory(ctx, p)
	},
	"rental.getRestoreCost": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalGetRestoreCostParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.GetRestoreCost(ctx, p)
	},
	"rental.restore": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p RentalRestoreParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Rental.Restore(ctx, p)
	},
	"email.getProviders": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Email.GetProviders(ctx)
	},
	"email.getData": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailGetDataParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.GetData(ctx, p)
	},
	"email.getCount": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailGetCountParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.GetCount(ctx, p)
	},
	"email.createOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailCreateOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.CreateOrder(ctx, p)
	},
	"email.getOrders": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Email.GetOrders(ctx)
	},
	"email.getMails": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailGetMailsParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.GetMails(ctx, p)
	},
	"email.getMailBody": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailGetMailBodyParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.GetMailBody(ctx, p)
	},
	"email.buyAlias": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailBuyAliasParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.BuyAlias(ctx, p)
	},
	"email.getMailsByAlias": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailGetMailsByAliasParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.GetMailsByAlias(ctx, p)
	},
	"email.prolongOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailProlongOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.ProlongOrder(ctx, p)
	},
	"email.cancelOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailCancelOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.CancelOrder(ctx, p)
	},
	"email.getArchivedOrders": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		return c.Email.GetArchivedOrders(ctx)
	},
	"email.calcRestoreOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailCalcRestoreOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.CalcRestoreOrder(ctx, p)
	},
	"email.restoreOrder": func(ctx context.Context, c *Client, raw json.RawMessage) (Response, error) {
		var p EmailRestoreOrderParams
		if err := json.Unmarshal(raw, &p); err != nil {
			return Response{}, err
		}
		return c.Email.RestoreOrder(ctx, p)
	},
}
