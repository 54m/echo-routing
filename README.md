# echo-routing / output

# How to use

code
```go
e := echo.New()

...

output.Do(e.Routes())
```

output
```commandline
+--------+-------------+--------------+
| Method | Path        | FunctionName |
+--------+-------------+--------------+
| GET    | /xxx        | router.XXX   |
+        +-------------+--------------+
|        | /yyy        | router.YYY   |
+--------+-------------+--------------+
| POST   | /zzz        | router.ZZZ   |
+--------+-------------+--------------+
| TOTAL  | 3 Endpoints |              |
+--------+-------------+--------------+
```

[example code](./example/main.go)

# LICENSE
[MIT](./LICENSE)
