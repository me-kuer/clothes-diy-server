# 用户表
create table users(
    id int auto_increment primary key,
    head_pic varchar(200) not null default "" comment "头像",
    nickname varchar(50) default "" comment "昵称",
    openid varchar(50) not null default "" comment "openid",
    unionid varchar(50) not null default "" comment "unionid",
    session_key varchar(100) not null default "" comment "session_key",
    register_time varchar(11) default "0" comment "注册时间"
);

# 图案表
create table picture(
    id int auto_increment primary key,
    src varchar(120) default "" comment "路径",
    name varchar(90) default "" comment "名称",
    create_time varchar(11) default "0" comment "添加时间"
);

# 商品表
create table goods(
    id int auto_increment primary key,
    name varchar(30) default "" comment "商品名称",
    price float(6,2) unsigned default 0 comment "商品价格",
    status tinyint(1) unsigned default 1 comment "状态（1上架，0下架）",
    create_time varchar(11) default "0" comment "添加时间"
);

# 商品图片表
create table goods_img(
    id int auto_increment primary key,
    goods_id int not null default 0 comment "商品id",
    front varchar(100) default "" comment "正面图片路径",
    contrary varchar(100) default "" comment "反面图片路径",
    color varchar(30) default "" comment "颜色"
);
# 订单表
create table orders(
    id int auto_increment primary key,
    user_id int not null default 0 comment "用户id",
    goods_id int not null default 0 comment "商品id",
    front varchar(100) default "" comment "正面图片路径",
    contrary varchar(100) default "" comment "反面图片路径",
    color varchar(30) default "" comment "颜色",
    size varchar(1) default "" comment "尺码",
    out_trade_no varchar(64) default "" comment "订单号",
    transaction_id varchar(64) default "" comment "微信商户平台返回的订单号",
    total float(6,2) unsigned default 0 comment "金额",
    description text comment "微信支付备注",
    consignee_name varchar(60) default "" comment "收货人姓名",
    consignee_tel varchar(11) default "" comment "收货人手机号",
    consignee_province varchar(90) default "" comment "省",
    consignee_city varchar(90) default "" comment "省",
    consignee_region varchar(90) default "" comment "省",
    consignee_detail text comment "详细地址",
    note text comment "备注信息",
    pay_time varchar(11) default "0" comment "支付时间",
    dispatchin_time varchar(11) default "0" comment "配送时间",
    complete_time varchar(11) default "0" comment "订单完成时间",
    status tinyint(1) default 0 comment "订单状态(0未付款, 1已付款, 2配送中, 3交易完成, 4交易关闭)"
);

# 管理员表
create table admin(
    id int auto_increment primary key,
    username varchar(20) not null default "" comment "用户名",
    password varchar(32) not null default "" comment "密码",
    last_login varchar(11) not null default "0" comment "上次登录时间",
    current_login varchar(11) not null default "0" comment "当前登录时间"
);

# 地址表
create table address(
    id int auto_increment primary key,
    name varchar(60) not null default "" comment "收货人姓名",
    tel varchar(11) not null default "" comment "收货人手机号",
    province varchar(90) default "" comment "省",
    city varchar(90) default "" comment "市",
    region varchar(90) default "" comment "区",
    detail text comment "详细地址",
    sex tinyint(1) unsigned default 1 comment "性别(1男，0女)"
);