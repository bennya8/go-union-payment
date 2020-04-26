# Alipay 配置说明

| KEY | VALUE说明 |
| --- | --- |
| use_sandbox | 是否使用沙盒模式 |
| app_id | 支付宝应用appId |
| sign_type | 验签方式、RSA、RSA2 |
| ali_public_key | 支付宝公钥字符串 |
| rsa_private_key | 自己生成的密钥字符串 |
| limit_pay | balance（余额）、moneyFund（余额宝）、debitCardExpress（借记卡快捷）、<br>creditCard（信用卡）、creditCardExpress（信用卡快捷）、<br>creditCardCartoon（信用卡卡通）、credit_group（信用支付类型（包含信用卡卡通、信用卡快捷、花呗、花呗分期）<br> 用户不可用指定渠道支付当有多个渠道时用“,”分隔
| notify_url | 与业务相关参数 |
| return_url | 与业务相关参数 |
| fee_type | 货币类型  当前仅支持该字段 CNY |