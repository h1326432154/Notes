# ElasticSearch

## 查询

### match_all 查询

match_all 查询简单的 匹配所有文档。在没有指定查询方式时，它是默认的查询

* 例

```json
{
    "match_all": {}
}
```

### match 查询

无论你在任何字段上进行的是全文搜索还是精确查询，match 标准查询

```json
{
    "match": {
        "tweet": "About Search"
    }
}
```

### multi_match 查询

可以在多个字段上执行相同的 match 查询：

```json
{
    "multi_match": {
        "query": "full text search",
        "fields": [
            "title",
            "body"
        ]
    }
}
```

### range 查询

查询找出那些落在指定区间内的数字或者时间：

* gt 大于
* lt 小于
* gte 大于等于
* lte 小于等于

```json
{
    "range": {
        "age": {
            "gte": 20,
            "lt": 30
        }
    }
}
```

### term 查询

term 查询被用于精确值 匹配，这些精确值可能是数字、时间、布尔或者那些 not_analyzed 的字符串

```json
{ "term": { "age":    26           }}
{ "term": { "date":   "2014-09-01" }}
{ "term": { "public": true         }}
{ "term": { "tag":    "full_text"  }}
```

### terms 查询

terms 查询和 term 查询一样，但它允许你指定多值进行匹配。如果这个字段包含了指定值中的任何一个值，那么这个文档满足条件

### exists 查询和 missing 查询

exists 查询和 missing 查询被用于查找那些指定字段中有值 (exists) 或无值 (missing) 的文档。这与SQL中的 IS_NULL (missing) 和 NOT IS_NULL (exists) 在本质上具有共性

```json
{
    "exists":   {
        "field":    "title"
    }
}
```
