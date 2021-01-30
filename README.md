# b2w-test

---

# routes
## GET
`/planets`
`/planet/id/:id`
`/planet/name/:name`
## POST
`/new/planet`
## REMOVE
`/delete/planet/:id`

# Data format

```json
    {
    "id": int //do not use,setted automatically, if used it will be overwriten
    "name": string,
    "climate": string,
    "terrain": string,
    "appearances": int //do not use, setted automatically, if used it will be overwriten
}
```