# converter generator

> designed for golang

# 背景

在你手写结构转换方法时, "你经历过绝望么" ？在落地ddd的过程中，由于需要在定义一堆VO、DO、DTO、PO，但是在全网找了一圈，在go生态中并不能找到一个像java中lombok或者BeanUtil这种可以自动进行结构转换的工具库，对于一些简单结构的转换，我们可以自己写转换方法，但是对于一些复杂结构，我们手写转换方法将会是一件及其痛苦的事情。所以想提供一个工具能自动生成两个结构之间的转换代码。

# 效果

例如以下的结构体

```go
type A struct {
   Str1              string
   Str2              *string
   Int               int
   Int1              *int
   StringStringMap1  map[string]string
   StringStringMap2 map[string]*string
}

type B struct {
   Str1              string
   Str2              string
   Int               int64
   Int1              *int64
   StringStringMap1  map[*string]string
   StringStringMap2 map[string]*string
}

type C struct {
   StructB B
   Data    map[string]A
}

type D struct {
   StructA A
   Data    map[string]B
}
```

通过 converter generator自动生成从D到C的结构转换方法如下

```cassandraql
package converts

import (
   test_data "github.com/jingleWang/converter_generator/test/data"
)

func ConvertDToC(in test_data.D) test_data.C {
   return test_data.C{
      Data: convertStringBMapToStringAMap(in.Data),
   }
}

func convertInt64PtrToInt(in *int64) int {
   if in == nil {
      return 0
   }
   return int(*in)
}

func convertStringBMapToStringAMap(in map[string]test_data.B) map[string]test_data.A {
   if in == nil {
      return nil
   }

   out := map[string]test_data.A{}
   for k, v := range in {
      key := k
      val := convertBToA(v)
      out[key] = val
   }
   return out
}

func convertStringToStringPtr(in string) *string {
   return &in
}

func convertInt64ToInt(in int64) int {

   return int(in)
}

func convertInt64PtrToIntPtr(in *int64) *int {
   if in == nil {
      return nil
   }
   out := convertInt64PtrToInt(in)
   return &out
}

func convertStringPtrToString(in *string) string {
   if in == nil {
      return ""
   }
   return string(*in)
}

func convertStringPtrStringMapToStringStringMap(in map[*string]string) map[string]string {
   if in == nil {
      return nil
   }

   out := map[string]string{}
   for k, v := range in {
      key := convertStringPtrToString(k)
      val := v
      out[key] = val
   }
   return out
}

func convertBToA(in test_data.B) test_data.A {
   return test_data.A{
      Str1:             in.Str1,
      Str2:             convertStringToStringPtr(in.Str2),
      Int:              convertInt64ToInt(in.Int),
      Int1:             convertInt64PtrToIntPtr(in.Int1),
      StringStringMap1: convertStringPtrStringMapToStringStringMap(in.StringStringMap1),
      StringStringMap2: in.StringStringMap2,
   }
}

```

conventer generator可以支持一些复杂结构转换的代码生成

# 用法

## Installation

```
go get github.com/jingleWang/converter_generator
```

## 编写脚本 scripts.go

```go
func main() {
   cg := converter_generator.ConverterGenerator{}
   cg.OutputDir("converts")  // 指定输出目录（必须）
   cg.PkgName("converts")    // 指定包名（必须）
   cg.FileName("xxx.go")     // 输出文件名（可选）

   // 指定需要转换的结构
   // D => C
   cg.Convert(data.D{}, data.C{}, converter_generator.ConversionConfig{
      Export: true,
   })
    // C=>D
    cg.Convert(data.C{}, data.D{}, converter_generator.ConversionConfig{
      Export: true,
   })

   cg.Generate()
}

```

## 运行脚本

```shell script
go run scripts.go
```

# 高级用法

## 结构别名指定

当你需要转换的两个结构拥有相同的结构名时,例如从pkg1.A转换到pkg2.A，将会生成一下方法名

```
convertAToA(xxx) xxx
```

此时你可能需要修改你的结构体名称，或者给结构体指定别名

```
cg.Alias(pkg1.A{}, "Pkg1A")
```

那么生成的方法名就会变成

```
convertPkg1AToA(xxx) xxx
```

## 字段映射

converter generator默认是通过字段名来对两个结构体直接做映射的，但在某些场景想需要将A结构的F1字段映射到B结构的F2字段，那么可以用如下方式做映射

```
cg.Convert(A{}, B{}, converter_generator.ConversionConfig{
			FieldMap: [][]string{
				{"F1", "F2"}
			}
})
```

那么生成的代码就会将B.F2赋值为A.F1



## 方法导出

默认生成的方法名首字母是小写的，可以通过以下方法将方法暴露

```
cg.Convert(A{}, B{}, converter_generator.ConversionConfig{
      Export: true,
})
```



## 指定方法名

当你想指定一个转换方法的方法名时，可以这样配置

```
cg.Convert(A{}, B{}, converter_generator.ConversionConfig{
      Export: true,
      Name: "ConvAB"
})
```

那么生成的方法名就会变成 `ConvAB` 了

