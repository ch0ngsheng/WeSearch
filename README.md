# WeSearch

```text
       ------------                     -------------       ----------
       |  manager |  -- rpc doc url ->  |  user-doc |   --  |  mysql |
       |  manager |                     |  user-doc |       |  mysql |
       ------------                     -------------       ----------
            |                                 |
        rpc | search                  produce | consume
            |                                 |
       -------------                    -------------
       | retriever |    <- doc url --   |   kafka   |
       | retriever |    -- abstract ->  |   kafka   |
       -------------                    -------------
             |       
     -----------------
     | elasticsearch |
     -----------------
```
