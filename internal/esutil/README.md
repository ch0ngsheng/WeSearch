# ES command
```shell
PUT wxindex
PUT wxindex/_mapping
{
    "_source": {
      "enabled": false
    },
    "properties": {
      "content": {
        "type": "text",
        "store":false
      }
    }
}

GET wxindex/_search

POST wxindex/_doc/1001
{
  "content":"some content"
}

{
  "_index": "wxindex",
  "_id": "1005",
  "_version": 1,
  "result": "created",
  "_shards": {
    "total": 2,
    "successful": 1,
    "failed": 0
  },
  "_seq_no": 8,
  "_primary_term": 1
}

GET wxindex/_search
{
  "query": {
    "bool": {
      "should":[
        {"match_phrase": {"content":"keyword1"}},
        {"match_phrase": {"content":"keyword2"}}
      ],
      "filter":{
          "ids":{"values":["1001","1002"]}
      }
    }
  },
  "min_score":0.001
}

{
  "took": 323,
  "timed_out": false,
  "_shards": {
    "total": 1,
    "successful": 1,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 2,
      "relation": "eq"
    },
    "max_score": 4.8995247,
    "hits": [
      {
        "_index": "wxindex",
        "_id": "1003",
        "_score": 4.8995247
      },
      {
        "_index": "wxindex",
        "_id": "1004",
        "_score": 2.1943364
      }
    ]
  }
}
```