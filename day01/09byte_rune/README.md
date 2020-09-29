# byte和rune类型

GO语言的字符有以下两种：

+ `uint8` 类型，或者叫 `byte` 型，代表了ASCII码的一个字符

+ `rune` 类型，代表一个UTF-8字符

当需要处理中文或者其它复合字符时，需要使用 `rune` 类型。它实际是一个 `int32`。

Go使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 类型进行默认字符串处理。性能和扩展性都有兼顾。