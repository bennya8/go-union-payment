# Cmb 配置说明

| KEY | VALUE说明 |
| --- | --- |
| use_sandbox | 是否使用 招商测试系统 |
| branch_no | 商户分行号，4位数字 |
| mch_id | 商户号，6位数字 |
| mer_key | 秘钥16位，包含大小写字母 数字 |
| cmb_pub_key | 招商的公钥，建议每天凌晨2:15发起查询招行公钥请求更新公钥。|
| op_pwd | 操作员登录密码 |
| sign_type |  签名算法,固定为“SHA-256” |
| limit_pay | 允许支付的卡类型,默认对支付卡种不做限制，储蓄卡和信用卡均可支付   A:储蓄卡支付，即禁止信用卡支付|
| notify_url | 支付成功的回调 |
| sign_notify_url | 成功签约结果通知地址 |
| sign_return_url | 成功签约结果通知地址 |
| return_url | 如果是h5支付，可以设置该值，返回到指定页面 |