# WeSearch

```text
                                         ----------
                                         |  mysql |
                                         |  mysql |
                                         ----------
                                              |
       ------------                     -------------             ------------- 
       |  manager | -- rpc create doc-> |  user-doc |  -produce-> |   kafka   |
       |  manager | -- rpc search doc-> |  user-doc |  <-consume  |   kafka   |
       ------------                     -------------             -------------
                                              |                         |
                                 rpc retrieve |                         | 
                                              |                         |
       -----------------                -------------   -----------------
       | elasticsearch |   <- store -   | retriever |      <- doc url -
       | elasticsearch |  -retrieve ->  | retriever |      - abstract ->
       -----------------                -------------

```
