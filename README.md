# gongo

🔵 Easy to use model-based Go ORM for MongoDB

``` 
connection, err := gongo.Connect(ctx, config)
dogModel := gongo.Model("Dogs", domain.Dog{}, connection)

max := &domain.Dog{Name: "Max"}
dogModel.New(ctx, max)
```
