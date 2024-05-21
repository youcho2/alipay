package alipay

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	kSuccess = []byte("success")
)

func NewRequest(method, url string, values url.Values) (*http.Request, error) {
	method = strings.ToUpper(method)
	var body io.Reader
	if method == http.MethodGet || method == http.MethodHead {
		if len(values) > 0 {
			if strings.Contains(url, "?") {
				url = url + "&" + values.Encode()
			} else {
				url = url + "?" + values.Encode()
			}
		}
	} else {
		body = strings.NewReader(values.Encode())
	}
	return http.NewRequest(method, url, body)
}

func (this *Client) NotifyVerify(partnerId, notifyId string) bool {
	var values = url.Values{}
	values.Add("service", "notify_verify")
	values.Add("partner", partnerId)
	values.Add("notify_id", notifyId)
	req, err := NewRequest(http.MethodGet, this.notifyVerifyHost, values)
	if err != nil {
		return false
	}

	rsp, err := this.Client.Do(req)
	if err != nil {
		return false
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return false
	}
	if string(data) == "true" {
		return true
	}
	return false
}

func (this *Client) GetTradeNotification(req *http.Request) (notification *TradeNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notification = &TradeNotification{}
	notification.AppId = req.FormValue("app_id")
	notification.AuthAppId = req.FormValue("auth_app_id")
	notification.NotifyId = req.FormValue("notify_id")
	notification.NotifyType = req.FormValue("notify_type")
	notification.NotifyTime = req.FormValue("notify_time")
	notification.TradeNo = req.FormValue("trade_no")
	notification.TradeStatus = TradeStatus(req.FormValue("trade_status"))
	notification.TotalAmount = req.FormValue("total_amount")
	notification.ReceiptAmount = req.FormValue("receipt_amount")
	notification.InvoiceAmount = req.FormValue("invoice_amount")
	notification.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	notification.SellerId = req.FormValue("seller_id")
	notification.SellerEmail = req.FormValue("seller_email")
	notification.BuyerId = req.FormValue("buyer_id")
	notification.BuyerLogonId = req.FormValue("buyer_logon_id")
	notification.FundBillList = req.FormValue("fund_bill_list")
	notification.Charset = req.FormValue("charset")
	notification.PointAmount = req.FormValue("point_amount")
	notification.OutTradeNo = req.FormValue("out_trade_no")
	notification.OutBizNo = req.FormValue("out_biz_no")
	notification.GmtCreate = req.FormValue("gmt_create")
	notification.GmtPayment = req.FormValue("gmt_payment")
	notification.GmtRefund = req.FormValue("gmt_refund")
	notification.GmtClose = req.FormValue("gmt_close")
	notification.Subject = req.FormValue("subject")
	notification.Body = req.FormValue("body")
	notification.RefundFee = req.FormValue("refund_fee")
	notification.Version = req.FormValue("version")
	notification.SignType = req.FormValue("sign_type")
	notification.Sign = req.FormValue("sign")
	notification.PassbackParams = req.FormValue("passback_params")
	notification.VoucherDetailList = req.FormValue("voucher_detail_list")
	notification.AgreementNo = req.FormValue("agreement_no")
	notification.ExternalAgreementNo = req.FormValue("external_agreement_no")

	//if len(noti.NotifyId) == 0 {
	//	return nil, errors.New("不是有效的 Notify")
	//}

	if err = this.VerifySign(req.Form); err != nil {
		return nil, err
	}
	return notification, err
}

// AckNotification
// Deprecated: use ACKNotification instead.
func (this *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

// ACKNotification 返回异步通知成功处理的消息给支付宝
func (this *Client) ACKNotification(w http.ResponseWriter) {
	ACKNotification(w)
}

// AckNotification
// Deprecated: use ACKNotification instead.
func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}

// ACKNotification 返回异步通知成功处理的消息给支付宝
func ACKNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}
