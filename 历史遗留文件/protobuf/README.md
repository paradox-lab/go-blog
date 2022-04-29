# protobuf
> Protocol Buffers (a.k.a., protobuf) are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data.

**github地址**: https://github.com/protocolbuffers/protobuf

这里我们把标准类型列举一下protobuf内置的标准类型以及跟各平台对应的关系：
```text
.proto | 说明	C++	Java	Python	Go	Ruby	C#	PHP
double | double	double	float	float64	Float	double	float
float | float	float	float	float32	Float	float	float
int32 |	使用变长编码，对负数编码效率低，如果你的变量可能是负数，可以使用sint32	int32	int	int	int32	Fixnum or Bignum (as required)	int	integer
int64 | 使用变长编码，对负数编码效率低，如果你的变量可能是负数，可以使用sint64	int64	long	int/long	int64	Bignum	long	integer/string
uint32 | 使用变长编码	uint32	int	int/long	uint32	Fixnum or Bignum (as required)	uint	integer
uint64 | 使用变长编码	uint64	long	int/long	uint64	Bignum	ulong	integer/string
sint32 | 使用变长编码，带符号的int类型，对负数编码比int32高效	int32	int	int	int32	Fixnum or Bignum (as required)	int	integer
sint64 | 使用变长编码，带符号的int类型，对负数编码比int64高效	int64	long	int/long	int64	Bignum	long	integer/string
fixed32 | 4字节编码， 如果变量经常大于 的话，会比uint64高效	uint64	long	int/long	uint64	Bignum	ulong	integer/string
sfixed32 | 4字节编码	int32	int	int	int32	Fixnum or Bignum (as required)	int	integer
sfixed64 | 8字节编码	int64	long	int/long	int64	Bignum	long	integer/string
bool | bool	boolean	bool	bool	TrueClass/FalseClass	bool	boolean
string | 必须包含utf-8编码或者7-bit ASCII text	string	String	str/unicode	string	String (UTF-8)	string	string
bytes |	任意的字节序列	string	ByteString	str	[]byte	String (ASCII-8BIT)	ByteString	string
```

## 在.proto文件定义服务
### 指定命名服务
> 要定义一个服务，你需要在你的.proto文件中指定一个命名服务
```text
service RouteGuide {
   // (Method definitions not shown)
}
```
### 定义rpc方法
> 指定它们的请求和响应类型。
> gRPC允许你定义四种服务方法。

#### 一个简单的RPC 
> 客户端使用存根向服务器发送请求并等待响应返回，就像普通的函数调用一样。
```text
// Obtains the feature at a given position.
rpc GetFeature(Point) returns (Feature) {}
```

#### 响应流RPC
> 客户端向服务器发送请求并获取流来读取消息序列的响应流RPC。客户端从返回的流中读取，直到没有更多的消息为止。
> 正如您在示例中看到的，您可以通过在响应类型之前放置stream关键字来指定一个响应流方法。
```text
// Obtains the Features available within the given Rectangle.  Results are
// streamed rather than returned at once (e.g. in a response message with a
// repeated field), as the rectangle may cover a large area and contain a
// huge number of features.
rpc ListFeatures(Rectangle) returns (stream Feature) {}
```

#### 请求流RPC
> 其中客户端编写消息序列并将它们发送到服务器，同样使用提供的流。一旦客户端完成了消息的写入，它将等待服务器读取所有消息并返回响应。
> 您可以通过在请求类型之前放置stream关键字来指定请求流方法。
```text
// Accepts a stream of Points on a route being traversed, returning a
// RouteSummary when traversal is completed.
rpc RecordRoute(stream Point) returns (RouteSummary) {}
```

#### 双向流RPC
> 双向流RPC，其中双方使用读写流发送消息序列。两个流独立运作,因此客户端和服务器可以读和写在他们喜欢的任何顺序:例如,
> 服务器可以等待收到所有客户端消息之前写的反应,也可以交替阅读一条消息然后写一个消息,或其他一些读写的结合。每个流中的消息顺序保持不变。
> 您可以通过在请求和响应之前放置stream关键字来指定这种类型的方法

> 你的.proto文件也包含协议缓冲区消息类型定义，用于我们的服务方法中使用的所有请求和响应类型——例如，这里是点消息类型:
```text
// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
message Point {
  int32 latitude = 1;
  int32 longitude = 2;
} 
```