package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1739527022)
	ginrpc.AddGenOne("Weixin.GetQrcode", "weixin.get_qrcode", []string{"post"}, []ginrpc.GenThirdParty{}, `获取微信二维码`)
	ginrpc.AddGenOne("Weixin.Oauth", "weixin.oauth", []string{"post"}, []ginrpc.GenThirdParty{}, `微信授权获取登录信息`)
	ginrpc.AddGenOne("Weixin.RefundPay", "weixin.refund_pay", []string{"post"}, []ginrpc.GenThirdParty{}, `退款`)
	ginrpc.AddGenOne("Weixin.UpdateUserInfo", "weixin.update_user_info", []string{"post"}, []ginrpc.GenThirdParty{}, `更新用户信息`)
	ginrpc.AddGenOne("Upload.UploadMult", "/upload", []string{"post"}, []ginrpc.GenThirdParty{}, `上传文件 多个`)
	ginrpc.AddGenOne("Login.Login", "/login/login", []string{"POST"}, []ginrpc.GenThirdParty{}, `登录`)
	ginrpc.AddGenOne("Login.RefreshToken", "login.refresh_token", []string{"post"}, []ginrpc.GenThirdParty{}, `刷新token管理员`)
	ginrpc.AddGenOne("Product.AddAddress", "product.add_address", []string{"post"}, []ginrpc.GenThirdParty{}, `添加或者修改地址`)
	ginrpc.AddGenOne("Product.AddToBuyTmpCart", "product.add_to_buy_tmp_cart", []string{"post"}, []ginrpc.GenThirdParty{}, `添加到直接购买虚拟购物车里面`)
	ginrpc.AddGenOne("Product.AddToCart", "product.add_to_cart", []string{"post"}, []ginrpc.GenThirdParty{}, `加入购物车`)
	ginrpc.AddGenOne("Product.ChangeCart", "product.change_cart", []string{"post"}, []ginrpc.GenThirdParty{}, `改变`)
	ginrpc.AddGenOne("Product.Favorite", "product.favorite", []string{"post"}, []ginrpc.GenThirdParty{}, `收藏 or 取消收藏`)
	ginrpc.AddGenOne("Product.GetAdInfo", "product.get_ad_info", []string{"post"}, []ginrpc.GenThirdParty{}, `获取广告主信息列表(可能更新)`)
	ginrpc.AddGenOne("Product.GetAddress", "product.get_address", []string{"post"}, []ginrpc.GenThirdParty{}, `获取地址`)
	ginrpc.AddGenOne("Product.GetAdminAdInfo", "product.get_admin_ad_info", []string{"post"}, []ginrpc.GenThirdParty{}, `admin 获取首页信息(轮播图广告，类型广告，主销广告,精选列表)`)
	ginrpc.AddGenOne("Product.GetBoutiqueGroup", "product.get_boutique_group", []string{"post"}, []ginrpc.GenThirdParty{}, `获取团购列表`)
	ginrpc.AddGenOne("Product.GetCartList", "product.get_cart_list", []string{"post"}, []ginrpc.GenThirdParty{}, `获取购物车内容`)
	ginrpc.AddGenOne("Product.GetFavorite", "product.get_favorite", []string{"post"}, []ginrpc.GenThirdParty{}, `获取收藏列表`)
	ginrpc.AddGenOne("Product.GetProductByType", "product.get_product_by_type", []string{"post"}, []ginrpc.GenThirdParty{}, `通过商品类型获取商品列表`)
	ginrpc.AddGenOne("Product.GetProductDetails", "product.get_product_details", []string{"post"}, []ginrpc.GenThirdParty{}, `获取订单详情`)
	ginrpc.AddGenOne("Product.GetProductList", "product.get_product_list", []string{"post"}, []ginrpc.GenThirdParty{}, `admin 获取商品列表`)
	ginrpc.AddGenOne("Product.ReplaceProduct", "product.replace_product", []string{"post"}, []ginrpc.GenThirdParty{}, `添加一个产品(可能更新)`)
	ginrpc.AddGenOne("Product.ReplaceSku", "product.replace_sku", []string{"post"}, []ginrpc.GenThirdParty{}, `添加一个产品Sku(可能更新)`)
	ginrpc.AddGenOne("Product.UpdateSelect", "product.update_select", []string{"post"}, []ginrpc.GenThirdParty{}, `修改商品推荐信息`)
	ginrpc.AddGenOne("Product.UpsetAdminAdInfo", "product.upset_admin_ad_info", []string{"post"}, []ginrpc.GenThirdParty{}, `admin 更新首页信息(轮播图广告，类型广告，主销广告,精选列表)(id为0,添加，id>0 删除)`)
	ginrpc.AddGenOne("Product.UpsetSkuPrice", "product.upset_sku_price", []string{"post"}, []ginrpc.GenThirdParty{}, `更新/设置产品sku价格`)
	ginrpc.AddGenOne("Order.AddBillShipment", "order.add_bill_shipment", []string{"post"}, []ginrpc.GenThirdParty{}, `订单添加运单号`)
	ginrpc.AddGenOne("Order.CheckPay", "order.check_pay", []string{"post"}, []ginrpc.GenThirdParty{}, `支付成功check一次`)
	ginrpc.AddGenOne("Order.DealOrder", "order.deal_order", []string{"post"}, []ginrpc.GenThirdParty{}, `账单/订单处理`)
	ginrpc.AddGenOne("Order.DealSystem", "order.deal_system", []string{"post"}, []ginrpc.GenThirdParty{}, `意见与反馈`)
	ginrpc.AddGenOne("Order.GetAfterMsg", "order.get_after_msg", []string{"post"}, []ginrpc.GenThirdParty{}, `获取回复信息`)
	ginrpc.AddGenOne("Order.GetBillInfoByAdmin", "order.get_bill_info_by_admin", []string{"post"}, []ginrpc.GenThirdParty{}, `获取订单信息`)
	ginrpc.AddGenOne("Order.GetMyOrders", "order.get_my_orders", []string{"post"}, []ginrpc.GenThirdParty{}, `获取所有订单`)
	ginrpc.AddGenOne("Order.GetMyOrdersConfig", "order.get_my_orders_config", []string{"post"}, []ginrpc.GenThirdParty{}, `获取订单配置`)
	ginrpc.AddGenOne("Order.GetOrdertrackInfo", "order.get_ordertrack_info", []string{"post"}, []ginrpc.GenThirdParty{}, `获取物流轨迹信息`)
	ginrpc.AddGenOne("Order.GetShipmentPost", "order.get_shipment_post", []string{"post"}, []ginrpc.GenThirdParty{}, `获取快递列表`)
	ginrpc.AddGenOne("Order.Pay", "order.pay", []string{"post"}, []ginrpc.GenThirdParty{}, `支付`)
	ginrpc.AddGenOne("Order.PlaceOrder", "order.place_order", []string{"post"}, []ginrpc.GenThirdParty{}, `下单`)
	ginrpc.AddGenOne("Order.ReckonFee", "order.reckon_fee", []string{"post"}, []ginrpc.GenThirdParty{}, `费用计算`)
	ginrpc.AddGenOne("Coupon.GetAdminCouponInfo", "coupon.get_admin_coupon_info", []string{"post"}, []ginrpc.GenThirdParty{}, `admin 获取优惠券列表`)
	ginrpc.AddGenOne("Coupon.GetMyCoupon", "coupon.get_my_coupon", []string{"post"}, []ginrpc.GenThirdParty{}, `获取我的优惠券`)
	ginrpc.AddGenOne("Coupon.GetPromotionCoupon", "coupon.get_promotion_coupon", []string{"post"}, []ginrpc.GenThirdParty{}, `获取促销优惠券`)
	ginrpc.AddGenOne("Coupon.GoGetCoupon", "coupon.go_get_coupon", []string{"post"}, []ginrpc.GenThirdParty{}, `一键领取优惠券`)
	ginrpc.AddGenOne("Coupon.UpsetAdminCouponInfo", "coupon.upset_admin_coupon_info", []string{"post"}, []ginrpc.GenThirdParty{}, `admin 更新/添加 优惠券列表`)
	ginrpc.AddGenOne("User.GetUserInfo", "user.get_user_info", []string{"post"}, []ginrpc.GenThirdParty{}, `获取用户相关信息`)
	ginrpc.AddGenOne("User.LinkUser", "user.link_user", []string{"post"}, []ginrpc.GenThirdParty{}, `关联用户`)
	ginrpc.AddGenOne("User.RequestVip", "user.request_vip", []string{"post"}, []ginrpc.GenThirdParty{}, `申请vip`)
	ginrpc.AddGenOne("Dist.GetDistList", "dist.get_dist_list", []string{"post"}, []ginrpc.GenThirdParty{}, `获取可分销列表`)
	ginrpc.AddGenOne("Dist.RequestDist", "dist.request_dist", []string{"post"}, []ginrpc.GenThirdParty{}, `申请分销`)
}
