#Example usage:

> echo '{"s":"ciao"}' | http POST :8000/uppercase

```json
{
    "v": "CIAO"
}
```

> echo '{"s":"ciao"}' | http POST :8000/count

```json
{
    "v": 4
}
```
