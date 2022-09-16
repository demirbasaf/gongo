# gongo

🔵 Easy to use model-based Go ORM for MongoDB

``` 
connection, err := gongo.Connect(ctx, config)
dogModel := gongo.Model("dogs", examples.Dog{}, connection)

max := examples.Dog{Name: "Max"}
dogModel.New(ctx, max)
```
