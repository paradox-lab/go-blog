# XPath
> XPath即为XML路径语言（XML Path Language），它是一种用来确定XML文档中某部分位置的语言。

## 语法
> https://www.w3school.com.cn/xpath/xpath_syntax.asp

下面列出了最有用的路径表达式:
表达式  |描述   
-     |-      
nodename|选取此节点的所有子节点
/    | 从根节点选取
//   | 从匹配选择的当前节点选择文档中的节点，而不考虑它们的位置
.    | 选取当前节点
..   | 选取当前节点的父节点
@    | 选取属性

## 实例 
```xml
<?xml version="1.0" encoding="ISO-8859-1"?>

<bookstore>

<book>
  <title lang="eng">Harry Potter</title>
  <price>29.99</price>
</book>

<book>
  <title lang="eng">Learning XML</title>
  <price>39.95</price>
</book>

</bookstore>
```
路径表达式|结果 |   
-|-      |
/bookstore/book[1]         |选取属于 bookstore 子元素的第一个 book 元素。|
/bookstore/book[last()]    |选取属于 bookstore 子元素的最后一个 book 元素。|
/bookstore/book[last()-1]  |选取属于 bookstore 子元素的倒数第二个 book 元素。|
/bookstore/book[position()<3] |选取最前面的两个属于 bookstore 元素的子元素的 book 元素。 |
//title[@lang]             |选取所有拥有名为 lang 的属性的 title 元素。 |
//title[@lang='eng']       |选取所有 title 元素，且这些元素拥有值为 eng 的 lang 属性。 |
/bookstore/book[price>35.00] |选取 bookstore 元素的所有 book 元素，且其中的 price 元素的值须大于 35.00。 |
/bookstore/book[price>35.00]/title |选取 bookstore 元素中的 book 元素的所有 title 元素，且其中的 price 元素的值须大于 35.00。 |
//price[contains(text(),'39.95')] |选取文本包含39.95的price节点元素 |
//price[text()='39.95']     |选取文本为39.95的price节点元素 |