# http

## http 报文结构

```go
/*
A HTTP Request

Request-Line
GET / HTTP/1.1

Request-Header     
Host: www.baidu.com
Connection: keep-alive
Cache-Control: max-age=0 ...

then request body
*/

```

## http包变量解析

```go
/*
Request代表了一个HTTP请求
1. 服务器接收的 received by a server
2. 客户端待发送的 to be sent

Request中的字段在客户端和服务器端使用时语义略有不同
除了下面字段记录的field
还要看文档中的 Request.Write & RoundTripper
*/

type request struct {

	/* 
	精确说明了 HTTP Method
	request from client, Method == "" means GET
	Go request 不支持 CONNECT
	*/
	Method string

	/*
	URL 精确说明了
	1. 对于服务器端来说，是被请求的URL，即请求是来自
	哪个URL
	2. 对于客户端来说，是要请求的URL

	对服务器请求来说，URL可以从RequestURL
	中的Request-Line解析出来
	除了Path 和 RawQuery 其他字段基本为空
	*/

	URL *url.URL

	/*
	type Header map[string][]string
	

	*/

	Header Header
}

```